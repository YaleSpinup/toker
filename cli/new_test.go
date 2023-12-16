package cli

import (
	"reflect"
	"testing"
)

func TestGenerateNewTokenCountIsCorrect(t *testing.T) {
	test := func(t *testing.T) {
		var actuals []int

		validInputs := []int{3, 5, 10}
		invalidInputs := []int{0, -1}
		expectedValidValues := []int{3, 5, 10}

		for _, input := range validInputs {
			value, err := generate(input)
			if err != nil {
				t.Errorf("error generating tokens: %s with input: %d", err, input)
			}

			actuals = append(actuals, len(value))
		}

		if actuals == nil {
			actuals = []int{}
		}

		if !reflect.DeepEqual(actuals, expectedValidValues) {
			t.Errorf("actual values: %+v do not match expected valid values: %+v", actuals, expectedValidValues)
			return
		}

		for _, input := range invalidInputs {
			value, err := generate(input)
			if err == nil {
				t.Errorf("input was expected to throw an error and it didn't. input: %d, actual: %d", input, value)
				return
			}
		}
	}

	t.Run("list of expected input counts matches generated counts", test)
}
