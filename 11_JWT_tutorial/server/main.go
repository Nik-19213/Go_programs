package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "✅ Authorized: This information is a secret")
	fmt.Println("✅ Authorized request received at homePage")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenHeader := r.Header.Get("Token")
		if tokenHeader == "" {
			http.Error(w, "❌ Not Authorized: Missing Token", http.StatusUnauthorized)
			return
		}

		// Parse the JWT Token
		token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return mySigningKey, nil
		})

		// Handle token errors
		if err != nil {
			http.Error(w, "❌ Invalid Token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// If token is valid, continue to the endpoint
		if token.Valid {
			endpoint(w, r)
		} else {
			http.Error(w, "❌ Invalid Token", http.StatusUnauthorized)
		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("Server is running on port 9000...")
	handleRequests()
}
