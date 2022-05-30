package handlers

import (
	"app/data"
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"net/http"
)

func (h *Handlers) Form(w http.ResponseWriter, r *http.Request) {
	vars := make(jet.VarMap)
	validator := h.APP.Validator(nil)
	vars.Set("validator", validator)
	vars.Set("user", data.User{})

	err := h.APP.Render.Page(w, r, "form", vars, nil)
	if err != nil {
		h.APP.ErrorLog.Println(err)
	}
}

func (h *Handlers) PostForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.APP.ErrorLog.Println(err)
		return
	}

	validator := h.APP.Validator(nil)

	validator.Required(r, "first_name", "last_name", "email")
	validator.IsEmail("email", r.Form.Get("email"))

	validator.Check(len(r.Form.Get("first_name")) > 1, "first_name", "Must be at least two characters")
	validator.Check(len(r.Form.Get("last_name")) > 1, "last_name", "Must be at least two characters")

	if !validator.Valid() {
		vars := make(jet.VarMap)
		vars.Set("validator", validator)
		var user data.User
		user.FirstName = r.Form.Get("first_name")
		user.LastName = r.Form.Get("last_name")
		user.Email = r.Form.Get("email")
		vars.Set("user", user)

		if err := h.APP.Render.Page(w, r, "form", vars, nil); err != nil {
			h.APP.ErrorLog.Println(err)
			return
		}
		return
	}

	fmt.Fprint(w, "valid data")
}