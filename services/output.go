package services

import (
	"fmt"
	"os"
)

var OutputType []string = []string{"CommandLine"}

func IsValidOutput(outputName string) bool {
	return contains(OutputType, outputName)
}

type OutputService interface {
	Write(value string)
}

type CommandLineOutputService struct {
}

func (e *CommandLineOutputService) Write(value string) {
	fmt.Println(value)
}

func GetOutput(value string) OutputService {
	if !IsValidOutput(value) {
		fmt.Printf("Output %s not recognized. Choose one of: %v\n", value, OutputType)
		os.Exit(1)
	}

	if value == "CommandLine" {
		return &CommandLineOutputService{}
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
