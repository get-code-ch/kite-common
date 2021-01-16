package kite_common

import (
	"fmt"
	"regexp"
)

func (a Address) String() string {
	return fmt.Sprintf("%s.%s.%s.%s.%s", a.Domain, a.Type, a.Host, a.Address, a.Id)
}

func (a *Address) StringToAddress(str string) {
	splitRe := regexp.MustCompile(`\.`)
	for idx, value := range splitRe.Split(str, -1) {
		if idx == 0 {
			a.Domain = value
		}
		if idx == 1 {
			a.Type = HostType(value)
			if err := a.Type.IsValid(); err != nil {
				a.Type = H_ANY
			}
		}
		if idx == 2 {
			a.Host = value
		}
		if idx == 3 {
			a.Address = value
		}
		if idx == 4 {
			a.Id = value
		}
	}
	if a.Domain == "" {
		a.Domain = "*"
	}
	if a.Type == "" {
		a.Type = H_ANY
	}
	if a.Host == "" {
		a.Host = "*"
	}
	if a.Address == "" {
		a.Address = "*"
	}
	if a.Id == "" {
		a.Id = "*"
	}
}
