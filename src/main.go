package main

import (
	"log"
	"net/http"
	"time"

	"api.jwt.auth/controllers"
	"api.jwt.auth/core/authentication"
	"github.com/urfave/negroni"
)

func main() {
	router := NewRouter()

	// On the default page we will simply serve our static index page.
	router.Handle("/", http.FileServer(http.Dir("./views/")))
	// For auth purposes
	router.Handle("/get-token", GetTokenHandler).Methods("GET")
	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	router.HandleFunc("/token-auth", controllers.Login).Methods("POST")
	router.Handle("/refresh-token-auth",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RefreshToken),
		)).Methods("GET")
	router.Handle("/logout",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Logout),
		)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

/* Set up a global string for our secret */
var mySigningKey = []byte("secret")

/* Handlers */
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims
	   claims := token.Claims.(jwt.MapClaims)

	   /* Set token claims */
	claims["admin"] = true
	claims["name"] = "Ado Kukic"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	/* Finally, write the token to the browser window */
	w.Write([]byte(tokenString))
})
