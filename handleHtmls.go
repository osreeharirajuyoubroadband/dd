package main

import (
	"net/http"
	//"fmt"
)

type params struct {
	Title  string
	Header string
	imgS   []string
}

func handleHtmls() {
	//fmt.Println("Before /home..")
	http.HandleFunc("/home", homeHndl)
	//fmt.Println("After /home..")
}

func homeHndl(w http.ResponseWriter, r *http.Request) {
	imgs := []string{"photos/home.png"}
	pars := params{"Home - ImageWeb", "Home", imgs}
	//fmt.Println(pars.Title)
	//fmt.Println(pars.Header)
	//fmt.Println(pars.imgS[0])
	htmlTpl.ExecuteTemplate(w, "home.html", pars)
}
