package pkg

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

func ExtractCSVRange(basePath string, outDir string, outName string, rng []int, header []string) (string, error) {
	var outPath string
	base := filepath.Base(basePath)
	ext := filepath.Ext(base)
	if outName == "" {
		autoName := NameGenerate(base[:len(base)-len(ext)], "csv", rng, 255)
		outPath = filepath.Join(outDir, autoName)
	} else {
		outPath = filepath.Join(outDir, outName)
	}

	file, err := os.Open(basePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(outDir, 0755); err != nil {
		return "", err
	}
	outFile, err := os.Create(outPath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	if header != nil && len(header) != 0 {
		if err := writer.Write(header); err != nil {
			return "", err
		}
	}

	for _, r := range rng {
		if r > 0 && r < len(rows) {
			row := rows[r-1]
			if header != nil && len(header) > len(row) {
				padded := make([]string, len(header))
				copy(padded, row)
				row = padded
			}
			if err := writer.Write(row); err != nil {
				return "", err
			}
		}
	}

	return outPath, nil
}
