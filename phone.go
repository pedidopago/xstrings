package xstrings

import (
	_ "embed"
	"encoding/json"
	"strings"
)

// JSON Data FROM:
// https://github.com/AfterShip/phone/blob/master/src/data/country_phone_data.ts

var (
	//go:embed phonedata.json
	phoneDataJSON []byte
	phoneData     []CountryPhoneData
)

type CountryPhoneData struct {
	Alpha2             string   `json:"alpha2"`
	Alpha3             string   `json:"alpha3"`
	CountryCode        string   `json:"country_code"`
	CountryName        string   `json:"country_name"`
	MobileBeginWith    []string `json:"mobile_begin_with"`
	PhoneNumberLengths []int    `json:"phone_number_lengths"`
}

type extractPhoneNumberDataOptions struct {
	country              string
	validateMobilePrefix bool
	strictDetection      bool
}

type ExtractPhoneNumberDataOption func(*extractPhoneNumberDataOptions)

func WithValidateMobilePrefix(validateMobilePrefix bool) ExtractPhoneNumberDataOption {
	return func(options *extractPhoneNumberDataOptions) {
		options.validateMobilePrefix = validateMobilePrefix
	}
}

// WithCountry tries to match the phone with an ISO3 or ISO2 country code
func WithCountry(country string) ExtractPhoneNumberDataOption {
	return func(options *extractPhoneNumberDataOptions) {
		options.country = country
	}
}

func WithStrictDetection(strictDetection bool) ExtractPhoneNumberDataOption {
	return func(options *extractPhoneNumberDataOptions) {
		options.strictDetection = strictDetection
	}
}

type ExtractPhoneNumberDataResult struct {
	IsValid     bool   `json:"is_valid"`
	PhoneNumber string `json:"phone_number"`
	CountryCode string `json:"country_code"`
	CountryISO2 string `json:"country_iso2"`
	CountryISO3 string `json:"country_iso3"`
	CountryName string `json:"country_name"`
}

// ExtractPhoneNumberData extracts the country information based on the phone number's country code.
//
// Example:
// 		ExtractPhoneNumberData("+1-555-555-5555")
// 		// returns {IsValid: true, PhoneNumber: "15555555555", CountryCode: "1", CountryISO2: "US", CountryISO3: "USA", CountryName: "United States"}
func ExtractPhoneNumberData(phoneNumber string, options ...ExtractPhoneNumberDataOption) ExtractPhoneNumberDataResult {
	opts := &extractPhoneNumberDataOptions{}
	for _, v := range options {
		v(opts)
	}
	opts.country = strings.TrimSpace(opts.country)
	hasPlusSign := strings.ContainsRune(phoneNumber, '+')
	processedPhoneNumber := NormalizeNumericStr(phoneNumber)
	invalidResult := ExtractPhoneNumberDataResult{
		IsValid: false,
	}

	if opts.country != "" {
		ccfind := findPhoneDataByCountry(opts.country)
		if ccfind == nil {
			return invalidResult
		}
		if ccfind.Alpha3 == "CIV" || ccfind.Alpha3 == "COG" {
			processedPhoneNumber = strings.TrimPrefix(processedPhoneNumber, "0")
		}
		// if input 89234567890, RUS, remove the 8
		if ccfind.Alpha3 == "RUS" && len(processedPhoneNumber) == 11 && strings.HasPrefix(processedPhoneNumber, "89") {
			processedPhoneNumber = strings.TrimPrefix(processedPhoneNumber, "8")
		}

		// if there's no plus sign and the phone number length is one of the valid length under country phone data
		// then assume there's no country code, hence add back the country code
		if !hasPlusSign {
			for _, l := range ccfind.PhoneNumberLengths {
				if l == len(processedPhoneNumber) {
					processedPhoneNumber = ccfind.CountryCode + processedPhoneNumber
					break
				}
			}
		}

		return ExtractPhoneNumberDataResult{
			IsValid:     true,
			PhoneNumber: processedPhoneNumber,
			CountryCode: ccfind.CountryCode,
			CountryISO2: ccfind.Alpha2,
			CountryISO3: ccfind.Alpha3,
			CountryName: ccfind.CountryName,
		}
	}
	if hasPlusSign || (len(processedPhoneNumber) >= 7 && processedPhoneNumber[0] != '0') {
		// if there is a plus sign but no country provided
		// try to find the country phone data by the phone number
		result, exact := findCountryPhoneDataByPhoneNumber(processedPhoneNumber, opts.validateMobilePrefix)
		if exact && result != nil {
			return ExtractPhoneNumberDataResult{
				IsValid:     true,
				PhoneNumber: processedPhoneNumber,
				CountryCode: result.CountryCode,
				CountryISO2: result.Alpha2,
				CountryISO3: result.Alpha3,
				CountryName: result.CountryName,
			}
		}
		if !exact && result != nil && !opts.strictDetection {
			// for some countries, the phone number usually includes one trunk prefix for local use
			// The UK mobile phone number ‘07911 123456’ in international format is ‘+44 7911 123456’, so without the first zero.
			// 8 (AAA) BBB-BB-BB, 0AA-BBBBBBB
			// the numbers should be omitted in international calls
			processedPhoneNumber = result.CountryCode + processedPhoneNumber[len(result.CountryCode)+1:]
			return ExtractPhoneNumberDataResult{
				IsValid:     true,
				PhoneNumber: processedPhoneNumber,
				CountryCode: result.CountryCode,
				CountryISO2: result.Alpha2,
				CountryISO3: result.Alpha3,
				CountryName: result.CountryName,
			}
		}
	}
	return invalidResult
}

