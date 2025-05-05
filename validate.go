package cep

import "regexp"

func Validate(cep string) (ok bool) {
	cepRegex := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	ok = cepRegex.MatchString(cep)
	return
}
