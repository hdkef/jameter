package usecase

import "regexp"

type Validator struct{}

func (v *Validator) OnlyWordsOrDigit(data string) (string, bool) {
	return "Only alphabet or digit allowed", !regexp.MustCompile(`[^A-Za-z0-9]+`).MatchString(data)
}

func (v *Validator) OnlyWords(data string) (string, bool) {
	return "Only alphabet allowed", regexp.MustCompile(`^[a-zA-Z]*$`).MatchString(data)
}
