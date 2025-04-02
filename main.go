package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	var path = r.URL.Path;
	if path!= "/hello" {
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"Method is not supported",http.StatusNotFound)
	}
	fmt.Fprintf(w,"Hello\n")
}

func formHandler(w http.ResponseWriter,r *http.Request){
	if err:=r.ParseForm();err!=nil{
		fmt.Fprintf(w,"ParseForm() err: %v",err)
		return 
	}
	fmt.Fprintf(w,"Post Request Successful\n")

	 name:= r.FormValue("name")
	 address:= r.FormValue("address")

	fmt.Fprintf(w,"Name is %v and address is %v\n",name,address)
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"));

	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Server is running on port 8080\n")

	if err:= http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal(err)
	}
}