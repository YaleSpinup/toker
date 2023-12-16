/*
Copyright © 2020 Yale University

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cli

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generate tokens",
	Long:  `Generate a list of tokens and optionally their encrypted values`,
	PreRunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		tokens, err := generate(count)
		if err != nil {
			return err
		}

		for _, t := range tokens {
			s := t.String()
			if enc {
				e, err := bcrypt.GenerateFromPassword([]byte(s), cost)
				if err != nil {
					return err
				}

				s = fmt.Sprintf("%s %s", s, e)
			}

			fmt.Println(s)
		}

		return nil
	},
}

// validateCost validates that the provided cost value is within the given size range
func validateCost(cost int) error {
	if cost < 4 || cost > 31 {
		return fmt.Errorf("invalid cost %d, min: 4, max 31, default: %d", cost, bcrypt.DefaultCost)
	}

	return nil
}

// generate generates a number of UUIDv4 values equal to count
func generate(count int) ([]uuid.UUID, error) {
	var tokens []uuid.UUID

	if count < 1 {
		return nil, fmt.Errorf("invalid count. count must be greater than 1")
	}

	for i := 0; i < count; i++ {
		t, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, t)
	}

	if tokens == nil {
		tokens = []uuid.UUID{}
	}

	return tokens, nil
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.PersistentFlags().IntVarP(&cost, "cost", "o", bcrypt.DefaultCost, "the cost for the generated hash (min 4, max 31)")
	newCmd.PersistentFlags().IntVarP(&count, "count", "c", 5, "the number of UUIDs to generate")
	newCmd.PersistentFlags().BoolVarP(&enc, "encrypt", "e", false, "also show encrypted values")
}
