package packages

import (
	"LinkaAja/src/models"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	@created at 25 May 2021
	Method for returning respose JSON
*/
func BasicResponse(w http.ResponseWriter, message string, code int, data interface{}) {
	res := models.BasicResponse{}

	res.Code = code
	res.Message = message
	res.Data = data

	response, err := json.Marshal(res)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "[SERVER] Failed Getting Response Data. Error Details : %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, string(response))
}
