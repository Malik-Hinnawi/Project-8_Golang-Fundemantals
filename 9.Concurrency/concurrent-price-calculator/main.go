package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	donesChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		donesChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(donesChans[i], errorChans[i])

	}

	for i := range taxRates {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-donesChans[i]:
			fmt.Println("Done!")
		}
	}

}
