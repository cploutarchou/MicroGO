package microGo

import (
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Validation struct {
	Data  url.Values
	Error map[string]string
}

func (m *MicroGo) Validator(data url.Values) *Validation {
	return &Validation{
		Error: make(map[string]string),
		Data:  data,
	}

}

func (v *Validation) Valid() bool {
	return len(v.Error) == 0
}

func (v *Validation) AddError(key, message string) {
	if _, exists := v.Error[key]; !exists {
		v.Error[key] = message
	}
}

func (v *Validation) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}
func (v *Validation) Required(r *http.Request, fields ...string) {
	for _, field := range fields {
		value := r.Form.Get(field)
		if strings.TrimSpace(value) == "" {
			v.AddError(field, "Required field cannot be empty!")
		}
	}
}

func (v *Validation) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validation) IsEmail(field, value string) {
	if !govalidator.IsEmail(value) {
		v.AddError(field, "No valid email address provided")
	}
}

func (v *Validation) IsInt(field, value string) {
	_, err := strconv.Atoi(value)
	if err != nil {
		v.AddError(field, "integer value required")
	}
}
func (v *Validation) IsFloat(field, value string) {
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		v.AddError(field, "float number required")
	}
}
func (v *Validation) IsIsoDate(field, value string) {
	_, err := time.Parse("2022-05-18", value)
	if err != nil {
		v.AddError(field, "No valid iso date provided. Valid Format YYYY-MM-DD : ")
	}
}
func (v *Validation) NoSpaces(field, value string) {
	if govalidator.HasWhitespace(value) {
		v.AddError(field, "Spaces are not allowed")
	}
}