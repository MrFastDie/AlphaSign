package Routes

import (
	"encoding/json"
	"github.com/MrFastDie/AlphaSign/AlphaProtocol"
	"github.com/MrFastDie/AlphaSign/Serial"
	"net/http"
	"strings"
	"unicode"
)

func WriteText(w http.ResponseWriter, r *http.Request) {
	type jsonRequestType struct {
		Message     string  `json:"message"`
		Position    *string `json:"position"`
		DefaultMode *string `json:"default-mode"`
		SpecialMode *string `json:"special-mode"`
	}

	// TODO Json Parser https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
	var jsonRequest jsonRequestType

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&jsonRequest)

	jsonRequest.Message = strings.Replace(jsonRequest.Message, "{RED}", (string)(AlphaProtocol.AddColor(AlphaProtocol.COLOR_RED)), -1)
	jsonRequest.Message = strings.Replace(jsonRequest.Message, "{GREEN}", (string)(AlphaProtocol.AddColor(AlphaProtocol.COLOR_GREEN)), -1)
	jsonRequest.Message = strings.Replace(jsonRequest.Message, "{YELLOW}", (string)(AlphaProtocol.AddColor(AlphaProtocol.COLOR_YELLOW)), -1)
	jsonRequest.Message = strings.Replace(jsonRequest.Message, "{RAINBOW1}", (string)(AlphaProtocol.AddColor(AlphaProtocol.COLOR_RAINNBOW_1)), -1)
	jsonRequest.Message = strings.Replace(jsonRequest.Message, "{RAINBOW2}", (string)(AlphaProtocol.AddColor(AlphaProtocol.COLOR_RAINNBOW_2)), -1)
	jsonRequest.Message = strings.Replace(jsonRequest.Message, "{COLORMIX}", (string)(AlphaProtocol.AddColor(AlphaProtocol.COLOR_COLOR_MIX)), -1)
	jsonRequest.Message = strings.Replace(jsonRequest.Message, "{DATE}", (string)(AlphaProtocol.PrepareDate()), -1)
	jsonRequest.Message = strings.Replace(jsonRequest.Message, "{TIME}", (string)(AlphaProtocol.PrepareTime()), -1)

	for i := 0; i < len(jsonRequest.Message); i++ {
		if jsonRequest.Message[i] > unicode.MaxASCII {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{\"error\": \"non-ascii character in message\"}"))

			return
		}
	}

	if nil == jsonRequest.Position && nil == jsonRequest.DefaultMode && nil == jsonRequest.SpecialMode {
		_, err = Serial.Port.Write(AlphaProtocol.SetText([]byte(jsonRequest.Message)))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	if nil == jsonRequest.Position || nil == jsonRequest.DefaultMode || nil == jsonRequest.SpecialMode {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"error\": \"position, default-mode and special-mode must be defined all together, or none\"}"))

		return
	}

	position, err := AlphaProtocol.GetPositionByString(*jsonRequest.Position)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"error\": \"position is not valid\"}"))

		return
	}

	defaultMode, err := AlphaProtocol.GetDefaultModeByString(*jsonRequest.DefaultMode)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"error\": \"default-mode is not valid\"}"))

		return
	}

	specialMode, err := AlphaProtocol.GetSpecialModeByString(*jsonRequest.SpecialMode)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"error\": \"special-mode is not valid\"}"))

		return
	}

	_, err = Serial.Port.Write(AlphaProtocol.SetSpecialText([]byte(jsonRequest.Message), position, defaultMode, specialMode))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
