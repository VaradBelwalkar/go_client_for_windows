package session_handling

import (
	"bufio"
	"fmt"
	"os"
	"net/http"
	"strings"
	"net/url"
	"time"
	"github.com/PuerkitoBio/goquery"
)




// This function logs into the server and preserves JWT for further communication
func Register() {
	colorReset := "\033[0m"
	//colorGreen := "\033[32m"
    colorRed := "\033[31m"
	// Create a new HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
 	
	//Do IO on file to retrieve username and password 

	//Do whenever submitting form data
	data := url.Values{}

	
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the username you want: ")
	tempName, _ := reader.ReadString('\n')
	username:=strings.ReplaceAll(tempName,"\r\n","")

	fmt.Print("Enter your password: ")
	tempPassword, _ := reader.ReadString('\n')
	password:=strings.ReplaceAll(tempPassword,"\r\n","")
	fmt.Print("confirm your password: ")
	tempPassword1, _ := reader.ReadString('\n')
	password1:=strings.ReplaceAll(tempPassword1,"\r\n","")

    for {
		if password == password1{
			break;
		}
        fmt.Println(string(colorRed),"Your Password does not match. Please try again",string(colorReset))
		fmt.Print(" Enter your password: ")
		tempPassword, _ = reader.ReadString('\n')
		password=strings.ReplaceAll(tempPassword,"\r\n","")
		fmt.Print("confirm your password: ")
		tempPassword1, _ = reader.ReadString('\n')
		password1=strings.ReplaceAll(tempPassword1,"\r\n","")
	}
	fmt.Print("Enter the email ID: ")
	tempEMAIL,_:=reader.ReadString('\n')
	EMAIL:=strings.ReplaceAll(tempEMAIL,"\r\n","")
	EMAIL=strings.ReplaceAll(EMAIL," ","")

	fmt.Print("Enter the server IP: ")
	tempIP,_:=reader.ReadString('\n')
	IP:=strings.ReplaceAll(tempIP,"\r\n","")
	IP=strings.ReplaceAll(IP," ","")
	fmt.Print("Enter the port: ")
	tempPort,_:=reader.ReadString('\n')
	port:=strings.ReplaceAll(tempPort,"\r\n","")
	port=strings.ReplaceAll(port," ","")
	//Request made to get the form required	
	urlString:=	"http://"+strings.ReplaceAll(IP, " ", "")+":"+strings.ReplaceAll(port, " ", "")+"/register"
	res,err:=http.Get(urlString)
	if err!=nil{
		fmt.Println(string(colorRed),"Something went wrong,\n Check ip address or port if configured correctly else might be server issue!",string(colorReset))
		return
	}
	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Find the hidden field with the name "csrf_token"
	csrfToken := doc.Find("input[name=csrf]").First().AttrOr("value", "")
		if csrfToken == "" {
		fmt.Println(string(colorRed),"CSRF token not found", string(colorReset))
		return
	}

	//Preparing the body of the POST request, which is nothing but form data being sent using appropriate header
	data.Add("username", username)
	data.Add("password", password) 
	data.Add("email", EMAIL) 
	data.Add("csrf",csrfToken)


	cookie := &http.Cookie{
        Name:   "csrftoken",
        Value:  csrfToken,
        MaxAge: 300,
    }
	req,err:= http.NewRequest("POST","http://"+strings.ReplaceAll(IP, " ", "")+":"+strings.ReplaceAll(port, " ", "")+"/register",strings.NewReader(data.Encode()))
	if err!=nil{
		fmt.Println(err)
		return 
	}
	req.AddCookie(cookie)
	//The header is set to this to recognise that the body of the request is holding form data
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")
	
	//Here the request is being actually sent
	//the response object will contain the JWT token
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	
	// Unmarshal the response into a Response struct
	// This is just temporary implementation which doesn't contribute to any functionality as 
	// we are not currently having any response info in the body as of now
	//var response http.Response
	//body, err := ioutil.ReadAll(res.Body)
	//err = json.Unmarshal(body, &response)
	//if err != nil {
	//	return
	//}

	//The JWT token
	 //Here you can access this token anywhere in this package
    

	//Here the 201 StatusCode means the resource is successfully created on the server
	if res.StatusCode == 200 {
		Store_credentials(username,password,IP,port)
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
		Verify_Register_OTP()

	} else if  res.StatusCode == 424{
		fmt.Println(string(colorRed),"The Email ID you provided is not authorized with the server, Please contact DYPLUG team for more details", string(colorReset))
		return
	} else if res.StatusCode == 406{
		fmt.Println(string(colorRed),"The Email ID you provided has already an account registered!", string(colorReset))
	} else if res.StatusCode == 409 {  //    409 StatusCode indicates a "Conflit" that server cannot create a resource because
											  //    it already exists
		fmt.Println(string(colorRed),"The username already exists! Please choose another username", string(colorReset))
		return

	} else if res.StatusCode == 400{ 
		fmt.Println(string(colorRed),"Something went wrong on your side!", string(colorReset))
		return
		
	} else if res.StatusCode == 500 {
		fmt.Println(string(colorRed),"Something went wrong on server side!", string(colorReset))
		return
	} else if res.StatusCode == 412 {
		fmt.Println(string(colorRed),"CSRF Authentication Failed!", string(colorReset)) 		// http.StatusPreconditionFailed
		return
	}else {
		fmt.Println("something went wrong!")
		return
	}


}


func Verify_Register_OTP(){
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

	fmt.Print(string(colorYellow),"Please enter the OTP sent to the given EMAIL ID: ",string(colorReset))
	tempOTP,_:=reader.ReadString('\n')
	OTP:=strings.ReplaceAll(tempOTP,"\r\n","")
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
	req,err= http.NewRequest("POST","http://"+user_credentials["ip"]+":"+user_credentials["port"]+"/regotphandler",strings.NewReader(data.Encode()))
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
	
	if res.StatusCode == 200 {
		//Meaning the registration is successful
		fmt.Println(string(colorGreen),"Registered successfully!", string(colorReset))
		
		return

	} else if res.StatusCode == 400{ 
		fmt.Println(string(colorRed),"Something went wrong on your side!", string(colorReset))
		return
		
	} else if res.StatusCode == 500 {
		fmt.Println(string(colorRed),"Something went wrong on server side!", string(colorReset))
		return
	} else if res.StatusCode == 412 {
		fmt.Println(string(colorRed),"CSRF Authentication Failed!", string(colorReset)) 		// http.StatusPreconditionFailed
		return
	}else {
		fmt.Println("something went wrong!")
		return
	}


}












