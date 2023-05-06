package xstrings

import (
	"strings"

	"github.com/forPelevin/gomoji"
)

func FirstName(v string) string {
	// remove special characters
	v = Clean(v)
	// remove emojis
	v = gomoji.RemoveEmojis(v)
	v = strings.Replace(v, ",", " ", -1)
	v = strings.Replace(v, "  ", " ", -1)
	v = strings.Replace(v, "  ", " ", -1)
	vs := strings.Split(v, " ")
	return vs[0]
}

func RemoveEmojis(v string) string {
	return gomoji.RemoveEmojis(v)
}

func ContainsEmoji(v string) bool {
	return gomoji.ContainsEmoji(v)
}

// func isMnOrDingbats(r rune) bool {
// 	if isMn(r) {
// 		return true
// 	}
// 	// dingbats
// 	if r >= 0x2700 && r <= 0x27BF {
// 		return true
// 	}
// 	switch r {
// 	case '.', ',', ';', ':', '!', '?', '(', ')', '[', ']', '{', '}', '/', '\\', '"', '\'', '\t', '\n', '\r', '\v', '\f', '\a', '\b', '\000':
// 		return true
// 	}
// 	return false
// }

var allowedRanges = []struct {
	from rune
	to   rune
}{
	{
		from: 0x0041,
		to:   0x005A,
	},
	{
		from: 0x0061,
		to:   0x007A,
	},
	{
		from: 0x00e0,
		to:   0x00e3,
	},
	{
		from: 0x00e7,
		to:   0x00ef,
	},
	{
		from: 0x00f2,
		to:   0x00f5,
	},
	{
		from: 0x00f9,
		to:   0x00fb,
	},
}

// TODO: continue at block u0400+ https://www.compart.com/en/unicode/block/U+0400
var inverseReplacements = []struct {
	matches     string
	replacement string
}{
	{
		matches:     "\u00c0\u00c1\u00c2\u00c3\u00c4\u00c5\u0100\u0101\u0102\u0103\u0104\u0105\u01cd\u01ce\u01de\u01df\u01e0\u01e1\u01fa\u01fb\u0200\u0201\u0202\u0203\u0226\u0227\u023a\u0245\u0250\u0251\u0252\u028c\u0386\u0391\u039b\u1d00",
		replacement: "A",
	},
	{
		matches:     "\u00c6\u01e2\u01e3\u01fc\u01fd",
		replacement: "AE",
	},
	{
		matches:     "\u0180\u0181\u0182\u0183\u0184\u0185\u0243\u0299\u0392",
		replacement: "B",
	},
	{
		matches:     "\u00c7\u0106\u0107\u0108\u0109\u010a\u010b\u010c\u010d\u0187\u0188\u023b\u023c\u03f2\u03f9\u1d04",
		replacement: "C",
	},
	{
		matches:     "\u00d0\u010e\u010f\u0110\u0111\u0189\u018a\u018b\u018c\u0221\u1d05",
		replacement: "D",
	},
	{
		matches:     "\u0238",
		replacement: "DB",
	},
	{
		matches:     "\u00c8\u00c9\u00ca\u00cb\u0112\u0113\u0114\u0115\u0116\u0117\u0118\u0119\u011a\u011b\u018e\u018f\u0190\u01a9\u01dd\u0204\u0205\u0206\u0207\u0228\u0229\u0246\u0247\u0388\u0395\u03a3\u1d07",
		replacement: "E",
	},
	{
		matches:     "\u0191\u0192",
		replacement: "F",
	},
	{
		matches:     "\u011c\u011d\u011e\u011f\u0120\u0121\u0122\u0123\u0193\u01e4\u01e5\u01e6\u01e7\u01f4\u01f5\u0261\u0262",
		replacement: "G",
	},
	{
		matches:     "\u0124\u0125\u0126\u0127\u021e\u021f\u029c\u02b0\u02b1\u0389\u0397",
		replacement: "H",
	},
	{
		matches:     "\u01f6",
		replacement: "HU",
	},
	{
		matches:     "\u00cc\u00cd\u00ce\u00cf\u0128\u0129\u012a\u012b\u012c\u012d\u012e\u012f\u0130\u0131\u0197\u019a\u01cf\u01d0\u0208\u0209\u020a\u020b\u026a\u038a\u0399",
		replacement: "I",
	},
	{
		matches:     "\u0132\u0133",
		replacement: "IJ",
	},
	{
		matches:     "\u0134\u0135\u01f0\u0237\u0248\u0249\u02b2\u037f",
		replacement: "J",
	},
	{
		matches:     "\u0136\u0137\u0138\u0198\u0199\u01e8\u01e9\u039a",
		replacement: "K",
	},
	{
		matches:     "\u0139\u013a\u013b\u013c\u013d\u013e\u013f\u0140\u0141\u0142\u0196\u0234\u023d\u029f",
		replacement: "L",
	},
	{
		matches:     "\u028d\u039c\u03fa\u03fb",
		replacement: "M",
	},
	{
		matches:     "\u00d1\u0143\u0144\u0145\u0146\u0147\u0148\u0149\u014a\u014b\u019d\u019e\u01f8\u01f9\u0220\u0235\u0274\u039d",
		replacement: "N",
	},
	{
		matches:     "\u00d2\u00d3\u00d4\u00d5\u00d6\u00d8\u014c\u014d\u014e\u014f\u0150\u0151\u0186\u019f\u01a0\u01a1\u01d1\u01d2\u01fe\u01ff\u020c\u020d\u020e\u020f\u022a\u022b\u022c\u022d\u022e\u022f\u0230\u0231\u0275\u038c\u0398\u039f\u1d0f",
		replacement: "O",
	},
	{
		matches:     "\u0152\u0153",
		replacement: "OE",
	},
	{
		matches:     "\u00de\u01a4\u01a5\u01f7\u03a1",
		replacement: "P",
	},
	{
		matches:     "\u01ea\u01eb\u01ec\u01ed\u024a\u024b",
		replacement: "Q",
	},
	{
		matches:     "\u0154\u0155\u0156\u0157\u0158\u0159\u01a6\u0210\u0211\u0212\u0213\u024c\u024d\u0280\u02b3\u02f9",
		replacement: "R",
	},
	{
		matches:     "\u015a\u015b\u015c\u015d\u015e\u015f\u0160\u0161\u017f\u01a7\u01a8\u0218\u0219\u023f\u02e2\ua731",
		replacement: "S",
	},
	{
		matches:     "\u00df",
		replacement: "SS",
	},
	{
		matches:     "\u0162\u0163\u0164\u0165\u0166\u0167\u01aa\u01ab\u01ac\u01ad\u01ae\u021a\u021b\u0236\u023e\u03a4",
		replacement: "T",
	},
	{
		matches:     "\u00d9\u00da\u00db\u00dc\u0168\u0169\u016a\u016b\u016c\u016d\u016e\u016f\u0170\u0171\u0172\u0173\u01af\u01b0\u01d3\u01d4\u01d5\u01d6\u01d7\u01d8\u01d9\u01da\u01db\u01dc\u0214\u0215\u0216\u0217\u0244",
		replacement: "U",
	},
	{
		matches:     "\u1d20",
		replacement: "V",
	},
	{
		matches:     "\u0174\u0175\u019c\u02b7",
		replacement: "W",
	},
	{
		matches:     "\u00d7\u02e3\u03a7",
		replacement: "X",
	},
	{
		matches:     "\u00dd\u0176\u0177\u0178\u01b3\u01b4\u0232\u0233\u024e\u024f\u028f\u02b8\u038e\u03a5",
		replacement: "Y",
	},
	{
		matches:     "\u0179\u017a\u017b\u017c\u017d\u017e\u01b5\u01b6\u0224\u0225\u0240\u0290\u0396",
		replacement: "Z",
	},
}

