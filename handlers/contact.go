package handlers

import (
	"log"
	"net/http"

	"github.com/tonoy30/practice-rest/models"
	"github.com/tonoy30/practice-rest/utils"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ServeTemplate(nil, "contact", w)
		return
	}
	contact := models.Contact{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	log.Printf("Form Value is: %v", contact)
	utils.ServeTemplate(struct{ Success bool }{true}, "contact", w)
}
