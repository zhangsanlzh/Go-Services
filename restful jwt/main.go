package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func GetTokenString() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	claims["iss"] = "http://localhost"

	token.Claims = claims
	str, _ := token.SignedString([]byte("Luskyle"))
	return str
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", Login).Methods("GET")
	http.ListenAndServe(":8000", r)
}

func ReturnHTML(w http.ResponseWriter, url string) {
	t, err := template.ParseFiles(url)
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var valiad = Valiad(r.FormValue("account"), r.FormValue("password"))
	if valiad {
		ReturnCookie(w, r, GetTokenString())
		ReturnHTML(w, "client/welcome.html")
	} else {
		ReturnHTML(w, "client/loginFailed.html")
	}
}

func ReturnCookie(w http.ResponseWriter, r *http.Request, token interface{}) {
	c := http.Cookie{
		Name:     "your_token",
		Value:    token.(string),
		HttpOnly: true,

		// 取消注释，就会在 postman 的 cookie 项中显示 token 值
		// 否则 token 值 只能在 header 中看到
		// Secure:   true,
		MaxAge: 300}
	http.SetCookie(w, &c)
}
