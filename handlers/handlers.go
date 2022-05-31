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
	res.Header().Set("Content-Type", "application/json")
	year := req.URL.Query().Get("year")
	if year == "" {
		winners, err := data.ListAllJSON()
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError) //500 code
			return
		}
		res.Write(winners)
	} else {
		filteredwinners, err := data.ListAllByYear(year)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest) // 400 code
			return
		}
		res.Write(filteredwinners)
	}
}

// AddNewWinner adds new winner to the list
func AddNewWinner(res http.ResponseWriter, req *http.Request) {
	accesstoken := req.Header.Get("X-ACCESS-TOKEN")
	istokenvalid := data.IsAccessTokenValid(accesstoken)
	if !istokenvalid {
		res.WriteHeader(http.StatusUnauthorized) //401
		return
	} else {
		err := data.AddNewWinner(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusUnprocessableEntity) //422
			return
		}
		res.WriteHeader(http.StatusCreated) // 201
	}
}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		ListWinners(res, req)
		break
	case http.MethodPost:
		AddNewWinner(res, req)
		break
	default:
		res.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}
