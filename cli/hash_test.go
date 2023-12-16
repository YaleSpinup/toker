package cli

import (
	"testing"
)

func TestHashCostIsValidSize(t *testing.T) {
	test := func(t *testing.T) {
		validInputs := [][]interface{}{
			{8, []string{"0ebc3ea8-b731-4b78-a0bb-45cb11aa3336"}},
			{22, []string{"6d6fc85d-8a3c-4af3-b515-57e6931bf5eb"}},
		}
		invalidInputs := [][]interface{}{
			{3, []string{"a623da14-e1df-409f-887b-488026b7b921"}},
			{32, []string{"7d723ae1-d96f-4e32-a60c-9911a52bc857"}},
			{-1, []string{"ca4bd336-7d27-42e6-9ec7-59fdaec480de"}},
		}

		for _, input := range validInputs {
			err := validateCostAndArgs(input[0].(int), input[1].([]string))
			if err != nil {
				t.Errorf(
					"validateCostAndArgs failed with error: %s. cost: %d, and args: %+v",
					err,
					input[0].(int),
					input[1].([]string),
				)
				return
			}
		}

		for _, input := range invalidInputs {
			err := validateCostAndArgs(input[0].(int), input[1].([]string))
			if err == nil {
				t.Errorf(
					"validateCostAndArgs passed when it should failed. cost: %d, and args: %+v",
					input[0].(int),
					input[1].([]string),
				)
				return
			}
		}
	}

	t.Run("count sizes provide expected results", test)
}

func TestHashArgsLengthIsValid(t *testing.T) {
	test := func(t *testing.T) {
		validInputs := [][]interface{}{
			{30, []string{"5aca0753-9b50-4de6-bcad-6118f6adc1bb"}},
		}
		invalidInputs := [][]interface{}{
			{10, []string{}},
			{25, []string{"e9b93e68-375b-4f32-be5c-d46bcdae3eb5", "18b771aa-7ef7-4d76-84ce-b32d7dde7b0f"}},
		}

		for _, input := range validInputs {
			err := validateCostAndArgs(input[0].(int), input[1].([]string))
			if err != nil {
				t.Errorf(
					"validateCostAndArgs failed with error: %s. cost: %d, and args: %+v",
					err,
					input[0].(int),
					input[1].([]string),
				)
			}
			return
		}

		for _, input := range invalidInputs {
			err := validateCostAndArgs(input[0].(int), input[1].([]string))
			if err == nil {
				t.Errorf(
					"validateCostAndArgs passed when it should failed: %s, cost: %d, and args: %+v",
					err,
					input[0].(int),
					input[1].([]string),
				)
			}
			return
		}
	}

	t.Run("args values lengths provide expected results", test)
}
