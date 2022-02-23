package main

import (
	"fmt"

	"github.com/dghubble/oauth1"


	"github.com/dghubble/go-twitter/twitter"
)

func main(){
	consumerKey:=""
	consumerSecret:=""
	accessToken:=""
	acessTokenSecret:=""
	config:=oauth1.NewConfig(consumerKey,consumerSecret)
	token:=oauth1.NewToken(accessToken,acessTokenSecret)

	httpClinet := config.Client(oauth1.NoContext,token)
	client := twitter.NewClient(httpClinet)


	user, _,err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})

	if err != nil{
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("Account: @%s (%s)\n", user.ScreenName, user.Name)

	_,_, err = client.Statuses.Update("working on a twitter bot with goland",nil)

	if err != nil{
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("tweet sent successfully")

}