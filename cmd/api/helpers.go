package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// readJSON reads data into the data param(it assumes data is a reference type)
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	/* we're not gonna handle anything bigger than 1 MB. With this, we ensure we don't get some malicious user trying to pass us a massive
	request body.*/
	maxBytes := 1048576

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	/* As strange as this might seem, we're gonna decode it again! But this time, into an empty struct without any fields. Why?
	Because we're gonna assume that anytime we try to decode a JSON file, we're only gonna have one entry. We don't want
	multiple entries in there.*/
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only have a single JSON value")
	}

	return nil
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) error {
	var payload struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	payload.Error = true
	payload.Message = err.Error()

	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

	return nil
}
