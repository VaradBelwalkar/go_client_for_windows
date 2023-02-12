package requests

import (
	"fmt"
	sh "github.com/VaradBelwalkar/go_client_for_windows/session_handling"
)


func Container_stop(containerName string){
    colorReset := "\033[0m"

    colorRed := "\033[31m"
    colorGreen := "\033[32m"
    colorYellow := "\033[33m"
	request_path:="/container/stop/"+containerName


	//resp is of type map[string]interface{}
	_,err := sh.GET_Request(request_path)

	if err!=200{
		if err == 504 || err==505{
			fmt.Println(string(colorGreen),"Container successfully stopped!",string(colorReset))
			return
		} else if err == 500{
			fmt.Println(string(colorRed),"Server error sdlkfjsdlkj!",string(colorReset))
			return
		} else if err == 502{
			return
		} else if err == 404{
			fmt.Println(string(colorYellow),"Such container doesn't exist!",string(colorReset))
			return
		} else if err == 401{
			fmt.Println(string(colorRed),"Something went wrong on your side!",string(colorReset))
			return
		}
	} else{
		fmt.Println(string(colorGreen),"Container successfully stopped!",string(colorReset))
		return
	}
return
	
}

