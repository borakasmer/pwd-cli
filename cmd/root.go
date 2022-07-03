/*
Copyright © 2022 NAME HERE bora@borakasmer.com

*/
package cmd

import (
	"fmt"
	"math/big"
	"unicode"

	//"math/rand"
	"crypto/rand"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pwd-cli",
	Short: "You can generate a password with string and number length parameters, isCaption parameter and finally has symbol parameter.",
	Long: `
You can generate specific password with parameters like
"pwd-cli -s 7 -n 3 -c"
"pwd-cli -s9 -n2 -x"
"pwd-cli"
"pwd-cli -n 10 -x"
"pwd-cli -s 10"
*** Password Length = s(string) + n(number) total length ***
------------------
Example:
"-s 7" Length of string characters of the new password parameter.
"-n 3" => Length of number characters of the new password parameter.
"-c" => Has password capital letters parameter.
"-x" => Has password symbol letters parameter.
"pwd-cli" command => Default "pwd-cli --string 8 --number 0 --capital 0 --symbol 0"`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		isCapital, err := cmd.Flags().GetBool("capital")
		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
		isSymbol, err := cmd.Flags().GetBool("symbol")
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
		generatePassword(PasswordParams{isCapital: isCapital, strLength: str, numLength: num, isSymbol: isSymbol})
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
	//rootCmd.Flags().IntP("str", "s", 8, "Length of string") //03.07.2022 (If only set "-n6", password has only 6 digit numbers)
	rootCmd.Flags().IntP("str", "s", 0, "Length of string")
	rootCmd.Flags().IntP("num", "n", 0, "Length of number")
	rootCmd.Flags().BoolP("capital", "c", false, "Is there any capital letter ?")
	rootCmd.Flags().BoolP("symbol", "x", false, "Is there any symbol ?")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type PasswordParams struct {
	strLength int
	numLength int
	isCapital bool
	isSymbol  bool
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
	symbolList := make(map[int]string)
	password := ""
	//If not number and string length not set.
	if params.strLength == 0 && params.numLength == 0 {
		params.strLength = 8
	}
	passwordLength := params.numLength + params.strLength

	//Set Symbol Characters of the Password if IsSymbol==true
	if params.isSymbol {
		passwordLength += params.strLength/3 + 1
		for i := 0; i < params.strLength/3+1; i++ {
			for {
				index := getRandomNumber(int64(passwordLength))
				//If you not get same RandomIndex for symbol character you will set RandomNumber for RandomIndex
				if _, ok := symbolList[int(index)]; ok {

				} else {
					symbolList[int(index)] = getRandomSymbol()
					break
				}
			}
		}
	}
	//Set Symbols Index Finish----------

	if params.numLength == 0 {
		//Set String Characters of the Password
		if params.isSymbol {
			for {
				for i2 := 0; i2 < passwordLength; i2++ {
					if _, ok := symbolList[i2]; ok {
						password += symbolList[i2]
					} else {
						password += getRandomString(1, params.isCapital)
					}
				}
				//If isCapital flag false or there is at least one capital letter in password, we will break the loop
				//If not set String Characters again until has a capital letter!
				if !params.isCapital || checkHasCapital(password) {
					break
				} else {
					password = ""
				}
			}
		} else {
			password = getRandomString(params.strLength, params.isCapital)
			if params.isCapital {
				//If isCapital flag true, we will check is there any capital letter in password.
				//If not set again until has a capital letter!
				for {
					if checkHasCapital(password) {
						break
					} else {
						password = getRandomString(params.strLength, params.isCapital)
					}
				}
			}
		}
		/*var password = getRandomString(params.strLength, params.isCapital)
		if params.isCapital {
			//If isCapital flag true, we will check is there any capital letter in password.
			//If not set again until has a capital letter!
			for {
				if checkHasCapital(password) {
					break
				} else {
					password = getRandomString(params.strLength, params.isCapital)
				}
			}
		}*/
		fmt.Println(password)
	} else if params.numLength > 0 {
		password := ""
		numberList := make(map[int]string)
		passwordLength := params.numLength + params.strLength

		if params.isSymbol {
			passwordLength += params.strLength/3 + 1
		}
		//Set Number Characters of the Password
		for i := 0; i < params.numLength; i++ {
			for {
				index := getRandomNumber(int64(passwordLength))
				//If you not get same RandomIndex for number character you will set RandomNumber for RandomIndex
				if _, ok := numberList[int(index)]; ok {

				} else if _, ok := symbolList[int(index)]; ok { //If this is not in one of the symbolIndex

				} else {
					numberList[int(index)] = strconv.Itoa(int(getRandomNumber(int64(10))))
					break
				}
			}
		}

		//Set String Characters of the Password
		//03.07.2022
		if params.strLength > 0 { // if setLength not set but numLength set => Password has only Number or NumberAndSymbol
			for {
				for i2 := 0; i2 < passwordLength; i2++ {
					if _, ok := numberList[i2]; ok {
						password += numberList[i2]
					} else if _, ok := symbolList[i2]; ok {
						password += symbolList[i2]
					} else {
						password += getRandomString(1, params.isCapital)
					}
				}
				//If isCapital flag false or there is at least one capital letter in password, we will break the loop
				//If not set String Characters again until has a capital letter!
				if !params.isCapital || checkHasCapital(password) {
					break
				} else {
					password = ""
				}
			}
		} else { //Password has only Numbers and Symbols 03.07.2022
			for i2 := 0; i2 < passwordLength; i2++ {
				if _, ok := numberList[i2]; ok {
					password += numberList[i2]
				} else if _, ok := symbolList[i2]; ok {
					password += symbolList[i2]
				}
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
var symbolRunes = []rune("~=+%^*/()[]{}/!@#$?|")

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

func getRandomSymbol() string {
	result, _ := rand.Int(rand.Reader, big.NewInt(int64(len(symbolRunes))))
	return string(symbolRunes[result.Int64()])
}

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

func checkHasCapital(password string) bool {
	for i := 0; i < len(password); i++ {
		if unicode.IsUpper(rune(password[i])) {
			return true
		}
	}
	return false
}
