package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	//"io"
	// "os"
	"bufio"
	"strings"
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
	http.HandleFunc("/nude",nudeDirtHndl)
	http.HandleFunc("/beauty",beautyHndl)
	http.HandleFunc("/nudedirt",nudeDirtHndl)
}

func homeHndl(w http.ResponseWriter, r *http.Request) {
	Imgs := []string{"photos/error.png"}
	pars := params{"Home - ImageWeb", "Home", Imgs}
	htmlTpl.ExecuteTemplate(w, "home.html", pars)
}

func beautyHndl(w http.ResponseWriter, r *http.Request) {
	var Imgs []string
	nudeFiles, err := ioutil.ReadFile("beautyUrls.txt")
	if err != nil {
		fmt.Printf("Error while reading file beautyUrls.txt")
	}
	nudeStr:=string(nudeFiles)
	sr := strings.NewReader(nudeStr)
	br := bufio.NewReader(sr)
	line, isPrefix, err := br.ReadLine()
	fmt.Println(isPrefix)
	for err == nil {
			Imgs=append(Imgs,string(line))
			line, isPrefix, err = br.ReadLine()
	}
	pars := params{"Beauty - ImageWeb", "Beauty girls", Imgs}
	htmlTpl.ExecuteTemplate(w, "beautyImgs.html", pars)
}
//*******************************************************************
func nudeDirtHndl(w http.ResponseWriter, r *http.Request) {
	var Imgs []string
	nudeFiles, err := ioutil.ReadFile("imgsUrls.txt")
	if err != nil {
		fmt.Printf("Error while reading file imgsUrls.txt")
	}
	nudeStr:=string(nudeFiles)
	sr := strings.NewReader(nudeStr)
	br := bufio.NewReader(sr)
	line, isPrefix, err := br.ReadLine()
	fmt.Println(isPrefix)
	for err == nil {
			Imgs=append(Imgs,string(line))
			line, isPrefix, err = br.ReadLine()
	}
	pars := params{"Nude1 - ImageWeb", "Nude girls", Imgs}
	htmlTpl.ExecuteTemplate(w, "nudeImgsDirt.html", pars)
}
