package toc_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/toc"
	"gopkg.in/yaml.v3"
)

func TestIntegration(t *testing.T) {
	t.Parallel()

	testsdata, err := os.ReadFile("testdata/tests.yaml")
	require.NoError(t, err)

	var tests []struct {
		Desc    string `yaml:"desc"`
		Give    string `yaml:"give"`
		Want    string `yaml:"want"`
		Title   string `yaml:"title"`
		ListID  string `yaml:"listID"`
		TitleID string `yaml:"titleID"`

		MinDepth int  `yaml:"minDepth"`
		MaxDepth int  `yaml:"maxDepth"`
		Compact  bool `yaml:"compact"`
	}
	require.NoError(t, yaml.Unmarshal(testsdata, &tests))

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Desc, func(t *testing.T) {
			t.Parallel()

			md := goldmark.New(
				goldmark.WithExtensions(&toc.Extender{
					Title:    tt.Title,
					MinDepth: tt.MinDepth,
					MaxDepth: tt.MaxDepth,
					Compact:  tt.Compact,
					ListID:   tt.ListID,
					TitleID:  tt.TitleID,
				}),
				goldmark.WithParserOptions(parser.WithAutoHeadingID()),
			)

			var buf bytes.Buffer
			require.NoError(t, md.Convert([]byte(tt.Give), &buf))
			require.Equal(t, tt.Want, buf.String())
		})
	}
}
