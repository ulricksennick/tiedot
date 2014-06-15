/* Prepare HTTP server for serving web control panel application. */
package webcp

import (
	"github.com/GeertJohan/go.rice"
	"github.com/HouzuoGuo/tiedot/tdlog"
	"html/template"
	"net/http"
)

var WebCp string

func RegisterWebCp() {
	if WebCp == "" || WebCp == "none" || WebCp == "no" || WebCp == "false" {
		tdlog.Noticef("Web control panel is disabled on your request")
		return
	}
	http.HandleFunc("/"+WebCp, handleWebCp)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.Write(rice.MustFindBox("static/img").MustBytes("favicon.ico"))
	})
	http.Handle("/"+WebCp+"/css/", http.StripPrefix("/"+WebCp+"/css/", http.FileServer(rice.MustFindBox("static/css").HTTPBox())))
	http.Handle("/"+WebCp+"/js/", http.StripPrefix("/"+WebCp+"/js/", http.FileServer(rice.MustFindBox("static/js").HTTPBox())))
	tdlog.Noticef("Web control panel is accessible at /%s", WebCp)
}

func handleWebCp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	templateBox, err := rice.FindBox("static/views")
	if err != nil {
		panic(err)
	}
	templatesString, err := templateBox.String("templates.html")
	if err != nil {
		panic(err)
	}
	viewString, err := templateBox.String("index.html")
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("index").Parse(viewString)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(w, map[string]interface{}{"root": WebCp, "templates": template.HTML(templatesString)}); err != nil {
		panic(err)
	}
}