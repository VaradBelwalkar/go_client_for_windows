package requests

import (
	"fmt"
	"os/exec"
	"os"
	"strings"
	sh "github.com/VaradBelwalkar/go_client_for_windows/session_handling"
)
func Downloads(fileOrFolder string,containerPath string,userPath,containerName string){
	colorReset := "\033[0m"
	colorRed := "\033[31m"
    colorYellow := "\033[33m"
	

	parts := strings.Split(containerName, "_")
	container_ip := parts[1]
	cmd := exec.Command("scp","-i",sh.ProjectPath+"\\keyForRemoteServer","root@"+container_ip+":"+containerPath,userPath)
	if fileOrFolder == "file"{

	} else if fileOrFolder == "folder"{
		cmd = exec.Command("scp","-r","-i",sh.ProjectPath+"\\keyForRemoteServer","root@"+container_ip+":"+containerPath,userPath)
	}


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



// 	download <file,folder> <path in container> <path in your computer> <containername> 