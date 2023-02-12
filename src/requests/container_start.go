package requests

import (
	"fmt"
	"log"
	"os/exec"
	"os"
	"io/ioutil"
	sh "github.com/VaradBelwalkar/go_client_for_windows/session_handling"
)


func Container_Start(imageName string){
    colorReset := "\033[0m"

    colorRed := "\033[31m"
    colorYellow := "\033[33m"
	request_path:="/container/resume/"+imageName


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
		}
	}
	user_credentials,err:=sh.Show_Credentials()
	if err!=nil{
		fmt.Println(string(colorYellow),"Please run change config to store your credentials",string(colorReset))
	}
	privateKey:=resp["privatekey"].(string)	
	port:=resp["port"].(string)
	 //define the path to the bash script
	scriptPath := sh.ProjectPath+"\\src\\connections\\ps_script.ps1"
	
	err = ioutil.WriteFile(sh.ProjectPath+"\\src\\connections\\keyForRemoteServer", []byte(privateKey), 0644)
    if err != nil {
        fmt.Println(string(colorRed),"Something went wrong while storing PrivateKey",string(colorReset))
		return
    }
	// Parameters to pass to the script
	
	// start the script
	cmd := exec.Command("powershell.exe", "-File",scriptPath,port,user_credentials["ip"],sh.ProjectPath+"\\src\\connections\\keyForRemoteServer")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	// start the script and wait for it to finish
	if err := cmd.Start(); err != nil {
		// handle error
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		// handle error
		log.Fatal(err)
	}
	
	
}

