/*
Copyright © 2022 NAME HERE bora@borakasmer.com

*/
package cmd

import (
	"fmt"
	"math/big"
	//"math/rand"
	"crypto/rand"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pwd-cli",
	Short: "You can generate a password with string and number length parameters and isCaption parameter.",
	Long: `
You can generate specific password with parameters like
"pwd-cli -s 7 -n 3 -c"
"pwd-cli"
"pwd-cli -n 10"
"pwd-cli -s 10"
*** Password Length = s(string) + n(number) total length ***
------------------
Example:
"-s 7" Length of string characters of the new password parameter.
"-n 3" => Length of number characters of the new password parameter.
"-c" => Has password capital letters parameter.
"pwd-cli" command => Default "pwd-cli --string 8 --iscapital 0"`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		isCapital, err := cmd.Flags().GetBool("capital")
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		str, err := cmd.Flags().GetInt("str")
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		generatePassword(PasswordParams{isCapital: isCapital, strLength: str, numLength: num})
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pwd-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().IntP("str", "s", 8, "Length of string")
	rootCmd.Flags().IntP("num", "n", 0, "Length of number")
	rootCmd.Flags().BoolP("capital", "c", false, "Is there any capital letter ?")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type PasswordParams struct {
	strLength int
	numLength int
	isCapital bool
}

/*
func generatePassword(params PasswordParams) {
	//fmt.Println(params)
	if params.strLength == 0 && params.numLength == 0 {
		params.strLength = 8
	}
	if params.numLength == 0 {
		fmt.Println(getRandomString(params.strLength, params.isCapital))
	} else if params.numLength > 0 {
		password := ""
		numberList := make(map[int]string)
		passwordLength := params.numLength + params.strLength
		for i := 0; i < params.numLength; i++ {
			for {
				index := getRandomNumber(passwordLength)
				if _, ok := numberList[index]; ok {

				} else {
					numberList[index] = strconv.Itoa(getRandomNumber(10))
					break
				}
			}
		}

		for i2 := 0; i2 < passwordLength; i2++ {
			if _, ok := numberList[i2]; ok {
				password += numberList[i2]
			} else {
				password += getRandomString(1, params.isCapital)
			}
		}
		fmt.Println(password)
	}
}
*/

func generatePassword(params PasswordParams) {
	if params.strLength == 0 && params.numLength == 0 {
		params.strLength = 8
	}
	if params.numLength == 0 {
		fmt.Println(getRandomString(params.strLength, params.isCapital))
	} else if params.numLength > 0 {
		password := ""
		numberList := make(map[int]string)
		passwordLength := params.numLength + params.strLength
		for i := 0; i < params.numLength; i++ {
			for {
				index := getRandomNumber(int64(passwordLength))
				if _, ok := numberList[int(index)]; ok {

				} else {
					numberList[int(index)] = strconv.Itoa(int(getRandomNumber(int64(10))))
					break
				}
			}
		}

		for i2 := 0; i2 < passwordLength; i2++ {
			if _, ok := numberList[i2]; ok {
				password += numberList[i2]
			} else {
				password += getRandomString(1, params.isCapital)
			}
		}
		fmt.Println(password)
	}
}

/*func getRandomNumber(length int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(length)
}*/

func getRandomNumber(length int64) int64 {
	result, _ := rand.Int(rand.Reader, big.NewInt(length))
	return result.Int64()
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
var letterRunes2 = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

/*
func getRandomString(length int, isCapital bool) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		if isCapital {
			b[i] = letterRunes2[rand.Intn(len(letterRunes2))]
		} else {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
	}
	return string(b)
}
*/

func getRandomString(length int, isCapital bool) string {
	b := make([]rune, length)
	for i := range b {
		if isCapital {
			result, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes2))))
			b[i] = letterRunes2[result.Int64()]
		} else {
			result, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
			b[i] = letterRunes[result.Int64()]
		}
	}
	return string(b)
}
