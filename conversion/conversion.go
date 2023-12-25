package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(input []string) ([]float64, error) {

	//prices := make([]float64, len(input))
	var prices []float64
	for _, line := range input {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, errors.New("faile to convert string to float!!")
		}
		//prices[index] = floatPrice
		prices = append(prices, floatPrice)
	}
	return prices, nil

}
