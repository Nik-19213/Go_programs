package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func homepage(w http.ResponseWriter, r *http.Request) {
	// Generate JWT Token
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintln(w, "Failed to generate token")
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:9000/", nil)
	if err != nil {
		fmt.Fprintf(w, "Error creating request: %s", err.Error())
		return
	}

	// ✅ Attach the token in the request header
	req.Header.Set("Token", validToken)

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error making request: %s", err.Error())
		return
	}

	defer res.Body.Close()
	// Read and print the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, "Error reading response: %s", err.Error())
		return
	}

	fmt.Fprint(w, string(body))
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %s", err.Error())
	}
	return tokenString, nil
}

func handleRequests() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("Client is running on port 9001...")
	handleRequests()
}
