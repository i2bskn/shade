package shade

import (
	"bytes"
	"html/template"
	"path/filepath"
	"sync"
)

const (
	defaultBaseDir = "templates"
	defaultLayout  = "layout.html"
)

// Default returns a Layout instance with default settings.
func Default() *Layout {
	return NewLayout(defaultBaseDir, defaultLayout)
}

// Layout is html/template wrapper with single layout file.
type Layout struct {
	baseDir   string
	layout    string
	templates map[string]*template.Template
	mutex     *sync.RWMutex
}

// NewLayout returns a Layout instance with custom settings.
func NewLayout(baseDir, layout string) *Layout {
	return &Layout{
		baseDir:   baseDir,
		layout:    layout,
		templates: make(map[string]*template.Template),
		mutex:     new(sync.RWMutex),
	}
}

func (l *Layout) templatePaths(name string) []string {
	return []string{
		filepath.Join(l.baseDir, l.layout),
		filepath.Join(l.baseDir, name),
	}
}

func (l *Layout) getTemplate(name string) (*template.Template, bool) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	t, ok := l.templates[name]
	return t, ok
}

func (l *Layout) loadTemplate(name string) (*template.Template, error) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	paths := l.templatePaths(name)
	t, err := template.ParseFiles(paths...)
	if err == nil {
		l.templates[name] = t
	}
	return t, err
}

// Template returns a template.Template instance with template cache.
func (l *Layout) Template(name string) (*template.Template, error) {
	if t, ok := l.getTemplate(name); ok {
		return t, nil
	}

	return l.loadTemplate(name)
}

// Render returns a string to render template.
func (l *Layout) Render(name string, data interface{}) (string, error) {
	t, err := l.Template(name)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	return buf.String(), err
}
