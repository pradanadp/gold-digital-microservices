package helper

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, i interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(i); err != nil {
		return errors.New("can't decode from request body. " + err.Error())
	}

	return nil
}

func WriteToResponseBody(w http.ResponseWriter, i interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(i); err != nil {
		return errors.New("can't encode to response body. " + err.Error())
	}

	return nil
}

func ReadFromResponseBody(r *http.Response, i interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(i); err != nil {
		err := errors.New("can't decode form response body. " + err.Error())
		log.Println(err.Error())
		return err
	}

	return nil
}
