package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"mIndex/analyzer/mkv"
	"os"
	"path/filepath"
)

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}

	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

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

type MediaFileInformation struct {
	path string
	size int64
	hash []byte
}

func (mfi *MediaFileInformation) String() string {
	return fmt.Sprintf("%s size: %s hash: %x", mfi.path, ByteCountIEC(mfi.size), mfi.hash)
}

func (crawler *mediaCrawler) handleFile(path string, info os.FileInfo, err error) error {
	if err != nil {
		return fmt.Errorf("error handling file: %w", err)
	}

	if info.IsDir() {
		return nil
	}

	ext := filepath.Ext(path)

	if ext == ".mkv" {
		mfi := MediaFileInformation{
			path: path,
			size: info.Size(),
		}

		if crawler.configuration.computeHash {
			hash, hashErr := hashFile(path)
			if hashErr != nil {
				return fmt.Errorf("error hashing file: %w", hashErr)
			}

			mfi.hash = hash
		}

		fmt.Println(mfi.String())
		mkv.Parse(path)
		// getTags(path)
	}

	return nil
}

type crawlerConfiguration struct {
	computeHash bool
}

type mediaCrawler struct {
	configuration *crawlerConfiguration
}

func main() {
	hashPtr := flag.Bool("hash", false, "compute hash for media file deduplication")

	flag.Parse()

	config := &crawlerConfiguration{
		computeHash: *hashPtr,
	}

	crawler := &mediaCrawler{
		configuration: config,
	}

	err := filepath.Walk("/Volumes/media/movies", crawler.handleFile)

	if err != nil {
		log.Fatal(err)
	}
}
