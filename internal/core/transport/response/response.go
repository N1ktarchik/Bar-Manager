package response

import (
	"N1ktarchik/Bar-Manager/internal/core/errors"
	"encoding/json"
	"net/http"
)

type JWTResponse struct {
	Token string `json:"token"`
}

func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"code":500, "error":"INTERNAL_SERVER_ERROR"}`))
		return
	}

	w.WriteHeader(code)
	_, _ = w.Write(resp)
}

func RespondWithJWT(w http.ResponseWriter, code int, jwt string) {
	RespondWithJSON(w, code, JWTResponse{Token: jwt})
}

func RespondWithArray(w http.ResponseWriter, code int, key string, data any) {
	RespondWithJSON(w, code, map[string]any{
		key: data,
	})
}

func RespondWithError(w http.ResponseWriter, err error) {
	if appErr, ok := errors.IsErrorApp(err); ok {
		RespondWithJSON(w, appErr.Code, err)
		return
	}

	RespondWithJSON(w, http.StatusInternalServerError, errors.INTERNAL_SERVER_ERR())
}
