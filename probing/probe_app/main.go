package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("vim-go")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s @ %v", "Hello Vimlesh @", time.Now())
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		w.Write([]byte("heakth Ok"))
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln("Couldn't start server")
	}

}
