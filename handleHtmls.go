package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	//"io"
	//"os"
)

type params struct {
	Title  string
	Header string
	ImgS   []string
}

func handleHtmls() {
	//fmt.Println("Before /home..")
	http.HandleFunc("/home", homeHndl)
	//fmt.Println("After /home..")
}

func homeHndl(w http.ResponseWriter, r *http.Request) {
	Imgs := []string{"photos/error.png"}
	// Read images from directory
	nudeFiles, err := ioutil.ReadDir("./photos/bareb")
	if err != nil {
		fmt.Printf("Error while reading images from directory")
	}

	for _,f:= range nudeFiles {
		fmt.Println(f.Name())
		Imgs=append(Imgs,f.Name())
	}
	//End: Read	
	pars := params{"Home - ImageWeb", "Home", Imgs}
	//fmt.Println(pars.Title)
	//fmt.Println(pars.Header)
	//fmt.Println(pars.imgS[0])
	htmlTpl.ExecuteTemplate(w, "home.html", pars)
}
