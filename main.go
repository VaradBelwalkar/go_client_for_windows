package main

import (
    "bufio"
    "fmt"
    "os"
	"os/exec"
    "strings"
    //"github.com/VaradBelwalkar/help"
    //"github.com/VaradBelwalkar/session_handling"
	//
    sh "github.com/VaradBelwalkar/go_client/session_handling"
    h "github.com/VaradBelwalkar/go_client/help"
    rq "github.com/VaradBelwalkar/go_client/src/requests"

)



func main() {
    colorReset := "\033[0m"

    colorRed := "\033[31m"
    colorGreen := "\033[32m"
    colorYellow := "\033[33m"
    //colorBlue := "\033[34m"
   // colorPurple := "\033[35m"
    //colorCyan := "\033[36m"
    //colorWhite := "\033[37m"
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
    //Login with the present credentials

    for {
		
        // Prompt the user for input
        fmt.Print("dyplug > ")
        // Read the user's input
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')
        input=strings.TrimSuffix(input, "\n")
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
        case "clear":
            cmd := exec.Command("clear")
            cmd.Stdout = os.Stdout
            cmd.Run()
        case "register":
            sh.Register()
        case "remove":
            if len(words)>=2{
            if words[1]!="account"{
                fmt.Println("Wrong input!")
                continue
            }
            sh.Remove_account()}
        
        case "login":
            sh.Login()
        case "exit":
            // Exit the program
            fmt.Println(string(colorGreen),"Exiting...", string(colorReset))
            return
        case "help":
            // Print the help
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
            h.Help()
		case "container":
            if len(words)>=2{
			switch words[1] {
				case "list":
                    if len(words)>=3{
					if words[2] == "images"{
						//Run appropriate function
					}else if words[2] == "containers"{
						//Run appropriate function
                    rq.Container_List()
					} else{
                        fmt.Println("No such thing")
                    }}

                case "run":
                    if len(words)>=3{
                    rq.Container_Run(words[2])}else{fmt.Println(string(colorYellow),"Pass the name of the image you want", string(colorReset))}

                case "start":if len(words)>=3{
                    rq.Container_Start(words[2])}else{fmt.Println(string(colorYellow),"Pass the name of the container you want to start", string(colorReset))}
                case "stop":
                    if len(words)>=3{
                    rq.Container_stop(words[2])}else{fmt.Println(string(colorYellow),"Pass the name of the container you want to stop", string(colorReset))}
                case "remove": 
                    if len(words)>=3{
                    rq.Container_Remove(words[2])}else{fmt.Println(string(colorYellow),"Pass the name of the container you want to remove", string(colorReset))}
                
                    default:
                        fmt.Println("dyplug: "+"'"+words[1]+"'"+" is not a command\n See'help'")
                }}
        case "set":
            switch words[1]{
            case "url":
                // Call apropriate method here
                sh.Set_url()

            case "port":
                sh.Set_port()
            }


        case "change":
            switch words[1]{
            case "config":
                sh.Setup()

            default:
                fmt.Println(string(colorRed),"Unknown Command, try running help",string(colorReset))
            }
        case "config":
            //Print configuration here
            resp,err:=sh.Show_Credentials()
            if err!=nil{
                fmt.Println("No data found!Please fill up the data by running change config")
                continue
            }
            for k,v:=range resp{
                fmt.Print(string(colorGreen)," "+k+": ", string(colorReset))
                fmt.Println(v)
            }
			
		case "upload":
			switch words[1] {}
        default:
            if len(input)==0{

            }else{
                
            fmt.Println(string(colorRed),"Unknown Command, try running help",string(colorReset))}
        }
        
    }
}
