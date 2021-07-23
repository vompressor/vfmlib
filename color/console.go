package color

import (
	"strconv"
	"strings"
)

type Att int
type Atts []Att

const escapeStart = "\x1b["
const escapeEnd = "m"

const Reset Att = 0

func NewAtt(att ...Att) Atts {
	return att
}

func (a Atts) ColorString(s string) string {
	if len(a) == 0 {
		return s
	}
	z := make([]string, len(a))
	for i, at := range a {
		z[i] = strconv.Itoa(int(at))
	}
	return escapeStart + strings.Join(z, ";") + escapeEnd + s + escapeStart + strconv.Itoa(int(Reset)) + escapeEnd
}
