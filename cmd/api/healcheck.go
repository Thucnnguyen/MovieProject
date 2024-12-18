package main

import (
	"encoding/json"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":     "available",
		"Enviroment": app.config.env,
		"Version":    version,
	}

	res, err := json.Marshal(data)

	if err != nil {
		app.serverErrorResponse(w, r)
		return
	}

	res = append(res, '\n')

	w.Header().Set("content-type", "application/json")

	w.Write(res)
}
