package pkg

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func NameGenerate(filename, ext string, rng []int, limit int, dir string) string {
	rangeString, err := RangeFormat(rng)
	if err != nil {
		return fmt.Sprintf("%s[].%s", filename, ext)
	}

	fullname := fmt.Sprintf("%s[%s].%s", filename, rangeString, ext)
	fullPath := filepath.Join(dir, fullname)

	// パス全体が制限以内ならそのまま返す
	if len([]byte(fullPath)) <= limit {
		return fullname
	}

	basePattern := fmt.Sprintf("%s[].%s", filename, ext)
	available := limit - len([]byte(filepath.Join(dir, basePattern)))

	if available > 0 {
		truncatedRange := truncateToByteLength(rangeString, available)

		idx := strings.LastIndex(truncatedRange, ",")
		if idx != -1 {
			truncatedRange = truncatedRange[:idx] + "他"
		} else {
			truncatedRange += "他"
		}

		return fmt.Sprintf("%s[%s].%s", filename, truncatedRange, ext)
	}

	filename = truncateToByteLength(filename, 30)
	return fmt.Sprintf("%s[他].%s", filename, ext)
}

func truncateToByteLength(s string, maxBytes int) string {
	result := ""
	count := 0
	for _, r := range s {
		b := utf8.RuneLen(r)
		if count+b > maxBytes {
			break
		}
		result += string(r)
		count += b
	}
	return result
}
