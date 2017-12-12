/*
Package bob implements Bob answers on remarks.
*/
package bob

import (
	"strings"
)

// Hey returns an answer on the given remark
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	isQuestion := strings.HasSuffix(remark, "?")
	isUppercase := strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
	isEmpty := remark == ""

	switch {
	case isQuestion && isUppercase:
		return "Calm down, I know what I'm doing!"
	case isUppercase:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	case isEmpty:
		return "Fine. Be that way!"
	}
	return "Whatever."
}
