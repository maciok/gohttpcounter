package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const COUNTER_COOKIE = "pingPongCounter"

func handlePing(w http.ResponseWriter, r *http.Request) {
	i := cookieHandler(r, w, func(i int) int {
		return i + 1
	})
	fmt.Fprintf(w, "Pong!\nCountner: %d", i)
}

func handlePong(w http.ResponseWriter, r *http.Request) {
	i := cookieHandler(r, w, func(i int) int {
		return i - 1
	})
	fmt.Fprintf(w, "Ping!\nCounter: %d", i)
}

func cookieHandler(r *http.Request, w http.ResponseWriter, counter func(int) int) int {
	c, err := r.Cookie(COUNTER_COOKIE)
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:   COUNTER_COOKIE,
			Value:  "0",
			MaxAge: 3600,
		}
	}
	i, _ := strconv.Atoi(c.Value)
	i = counter(i)
	c.Value = strconv.FormatInt(int64(i), 10)
	http.SetCookie(w, c)
	return i
}

func main() {
	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/pong", handlePong)
	http.ListenAndServe(":8080", nil)
}
