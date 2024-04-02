/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"math/rand"

	"github.com/spf13/cobra"
)

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "generate a password",
	Long: `generate some passwords with options for length, digits and special characters. For example:
	
	passwordGen password -l 20 -d=false -s=false  // => CaroJKiiLJAwiBytqwIc
	
	this will generate a password of length 20 without digits and without special characters. 
	
	Default settings are: length 18, digits true, special characters true.

	passwordGen password // => r[+)<=[U1tbY37a
	`,

	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(passwordCmd)

	passwordCmd.Flags().IntP("length", "l", 12, "Length of the password")
	passwordCmd.Flags().BoolP("digits", "d", true, "Include digits in the password")
	passwordCmd.Flags().BoolP("special-chars", "s", true, "Include special characters in the password")
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	hasDigits, _ := cmd.Flags().GetBool("digits")
	hasSpecialChars, _ := cmd.Flags().GetBool("special-chars")

	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if hasDigits {
		charSet += "0123456789"
	}
	if hasSpecialChars {
		charSet += "!@#$%^&*()_+[]{}|;:,.<>?-="
	}

	password := make([]byte, length)

	for i := range password {
		password[i] = charSet[rand.Intn(len(charSet))]
	}

	cmd.Println(string(password))
}
