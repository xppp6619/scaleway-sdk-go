package scw

import (
	"encoding/json"
	"testing"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/testhelpers"
)

func TestParseZone(t *testing.T) {
	tests := []struct {
		input    string
		err      error
		expected Zone
	}{
		{
			input:    "fr-par-1",
			expected: ZoneFrPar1,
		},
		{
			input:    "par1",
			expected: ZoneFrPar1,
		},
		{
			input:    "ams1",
			expected: ZoneNlAms1,
		},
		{
			input:    "xx-xxx-1",
			expected: "xx-xxx-1",
		},
		{
			input:    "fr-par",
			expected: "",
			err:      errors.New("bad zone format, available zones are: fr-par-1, fr-par-2, nl-ams-1"),
		},
		{
			input:    "fr-par-n",
			expected: "",
			err:      errors.New("bad zone format, available zones are: fr-par-1, fr-par-2, nl-ams-1"),
		},
		{
			input:    "fr-par-0",
			expected: "",
			err:      errors.New("bad zone format, available zones are: fr-par-1, fr-par-2, nl-ams-1"),
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			z, err := ParseZone(test.input)
			testhelpers.Equals(t, test.err, err)
			testhelpers.Equals(t, test.expected, z)
		})
	}
}

func TestZoneJSONUnmarshall(t *testing.T) {
	t.Run("test with zone", func(t *testing.T) {
		input := `{"Test": "par1"}`
		value := struct{ Test Zone }{}

		err := json.Unmarshal([]byte(input), &value)
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, ZoneFrPar1, value.Test)
	})

	t.Run("test with region", func(t *testing.T) {
		input := `{"Test": "par1"}`
		value := struct{ Test Region }{}

		err := json.Unmarshal([]byte(input), &value)
		testhelpers.AssertNoError(t, err)

		testhelpers.Equals(t, RegionFrPar, value.Test)
	})
}

func TestParseRegion(t *testing.T) {
	tests := []struct {
		input    string
		err      error
		expected Region
	}{
		{
			input:    "fr-par",
			expected: RegionFrPar,
		},
		{
			input:    "par1",
			expected: RegionFrPar,
		},
		{
			input:    "ams1",
			expected: RegionNlAms,
		},
		{
			input:    "xx-xxx",
			expected: "xx-xxx",
		},
		{
			input:    "fr-par-1",
			expected: "",
			err:      errors.New("bad region format, available regions are: fr-par, nl-ams"),
		},
		{
			input:    "fr-pa1",
			expected: "",
			err:      errors.New("bad region format, available regions are: fr-par, nl-ams"),
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			r, err := ParseRegion(test.input)
			testhelpers.Equals(t, test.err, err)
			testhelpers.Equals(t, test.expected, r)
		})
	}
}
