package faker

import (
	"math/rand"
	"strings"
)

// NewRandomCPF returns a new valid CPF
func NewRandomCPF() string {
	return newRandomCPF()
}

func NewRandomBRPhone(mobile bool) string {
	outp := strings.Builder{}
	outp.WriteString("55")
	rsrc := "0123456789"
	outp.WriteByte(rsrc[1+rand.Intn(len(rsrc)-1)])
	outp.WriteByte(rsrc[1+rand.Intn(len(rsrc)-1)])
	if mobile {
		outp.WriteString("9")
	}
	for i := 0; i < 8; i++ {
		outp.WriteByte(rsrc[rand.Intn(len(rsrc))])
	}
	return outp.String()
}

func newRandomCPF() string {
	dgts := make([]int, 9)
	dcode := make([]int, 2)
	for i := range dgts {
		dgts[i] = rand.Intn(10)
	}
	sum := 0
	for i := 10; i > 1; i-- {
		sum += dgts[10-i] * i
	}
	if sum%11 < 2 {
		dcode[0] = 0
	} else {
		dcode[0] = 11 - sum%11
	}
	dgts = append(dgts, dcode[0])
	sum = 0
	for i := 11; i > 1; i-- {
		sum += dgts[11-i] * i
	}
	if sum%11 < 2 {
		dcode[1] = 0
	} else {
		dcode[1] = 11 - sum%11
	}
	dgts = append(dgts, dcode[1])
	sb := strings.Builder{}
	for _, v := range dgts {
		sb.WriteRune(runeIV(v))
	}
	return sb.String()
}

func runeIV(r int) rune {
	const runeTable = "0123456789"
	if r >= 0 && r < 10 {
		return rune(runeTable[r])
	}
	return 0
}
