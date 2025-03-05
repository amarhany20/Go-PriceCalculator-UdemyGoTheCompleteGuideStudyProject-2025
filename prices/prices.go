package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64           `json:"tax_rate"`
	InputPrices       []float64         `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`

	IOManager filemanager.FileManager `json:"-"`
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:         fm,
		TaxRate:           taxRate,
		InputPrices:       []float64{},
		TaxIncludedPrices: make(map[string]string),
	}
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	prices, err := conversion.StringToFloat64(lines)
	if err != nil {
		fmt.Println("Error converting strings to float64:", err)
		return
	}

	job.InputPrices = prices

}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	prices := job.InputPrices
	result := make(map[string]string)
	for _, price := range prices {
		taxIncludedPrice := price * (1 + job.TaxRate)

		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result
	err := job.IOManager.WriteJSON(job)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}
}
