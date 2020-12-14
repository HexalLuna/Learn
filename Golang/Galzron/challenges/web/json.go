package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type response1 struct {
    Page   int
    Produit []string
}

type response2 struct {
    Page   int      `json:"page"`
    Produit []string `json:"produit"`
}

func Json() {
	//json.Marshal() will be very util for this challenge
}