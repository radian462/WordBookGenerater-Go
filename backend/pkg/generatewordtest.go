package pkg

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"github.com/xuri/excelize/v2"
	"path/filepath"
	"strings"
)

func FormatString(str string, maxWidth int) string {
	if str == "" {
		return ""
	}

	w := runewidth.StringWidth(str)
	if w <= maxWidth {
		return str
	}

	lines := []string{""}
	var nowWidth int

	for _, r := range str {
		linesLastIdx := len(lines) - 1
		rw := runewidth.RuneWidth(r)
		if nowWidth+rw <= maxWidth {
			lines[linesLastIdx] += string(r)
			nowWidth += rw
		} else {
			lines = append(lines, string(r))
			nowWidth = rw
		}
	}

	// 行が3を超えるときは強制2分割
	if len(lines) >= 3 {
		totalWidth := runewidth.StringWidth(str)
		targetWidth := totalWidth / 2

		lines = []string{"", ""}
		for _, r := range str {
			if runewidth.StringWidth(lines[0]) < targetWidth {
				lines[0] += string(r)
			} else {
				lines[1] += string(r)
			}
		}
	}

	return strings.Join(lines, "\n")
}

func GenerateWordTest(template string, outDir string, outName string, words, answers []string) (string, error) {
	if len(words) > 50 {
		words = words[:50]
	}
	if len(answers) > 50 {
		answers = answers[:50]
	}

	length := max(len(words), len(answers))

	f, err := excelize.OpenFile(template)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("Failed to close file: %v\n", err)
		}
	}()

	wordsPerColumn := 25
	TestsheetName := f.GetSheetName(0)
	AnswersheetName := f.GetSheetName(1)

	colMaxWidths := make(map[string]float64)

	for i := 0; i < length; i++ {
		colOffset := (i / wordsPerColumn) * 4
		col := colOffset + 1
		row := (i % wordsPerColumn) + 1

		colStr, _ := excelize.ColumnNumberToName(col)
		cell1 := fmt.Sprintf("%s%d", colStr, row)
		f.SetCellValue(TestsheetName, cell1, i+1)
		f.SetCellValue(AnswersheetName, cell1, i+1)

		colStr2, _ := excelize.ColumnNumberToName(col + 1)
		cell2 := fmt.Sprintf("%s%d", colStr2, row)
		if i < len(words) {
			f.SetCellValue(TestsheetName, cell2, words[i])
			f.SetCellValue(AnswersheetName, cell2, words[i])
		}

		colStr3, _ := excelize.ColumnNumberToName(col + 2)
		cell3 := fmt.Sprintf("%s%d", colStr3, row)
		if i < len(answers) {
			ans := FormatString(answers[i], 25)
			f.SetCellValue(AnswersheetName, cell3, ans)

			for _, l := range strings.Split(ans, "\n") {
				w := float64(runewidth.StringWidth(l))
				if w > colMaxWidths[colStr3] {
					colMaxWidths[colStr3] = w
				}
			}
		}
	}

	for colStr, w := range colMaxWidths {
		f.SetColWidth(AnswersheetName, colStr, colStr, w)
	}

	outPath := filepath.Join(outDir, outName)
	if err := f.SaveAs(outPath); err != nil {
		return "", err
	}

	return outPath, nil
}
