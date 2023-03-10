package session_handling



import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
	"net/http"

)


//JWT will be in the following format

//jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"






//The JWT standard defines a specific structure for the encoded token, which consists of three parts separated by dots:

//The first part is the header, which contains information about how the JWT is encoded.
//The second part is the payload, which contains the claims. 
//Claims are statements about an entity (typically, the user) and additional data. Claims are encoded as a JSON object.
//The third part is the signature, which is used to verify that the sender of the JWT is who it says
//it is and to ensure that the message wasn't changed along the way.


//To encode a JWT, you typically need to do the following:

//Create the header and payload as JSON objects.
//Base64-encode the header and payload.
//Concatenate the encoded header, payload, and a secret key with dots (.) to form the JWT.



// Standerdized and returns type of map[string]interface{}
// To be used after successful login and JWT retrieval
func GET_Request(request_path string) (map[string]interface{},int) {
    colorReset := "\033[0m"

    colorRed := "\033[31m"
	_, ok := os.LookupEnv("JWT")
	if ok==false{
			fmt.Println(string(colorRed),"Please login",string(colorReset))
			return nil,502
	}
	JWT:=os.Getenv("JWT")

	_, ok = os.LookupEnv("session")
	if ok==false{
		fmt.Println(string(colorRed),"Please login",string(colorReset))
		return nil,502	
	}
  	cookieValue:=os.Getenv("session")
	cookie := &http.Cookie{
        Name:   "session",
        Value:  cookieValue,
        MaxAge: 300,
    }
	credHolder,err:=Show_Credentials()
	if err!=nil{
		fmt.Println(string(colorRed),"Something went wrong while reading credentials file!",string(colorReset))
		return nil,502
	}

	req, err := http.NewRequest("GET","http://"+credHolder["ip"]+":"+credHolder["port"]+request_path,nil)
	client:=&http.Client{}
	req.Header.Set("Authorization","Bearer "+JWT) // JWT must be available
	req.AddCookie(cookie)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(string(colorRed),"Error while receiving response",string(colorReset))
		return nil,500
	}
	defer res.Body.Close()

	if res.StatusCode!=200{
		if res.StatusCode == 401{
		fmt.Println(string(colorRed),"Please login again!",string(colorReset))
		return nil,401
		}	else {
			return nil, res.StatusCode
		}
	}	

	// Read the response body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil,504 	//response body isn't there
	}
	// Unmarshal the response body into a map interface 
	var response map[string]interface{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil,505   // body is there but cannot unmarshal it
	}
	return response,200

}



// To be used after successful login and JWT retrieval
func POST_Request(request_path string, data map[string]interface{}) (map[string]interface{},int) {
	colorReset := "\033[0m"

    colorRed := "\033[31m"
	b, err := json.Marshal(data)
	client:=&http.Client{}
	if err != nil {
		fmt.Println(string(colorRed),"something went wrong",string(colorReset))
	}
	//Change URL here

	credHolder,err:=Show_Credentials()
	if err!=nil{
		fmt.Println(string(colorRed),"Something went wrong while reading credentials file!",string(colorReset))
		return nil,502
	}

	req, err := http.NewRequest("POST","http://"+credHolder["ip"]+":"+credHolder["port"]+request_path, bytes.NewBuffer(b))
	if err != nil {
		return nil,502
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Add the JWT to the request header
		_, ok := os.LookupEnv("JWT")
	if ok==false{
		fmt.Println(string(colorRed),"Please login",string(colorReset))
		return nil,502	
	}
	JWT:=os.Getenv("JWT")

	_, ok = os.LookupEnv("session")
	if ok==false{
		fmt.Println(string(colorRed),"Please login",string(colorReset))
		return nil,502	
	}
  	cookieValue:=os.Getenv("session")
	cookie := &http.Cookie{
        Name:   "session",
        Value:  cookieValue,
        MaxAge: 300,
    }

	req.AddCookie(cookie)
	req.Header.Set("Authorization", "Bearer "+JWT)

	// Add the cookie to the request
	//req.AddCookie(&http.Cookie{Name: "session_id", Value: "123456"})

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(string(colorRed),"Server not responding!",string(colorReset))
		return nil,502
	}
	defer res.Body.Close()

	if res.StatusCode ==200{}else if res.StatusCode==401{
		fmt.Println(string(colorRed),"Please login again!",string(colorReset))
	}else{
		return nil,502
	}
	// Read the response
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil,200
	}

	// Unmarshal the response into a Response struct
	var response map[string]interface{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil,500
	}

	return response,200

	
}


