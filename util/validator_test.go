package validator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateProduceCode(t *testing.T) {
	cases := []struct {
		name        string
		produceCode string

		expectedResp error
	}{

		{
			name:        "valid",
			produceCode: "SDFS-SDFS-ABCS-ABCD",
		},
		{
			name:         "invalid non alphanumeric",
			produceCode:  "SDFS-SDF*-ABCS-ABCD",
			expectedResp: errors.New("invalid produce code"),
		},
		{
			name:         "invalid non alphanumeric",
			produceCode:  "SDFS-SDF*-ABCS-ABCD",
			expectedResp: errors.New("invalid produce code"),
		},
		{
			name:         "invalid without dashes",
			produceCode:  "SDFS1SDF1eABCSxABCD",
			expectedResp: errors.New("invalid produce code"),
		},
	}

	for _, c := range cases {
		actualErr := ValidateProduceCode(c.produceCode)
		assert.Equal(t, c.expectedResp, actualErr, c.name)
	}
}

func TestValidatePrice(t *testing.T) {
	cases := []struct {
		name  string
		price float64

		expectedResp error
	}{

		{
			name:  "valid",
			price: float64(12.12),
		},
		{
			name:         "invalid",
			price:        float64(12.1212),
			expectedResp: errors.New("invalid price"),
		},
	}

	for _, c := range cases {
		actualErr := ValidatePrice(c.price)
		assert.Equal(t, c.expectedResp, actualErr, c.name)
	}
}

func TestValidateName(t *testing.T) {
	cases := []struct {
		name        string
		productName string

		expectedResp error
	}{

		{
			name:        "valid",
			productName: "lettuce",
		},
		{
			name:        "spaces",
			productName: "butter lettuce",
		},
		{
			name:        "double space",
			productName: "pricey butter lettuce",
		},
		{
			name:         "nonalphanumeric",
			productName:  "sdlkfj@sdf",
			expectedResp: errors.New("name must be alphanumeric"),
		},
	}

	for _, c := range cases {
		actualErr := ValidateName(c.productName)
		assert.Equal(t, c.expectedResp, actualErr, c.name)
	}
}
