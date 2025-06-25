package pkg

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func writeRange(builder *strings.Builder, start, end, limit int) {
	if start == end {
		builder.WriteString(strconv.Itoa(start))
	} else if start+1 == end {
		builder.WriteString(fmt.Sprintf("%d,%d", start, end))
	} else {
		builder.WriteString(fmt.Sprintf("%d~%d", start, end))
	}
}

func RangeFormat(rng []int, limit int) (string, error) {
	if len(rng) == 0 {
		return "", nil
	}

	sort.Ints(rng)

	var builder strings.Builder
	start, end := rng[0], rng[0]

	for i := 1; i < len(rng); i++ {
		if rng[i] == end+1 {
			end = rng[i]
		} else {
			writeRange(&builder, start, end, limit)
			builder.WriteString(",")

			start, end = rng[i], rng[i]
		}
	}

	writeRange(&builder, start, end, limit)

	return builder.String(), nil
}
