package cli

import "testing"

func TestCompareArgsLengthIsValid(t *testing.T) {
	test := func(t *testing.T) {
		validInputs := [][]string{
			{"5aca0753-9b50-4de6-bcad-6118f6adc1bb", "18b771aa-7ef7-4d76-84ce-b32d7dde7b0f"},
		}
		invalidInputs := [][]string{
			{},
			{"e9b93e68-375b-4f32-be5c-d46bcdae3eb5"},
			{
				"8954a500-2477-45f4-8d8f-73b616bd6953",
				"af70e31e-d05a-49c0-9aa4-86a61f106c1a",
				"d1a116e0-ea75-41ab-b71c-f5ae59263fed",
			},
		}

		for _, input := range validInputs {
			err := validateArgs(input)
			if err != nil {
				t.Errorf("validateArgs failed with error: %s. args: %+v", err, input)
			}
			return
		}

		for _, input := range invalidInputs {
			err := validateArgs(input)
			if err == nil {
				t.Errorf("validateArgs passed when it should failed: %s. args: %+v", err, input)
			}
			return
		}
	}

	t.Run("provided cost sizes return expected and valid results", test)
}
