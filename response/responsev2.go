package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"nicessm-api-service/constants"
)

//ResponseSruct :""
type ResponseSruct struct {
	Response Response `json:"response,omitempty"`
}

//Response : ""
type Response struct {
	StatusCode int                    `json:"statusCode,omitempty"`
	Message    string                 `json:"message,omitempty"`
	Data       map[string]interface{} `json:"data"`
	Error      *ErrorStruct           `json:"-"`
}

//ErrorStruct : ""
type ErrorStruct struct {
	ErrorCode        int                    `json:"errorCode,omitempty"`
	ErrorDescription string                 `json:"errorDescription,omitempty"`
	ErrorLabel       string                 `json:"errorLabel,omitempty"`
	Data             map[string]interface{} `json:"data,omitempty"`
}

// With200V2 : Send Response with 200
// Content Type will be JSON
func With200V2(w http.ResponseWriter, msg string, data map[string]interface{}, platform string) {
	response := new(ResponseSruct)
	response.Response.StatusCode = 200
	response.Response.Message = msg
	response.Response.Data = data
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dataB)
}

// With201dV2 : use this function, if you have created a new document
// with some data
func With201dV2(w http.ResponseWriter, msg string, data map[string]interface{}, platform string) {
	response := new(ResponseSruct)
	response.Response.StatusCode = 201
	response.Response.Message = msg
	response.Response.Data = data
	w.Header().Add("Content-Type", "application/json")
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(dataB)
}

// With201mV2 : use this function, if you have created a new document
// with some string msg
func With201mV2(w http.ResponseWriter, msg string, platform string) {
	response := new(ResponseSruct)
	response.Response.StatusCode = 201
	response.Response.Message = msg
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(dataB)
}

// With500V2 : send response with 500 status, Internal Server Error
func With500V2(w http.ResponseWriter, platform string) {
	response := new(ResponseSruct)
	response.Response.StatusCode = 500
	response.Response.Message = "Internal Server Error"
	response.Response.Error = new(ErrorStruct)
	response.Response.Error.ErrorCode = 500
	response.Response.Error.ErrorDescription = "Internal Server Error"
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if platform == constants.PLATFORMMOBILE {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
	w.Write(dataB)
}

// With500mV2 : send response with 500 status and message
func With500mV2(w http.ResponseWriter, message string, platform string) {
	response := new(ResponseSruct)
	response.Response.StatusCode = 500
	response.Response.Message = message
	response.Response.Error = new(ErrorStruct)
	response.Response.Error.ErrorCode = 500
	response.Response.Error.ErrorDescription = message
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if platform == constants.PLATFORMMOBILE {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
	w.Write(dataB)
}

// With500dV2 : send response with 500 status, Internal Server Error
func With500dV2(w http.ResponseWriter, msg string, data map[string]interface{}, platform string) {
	response := new(ResponseSruct)
	response.Response.StatusCode = 500
	response.Response.Message = msg
	response.Response.Error = new(ErrorStruct)
	response.Response.Error.ErrorCode = 500
	response.Response.Error.ErrorDescription = msg
	response.Response.Error.Data = data
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if platform == constants.PLATFORMMOBILE {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
	w.Write(dataB)

}

// With400V2 : send response with 400 status and message
func With400V2(w http.ResponseWriter, message string, platform string) {
	response := new(ResponseSruct)
	response.Response.StatusCode = 400
	response.Response.Message = message
	response.Response.Error = new(ErrorStruct)
	response.Response.Error.ErrorCode = 400
	response.Response.Error.ErrorDescription = message
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if platform == constants.PLATFORMMOBILE {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
	}
	w.Write(dataB)
}

// With404mV2 : send response with 404 status and message
func With404mV2(w http.ResponseWriter, message string, platform string) {
	response := new(ResponseSruct)
	response.Response.StatusCode = 404
	response.Response.Message = message
	response.Response.Error = new(ErrorStruct)
	response.Response.Error.ErrorCode = 404
	response.Response.Error.ErrorDescription = message
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if platform == constants.PLATFORMMOBILE {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(404)
	}
	w.Write(dataB)
}

// With404V2 : send response with 404 status and default message
func With404V2(w http.ResponseWriter, platform string) {
	response := new(ResponseSruct)
	response.Response.StatusCode = 404
	response.Response.Message = "Resource not found"
	response.Response.Error = new(ErrorStruct)
	response.Response.Error.ErrorCode = 404
	response.Response.Error.ErrorDescription = "Resource not found"
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if platform == constants.PLATFORMMOBILE {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(404)
	}
	w.Write(dataB)
}

// With403mV2 : send response with 403 status and  message
func With403mV2(w http.ResponseWriter, message string, platform string) {
	response := new(ResponseSruct)
	response.Response.StatusCode = http.StatusForbidden
	response.Response.Message = message
	response.Response.Error = new(ErrorStruct)
	response.Response.Error.ErrorCode = http.StatusForbidden
	response.Response.Error.ErrorDescription = message
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	log.Println("platform==>", platform)
	if platform == constants.PLATFORMMOBILE {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
	w.Write(dataB)
}

// With401mV2 : send response with 401 status and  message
func With401mV2(w http.ResponseWriter, message string, platform string) {
	response := new(ResponseSruct)
	response.Response.StatusCode = http.StatusUnauthorized
	response.Response.Message = message
	response.Response.Error = new(ErrorStruct)
	response.Response.Error.ErrorCode = http.StatusUnauthorized
	response.Response.Error.ErrorDescription = message
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if platform == constants.PLATFORMMOBILE {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
	w.Write(dataB)
}

// With409mV2 : send response with 409 status and  message
func With409mV2(w http.ResponseWriter, message string, platform string) {
	response := new(ResponseSruct)
	response.Response.StatusCode = http.StatusConflict
	response.Response.Message = message
	response.Response.Error = new(ErrorStruct)
	response.Response.Error.ErrorCode = http.StatusConflict
	response.Response.Error.ErrorDescription = message
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if platform == constants.PLATFORMMOBILE {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusConflict)
	}
	w.Write(dataB)
}

// With412mV2 : send response with 412 status and  message
func With412mV2(w http.ResponseWriter, message string, platform string) {
	response := new(ResponseSruct)
	response.Response.StatusCode = http.StatusPreconditionFailed
	response.Response.Message = message
	response.Response.Error = new(ErrorStruct)
	response.Response.Error.ErrorCode = http.StatusPreconditionFailed
	response.Response.Error.ErrorDescription = message
	dataB, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid Data")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if platform == constants.PLATFORMMOBILE {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusPreconditionFailed)
	}
	w.Write(dataB)
}
