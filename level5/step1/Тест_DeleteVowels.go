package main

import (
	"strings"
	"testing"
)

func TestDeleteVowels(t *testing.T) {
	var flag bool

	for _, letter := range DeleteVowels("auioe") {
		if strings.ContainsRune("auioe", letter) {
			flag = true
			break
		}
	}

	if flag {
		t.Errorf("Test sum failed")
	}
}
