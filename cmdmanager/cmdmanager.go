package cmdmanager

import (
	"fmt"
)

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	var inputs []string
	for {
		var input string
		fmt.Println("prices:")
		fmt.Scan(&input)
		if input == "-1" {
			break
		}
		inputs = append(inputs, input)
	}
	return inputs, nil
}

func (cmd CMDManager) WriteResults(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
