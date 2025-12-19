package render

import (
	"bytes"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/htet-29/home-pos/internal/config"
	"github.com/htet-29/home-pos/ui"
)

var functions = template.FuncMap{}

func Template(w http.ResponseWriter, r *http.Request, status int, page string, app *config.AppConfig) {
	templateCache, err := newTemplateCache()
	if err != nil {
		app.ServerError(w, r, err)
		os.Exit(1)
	}

	ts, ok := templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist.\n", page)
		app.ServerError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	err = ts.ExecuteTemplate(buf, "base", app.TemplateData)
	if err != nil {
		app.ServerError(w, r, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.gohtml")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.gohtml",
			"html/partials/*.gohtml",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
