package utils

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/manifoldco/promptui"
)

// PromptText run prompt for given label and validate if required is true
func PromptText(label string, required bool) (string, error) {
	textPrompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			if required {
				return validation.Validate(s, validation.Required)
			}
			return nil
		},
	}
	return textPrompt.Run()
}

// PromptTextMasked run prompt with masked text for given label and validate if required
func PromptTextMasked(label string, required bool) (string, error) {
	textPrompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			if required {
				return validation.Validate(s, validation.Required)
			}
			return nil
		},
		Mask: '*',
	}
	return textPrompt.Run()
}
