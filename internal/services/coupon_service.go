package services

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

type CouponService interface {
	ValidateCoupon(code string) (bool, error)
}

type couponService struct {
	coupons map[string]int // map[coupon]count
}

func NewCouponService(filePaths string) CouponService {
	return &couponService{
		coupons: make(map[string]int),
	}
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

	// Check if the coupon is valid by its occurrence across files
	count := cs.coupons[code]
	return count >= 2, nil // Coupon is valid if it appears in at least two files
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
