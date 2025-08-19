package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
	port string
}
type ResponseMessage struct {
	Result string `json:"result"`
	Error  string `json:"error"`
}

func safeAdd64(a, b int64) (int64, error) {
	if (b > 0 && a > math.MaxInt64-b) || (b < 0 && a < math.MinInt64-b) {
		return 0, fmt.Errorf("overflow on addition")
	}
	return a + b, nil
}

func safeSub64(a, b int64) (int64, error) {
	if (b > 0 && a < math.MinInt64+b) || (b < 0 && a > math.MaxInt64+b) {
		return 0, fmt.Errorf("overflow on subtraction")
	}
	return a - b, nil
}

func (s *Server) add(w http.ResponseWriter, r *http.Request) {
	numsParam := r.URL.Query().Get("numbers")
	if numsParam == "" {
		response := ResponseMessage{
			Result: "",
			Error:  "'numbers' parameter missing",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	nums := strings.Split(numsParam, ",")

	first, err := strconv.ParseInt(strings.TrimSpace(nums[0]), 10, 64)
	if err != nil {
		response := ResponseMessage{
			Result: "",
			Error:  "invalid number: " + nums[0] + " " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	result := first

	for _, n := range nums[1:] {
		val, err := strconv.ParseInt(strings.TrimSpace(n), 10, 64)
		if err != nil {
			response := ResponseMessage{
				Result: "",
				Error:  "invalid number: " + n + " " + err.Error(),
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		result, err = safeAdd64(result, val)
		if err != nil {
			response := ResponseMessage{
				Result: "",
				Error:  "Overflow",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	response := ResponseMessage{
		Result: fmt.Sprintf("The result of your query is: %d", result),
		Error:  "",
	}
	json.NewEncoder(w).Encode(response)
}

func (s *Server) sub(w http.ResponseWriter, r *http.Request) {
	numsParam := r.URL.Query().Get("numbers")
	if numsParam == "" {
		response := ResponseMessage{
			Result: "",
			Error:  "missing numbers",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	parts := strings.Split(numsParam, ",")

	first, err := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
	if err != nil {
		response := ResponseMessage{
			Result: "",
			Error:  "invalid number: " + parts[0] + " " + err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	res := first

	for _, p := range parts[1:] {
		n, err := strconv.ParseInt(strings.TrimSpace(p), 10, 64)
		if err != nil {
			response := ResponseMessage{
				Result: "",
				Error:  "invalid number: " + p + " " + err.Error(),
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		res, err = safeSub64(res, n)
		if err != nil {
			response := ResponseMessage{
				Result: "",
				Error:  "Overflow",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	response := ResponseMessage{
		Result: fmt.Sprintf("The result of your query is: %d", res),
		Error:  "",
	}
	json.NewEncoder(w).Encode(response)
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Start() {
	http.HandleFunc("/add", s.add)
	http.HandleFunc("/sub", s.sub)
	http.ListenAndServe(":"+s.port, nil)
}

func main() {
	server := NewServer("8000")
	server.Start()

}
