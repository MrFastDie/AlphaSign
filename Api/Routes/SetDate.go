package Routes

import (
	"encoding/json"
	"github.com/MrFastDie/AlphaSign/AlphaProtocol"
	"github.com/MrFastDie/AlphaSign/Serial"
	"net/http"
	"time"
)

func SetDate(w http.ResponseWriter, r *http.Request) {
	type jsonRequestType struct {
		Date string `json:"date"`
	}

	// TODO Json Parser https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
	var jsonRequest jsonRequestType

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&jsonRequest)

	dateTime, err := time.Parse("2006-01-02T15:04", jsonRequest.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"error\": \"date is not valid must be like '2022-04-23T00:09'\"}"))

		return
	}

	_, err = Serial.Port.Write(AlphaProtocol.SetDateTime(dateTime))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
