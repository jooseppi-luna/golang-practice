package convert_string_to_float

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStringToFloat(t *testing.T) {
	tests := []struct{
		name string
		numberString string
		numberNumber float64
	}{
		{
			name: "base case",
			numberString: "12.3",
			numberNumber: 12.3,
		},
		{
			name: "zero no decimal",
			numberString: "0",
			numberNumber: 0.0,
		},
		{
			name: "less than one",
			numberString: "0.352",
			numberNumber: 0.352,
		},
		{
			name: "long number",
			numberString: "126476845677688384582.3447967727739576",
			numberNumber: 126476845677688384582.3447967727739576,
		},
		{
			name: "negative number",
			numberString: "-12.3",
			numberNumber: -12.3,
		},
		{
			name: "no decimal",
			numberString: "356",
			numberNumber: 356.0,
		},
		{
			name: "zeroes at end",
			numberString: "356000",
			numberNumber: 356000.0,
		},
		{
			name: "zeroes in middle",
			numberString: "3560001",
			numberNumber: 3560001.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numberNumber := StringToFloat(tt.numberString)
			assert.Equal(t, tt.numberNumber, numberNumber)
		})
	}
}

