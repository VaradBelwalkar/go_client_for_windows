package requests

import (
	"fmt"
	"os/exec"
	"os"
	"bufio"
	"io/ioutil"
	sh "github.com/VaradBelwalkar/go_client_for_windows/session_handling"
)


func Container_Start(imageName string,browser bool,given_port string){
    colorReset := "\033[0m"
	colorBlue := "\033[34m"
    colorRed := "\033[31m"
	colorGreen := "\033[32m"
    colorYellow := "\033[33m"
	request_path:="/container/resume/"+imageName
	reader := bufio.NewReader(os.Stdin)

	//resp is of type map[string]interface{}
	resp,status:= sh.GET_Request(request_path)  

	if status!=200 {
		  if status == 500{
		fmt.Println(string(colorRed),"Server error!",string(colorReset))
		return
		} else if status == 502{
			return
		}else if status == 401{
			fmt.Println(string(colorRed),"Something went wrong on your side!",string(colorReset))
			return
		} else if status == 404{
			fmt.Println(string(colorYellow),"No such container!",string(colorReset))
			return
		}
	}
	user_credentials,err:=sh.Show_Credentials()
	if err!=nil{
		fmt.Println(string(colorYellow),"Please run change config to store your credentials",string(colorReset))
	}
	privateKey:=resp["privatekey"].(string)	
	container_ip:=resp["container_ip"].(string)
	 //define the path to the bash script
	
	err = ioutil.WriteFile(sh.ProjectPath+"\\keyForRemoteServer", []byte(privateKey), 0600)
    if err != nil {
        fmt.Println(string(colorRed),"Something went wrong while storing PrivateKey",string(colorReset))
		return
    }
	// Parameters to pass to the script
	fmt.Println(string(colorBlue),"Enter the following line in VSCode for remote development\n",string(colorReset))
	fmt.Println(string(colorGreen),"ssh "+"-i "+sh.ProjectPath+"\\keyForRemoteServer"+" root@"+container_ip,string(colorReset))
	fmt.Print("\nPress enter when copied: ")
	_, _ = reader.ReadString('\n')
	// start the script
	cmd := exec.Command("ssh","-i",sh.ProjectPath+"\\keyForRemoteServer","root@"+container_ip)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	// start the script and wait for it to finish
	if err := cmd.Start(); err != nil {
		// handle error
		fmt.Println(string(colorRed),"Something went wrong,\n Check ip address or port if configured correctly else might be server issue!",string(colorReset))
		return
	}
	if err := cmd.Wait(); err != nil {
		// handle error
		fmt.Println(string(colorRed),"Something went wrong,\n Check ip address or port if configured correctly else might be server issue!",string(colorReset))
	}
	
	
}

