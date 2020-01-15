package kalkan_test

import (
	"testing"

	"github.com/danikarik/kalkan"
	"github.com/stretchr/testify/require"
)

func TestSHA256Hasher(t *testing.T) {
	testCases := []struct {
		Name   string
		Input  string
		Output string
	}{
		{
			Name:   "TestSample1",
			Input:  "Ly210Ba4Ju",
			Output: "ccd546a4e3d159d99e0d043d0d0f74430403500626376b00b8d86d3e7718c9e6",
		},
		{
			Name:   "TestSample2",
			Input:  "kAZJaFCiGy",
			Output: "7d2bcd1592301e5336ce79cd1c46bdd817cb7a3805f4e9a7b7658e0cd8e4ab70",
		},
		{
			Name:   "TestSample3",
			Input:  "hsCtRn1349",
			Output: "8ad8d8c7fed2c49dc32295d6abce941e3245448e33a605d2d31a59b398411e00",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var (
				require = require.New(t)
				hasher  = kalkan.NewSHA256()
			)

			out, err := hasher.WriteString(tc.Input)
			require.NoError(err)
			require.Equal(tc.Output, out)
		})
	}
}
