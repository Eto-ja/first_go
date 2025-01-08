package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Eto-ja/rpn/pkg/rpn"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result"`
}

type Errorr struct {
	Error string `json:"error"`
}

func calculateHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := Errorr{
			Error: fmt.Sprintf("%v", "Internal server error"),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		// http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		err := Errorr{
			Error: fmt.Sprintf("%v", "Internal server error"),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		// http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	body, _ := io.ReadAll(r.Body)

	var req Request
	err := json.Unmarshal(body, &req)
	if err != nil {
		err := Errorr{
			Error: fmt.Sprintf("%v", "Internal server error"),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		// http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	res, err := rpn.Calc(req.Expression)
	if err != nil {
		err1 := Errorr{
			Error: fmt.Sprintf("%v", "Expression Is not Valid"),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(err1)
		// http.Error(w, "Expression Is not Valid", http.StatusUnprocessableEntity)
		return
	}
	resp := Response{
		Result: fmt.Sprintf("%v", res),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func main() {
	http.HandleFunc("/api/v1/calculate", calculateHandle)
	fmt.Println("Сервер запущен...")
	http.ListenAndServe("localhost:8181", nil)
}
