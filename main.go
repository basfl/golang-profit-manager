package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	// results := make(map[float64][]float64)
	for index, taxrate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fmt := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxrate*100))
		//cmd := cmdmanager.New()
		// cmd.WriteResults("something")
		//priceJob := prices.NewTaxIncludedPriceJob(cmd, taxrate)
		priceJob := prices.NewTaxIncludedPriceJob(fmt, taxrate)
		go priceJob.Process(doneChans[index], errorChans[index])
		//priceJob.LoadDate()

	}
	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("DONE!")
		}
	}
	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
	fmt.Println("****")

}
