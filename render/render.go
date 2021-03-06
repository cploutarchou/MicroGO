package render

import (
	"errors"
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews   *jet.Set
	Session    *scs.SessionManager
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float64
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
	Error           string
	Flash           string
}

func (r *Render) DefaultData(templateData *TemplateData, request *http.Request) *TemplateData {
	templateData.Secure = r.Secure
	templateData.ServerName = r.ServerName
	templateData.Port = r.Port
	templateData.CSRFToken = nosurf.Token(request)
	if r.Session.Exists(request.Context(), "userID") {
		templateData.IsAuthenticated = true
	}
	templateData.Error = r.Session.PopString(request.Context(), "error")
	templateData.Flash = r.Session.PopString(request.Context(), "flash")
	return templateData
}

// Page The page render function. You can use it to render pages using go or jet templates.
func (r *Render) Page(writer http.ResponseWriter, request *http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(r.Renderer) {
	case "go":
		return r.GoPage(writer, request, view, data)
	case "jet":
		return r.JetPage(writer, request, view, variables, data)
	default:
	}
	return errors.New("No rendering engine available. Please fill the required value (go or jet) in .env file ")
}

// GoPage The default go template engine renderer function.
func (r *Render) GoPage(writer http.ResponseWriter, request *http.Request, view string, data interface{}) error {
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", r.RootPath, view))
	if err != nil {
		return err
	}

	td := &TemplateData{}
	if data != nil {
		td = data.(*TemplateData)
	}
	td = r.DefaultData(td, request)
	err = tmpl.Execute(writer, &td)
	if err != nil {
		return err
	}
	return nil
}

//JetPage The jet engine template renderer function.
func (r *Render) JetPage(writer http.ResponseWriter, request *http.Request, view string, variables, data interface{}) error {
	var vars jet.VarMap
	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}
	td := &TemplateData{}

	if data != nil {
		td = data.(*TemplateData)
	}

	td = r.DefaultData(td, request)
	t, err := r.JetViews.GetTemplate(fmt.Sprintf("%s.jet", view))
	if err != nil {
		log.Println(err)
		return err
	}
	if err = t.Execute(writer, vars, td); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
