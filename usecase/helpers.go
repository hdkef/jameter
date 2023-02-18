package usecase

import "regexp"

func OnlyWordsOrDigit(name string) (string, bool) {
	return "only words or digit allowed", !regexp.MustCompile(`[^A-Za-z0-9]+`).MatchString(name)
}
