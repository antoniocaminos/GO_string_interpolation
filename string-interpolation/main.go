package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsADog        bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.UserName = readString("cual es tu nombre?")

	user.Age = readInt("cuantos años tenes?")

	user.FavouriteNumber = readFloat("numero favorito")

	user.OwnsADog = readBool("tiene pesho y/n")

	fmt.Printf("tu nombre es%s. y tenes %d años y tu numero favorito es %.2f. tiene pesho %t",
		user.UserName,
		user.Age,
		user.FavouriteNumber,
		user.OwnsADog)

}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("el nombre genio, con letras")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("una edad de verdad necesito")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("ingrese un numero")
		} else {
			return num
		}
	}
}

func readBool(s string) bool {
	for {
		err := keyboard.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			_ = keyboard.Close()
		}()
		for {
			fmt.Println(s)
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				log.Fatal(err)
			}
			if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
				fmt.Println("y para si , n para no")
			} else if char == 'n' || char == 'N' {
				return false
			} else if char == 'y' || char == 'Y' {
				return true
			}
		}
	}
}
