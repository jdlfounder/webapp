package main
import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := "0.0.0.0:8087"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "%s: %s", hostname, time.Now().Format("2023-01-12 17:165:00"))
	})
	http.ListenAndServe(addr, nil)
}

