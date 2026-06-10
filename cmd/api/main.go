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
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeRoute)
	mux.HandleFunc("/courses", ourCourses)
	http.ListenAndServe(":8080", mux)
	
}

func homeRoute(w http.ResponseWriter, r *http.Request){
	var welcome ="welcome to the lang infrastructure"
	fmt.Fprintln(w, welcome)
}

func ourCourses(w http.ResponseWriter, r *http.Request){
  oc, err := w.Write([]byte("hey"))
  if err != nil{
	panic(err)
  }
  fmt.Println(oc)
}