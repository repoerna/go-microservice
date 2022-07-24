package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductSKUValidationSuccess(t *testing.T) {
	p := &Product{
		Name:  "tets",
		Price: 1.00,
		SKU:   "qwe-qwer-qwer",
	}

	err := p.Validate()

	assert.Nil(t, err)
}

func TestProductSKUValidationError(t *testing.T) {
	p := &Product{
		Name:  "tets",
		Price: 1.00,
		SKU:   "123-qwer-qwer",
	}

	err := p.Validate()

	assert.Error(t, err)
}

func TestProductMissingNameError(t *testing.T) {
	p := &Product{
		Price: 1.00,
		SKU:   "qwer-qwer-qwer",
	}

	err := p.Validate()

	assert.Len(t, err, 1)

	assert.Error(t, err)
}
