package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// With200 : Send Response with 200
// Content Type will be JSON
func With200(w http.ResponseWriter, data interface{}) {
	dataB, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataB)
}

// With200m : Send Response with 200
// Content Type will be text/plain
func With200m(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, `{"message": "%s"}`, message)
}

// With201 : use this function, if you have created a new document
// with some data
func With201(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	dataB, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(dataB)
}

// With201m : use this function, if you have created a new document
// with some string msg
func With201m(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, `{"message": "%s"}`, message)
}

// With500 : send response with 500 status, Internal Server Error
func With500(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(500)
	fmt.Fprintf(w, `{"message": "Internal Server Error"}`)
}

// With500m : send response with 500 status and message
func With500m(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(500)
	fmt.Fprintf(w, `{"message": "%s"}`, message)
}

// With500d : "send response with 500 status, Internal Server Error"
func With500d(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	dataB, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.WriteHeader(500)
	w.Write(dataB)
}

// With400 : send response with 400 status and message
func With400(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(400)
	fmt.Fprintf(w, `{"message": "%s"}`, message)
}

// With400ui : send response with 400 status and message
func With400ui(w http.ResponseWriter, message string, show bool) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(400)
	fmt.Fprintf(w, `{"message": "%s","show" :%v }`, message, show)
}

// With404m : send response with 404 status and message
func With404m(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(404)
	fmt.Fprintf(w, `{"message": "%s"}`, message)
}

// With404 : send response with 404 status and default message
func With404(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, `{"message": "%s"}`, "Resource not found")
}

// With403m : send response with 403 status and  message
func With403m(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintf(w, `{"message": "%s"}`, message)
}

// With401m : send response with 401 status and  message
func With401m(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, `{"message": "%s"}`, message)
}

// With409m : send response with 401 status and  message
func With409m(w http.ResponseWriter, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusConflict)
	fmt.Fprintf(w, `{"message": "%s"}`, message)
}
