package Routes

import (
	"encoding/json"
	"github.com/MrFastDie/AlphaSign/AlphaProtocol"
	"github.com/MrFastDie/AlphaSign/Serial"
	"net/http"
)

func SpecialSign(w http.ResponseWriter, r *http.Request) {
	type jsonRequestType struct {
		SpecialSign string `json:"special-sign"`
		Position    string `json:"position"`
	}

	// TODO Json Parser https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
	var jsonRequest jsonRequestType

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&jsonRequest)

	specialSign, err := AlphaProtocol.GetSpecialSignByString(jsonRequest.SpecialSign)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"error\": \"special-sign is not valid\"}"))

		return
	}

	position, err := AlphaProtocol.GetPositionByString(jsonRequest.Position)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"error\": \"position is not valid\"}"))

		return
	}

	_, err = Serial.Port.Write(AlphaProtocol.SetSpecialSign(position, specialSign))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
