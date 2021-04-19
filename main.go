package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	log "github.com/sirupsen/logrus"
)

var (
	ready = false
	addr = flag.String("listen-address", ":8000", "The address to listen on for HTTP requests.")
	appVersion = "v0.0.1"
	instanceNum int
)

func main() {
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
	instanceNum = rand.Intn(1000)

	log.Info("Starting gitlab-k8s application...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		fmt.Fprintf(w, "Hello Golang In Gitlab CI!!\nI'm instance %d running version %s at %s\n", instanceNum, appVersion, t.Format("2006-01-02 15:04:05"))
		hostname, _ := os.Hostname()
		w.Write([]byte("Hostname: " + hostname + "\n"))
		w.Write([]byte("Version Info:\n"))
		w.Write([]byte(version.Print("app") + "\n"))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if ready {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500"))
		}
	})

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		<-time.After(5 * time.Second)
		ready = true
		log.Info("Application is ready!")
	}()

	//http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w,"%s\n",appVersion)
	//})

	log.Info("Listen on " + *addr)
	log.Fatal(http.ListenAndServe(*addr, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
