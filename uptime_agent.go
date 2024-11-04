package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	log.SetPrefix("uptime: ")
	if os.Getenv("CONN_STR") == "" {
		envErr := godotenv.Load()
		if envErr != nil {
			log.Fatal("Error loading .env file")
		}
	}

	var f func()
	var t *time.Timer

	f = func() {
		t = time.AfterFunc(time.Duration(5)*time.Second, f)
		url := os.Getenv("CONN_STR") + "/node"
		r, err := http.NewRequest("POST", url, nil)
		if err != nil {
			panic(err)
		}
		client := &http.Client{}
		res, err := client.Do(r)
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()
	}

	t = time.AfterFunc(time.Duration(5)*time.Second, f)

	defer t.Stop()

	for {
	}
}
