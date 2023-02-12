package requests

import (
	"fmt"
	sh "github.com/VaradBelwalkar/go_client_for_windows/session_handling"
)

func Container_List(){
    colorReset := "\033[0m"
    colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorRed := "\033[31m"
	request_path:="/container/list/containers"
	//resp is of type map[string]interface{}
	// Here the object is derived from the JSON received from the response
	resp,err:= sh.GET_Request(request_path)  

	if err == 500 || err == 504 || err == 505{
			fmt.Println(string(colorRed),"Server error!",string(colorReset))
		return
	} else if err == 502{
		return
	} else if err == 401{
		fmt.Println(string(colorRed),"Something went wrong on your side!",string(colorReset))
		return
	} 
	resultInfo:=make(map[string]string)
	for k,v:=range resp{
		resultInfo[k]=v.(string)
	}

	if len(resultInfo) == 0{
		fmt.Println(string(colorYellow),"You don't have any running containers!",string(colorReset))
		return
	}

	for k, v := range resultInfo { 
		fmt.Println(string(colorGreen),k,string(colorReset),string(colorYellow),v,string(colorReset))
		
	}	
	
}

