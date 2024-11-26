package main

import "net/http"


func errorHandler (w http.ResponseWriter , r *http.Request){
   ResponseWithError(w , 400 , "something went wrong")
}