package main

import (
    "net/http"
)

func response(res http.ResponseWriter,req *http.Request){
    res.Write([]byte("Hello world"))
}

func main(){
    http.HandleFunc("/",response)
    http.ListenAndServe(":2017",nil)
}