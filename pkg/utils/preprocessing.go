package utils

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
)

const (
	MinLength = 8
	MaxLength = 10
)

func PreprocessCoupons(inputFiles []string, outputFile string) error {
	globalCount := make(map[string]int)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, filePath := range inputFiles {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			localSet := make(map[string]bool)

			f, err := os.Open(path)
			if err != nil {
				fmt.Printf("Error opening %s: %v\n", path, err)
				return
			}
			defer f.Close()

			gz, err := gzip.NewReader(f)
			if err != nil {
				fmt.Printf("Error reading gzip %s: %v\n", path, err)
				return
			}
			defer gz.Close()

			scanner := bufio.NewScanner(gz)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if len(line) >= MinLength && len(line) <= MaxLength {
					localSet[line] = true
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Printf("Scanner error in %s: %v\n", path, err)
				return
			}

			mu.Lock()
			for code := range localSet {
				globalCount[code]++
			}
			mu.Unlock()
		}(filePath)
	}

	wg.Wait()

	validCoupons := make(map[string]int)
	for code, count := range globalCount {
		if count >= 2 {
			validCoupons[code] = 2
		}
	}

	out, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("could not create output JSON: %v", err)
	}
	defer out.Close()

	if err := json.NewEncoder(out).Encode(validCoupons); err != nil {
		return fmt.Errorf("could not encode JSON: %v", err)
	}

	fmt.Printf("✅ Preprocessing done. %d valid codes saved to %s\n", len(validCoupons), outputFile)
	return nil
}
