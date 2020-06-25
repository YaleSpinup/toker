/*
Copyright © 2020 E Camden Fisher <fish@fishnix.net>

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
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

// hashCmd represents the hash command
var hashCmd = &cobra.Command{
	Use:   "hash password",
	Short: "Hash a password",
	Long:  `Encrypts the password using bcrypt`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cost < 4 || cost > 31 {
			return fmt.Errorf("invalid cost %d, min: 4, max 31, default: %d", cost, bcrypt.DefaultCost)
		}

		if len(args) != 1 {
			return fmt.Errorf("expected one argument, got %d", len(args))
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		e, err := bcrypt.GenerateFromPassword([]byte(args[0]), cost)
		if err != nil {
			return err
		}

		fmt.Println(string(e))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(hashCmd)
	hashCmd.PersistentFlags().IntVarP(&cost, "cost", "o", bcrypt.DefaultCost, "the cost for the generated hash (min 4, max 31)")
}
