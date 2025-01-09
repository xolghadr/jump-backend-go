package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func parseNumbers(r *http.Request) ([]int, error) {
	values := r.URL.Query()
	str := values.Get("numbers")

	if str == "" {
		return nil, errors.New("'numbers' paremeter missing")
	}

	numbers := strings.Split(str, ",")
	var result []int

	for _, w := range numbers {
		i, err := strconv.Atoi(w)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

func sum(nums []int) (int, error) {
	var result int
	for _, v := range nums {
		if v > 0 {
			if result > math.MaxInt64-v {
				return 0, errors.New("Overflow")
			}
		} else {
			if result < math.MaxInt64-v {
				return 0, errors.New("Overflow")
			}
		}
		result += v
	}
	return result, nil
}

func (s *Server) additionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	numbers, err := parseNumbers(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Result string `json:"result"`
			Error  string `json:"error"`
		}{
			"",
			err.Error(),
		})
		return
	}

	result, err := sum(numbers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Result string `json:"result"`
			Error  string `json:"error"`
		}{
			"",
			err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Result string `json:"result"`
		Error  string `json:"error"`
	}{
		fmt.Sprintf("The result of your query is: %d", result),
		"",
	})
	return
}

func sub(nums []int) (int, error) {
	var result int = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			if result < -math.MaxInt64+nums[i] {
				return 0, errors.New("Overflow")
			}
		} else {
			if result > -math.MaxInt64+nums[i] {
				return 0, errors.New("Overflow")
			}
		}

		result -= nums[i]
	}
	return result, nil
}

func (s *Server) subtractionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	numbers, err := parseNumbers(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Result string `json:"result"`
			Error  string `json:"error"`
		}{
			"",
			err.Error(),
		})
		return
	}

	result, err := sub(numbers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(struct {
			Result string `json:"result"`
			Error  string `json:"error"`
		}{
			"",
			err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Result string `json:"result"`
		Error  string `json:"error"`
	}{
		fmt.Sprintf("The result of your query is: %d", result),
		"",
	})
	return
}
