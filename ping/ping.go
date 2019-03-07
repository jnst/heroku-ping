package ping

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Start starts ping continually.
func Start() error {
	url, ok := os.LookupEnv("HEROKU_URL")
	if !ok {
		return errors.New("missing HEROKU_URL")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		return errors.New("missing PORT")
	}

	go ListenAndServe(port)

	for {
		Ping(url)
		time.Sleep(10 * time.Minute)
	}
}

// Ping sends ping to heroku app.
func Ping(url string) {
	resp, err := http.Get(fmt.Sprintf("%sping", url))
	if err != nil {
		fmt.Printf("[ping] error: %v\n", err)
		return
	}
	defer resp.Body.Close()
}

// ListenAndServe listens and serve http request.
func ListenAndServe(port string) {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s %s\n", r.Method, r.URL.Path, r.Proto, r.Header.Get("User-Agent"))
		_, err := fmt.Fprintf(w, "ok")
		if err != nil {
			log.Printf("[ping] error: %v", err)
		}
	})

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("[ping] error: %v", err)
	}
}
