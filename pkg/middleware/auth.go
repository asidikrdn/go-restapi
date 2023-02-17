package middleware

import (
	"context"
	"encoding/json"
	"go-restapi-boilerplate/dto"
	jwtToken "go-restapi-boilerplate/pkg/jwt"
	"net/http"
)

func UserAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// mengambil token
		token := r.Header.Get("Authorization")
		if token == "" {
			response := dto.ErrorResult{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		// validasi token dan mengambil claims
		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			response := dto.ErrorResult{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		// set up context value and send it to next handler
		ctx := context.WithValue(r.Context(), "userData", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// mengambil token
		token := r.Header.Get("Authorization")
		if token == "" {
			response := dto.ErrorResult{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		// validasi token dan mengambil claims
		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			response := dto.ErrorResult{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		// validate is it admin
		if claims["role"].(string) != "superadmin" && claims["role"].(string) != "admin" {
			response := dto.ErrorResult{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized, you're not administrator",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		// set up context value and send it to next handler
		ctx := context.WithValue(r.Context(), "userData", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SuerAdminAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// mengambil token
		token := r.Header.Get("Authorization")
		if token == "" {
			response := dto.ErrorResult{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		// validasi token dan mengambil claims
		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			response := dto.ErrorResult{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		// validate is it superadmin
		if claims["role"].(string) != "superadmin" {
			response := dto.ErrorResult{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized, you're not Super Administrator",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		// set up context value and send it to next handler
		ctx := context.WithValue(r.Context(), "userData", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
