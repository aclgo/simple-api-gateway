package load

import (
	"html/template"
	"net/http"
)

type Pages struct {
	tmpl    *template.Template
	pathCss string
}

func (p *Pages) Login(w http.ResponseWriter, r *http.Request) {
	if err := p.tmpl.ExecuteTemplate(w, "login.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func (p *Pages) Home(w http.ResponseWriter, r *http.Request) {
	if err := p.tmpl.ExecuteTemplate(w, "home.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func (p *Pages) Unauthorized(w http.ResponseWriter, r *http.Request) {
	if err := p.tmpl.ExecuteTemplate(w, "unauthorized.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func (p *Pages) ConfirmSignup(w http.ResponseWriter, r *http.Request) {
	if err := p.tmpl.ExecuteTemplate(w, "confirm_signup.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func (p *Pages) ResetPass(w http.ResponseWriter, r *http.Request) {
	if err := p.tmpl.ExecuteTemplate(w, "resetpass.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func (p *Pages) NewPass(w http.ResponseWriter, r *http.Request) {
	if err := p.tmpl.ExecuteTemplate(w, "newpass.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
