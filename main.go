package main

import (
    "bufio"
    "fmt"
    "os"
	"os/exec"
    "strings"
	"github.com/VaradBelwalkar/go_client"
)



func main() {
	cmd := exec.Command("cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
    for {
		
        // Prompt the user for input
        fmt.Print("go > ")

        // Read the user's input
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println(err)
            continue
        }

        // Strip leading and trailing whitespace from the input
        input = strings.TrimSpace(input)

        // Split the input into separate words
        words := strings.Split(input, " ")

        // Check the first word to see which command the user entered
        switch words[0] {
        case "exit":
            // Exit the program
            fmt.Println("Exiting...")
            return
        case "help":
            // Print the help
			cmd := exec.Command("cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
            fmt.Println(Help)
		case "container":
			switch words[1] :
				case "run":
					//Run appropriate function by passing value of words[2] (container requested)
				case "list":
					if words[2] == "images"{
						//Run appropriate function
					}else{
						//Run appropriate function
					}
			

			
		case "upload":
			switch words[1] {}
        default:
            // Print an error message
            fmt.Println("Unknown command")
        }
    }
}
