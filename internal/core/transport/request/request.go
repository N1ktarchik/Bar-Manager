package request

import (
	"N1ktarchik/Bar-Manager/internal/core/errors"
	"encoding/json"

	"io"
	"net/http"
)

func DecodeAndValidate(r *http.Request, userData any) error {
	reqData, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.INTERNAL_SERVER_ERR()
	}

	if len(reqData) == 0 {
		return errors.BAD_REQUEST_ERR()
	}

	if err := json.Unmarshal(reqData, &userData); err != nil {
		return errors.INTERNAL_SERVER_ERR()
	}

	return nil
}
