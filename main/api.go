package main

import "net/http"


func apiAppAdd(rw http.ResponseWriter, req *http.Request){
	req.ParseForm()
	req.Form
	
}

func apiAppDel(rw http.ResponseWriter, req *http.Request){
	
}

func apiAppUpdate(rw http.ResponseWriter, req *http.Request){
	
}

func apiAppList(rw http.ResponseWriter, req *http.Request){
	SuccResponse(rw, host2app)
}
