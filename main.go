package main

import (
	"fmt"
	"strings"

	"github.com/mbndr/figlet4go"

	"maxnormand/get-my-info/api"
)

func renderWelcomeText() string {
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	options.FontName = "larry3d"

	renderStr, _ := ascii.Render("Good Morning!")
	return renderStr
}

func printApiCall(text string) string {
	message := fmt.Sprintf("Going to get your %s info...", text)
	return message
}

func main() {
	fmt.Print(renderWelcomeText())
	// putting the application in a for loop makes it recursively run
	for {
		fmt.Print("Enter 'chess' to get the daily chess Puzzle" +
			"\nor 'cat' to get cat facts" +
			"\nor 'exit' to leave the application:\n")

		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			panic(err)
		}

		switch strings.ToLower(input) {
		case "chess":
			// TODO: could have a handle API call function that does this
			fmt.Println(printApiCall("Chess"))
			res, err := api.GetChessPuzzle()
			if err != nil {
				panic(err)
			}
			fmt.Println(res)
			fmt.Println()
			fmt.Println()
		case "cat":
			fmt.Println(printApiCall("Cat"))
			catFacts, err := api.GetCatFacts()
			if err != nil {
				panic(err)
			}
			fmt.Println(catFacts)
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid input. Please enter 'chess' or 'cat'.")
		}
	}
}
