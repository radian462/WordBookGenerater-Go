package pkg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func RangeParse(rng string) ([]int, error) {
	for _, r := range rng {
		if !unicode.IsDigit(r) && r != ',' && r != '-' && r != '~' {
			return nil, errors.New(fmt.Sprintf("無効な文字が含まれています: '%c'", r))
		}
	}

	parts := strings.Split(rng, ",")
	var result []int

	for _, part := range parts {
		if strings.Contains(part, "-") || strings.Contains(part, "~") {
			var rangeParts []string
			if strings.Contains(part, "-") {
				rangeParts = strings.Split(part, "-")
			} else {
				rangeParts = strings.Split(part, "~")
			}

			start, err1 := strconv.Atoi(rangeParts[0])
			end, err2 := strconv.Atoi(rangeParts[1])

			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("数値変換エラー")
			}

			if start > end {
				start, end = end, start
			}

			for i := start; i <= end; i++ {
				result = append(result, i)
			}
		} else {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("数値変換エラー")
			}
			result = append(result, num)
		}
	}
	return result, nil

}
