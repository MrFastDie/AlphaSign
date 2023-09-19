package main

import (
	"github.com/MrFastDie/AlphaSign/Api"
	_ "github.com/MrFastDie/AlphaSign/Serial"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", Api.GetRouter())
}
