package errstack

import (
	"net/http"
	"fmt"
)


// HTTPError replies with the code proper to the error type and
// the error marshaled to JSON.
func HTTPError(err E, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if err.Inf() {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
	bb, jsonErr := err.MarshalJSON()
	if jsonErr != nil {
		return fmt.Errorf("Failed to marshall error to JSON: %s", jsonErr.Error())
	}
	w.Write(bb)
	return nil
}

