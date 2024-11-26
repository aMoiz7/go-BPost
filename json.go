package main


import (
	"encoding/json"
	"log"
	"net/http"
)

// ResponseJson sends a JSON response with the provided status code and payload

func ResponseWithError(w http.ResponseWriter , code int ,  msg string){
       if(code > 499 ){
		 log.Println("response with 5xx error: ",msg)
	   }
          
	   type errResponse struct{
		Error string `json:"error"`
	   }

	   ResponseJson(w ,code , errResponse{
		Error:msg,
	   })

}


func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {
	// Marshal the payload into JSON
	data, err := json.Marshal(payload)
	if err != nil {
		// Log the error if JSON marshaling fails
		log.Printf("failed to marshal JSON response: %v", payload)
		w.WriteHeader(http.StatusInternalServerError) // 500
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Add("Content-Type", "application/json")

	// Set the appropriate HTTP status code
	w.WriteHeader(code)

	// Write the JSON response
	w.Write(data)
}
