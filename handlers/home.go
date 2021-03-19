package handlers

import (
	"net/http"

	"github.com/tonoy30/practice-rest/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	utils.ServeTemplate(nil, "index", w)
}
