package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_It_Gets_An_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func Test_If_It_Gets_An_Error_If_Price_Is_Invalid(t *testing.T) {
	order := Order{ID: "1"}
	assert.Error(t, order.Validate(), "invalid price")
}

func Test_If_It_Gets_An_Error_If_Tax_Is_Invalid(t *testing.T) {
	order := Order{ID: "1", Price: 1.0}
	assert.Error(t, order.Validate(), "invalid tax")
}

func Test_CreateNewOrder_OK(t *testing.T) {
	order := &Order{
		ID:         "1",
		Price:      1.0,
		Tax:        0.5,
		FinalPrice: 0,
	}

	want, err := NewOrder("1", 1.0, 0.5)
	assert.Nil(t, err)
	assert.Equal(t, want, order)
}

func Test_CreateNewOrder_ValidateError(t *testing.T) {
	_, err := NewOrder("1", -1, 0.5)
	assert.NotNil(t, err)
}

func Test_CalculateFinalPrice_OK(t *testing.T) {
	order := &Order{
		ID:    "1",
		Price: 1.0,
		Tax:   0.5,
	}

	assert.Equal(t, order.CalculateFinalPrice(), nil)
	assert.Equal(t, order.FinalPrice, 1.5)
}

func Test_CalculateFinalPrice_ValidateErro(t *testing.T) {
	order := &Order{
		ID:  "1",
		Tax: 0.5,
	}

	assert.Equal(t, order.CalculateFinalPrice().Error(), "invalid price")
}
