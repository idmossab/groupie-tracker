package main

import(
	"fmt"
	"net/http"
)

func errorHandler(w http.ResponseWriter,status int){
	w.WriteHeader(status) // Sends the HTTP status code (e.g., 200 for success, 404 for not found) to the client before writing the response body
	switch status{
	case http.StatusNotFound:
			fmt.Fprint(w,"Page not found (404)")
	case http.StatusMethodNotAllowed:
			fmt.Fprint(w,"Methode not allowed (405)")
	case http.StatusBadRequest:
			fmt.Fprint(w,"Bad request (400)")
	case http.StatusInternalServerError:
			fmt.Fprint(w,"Internal server error  (500)")
	default:
		fmt.Fprintf(w,"Error %d",status)
	}
}

