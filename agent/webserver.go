package agent

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/areYouLazy/minemeld-agent/log"
	"github.com/gorilla/mux"
)

//WebServerInit starts API WebServer
func WebServerInit() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/check-ipv4/{address}", HandleCheckIPv4).Methods("GET")
	router.HandleFunc("/api/v1/check-ipv6/{address}", HandleCheckIPv6).Methods("GET")
	router.HandleFunc("/api/v1/check-fqdn/{address}", HandleCheckFqdn).Methods("GET")

	router.Use(loggingMiddleware)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + strconv.Itoa(*Opt.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Info("WebServe listening on %s", log.Bold(srv.Addr))
	go func() {
		srv.ListenAndServe()
	}()

	//Shutdown procedure from https://github.com/gorilla/mux
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info("Shutting down WebServer...")
	os.Exit(0)

}

//loggingMiddleware is a middleware for WebServer to log call info
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("%s %s %s %d", r.Method, r.RemoteAddr, r.URL, r.ContentLength)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

//HandleCheckIPv4 API to expose CheckIPv4 routine
func HandleCheckIPv4(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	address := vars["address"]

	res := CheckIPv4(address)

	var payload string

	if res == true {
		payload = fmt.Sprintf("Address %s is in list", log.Bold(address))
	} else {
		payload = fmt.Sprintf("Address %s is not in list", log.Bold(address))
	}

	log.Debug(payload)
	w.Write([]byte(payload))
}

//HandleCheckIPv6 API to expose CheckIPv6 routine
func HandleCheckIPv6(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)

	// address := vars["address"]

	// res := CheckIPv6(address)

	// var payload string
	// if res == true {
	// 	payload = fmt.Sprintf("Address %s is in list", log.Bold(address))
	// } else {
	// 	payload = fmt.Sprintf("Address %s is not in list", log.Bold(address))
	// }

	// log.Debug(payload)
	// w.Write([]byte(payload))

	w.Write([]byte("Not implemented"))
}

//HandleCheckFqdn API to expose CheckFqdn routine
func HandleCheckFqdn(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	address := vars["address"]

	res := CheckFQDN(address)

	var payload string
	if res == true {
		payload = fmt.Sprintf("Address %s is in list", log.Bold(address))
	} else {
		payload = fmt.Sprintf("Address %s is not in list", log.Bold(address))
	}

	log.Debug(payload)
	w.Write([]byte(payload))
}
