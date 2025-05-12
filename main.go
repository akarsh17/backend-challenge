package main

import (
	"backend-challenge/api/routes"
	"backend-challenge/internal/services"
	"backend-challenge/pkg/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Optional: support preprocessing with a flag

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting the current working directory:", err)
	}
	coupon1 := filepath.Join(dir, "data/couponbase1.gz")
	coupon2 := filepath.Join(dir, "data/couponbase2.gz")
	coupon3 := filepath.Join(dir, "data/couponbase3.gz")
	jsonFile := filepath.Join(dir, "data/valid_coupons.json")
	err = utils.PreprocessCoupons(
		[]string{coupon1, coupon2, coupon3},
		jsonFile,
	)
	if err != nil {
		log.Fatalf("❌ Preprocessing failed: %v", err)
		return
	}

	// API mode
	services.NewCouponServiceFromJSON("valid_coupons.json")
	r := routes.SetupRouter()
	log.Println("Starting server on :8080")
	r.Run(":8080")
}
