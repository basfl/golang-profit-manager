package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadDate() error {

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println("filemanager error")
		fmt.Println(err)
		return err
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("conversion failed")
		fmt.Println(err)
		return err
	}
	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errChan chan error) {
	err := job.LoadDate()
	if err != nil {
		errChan <- err
		return
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {

		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)

	}

	//fmt.Println(result)
	job.TaxIncludedPrices = result
	// err := job.IOManager.WriteResults(job)
	// if err != nil {
	// 	fmt.Println("Failed to convert/write into the json file")
	// 	fmt.Println(err)
	// 	return
	// }
	job.IOManager.WriteResults(job)
	doneChan <- true

}
func NewTaxIncludedPriceJob(io iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		IOManager:   io,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}

}
