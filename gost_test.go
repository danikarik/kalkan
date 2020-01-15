package kalkan_test

import (
	"testing"

	"github.com/danikarik/kalkan"
	"github.com/stretchr/testify/require"
)

func TestGOSTHasher(t *testing.T) {
	testCases := []struct {
		Name   string
		Input  string
		Output string
	}{
		{
			Name:   "TestSample1",
			Input:  "Ly210Ba4Ju",
			Output: "e0c9ce26ca2f0b39a3f6e3ce05b3220375ab36135cb481735ada1b1d07feea30",
		},
		{
			Name:   "TestSample2",
			Input:  "kAZJaFCiGy",
			Output: "3c37c64d0c5e438cde4aa949f07b618299c5e24460e1656a836ffee01e835083",
		},
		{
			Name:   "TestSample3",
			Input:  "hsCtRn1349",
			Output: "2fb528f4703c27c0883fe17aaadfa0240f56c65473b1b37b8572100ba4469c6b",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var (
				require = require.New(t)
				hasher  = kalkan.NewGOST()
			)

			out, err := hasher.WriteString(tc.Input)
			require.NoError(err)
			require.Equal(tc.Output, out)
		})
	}
}
