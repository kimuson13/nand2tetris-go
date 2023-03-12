package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateHack(path string, lines []string) error {
	dir, fileName := filepath.Split(path)
	ext := filepath.Ext(path)

	hackFileName := strings.ReplaceAll(fileName, ext, ".hack")
	hackFilePath := filepath.Join(dir, hackFileName)

	f, err := os.Create(hackFilePath)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("create hack file error: %w", err)
	}

	body := genBody(lines)
	if _, err := f.Write(body); err != nil {
		os.Remove(f.Name())
		return fmt.Errorf("create hack file error: %w", err)
	}

	return nil
}

func genBody(lines []string) []byte {
	var s string
	for _, line := range lines {
		if line != "" {
			s += fmt.Sprintf("%s\r\n", line)
		}
	}

	return []byte(s)
}
