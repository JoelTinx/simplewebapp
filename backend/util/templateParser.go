package util

import (
  "html/template"
  "image"
  "image/color"
  "image/jpeg"
  "image/png"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "path/filepath"
  "strconv"
  "strings"
)

var templates map[string]*template.Template

func init()  {
  templates = parseTemplates()
}

func parseTemplates() map[string]*template.Template {
  result := make(map[string]*template.Template)
  layout, _ := template.ParseFiles("templates/layout.tmpl")
  dir, _ := os.Open("templates/index")
  defer dir.Close()

  fis, _ := dir.Readdir(-1)

  for _, fi := range fis {
    if fi.IsDir() {
      continue
    }
    f, _ := os.Open("templates/index/" + fi.Name())
    content, _ := ioutil.ReadAll(f)
    f.Close()
    tmpl, _ := layout.Clone()
    tmpl.Parse(string(content))
    result[fi.Name()] = tmpl
  }
  return result
}
