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
	http.HandleFunc("/nude",nudeHndl)
	http.HandleFunc("/beauty",beautyHndl)
}

func homeHndl(w http.ResponseWriter, r *http.Request) {
	Imgs := []string{"photos/error.png"}
	pars := params{"Home - ImageWeb", "Home", Imgs}
	htmlTpl.ExecuteTemplate(w, "home.html", pars)
}

func nudeHndl(w http.ResponseWriter, r *http.Request) {
	var Imgs []string
	nudeFiles, err := ioutil.ReadDir("./photos/bareb")
	if err != nil {
		fmt.Printf("Error while reading images from directory")
	}

	for _,f:= range nudeFiles {
		Imgs=append(Imgs,f.Name())
	}
	pars := params{"Nude - ImageWeb", "Nude girls", Imgs}
	htmlTpl.ExecuteTemplate(w, "nudeImgs.html", pars)
}

func beautyHndl(w http.ResponseWriter, r *http.Request) {
	var Imgs []string
	nudeFiles, err := ioutil.ReadDir("./photos/beauty")
	if err != nil {
		fmt.Printf("Error while reading images from directory")
	}

	for _,f:= range nudeFiles {
		Imgs=append(Imgs,f.Name())
	}
	pars := params{"Beauty - ImageWeb", "Beautiful girls", Imgs}
	htmlTpl.ExecuteTemplate(w, "beautyImgs.html", pars)
}
