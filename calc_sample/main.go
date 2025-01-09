package main

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{port: port}
}

func (s *Server) Start() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/sub", subHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%s", s.port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	strNumbers := r.URL.Query().Get("numbers")
	sum := 0
	for _, strNumber := range strings.Split(strNumbers, ",") {
		number, err := strconv.Atoi(strNumber)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf(`"result": "", "error": "%s"`, err.Error())))
		}
		sum, err = add(sum, number)
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf(`{"result": "", "error": "%s"}`, err.Error())))
		}
	}
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(fmt.Sprintf(`{"result": "The result of your query is: %d", "error": ""}`, sum)))
}

func subHandler(w http.ResponseWriter, r *http.Request) {
	strNumbers := r.URL.Query().Get("numbers")
	sum := 0
	var err error
	for i, strNumber := range strings.Split(strNumbers, ",") {
		number, err2 := strconv.Atoi(strNumber)
		if err2 != nil {
			err = errors.New(err2.Error())
			break
		}
		if i == 0 {
			sum = number
		} else {
			sum, err2 = subtract(sum, number)
		}
		if err2 != nil {
			err = errors.New(err2.Error())
			break
		}
	}
	w.Header().Set("Content-Type", "text/plain")
	var response string
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = fmt.Sprintf(`{"result": "", "error": "%s"}`, err.Error())
	} else {
		response = fmt.Sprintf(`{"result": "The result of your query is: %d", "error": ""}`, sum)
	}
	_, _ = w.Write([]byte(response))
}

func add(a, b int) (int, error) {
	if b >= 0 && math.MaxInt-b < a {
		return 0, errors.New("Overflow")
	}
	if b < 0 && math.MinInt-b > a {
		return 0, errors.New("Overflow")
	}
	return a + b, nil
}

func subtract(a, b int) (int, error) {
	if b >= 0 && math.MinInt+b > a {
		return 0, errors.New("Overflow")
	}
	if b < 0 && math.MaxInt+b < a {
		return 0, errors.New("Overflow")
	}
	return a - b, nil
}
