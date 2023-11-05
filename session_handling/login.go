package session_handling

import (
	"fmt"
	"os"
	"bufio"
	"net/url"
	"strings"
	"net/http"
	"time"
	"github.com/PuerkitoBio/goquery"
	//"github.com/VaradBelwalkar/go_client"
)

		//Content-Type is one of the headers available in http request

//	Content-Type -------> application/json  

//To pass normal json data 


//Content-Type --------> application/x-www-form-urlencoded

//To recogise on server side that the request is holding form data (data submitted through form)

//Here the data is submitted as a string(obviously)as it is NOT mentioned like this "application/json"
//In a URL-encoded query string, the characters in the string are encoded so that they can be safely transmitted over the Internet. 
//The encoding replaces certain characters with a percent sign followed by a two-digit hexadecimal number. 
//or example, a space is encoded as "%20" and a plus sign is encoded as "%2B".

//For example, consider an HTML form with two text fields, 
//"name" and "email", and a submit button. 
//When the user fills out the form and clicks the submit button, the form data is sent to the server in the body of a POST request. 
//The "Content-Type" header of the request would be set to "application/x-www-form-urlencoded" and the body of the 
//request would look something like this:

//name=John+Doe&email=johndoe%40example.com

//This format is useful because it is simple and easy to parse on the server side.
//However, it has a limitation in that it can only transmit ASCII characters and does not support file uploads. 
//For more advanced functionality, other media types such as "multipart/form-data" or "application/json" may be used instead.









		// Authorization is one of the headers available in the http request
	
//  Authorization -----------> "Bearer <JWT-Token>"

//This indicates the authorization method, which is JWT one



	//Bearer is used indicate JWT based authentication while Basic is used to indicate username:password based authentication

//Basic authentication: Basic authentication is a simple authentication scheme that uses a username and password to authenticate requests. 
// To use basic authentication, you can set the value of the Authorization header to "Basic <credentials>",
// where <credentials> is the base64-encoded string of the username and password separated by a colon. For example:
// Authorization: Basic QWxhZGRpbjpvc


				//CSRF Token 

//Cross-Site Request Forgery (CSRF) tokens are usually passed in a header or as a request parameter in the query string.
//The specific location of the CSRF token will depend on the requirements of the server and the client.
//Here we are passing CSRF Token as header in the request field








									//Understand how CSRF works


//First we make get request to the /login page to the server, then server sends the form with the CSRF token embedded into it
//You don't need to retrieve that CSRF token any way, as it is in the form itself and when you submit it ,
//with the POST request, the CSRF token gets sent with the Form data 
//You just need to retrieve the Form from html and fill in appropriate values in it and then just submit it 

//It is important to make get request and get the form as ,

//Directly sending POST request without CSRF token will not work

//CSRF tokens can be passed as header from the server but,
//They are mostly embedded in the html documents within form field like

// 					<form action="/login" method="post">
// 					  <!-- Other form fields go here -->
// 					  <input type="hidden" name="csrf_token" value="abc123">
// 					  <button type="submit">Log In</button>
// 					</form>


//It is common for servers to send the JSON Web Token (JWT) to the client in the response body of the login request. 
//The client can then extract the JWT from the response and store it locally for use in authenticating subsequent requests.

//For example, the server may return the JWT in the response body as a JSON object:

//   {
//  	"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
//   }
  
//The client can then parse the JSON response and extract the "access_token" field to obtain the JWT.

//Alternatively, the server may send the JWT in the "Authorization" header of the response. For example:

// Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c



//We will assume server is going to send the JWT token through header 

// This function logs into the server and preserves JWT for further communication
func Login() {
	colorReset := "\033[0m"
	colorYellow := "\033[33m"
    colorRed := "\033[31m"
	// Create a new HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
 
	//Do IO on file to retrieve username and password 

	user_credentials,err:=Show_Credentials()
	if err!=nil{
		fmt.Println(string(colorYellow),"Please run change config to store your credentials",string(colorReset))
		return
	}
	//Do whenever submitting form data
	data := url.Values{}

	

	//Request made to get the form required
	resp,err:=http.Get("http://"+user_credentials["ip"]+":"+user_credentials["port"]+"/login")
	if err!=nil{
		fmt.Println(string(colorRed),"Something went wrong,\n Check ip address or port if configured correctly else might be server issue!",string(colorReset))
		return
	}
	defer resp.Body.Close()
	

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(string(colorRed),"something went wrong",string(colorReset))
		return
	}

	// Find the hidden field with the name "csrf_token"
	csrfToken := doc.Find("input[name=csrf]").First().AttrOr("value", "")
	if csrfToken == "" {
		fmt.Println("CSRF token not found")
		return 
	}

	//Preparing the body of the POST request, which is nothing but form data being sent using appropriate header
	data.Add("username", string(user_credentials["username"]))
	data.Add("password", string(user_credentials["password"])) //To be retrieved 
	data.Add("csrf",csrfToken)

	cookie := &http.Cookie{
        Name:   "csrftoken",
        Value:  csrfToken,
        MaxAge: 300,
    }
	req,err:= http.NewRequest("POST","http://"+user_credentials["ip"]+":"+user_credentials["port"]+"/login",strings.NewReader(data.Encode()))
	if err!=nil{
		return 
	}
	req.AddCookie(cookie)
	//The header is set to this to recognise that the body of the request is holding form data
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	
	//Here the request is being actually sent
	//the response object will contain the JWT token
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(string(colorRed),"Something went wrong!",string(colorReset))
		return 
	}
	defer res.Body.Close()
	
	//We can get here statuses only 403 or 208 
	if res.StatusCode!=200{
	if res.StatusCode==401 {
		fmt.Println(string(colorRed),"Wrong username or password!",string(colorReset))
		return
	} else if res.StatusCode == 404{
		fmt.Println(string(colorRed),"User doesn't exist!\n",string(colorReset),string(colorYellow),"Please correct your username or if not registered, please register first",string(colorReset))
		return 
	}}

	var sessionID string
	for _, cookie := range res.Cookies() {
		if cookie.Name == "session" {
			sessionID = cookie.Value
			break
		}
	}

	//The JWT token
	JWT:= res.Header.Get("authorization")    //Here you can access this token anywhere in this package
	splitToken:=strings.Split(JWT, "Bearer ")
	tokenString:=splitToken[1]
	os.Setenv("JWT",tokenString)
	os.Setenv("session",sessionID)
	Verify_OTP()
	
