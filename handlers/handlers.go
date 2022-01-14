package handlers

import (
	"golang-fifa-world-cup-web-service/data"
	"net/http"
)

// RootHandler returns an empty body status code
func RootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)
}

// ListWinners returns winners from the list
func ListWinners(res http.ResponseWriter, req *http.Request) {
	year := req.URL.Query().Get("year")
	res.Header().Set("Content-Type", "application/json")
	if year == "" {
		content, err := data.ListAllJSON()
		if err != nil {
			// TODO
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Write(content)
	} else if year != "" {
		content, err := data.ListAllByYear(year)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		res.Write(content)
	}
}

// AddNewWinner adds new winner to the list
func AddNewWinner(res http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("X-ACCESS-TOKEN")
	isTokenValid := data.IsAccessTokenValid(accessToken)
	if !isTokenValid {
		res.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		res.WriteHeader(http.StatusCreated)
		err := data.AddNewWinner(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
	}

}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {

}
