package kalkan_test

import (
	"testing"

	"github.com/danikarik/kalkan"
	"github.com/stretchr/testify/require"
)

func TestSHA1Hasher(t *testing.T) {
	testCases := []struct {
		Name   string
		Input  string
		Output string
	}{
		{
			Name:   "TestSample1",
			Input:  "Ly210Ba4Ju",
			Output: "0132a1b590991bd1d11d71326c5c2a6c91991534",
		},
		{
			Name:   "TestSample2",
			Input:  "kAZJaFCiGy",
			Output: "68446cb66a502c93c5cb3a3b1d2dfad06e684a4b",
		},
		{
			Name:   "TestSample3",
			Input:  "hsCtRn1349",
			Output: "b8e79e25572bc726706f9b6177e746c515466912",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var (
				require = require.New(t)
				hasher  = kalkan.NewSHA1()
			)

			out, err := hasher.WriteString(tc.Input)
			require.NoError(err)
			require.Equal(tc.Output, out)
		})
	}
}
