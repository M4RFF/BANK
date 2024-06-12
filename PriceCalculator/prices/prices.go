package prices

import (
	"bufio"
	"fmt"
	"os"
	"price-calculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt") // allows us to open this file

	if err != nil {
		fmt.Println("could not open a file")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file) // allows us to go line by line in this "file"

	// we move farword one line of the line where we were
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("reading the file content failed")
		fmt.Println(err)
		file.Close()
		return
	}

	// convert values string to float64
	prices, err := conversion.StringsToFloats(lines)

	// for lineIndex, lineValue := range lines {
	// 	floatPrice, err := strconv.ParseFloat(lineValue, 64)

	if err != nil {
		// fmt.Println("converting price to float failed")
		fmt.Println(err)
		file.Close()
		return
	}

	// prices[lineIndex] = floatPrice

	job.InputPrices = prices
	file.Close()
}

func (job TaxIncludedPriceJob) Process() {
	job.LoadData()

	// result := make(map[string]float64)
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)

}

func NewTaxtIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
