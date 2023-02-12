package requests

import (
	"fmt"
	sh "github.com/VaradBelwalkar/go_client_for_windows/session_handling"
)


func Container_Remove(containerName string){
    colorReset := "\033[0m"
	colorYellow := "\033[33m"
    colorRed := "\033[31m"
    colorGreen := "\033[32m"
	request_path:="/container/remove/"+containerName


	//resp is of type map[string]interface{}
	_,err:= sh.GET_Request(request_path)
	if err!=200{

		if err == 404{
			fmt.Println(string(colorYellow),"No such container!",string(colorReset))
			return
		} else if err == 500{			
		fmt.Println(string(colorRed),"Server error!",string(colorReset))
		return
	} else if err ==502{
		return
	}else if err == 401{
		fmt.Println(string(colorRed),"Something went wrong on your side!",string(colorReset))
		return
	} else if err == 504 || err==505{
		fmt.Println(string(colorGreen),"Container Removed successfully!",string(colorReset))
	}
	} else{

		fmt.Println(string(colorGreen),"Container Removed successfully!",string(colorReset))
	}


	//Here the backend is going to give JSON response containing info with 4 values,
	// 200 == OK
	// 300 == 


//Handle the response here
	
	
}

