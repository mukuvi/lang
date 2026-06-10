package main

import (
	"fmt"
	"net/http"
)

type Course struct {
	Name string `json:"coursename"`
	Price int `json:"courseprice"`
	Tags []string `json:"tags"`

}
func main(){
	http.HandleFunc("/", homeRoute)
	http.ListenAndServe(":8080",nil)
}

func homeRoute(w http.ResponseWriter, r *http.Request){
	var welcome ="welcome to the lang infrastructure"
	fmt.Fprintln(w, welcome)
}