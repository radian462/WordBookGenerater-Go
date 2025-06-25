package pkg

import (
    "fmt"
    "strings"
)

func NameGenerate(filename, ext string, rng []int, limit int) string {
    rangeString, err := RangeFormat(rng)
    if err != nil {
        return fmt.Sprintf("%s [].%s", filename, ext)
    }

    fullname := fmt.Sprintf("%s [%s].%s", filename, rangeString, ext)

    if len(fullname) > limit {
        rangeLimit := limit - len(fmt.Sprintf("%s [].%s", filename, ext))
        if rangeLimit < 0 {
            rangeLimit = 0
        }

        limitedRange := string([]rune(rangeString)[:rangeLimit])

        idx := strings.LastIndex(limitedRange, ",")
        if idx != -1 {
            limitedRange = limitedRange[:idx] + "他"
        } else {
            limitedRange += "他"
        }

        return fmt.Sprintf("%s [%s].%s", filename, limitedRange, ext)
    }

    return fullname
}
