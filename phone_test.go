package xstrings_test

import (
	"testing"

	"github.com/pedidopago/xstrings"
	"github.com/stretchr/testify/assert"
)

func TestExtractPhoneNumberData(t *testing.T) {
	res := xstrings.ExtractPhoneNumberData("+5511984539231", xstrings.WithValidateMobilePrefix(true))
	assert.True(t, res.IsValid)
	assert.Equal(t, "55", res.CountryCode)
	assert.Equal(t, "BR", res.CountryISO2)
	assert.Equal(t, "BRA", res.CountryISO3)
	assert.Equal(t, "Brazil", res.CountryName)

	res = xstrings.ExtractPhoneNumberData("+1-202-555-0102", xstrings.WithValidateMobilePrefix(true))
	assert.True(t, res.IsValid)
	assert.Equal(t, "1", res.CountryCode)
	assert.Equal(t, "US", res.CountryISO2)
	assert.Equal(t, "USA", res.CountryISO3)
	assert.Equal(t, "United States", res.CountryName)

	res = xstrings.ExtractPhoneNumberData("55 23 - 984539231")
	assert.True(t, res.IsValid)
	assert.Equal(t, "55", res.CountryCode)
	assert.Equal(t, "BR", res.CountryISO2)
	assert.Equal(t, "BRA", res.CountryISO3)
	assert.Equal(t, "Brazil", res.CountryName)
	assert.Equal(t, "5523984539231", res.PhoneNumber)
}
