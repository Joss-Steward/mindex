package analyzer

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func hashFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open '%s': %w", path, err)
	}

	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, fmt.Errorf("could not hash '%s': %w", path, err)
	}

	return hash.Sum(nil), nil
}
