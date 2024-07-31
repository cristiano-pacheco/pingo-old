// Package response contains sintax sugar functions to handle the JSON response.
package response

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/cristiano-pacheco/pingo/internal/infra/http/request"
	"github.com/cristiano-pacheco/pingo/internal/infra/validator"
)

type Envelope map[string]any

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the request resouce could not be found"
	ErrorResponse(w, r, http.StatusNotFound, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	LogError(r, err)
	message := "the server encountered a problem and could not process your request"
	ErrorResponse(w, r, http.StatusInternalServerError, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

func ValidationFailedResponse(w http.ResponseWriter, r *http.Request, vr *validator.ValidationResult) {
	if vr == nil {
		ServerErrorResponse(w, r, nil)
		return
	}

	envelope := Envelope{"data": vr}

	err := JSONResponse(w, http.StatusUnprocessableEntity, envelope, nil)
	if err != nil {
		LogError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func InvalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	ErrorResponse(w, r, http.StatusUnprocessableEntity, message)
}

func UnauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	ErrorResponse(w, r, http.StatusUnauthorized, message)
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	envelope := Envelope{"error": message}

	err := JSONResponse(w, status, envelope, nil)
	if err != nil {
		LogError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func JSONResponse(w http.ResponseWriter, status int, data Envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func EmptyOKResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func NoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func LogError(r *http.Request, err error) {
	method := r.Method
	uri := r.URL.RequestURI()
	logger := r.Context().Value(request.LoggerContextKey).(*slog.Logger)

	logger.Error(err.Error(), "method", method, "uri", uri)
}
