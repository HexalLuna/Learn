package main

import (
	"net/http"
	"encoding/json"
)

type Talk struct {
	name string
	length int
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	talk := Talk{name: "test format json", length: 45}
	payload, err := json.Marshal(talk)

	if err != nil {
		http.Error(writer, err.Error(), 500)
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.Write(payload)
	writer.Write([]byte("Ceci est un web-server en go"))
}

func webServer()  {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

/*
* Pour ce connecter sur le server entrez dans votre navigateur web http://localhost:8080
*
* Si vous êtes sur MacOS pour run votre fichier go sur Windows voici une commande a effectuer:
* 
* GOOS=windows GOARCH=386 go build
* 
* Ainsi votre fichier sera sous .exe et donc executable sur une autre os que celle sur laquelle vous travaillez.
*/