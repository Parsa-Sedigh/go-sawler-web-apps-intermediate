package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// templateData contains all kinds of things we think we might at some point need to pass to a template
type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
	CSSVersion      string
}

// at some point we want to pass functions to our templates
var functions = template.FuncMap{}

//go:embed templates
var templateFS embed.FS

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	td.API = app.config.api

	return td
}

/*
	renderTemplate handles multiple situations where:

- we have a very simple template that has no partials
- we have a template that has partials
- we have a template that's being passed no data
- we have a template that's being passed some data
*/
func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error

	/* We want to be able to pass a name like "terminal" and not sth like "terminal.page.gohtml" . So we need to build the name that that
	template will have:*/
	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)

	// if templateInMap is true, we can use it but if templateInMap is false, we have to build it
	_, templateInMap := app.templateCache[templateToRender]

	/* When we're in development, we never want to use the template cache. We don't want to be always stopping and starting our
	application every time we make a change to the template files, instead we want it to happen automatically.  */
	if app.config.env == "production" && templateInMap {
		t = app.templateCache[templateToRender]
	} else { // if we're not in production or it doesn't exist in that map, we need to build it
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	if td == nil {
		td = &templateData{}
	}

	td = app.addDefaultData(td, r)

	err = t.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

func (app *application) parseTemplate(partials []string, page, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	// build partials
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.gohtml", x)
		}
	}

	if len(partials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.gohtml", strings.Join(partials, ","), templateToRender)
	} else {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.gohtml", templateToRender)
	}

	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	app.templateCache[templateToRender] = t

	return t, nil
}
