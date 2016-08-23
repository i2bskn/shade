package shade

import (
	"html/template"
	"strings"
	"sync"
)

const nameDelimiter = ";"

var (
	templates = make(map[string]*template.Template)
	mutex     = new(sync.RWMutex)
)

func getTemplate(paths ...string) (*template.Template, bool) {
	mutex.RLock()
	defer mutex.RUnlock()
	t, ok := templates[strings.Join(paths, nameDelimiter)]
	return t, ok
}

func loadTemplate(paths ...string) (*template.Template, error) {
	mutex.Lock()
	defer mutex.Unlock()

	t, err := template.ParseFiles(paths...)
	if err == nil {
		templates[strings.Join(paths, nameDelimiter)] = t
	}
	return t, err
}

func Template(paths ...string) (*template.Template, error) {
	if t, ok := getTemplate(paths...); ok {
		return t, nil
	}

	return loadTemplate(paths...)
}