//Login completed

return 

}



func Verify_OTP(){
	run_again:
    colorReset := "\033[0m"
	colorYellow := "\033[33m"
    colorRed := "\033[31m"
	colorGreen := "\033[32m"
	data := url.Values{}
	reader := bufio.NewReader(os.Stdin)
	user_credentials,err:=Show_Credentials()
	if err!=nil{
		fmt.Println(string(colorYellow),"Please run change config to store your credentials",string(colorReset))
		return
	}
	_, ok := os.LookupEnv("JWT")
	if ok==false{
			fmt.Println(string(colorRed),"Please login",string(colorReset))
			return
	}
	JWT:=os.Getenv("JWT")

	_, ok = os.LookupEnv("session")
	if ok==false{
		fmt.Println(string(colorRed),"Please login",string(colorReset))
		return 
	}
  	cookieValue:=os.Getenv("session")
	cookie := &http.Cookie{
        Name:   "session",
        Value:  cookieValue,
        MaxAge: 300,
    }

	req, err := http.NewRequest("GET","http://"+user_credentials["ip"]+":"+user_credentials["port"]+"/otphandler",nil)
	client:=&http.Client{}
	req.Header.Set("Authorization","Bearer "+JWT) // JWT must be available
	req.AddCookie(cookie)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(string(colorRed),"Error while receiving response",string(colorReset))
		return 
	}
	defer res.Body.Close()

	if res.StatusCode!=200{
		if res.StatusCode == 401{
		fmt.Println(string(colorRed),"Please login again!",string(colorReset))
		return 
		}	else {
			return 
		}
	}	

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(string(colorRed),"something went wrong",string(colorReset))
		return
	}

	// Find the hidden field with the name "csrf_token"
	csrfToken := doc.Find("input[name=csrf]").First().AttrOr("value", "")
	if csrfToken == "" {
		fmt.Println("CSRF token not found")
		return 
	}

	fmt.Print(string(colorYellow),"Please enter the OTP sent to your registered EMAIL ID: ",string(colorReset))
	tempOTP,_:=reader.ReadString('\n')
	OTP:=strings.ReplaceAll(tempOTP,"\n","")
	OTP=strings.ReplaceAll(OTP," ","")
	//Preparing the body of the POST request, which is nothing but form data being sent using appropriate header
	data.Add("otp", OTP)
	data.Add("csrf",csrfToken)

	csrfCookie := &http.Cookie{
        Name:   "csrftoken",
        Value:  csrfToken,
        MaxAge: 30000,
    }
	sessionCookie:=&http.Cookie{
        Name:   "session",
        Value:  cookieValue,
        MaxAge: 30000,
    }
	req,err= http.NewRequest("POST","http://"+user_credentials["ip"]+":"+user_credentials["port"]+"/otphandler",strings.NewReader(data.Encode()))
	if err!=nil{
		return 
	}
	req.AddCookie(csrfCookie)
	req.AddCookie(sessionCookie)
	req.Header.Set("Authorization","Bearer "+JWT)
	//The header is set to this to recognise that the body of the request is holding form data
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	
	//Here the request is being actually sent
	//the response object will contain the JWT token
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(string(colorRed),"Something went wrong!",string(colorReset))
		return 
	}
	defer res.Body.Close()
	
	//We can get here statuses only 403 or 208 
	if res.StatusCode!=200{
	if res.StatusCode==401 {
		fmt.Println(string(colorRed),"Wrong OTP!",string(colorReset))
		goto run_again
	} else if res.StatusCode == 404{
		fmt.Println(string(colorRed),"User doesn't exist!\n",string(colorReset),string(colorYellow),"Please correct your username or if not registered, please register first",string(colorReset))
		return 
	}}

	var sessionID string
	for _, cookie := range res.Cookies() {
		if cookie.Name == "session" {
			sessionID = cookie.Value
			break
		}
	}

	//The JWT token
	JWT= res.Header.Get("authorization")    //Here you can access this token anywhere in this package
	splitToken:=strings.Split(JWT, "Bearer ")
	tokenString:=splitToken[1]
	os.Setenv("JWT",tokenString)
	os.Setenv("session",sessionID)
	fmt.Println(string(colorGreen),"Login Completed!",string(colorReset))
	return


}


