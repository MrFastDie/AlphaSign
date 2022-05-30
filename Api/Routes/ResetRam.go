package Routes

import (
	"github.com/MrFastDie/AlphaSign/AlphaProtocol"
	"github.com/MrFastDie/AlphaSign/Serial"
	"net/http"
)

func ResetRam(w http.ResponseWriter, r *http.Request) {
	_, err := Serial.Port.Write(AlphaProtocol.SetClearMemory())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
