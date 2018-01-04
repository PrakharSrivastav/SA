package main

import (
	"github.com/SA/config"
	"gopkg.in/resty.v1"
	"fmt"
)

type tokenRequest struct {
	Grant_type string `json:"grant_type"`
	Client_id  string `json:"client_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Connection string `json:"connection"`
	Scope      string `json:"scope"`
}

type tokenSuccess struct {
	Id_token     string `json:"id_token"`
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
}

type tokenFailure struct {
	Error             string `json:"error"`
	Error_description string `json:"error_description"`
}

func main() {

	tokenDetails, failureDetails, err := getAccessToken()
	if err != nil {
	}
	if failureDetails != nil {
	}
	token := *tokenDetails
	fmt.Println(token.Access_token)
	fmt.Println(token.Id_token)
	fmt.Println(token.Token_type)
}

func getAccessToken() (*tokenSuccess, *tokenFailure, error) {
	configs, err := config.Config()
	if err != nil {
		panic(err)
	}
	body := tokenRequest{
		Grant_type: configs["auth0granttype"].(string),
		Client_id:  configs["auth0clientd"].(string),
		Connection: configs["auth0connection"].(string),
		Password:   configs["auth0password"].(string),
		Username:   configs["auth0username"].(string),
		Scope:      configs["auth0scope"].(string),
	}

	response, err := resty.R().
		SetBody(body).
		SetResult(tokenSuccess{}).
		SetError(tokenFailure{}).
		Post(configs["auth0host"].(string))

	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	fmt.Println("response code is ", response.StatusCode())
	if response.StatusCode() == 200 {
		r := response.Result().(*tokenSuccess)
		//fmt.Printf("Access Token %v	\n", r.Access_token)
		//fmt.Printf("Id Token %v	\n", r.Id_token)
		//fmt.Printf("Token Type %v	\n", r.Token_type)
		return r, nil, nil
	} else {
		err := response.Error().(*tokenFailure)
		//fmt.Printf("Error %v \n", err.Error)
		//fmt.Printf("Error Description %v \n ", err.Error_description)
		return nil, err, nil
	}
}
