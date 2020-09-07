package main

import ("fmt"
"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//var mySigningKey = os.Get("MY_JWT_TOKEN")
var mySigningKey = []byte("mysupersecrets")

//func generates a jwt token after request is received and will return a response
func homePage(w http.ResponseWriter, r *http.Request)  {
	validateToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:9000/", nil)//creates a new request using GET
	req.Header.Set("Token", validateToken)//generates a token
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, string(body))

}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["users"] = "Reticent93"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func handleRequests()  {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}


func main() {
	fmt.Println("My Simple Client")
	handleRequests()
}
