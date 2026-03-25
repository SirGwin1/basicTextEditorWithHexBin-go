package functionality

import (
	"strconv"
	"strings"
)

func BaseConversion(str string) string {
	sli := strings.Fields(str)

	for i := 0; i < len(sli); i++ {
		if sli[i] == "(hex)" && i > 0 {
			result, err := strconv.ParseInt(sli[i-1], 16, 64)
			if err == nil {
				sli[i-1] = strconv.Itoa(int(result))
				sli = append(sli[:i], sli[i+1:]...)
				i-- // adjust index after removal
			}
		} else if sli[i] == "(bin)" && i > 0 {
			result, err := strconv.ParseInt(sli[i-1], 2, 64)
			if err == nil {
				sli[i-1] = strconv.Itoa(int(result))
				sli = append(sli[:i], sli[i+1:]...)
				i-- // adjust index after removal
			}
		}
	}

	return strings.Join(sli, " ")
}
