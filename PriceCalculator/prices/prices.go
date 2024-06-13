package prices

import (
	"fmt"
	"price-calculator/conversion"
	"price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json: "-"`
	TaxRate           float64             `json: "tax_rate"`
	InputPrices       []float64           `json: "input_prices"`
	TaxIncludedPrices map[string]string   `json: "tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	// convert values string to float64
	prices, err := conversion.StringsToFloats(lines)

	// for lineIndex, lineValue := range lines {
	// 	floatPrice, err := strconv.ParseFloat(lineValue, 64)

	if err != nil {
		// fmt.Println("converting price to float failed")
		fmt.Println(err)
		// file.Close()
		return
	}

	// prices[lineIndex] = floatPrice

	job.InputPrices = prices
	// file.Close()
}

func (job TaxIncludedPriceJob) Process() {
	job.LoadData()

	// result := make(map[string]float64)
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	// fmt.Println(result)

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)
}

func NewTaxtIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
