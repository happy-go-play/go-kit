package larkx

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildMarkdownCardV2(t *testing.T) {
	cardJSON, err := buildMarkdownCardV2("**hello**", "title", "purple")
	require.NoError(t, err)

	var card map[string]any
	require.NoError(t, json.Unmarshal([]byte(cardJSON), &card))
	require.Equal(t, "2.0", card["schema"])

	header, ok := card["header"].(map[string]any)
	require.True(t, ok)
	title, ok := header["title"].(map[string]any)
	require.True(t, ok)
	require.Equal(t, "plain_text", title["tag"])
	require.Equal(t, "title", title["content"])
	require.Equal(t, "purple", header["template"])

	body, ok := card["body"].(map[string]any)
	require.True(t, ok)
	elements, ok := body["elements"].([]any)
	require.True(t, ok)
	require.Len(t, elements, 1)

	element, ok := elements[0].(map[string]any)
	require.True(t, ok)
	require.Equal(t, "markdown", element["tag"])
	require.Equal(t, "**hello**", element["content"])
}

func TestBuildMarkdownCardV2WithoutTitle(t *testing.T) {
	cardJSON, err := buildMarkdownCardV2("hello", "", "")
	require.NoError(t, err)

	var card map[string]any
	require.NoError(t, json.Unmarshal([]byte(cardJSON), &card))
	require.Equal(t, "2.0", card["schema"])
	require.NotContains(t, card, "header")
}
