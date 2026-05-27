//go:build windows

package main

import (
	"context"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"runtime"
	"time"

	webview "github.com/jchv/go-webview2"
)

type app struct {
	templates *template.Template
}

type indexData struct {
	Title string
}

func main() {
	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("Go WebView2 + HTMX Demo")
	w.SetSize(980, 700, webview.HintNone)

	server, baseURL, err := startStaticServer()
	if err != nil {
		panic(err)
	}
	defer shutdownServer(server)

	w.Navigate(baseURL)
	w.Run()
}

func startStaticServer() (*http.Server, string, error) {
	templates, err := template.ParseGlob("web/templates/*.gohtml")
	if err != nil {
		return nil, "", err
	}

	application := &app{templates: templates}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, "", err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", application.handleIndex)
	mux.HandleFunc("POST /greet", application.handleGreet)
	mux.HandleFunc("GET /runtime", application.handleRuntime)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	server := &http.Server{Handler: mux}

	go func() {
		_ = server.Serve(listener)
	}()

	return server, "http://" + listener.Addr().String(), nil
}

func shutdownServer(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)
}

func (a *app) handleIndex(w http.ResponseWriter, r *http.Request) {
	if err := a.templates.ExecuteTemplate(w, "index.gohtml", indexData{Title: "Go + HTMX + WebView2"}); err != nil {
		http.Error(w, "failed to render index template", http.StatusInternalServerError)
	}
}

func (a *app) handleGreet(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form payload", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		name = "friend"
	}

	msg := greet(name)
	if err := a.templates.ExecuteTemplate(w, "greet_fragment.gohtml", map[string]string{"Message": msg}); err != nil {
		http.Error(w, "failed to render greet fragment", http.StatusInternalServerError)
	}
}

func (a *app) handleRuntime(w http.ResponseWriter, r *http.Request) {
	if err := a.templates.ExecuteTemplate(w, "runtime_fragment.gohtml", map[string]string{"Runtime": getRuntime()}); err != nil {
		http.Error(w, "failed to render runtime fragment", http.StatusInternalServerError)
	}
}

func greet(name string) string {
	return fmt.Sprintf("Hello %s. Native Go says hi at %s", name, time.Now().Format(time.RFC1123))
}

func getRuntime() string {
	return fmt.Sprintf("Running on %s/%s", runtime.GOOS, runtime.GOARCH)
}
