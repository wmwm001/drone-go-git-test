package main

import (
	"fmt"
	"log"
	"net/http"
)

// HelloWorld for hello world
func HelloWorld() string {
	return "Hello World, golang workshop"
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("http request....")
	_, _ = fmt.Fprintf(w, "I love %s!", r.URL.Path[1:])
}

func pinger(port string) error {
	resp, err := http.Get("http://localhost:" + port)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("server returned not-200 status code")
	}

	return nil
}

func main() {
	fmt.Println("hello World Drone Go")
}