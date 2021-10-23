package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"nicessm-api-service/response"
	"nicessm-api-service/services"
	"time"
)

// Adapter :
type Adapter func(http.Handler) http.Handler

// AllowCors :
func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("options called")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")

		if r.Method == "OPTIONS" {
			fmt.Println("options return")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}

//Log : ""
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		platform := r.URL.Query().Get("platform")
		log.Println("platform ==>", platform)
		next.ServeHTTP(w, r)
		duration := time.Since(t)
		log.Println("API ==>", r.RequestURI, " Time taken ===> ", duration.Minutes(), "m")
		fmt.Println()
		fmt.Println()
		fmt.Println()
	})
}

//JWT : ""
func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Authorization := r.Header.Get("Authorization")
		platform := r.Header.Get("platform")
		if Authorization == "" {
			response.With403mV2(w, "Token Not Provided", platform)
			return
		}
		auth, stat, err := services.ValidateTokenV2(Authorization)
		if err != nil {
			response.With500mV2(w, err.Error(), platform)
			return
		}
		if !stat {
			response.With403mV2(w, "UnAuthorized User", platform)
			return
		}
		c := context.Background()
		context.WithValue(c, "Authorization", auth)
		next.ServeHTTP(w, r)

	})
}
