package main

import (
	"empstore/empservice"
	"fmt"
	"net/http"
	"os"
)

func port() string {
	port := os.Getenv("EMP_WEBSERVER_PORT")
	if port == "" {
		return ":8080"
	}
	return ":" + port
}

func main() {
	fmt.Println("Emp Webserver running....")
	fmt.Printf("Port: %v", port())

	http.HandleFunc("/api/Add", empservice.AddHandler)
	http.HandleFunc("/api/Search", empservice.SearchHandler)
	http.HandleFunc("/api/List", empservice.ListHandler)
	http.HandleFunc("/api/Update", empservice.UpdateHandler)
	http.HandleFunc("/api/Delete", empservice.DeleteHandler)
	http.HandleFunc("/api/Restore", empservice.RestoreHandler)

	http.ListenAndServe(port(), nil)
}
