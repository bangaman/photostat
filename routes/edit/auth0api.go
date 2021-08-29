package edit

import (
	"os"
	"strings"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"html/template"
)


func(d *HomePage )  Auth0UpdateUsername(name, id string){

	url := "https://dev-vwwer09i.us.auth0.com/api/v2/users/"+id
	payload := strings.NewReader("{\"username\": \""+name+"\"}")
	req, _ := http.NewRequest("PATCH", url, payload)
	req.Header.Add("authorization", "Bearer "+os.Getenv("Api_token"))

	req.Header.Add("content-type", "application/json")

	res, _:= http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _:= ioutil.ReadAll(res.Body)

	/**
	 *	- reason for unmarshalling
	 *	- to check if status is equals succcess or not(404)
	**/

	// - jsonMap holds a map of strings
	jsonMap := make(map[string]interface{})

	// - unmarshaling body => first argument is bytes and second is referencing the jsonMap (map)
	err := json.Unmarshal([]byte(body), &jsonMap)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonMap)
	
	if jsonMap["error"] == nil{
		//if success reach this execution it means a 404 error didnt occur
		d.UsernameSuccess = true
	}else{
		if jsonMap["message"] == "The specified new username already exists"{
			d.UsernameErr = template.HTML("<p style='color:red; font-size:12px;'>Username taken already</p>")
		}else{
			d.UsernameErr =  template.HTML("<p style='color:red; font-size:12px;'>something went wrong</p>")
		}
	}
}


	
func(d *HomePage )  Auth0UpdateEmail(email, id string){

	url := "https://dev-vwwer09i.us.auth0.com/api/v2/users/"+id
	
	payload := strings.NewReader("{\"email\": \""+email+"\"}")
	req, _ := http.NewRequest("PATCH", url, payload)
	req.Header.Add("authorization", "Bearer "+os.Getenv("Api_token"))

	req.Header.Add("content-type", "application/json")

	res, _:= http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _:= ioutil.ReadAll(res.Body)

	/**
	 *	- reason for unmarshalling
	 *	- to check if status is equals succcess or not(404)
	**/

	// - jsonMap holds a map of strings
	jsonMap := make(map[string]interface{})

	// - unmarshaling body => first argument is bytes and second is referencing the jsonMap (map)
	err := json.Unmarshal([]byte(body), &jsonMap)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonMap)

	if jsonMap["error"] == nil{
		//if success reach this execution it means a 404 error didnt occur
		d.EmailSuccess = true
	}else{
		if jsonMap["message"] == "The specified new email already exists"{
			d.EmailErr = template.HTML("<p style='color:red; font-size:12px;'>Email taken already</p>")
		}else{
			d.EmailErr =  template.HTML("<p style='color:red; font-size:12px;'>something went wrong</p>")
		}
	}
	
}
