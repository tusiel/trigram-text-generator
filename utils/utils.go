package utils

import "regexp"

func RemoveUnwantedCharacters(word string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9!.?,:;\\s]+")
	if err != nil {
		return "", err
	}

	return reg.ReplaceAllString(word, ""), nil
}

func IsSentenceEnd(s string) (bool, error) {
	match, err := regexp.MatchString("([.!?]$)", s)
	if err != nil {
		return false, err
	}

	return match, err
}
