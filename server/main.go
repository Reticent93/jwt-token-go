package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

var mySigninKey = []byte("mysupersecrets")

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Super Secret Info")
}

//Middleware function takes in request and validates token
func isAuthorized(endpoint func(w http.ResponseWriter, r *http.Request))  http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {//If request has value of "Token
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error){
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return mySigninKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
			if token.Valid {//If token is valid
				endpoint(w,r)//We call our homepage func
			}
		}else {
			fmt.Fprintf(w, "Not Authorized")//This is purposely vague as to not give away too much details
		}
	})
}

func handleRequests()  {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("My Simple Server")
	handleRequests()
}
