package services

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

type CouponService interface {
	ValidateCoupon(code string) (bool, error)
}

type couponService struct {
	once      sync.Once
	coupons   map[string]int // map[coupon]count
	filePaths []string
}

func NewCouponService(filePaths []string) CouponService {
	return &couponService{
		filePaths: filePaths,
		coupons:   make(map[string]int),
	}
}

func (cs *couponService) loadCoupons() error {
	var loadErr error
	cs.once.Do(func() {
		couponCount := make(map[string]int)
		mu := sync.Mutex{}
		wg := sync.WaitGroup{}
		errChan := make(chan error, len(cs.filePaths))

		for _, filePath := range cs.filePaths {
			wg.Add(1)
			go func(path string) {
				defer wg.Done()

				localSet := make(map[string]bool)
				file, err := os.Open(path)
				if err != nil {
					errChan <- err
					return
				}
				defer file.Close()

				gz, err := gzip.NewReader(file)
				if err != nil {
					errChan <- err
					return
				}
				defer gz.Close()

				scanner := bufio.NewScanner(gz)
				for scanner.Scan() {
					line := strings.TrimSpace(scanner.Text())
					if len(line) >= 8 && len(line) <= 10 {
						localSet[line] = true
					}
				}
				if err := scanner.Err(); err != nil {
					errChan <- err
					return
				}

				// Lock and update shared map
				mu.Lock()
				for coupon := range localSet {
					couponCount[coupon]++
				}
				mu.Unlock()
			}(filePath)
		}

		wg.Wait()
		close(errChan)

		// Handle any errors
		for err := range errChan {
			loadErr = err
			return
		}

		// Final filter: only keep coupons in 2 or more files
		for coupon, count := range couponCount {
			if count >= 2 {
				cs.coupons[coupon] = count
			}
		}
	})

	return loadErr
}

// ValidateCoupon validates a coupon code based on its length, pattern, and file occurrence
func (cs *couponService) ValidateCoupon(code string) (bool, error) {
	// Check if coupon code length is valid
	if len(code) < 8 || len(code) > 10 {
		return false, nil
	}

	// Regular expression for valid promo codes (e.g., upper-case alphanumeric only)
	validCouponRegex := "^[A-Z0-9]+$"
	matched, err := regexp.MatchString(validCouponRegex, code)
	if err != nil {
		return false, err
	}
	if !matched {
		return false, nil // Invalid promo code if it doesn't match the valid pattern
	}

	// Load coupons from files if not already loaded
	if err := cs.loadCoupons(); err != nil {
		fmt.Printf("Error: %v\n", err)
		return false, err
	}

	// Check if the coupon is valid by its occurrence across files
	count := cs.coupons[code]
	return count >= 1, nil // Coupon is valid if it appears in at least two files
}

func NewCouponServiceFromJSON(jsonPath string) CouponService {
	file, err := os.Open(jsonPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to load %s: %v", jsonPath, err))
	}
	defer file.Close()

	coupons := make(map[string]int)
	if err := json.NewDecoder(file).Decode(&coupons); err != nil {
		panic(fmt.Sprintf("Failed to parse JSON: %v", err))
	}

	return &couponService{
		coupons: coupons,
	}
}
