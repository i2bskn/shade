package shade

import (
	"bytes"
	"html/template"
	"path/filepath"
	"testing"
)

const (
	testBaseDir  = "testdata"
	testLayout   = "layout.html"
	testTemplate = "index.html"
)

func TestDefault__baseDir(t *testing.T) {
	layout := Default()
	if layout.baseDir != defaultBaseDir {
		t.Fatalf("Expected %v, but %v", defaultBaseDir, layout.baseDir)
	}
}

func TestDefault__layout(t *testing.T) {
	layout := Default()
	if layout.layout != defaultLayout {
		t.Fatalf("Expected %v, but %v", defaultLayout, layout.layout)
	}
}

func TestNewLayout__baseDir(t *testing.T) {
	layout := NewLayout(testBaseDir, defaultLayout)
	if layout.baseDir != testBaseDir {
		t.Fatalf("Expected %v, but %v", testBaseDir, layout.baseDir)
	}
}

func TestNewLayout__layout(t *testing.T) {
	layout := NewLayout(defaultBaseDir, testLayout)
	if layout.layout != testLayout {
		t.Fatalf("Expected %v, but %v", testLayout, layout.layout)
	}
}

func TestTemplatePaths__pathsSize(t *testing.T) {
	layout := Default()
	pathsSize := len(layout.templatePaths(testTemplate))
	expected := 2
	if pathsSize != expected {
		t.Fatalf("Expected %v, but %v", expected, pathsSize)
	}
}

func TestTemplatePaths__layoutPath(t *testing.T) {
	layout := Default()
	paths := layout.templatePaths(testTemplate)
	expected := filepath.Join(layout.baseDir, layout.layout)
	if paths[0] != expected {
		t.Fatalf("Expected %v, but %v", expected, paths[0])
	}
}

func TestTemplatePaths__templatePath(t *testing.T) {
	layout := Default()
	paths := layout.templatePaths(testTemplate)
	expected := filepath.Join(layout.baseDir, testTemplate)
	if paths[1] != expected {
		t.Fatalf("Expected %v, but %v", expected, paths[1])
	}
}

func TestGetTemplate__notExist(t *testing.T) {
	layout := Default()
	if _, ok := layout.getTemplate(testTemplate); ok {
		t.Fatalf("Expected %v, but %v", false, true)
	}
}

func TestGetTemplate__exist(t *testing.T) {
	layout := NewLayout(testBaseDir, testLayout)
	layout.loadTemplate(testTemplate)
	if _, ok := layout.getTemplate(testTemplate); ok != true {
		t.Fatalf("Expected %v, but %v", true, false)
	}
}

func TestLoadTemplate(t *testing.T) {
	layout := NewLayout(testBaseDir, testLayout)
	layout.loadTemplate(testTemplate)
	_, ok := layout.templates[testTemplate]
	if ok != true {
		t.Fatalf("Expected %v, but %v", true, false)
	}
}

func TestTemplate__cached(t *testing.T) {
	layout := NewLayout(testBaseDir, testLayout)
	layout.loadTemplate(testTemplate)
	_, err := layout.Template(testTemplate)
	if err != nil {
		t.Fatalf("Expected not error, but %v", err)
	}
}

func TestTemplate__notCached(t *testing.T) {
	layout := NewLayout(testBaseDir, testLayout)
	_, err := layout.Template(testTemplate)
	if err != nil {
		t.Fatalf("Expected not error, but %v", err)
	}
}

func TestRender__exist(t *testing.T) {
	layout := NewLayout(testBaseDir, testLayout)
	paths := layout.templatePaths(testTemplate)
	tpl, _ := template.ParseFiles(paths...)
	buf := new(bytes.Buffer)
	tpl.Execute(buf, nil)
	expected := buf.String()
	body, err := layout.Render(testTemplate, nil)
	if err != nil {
		t.Fatalf("Expected not error, but %v", err)
	}

	if body != expected {
		t.Fatalf("Expected %v, but %v", expected, body)
	}
}

func TestRender__notExist(t *testing.T) {
	layout := NewLayout(testBaseDir, testLayout)
	body, err := layout.Render("notexist.html", nil)
	if err == nil {
		t.Fatalf("Expected %v, but not error", err)
	}

	if body != "" {
		t.Fatalf("Expected empty, but %v", body)
	}
}
