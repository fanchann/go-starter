package helpers

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/manifoldco/promptui"
)

var (
	genFailed = "generate failed"
)

func PromptNamePackage() string {
	validate := func(input string) error {
		matched, err := regexp.MatchString("^github\\.com\\/[a-zA-Z]+\\/[a-zA-Z]+", input)
		if err != nil {
			return errors.New("[*] Package name must start with github.com/{username}/{package}")
		}

		if !matched {
			return errors.New("[*] Package name must start with github.com/{username}/{package}")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Name Package",
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf(genFailed, err)
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
		panic(genFailed)
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
		panic(genFailed)
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
		panic(genFailed)
	}

	port, _ := strconv.Atoi(result)
	return port
}
