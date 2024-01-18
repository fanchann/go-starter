package helpers

import (
	"errors"
	"log"
	"strconv"

	"github.com/manifoldco/promptui"
)

func PromptNamePackage() string {
	validate := func(input string) error {
		if len(input) <= 3 {
			return errors.New("name package must be longer than 3 characters")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Name Package",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Failed to prompt for package name: %v", err)
		return ""
	}

	return result
}

func PromptArrayString(label string, item []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: item,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Error : %v", err)
	}

	return result
}

func PromptString(label string) string {
	context := func(input string) error {
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: context,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Error : %v", err)

	}

	return result
}

func PromptInteger(label string) int {
	num := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("invalid")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: num,
	}

	result, err := prompt.Run()

	if err != nil {
		log.Fatalf("Error : %v", err)

	}

	port, _ := strconv.Atoi(result)
	return port
}