// transpiled from https://github.com/AfterShip/phone/blob/da5e654453d309f468f654a21694648fa3c30b11/src/lib/utility.ts#L91
func findCountryPhoneDataByPhoneNumber(phoneNumber string,
	validateMobilePrefix bool) (result *CountryPhoneData, isExactMatch bool) {
	phoneNumber = NormalizeNumericStr(phoneNumber)
	var possibleMatch *CountryPhoneData
	for _, v := range phoneData {
		if !strings.HasPrefix(phoneNumber, v.CountryCode) {
			continue
		}
		if matchExactCountryPhoneData(phoneNumber, validateMobilePrefix, v) {
			copy1 := v
			return &copy1, true
		}
		if matchPossibleCountryPhoneData(phoneNumber, validateMobilePrefix, v) {
			copy1 := v
			possibleMatch = &copy1
		}
	}
	if possibleMatch != nil {
		return possibleMatch, false
	}
	return nil, false
}

func matchExactCountryPhoneData(phoneNumber string, validateMobilePrefix bool, d CountryPhoneData) bool {
	// check if the phone number length match any one of the length config
	phoneNumberLengthMatched := false
	plen := len(phoneNumber)
	for _, v := range d.PhoneNumberLengths {
		if plen == v+len(d.CountryCode) || plen == v {
			phoneNumberLengthMatched = true
			break
		}
	}
	if !phoneNumberLengthMatched {
		return false
	}
	// if no need to validate mobile prefix or the country data does not have mobile begin with
	// pick the current one as the answer directly
	if len(d.MobileBeginWith) < 1 || !validateMobilePrefix {
		return true
	}
	// if the mobile begin with is correct, pick as the correct answer
	for _, v := range d.MobileBeginWith {
		if strings.HasPrefix(phoneNumber[len(d.CountryCode):], v) {
			return true
		}
	}
	return false
}

func matchPossibleCountryPhoneData(phoneNumber string, validateMobilePrefix bool, d CountryPhoneData) bool {
	// check if the phone number length match any one of the length config
	phoneNumberLengthMatched := false
	plen := len(phoneNumber)
	cclen := len(d.CountryCode)
	for _, v := range d.PhoneNumberLengths {
		// the phone number must include the country code
		// countryPhoneDatum.phone_number_lengths is the length without country code
		// + 1 is assuming there is an unwanted trunk code prepended to the phone number
		if cclen+v+1 == plen {
			phoneNumberLengthMatched = true
			break
		}
	}
	if !phoneNumberLengthMatched {
		return false
	}
	// if no need to validate mobile prefix or the country data does not have mobile begin with
	// pick the current one as the answer directly
	if len(d.MobileBeginWith) < 1 || !validateMobilePrefix {
		return true
	}
	// if the mobile begin with is correct, pick as the correct answer
	// match another \d for the unwanted trunk code prepended to the phone number
	for _, v := range d.MobileBeginWith {
		if strings.HasPrefix(phoneNumber[len(d.CountryCode):], v) {
			return true
		}
		if strings.HasPrefix(phoneNumber[len(d.CountryCode)+1:], v) {
			return true
		}
	}
	return false
}

func findPhoneDataByCountry(country string) *CountryPhoneData {
	if len(country) == 3 {
		cc := strings.ToUpper(country)
		for _, v := range phoneData {
			vcopy := v
			if v.Alpha3 == cc {
				return &vcopy
			}
		}
	}
	if len(country) == 2 {
		cc := strings.ToUpper(country)
		for _, v := range phoneData {
			vcopy := v
			if v.Alpha2 == cc {
				return &vcopy
			}
		}
	}
	for _, v := range phoneData {
		vcopy := v
		if v.CountryName == country {
			return &vcopy
		}
	}
	return nil
}

func init() {
	phoneData = make([]CountryPhoneData, 0)
	if err := json.Unmarshal(phoneDataJSON, &phoneData); err != nil {
		panic(err)
	}
}