var conditionalReplacements = []struct {
	from        rune
	to          rune
	replacement func(rune) string
}{
	{
		// https://www.compart.com/en/unicode/block/U+1D400
		from: 0x1D400,
		to:   0x1D419,
		replacement: func(r rune) string {
			return string(r - 0x1D400 + 0x41)
		},
	},
	{
		// https://www.compart.com/en/unicode/block/U+1D400
		from: 0x1D41a,
		to:   0x1D433,
		replacement: func(r rune) string {
			return string(r - 0x1D41a + 0x61)
		},
	},
	{
		// https://www.compart.com/en/unicode/block/U+1D400
		from: 0x1D434,
		to:   0x1D44d,
		replacement: func(r rune) string {
			return string(r - 0x1D434 + 0x41)
		},
	},
	{
		// https://www.compart.com/en/unicode/block/U+1D400
		from: 0x1d44e,
		to:   0x1d467,
		replacement: func(r rune) string {
			return string(r - 0x1d44e + 0x61)
		},
	},
	{
		// https://www.compart.com/en/unicode/block/U+1D400
		from: 0x1d468,
		to:   0x1d481,
		replacement: func(r rune) string {
			return string(r - 0x1d468 + 0x41)
		},
	},
	{
		// https://www.compart.com/en/unicode/block/U+1D400
		from: 0x1d482,
		to:   0x1d49b,
		replacement: func(r rune) string {
			return string(r - 0x1d482 + 0x61)
		},
	},
	//TODO: continue from 0x1d49b+1
}

var invreplMap map[rune]string

func unicodeNameRemap(name string) string {
	sb := new(strings.Builder)
	for _, r := range name {
		if v, ok := invreplMap[r]; ok {
			sb.WriteString(v)
		} else {
			sb.WriteRune(' ')
		}
	}
	return sb.String()
}

func NormalizeForNameExcludingInvalidChars(v string) string {

	// t := transform.Chain(norm.NFKD, runes.Remove(containsRuneFunc(isMnOrDingbats)), norm.NFC)
	// t := transform.Chain(norm.NFKD, transform.RemoveFunc(isMn), norm.NFC)
	// v, _, _ = transform.String(t, v)
	vf := strings.TrimSpace(NormalizeForName(RemoveEmojis(v)))

	vf = unicodeNameRemap(vf)
	fields := strings.Fields(vf)
	for i, f := range fields {
		if len(f) > 1 {
			fields[i] = strings.ToUpper(f[0:1]) + strings.ToLower(f[1:])
		} else {
			fields[i] = strings.ToUpper(f)
		}
	}
	vf = strings.Join(fields, " ")
	return vf
}

func init() {
	invreplMap = make(map[rune]string)
	for _, r := range inverseReplacements {
		for _, r2 := range r.matches {
			invreplMap[r2] = r.replacement
		}
	}
	for _, item := range allowedRanges {
		for i := item.from; i <= item.to; i++ {
			invreplMap[i] = string(i)
		}
	}
	for _, item := range conditionalReplacements {
		for i := item.from; i <= item.to; i++ {
			invreplMap[i] = item.replacement(i)
		}
	}
}
