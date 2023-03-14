package tools

import (
	"strings"
)

func ContainsString(str string, arr []string) bool {
	for _, s := range arr {
		if strings.Contains(str, s) {
			return true
		}
	}

	return false
}
