package pkg

import (
    "path/filepath"
    "os"
    "encoding/csv"
)

func ExtractCSVRange(basePath string, outDir string, outName string, rng []int, header []string) (string, error) {
    var outPath string
    if outName == "" {
        autoName := NameGenerate(filepath.Base(basePath), "csv", rng, 255)
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

    if header != nil && len(header) != 0{
        if err := writer.Write(header); err != nil {
            return "", err
        }
    }

    for _, r := range rng {
        if r > 0 && r < len(rows) {
            if err := writer.Write(rows[r-1]); err != nil {
                return "", err
            }
        }
    }

    return outPath, nil
}