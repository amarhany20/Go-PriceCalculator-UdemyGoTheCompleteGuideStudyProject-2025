package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("results.txt", fmt.Sprintf("results_%.2f.txt", taxRate*100))
		job := prices.NewTaxIncludedPriceJob(fm, taxRate)
		job.Process()
		fmt.Printf("Processed tax rate: %.2f%%, output file: results_%.2f.txt\n", taxRate*100, taxRate*100)
	}
}
