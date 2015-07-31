package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceTitleNotNull(t *testing.T) {

	res := &Resource{Title: "", Link: "www.google.com"}
	validationErros := make([]map[string][]string, 0)
	validationErros = res.Validate()
	actual := make(map[string][]string)
	actual = validationErros[0]
	expected := make(map[string][]string)
	expected["title"] = []string{"title is required"}
	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, expected["title"][0], actual["title"][0])
}
func TestResourceLinkNotNull(t *testing.T) {

	res := &Resource{Title: "Google", Link: ""}
	validationErros := make([]map[string][]string, 0)
	validationErros = res.Validate()
	actual := make(map[string][]string)
	actual = validationErros[0]
	expected := make(map[string][]string)
	expected["link"] = []string{"link is required"}
	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, expected["link"][0], actual["link"][0])
}
