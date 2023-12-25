package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fileManager FileManager) ReadLines() ([]string, error) {

	file, err := os.Open(fileManager.InputFilePath)
	if err != nil {

		return nil, errors.New("Could not open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading the file content failed ")
		fmt.Println(err)
		//	file.Close()
		return nil, errors.New("Reading the file content failed ")
	}
	//defer file.Close()
	return lines, nil

}

func (fileManger FileManager) WriteResults(data interface{}) error {
	file, err := os.Create(fileManger.OutputFilePath)
	defer file.Close()
	if err != nil {

		return errors.New("Could not open file")
	}
	//adding delay to simulate for channels
	time.Sleep(3 * time.Second)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("Encoding failed")
	}
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
