//go:build windows

package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"runtime"
	"time"

	webview "github.com/jchv/go-webview2"
)

func main() {
	w := webview.New(true)
	defer w.Destroy()

	w.SetTitle("Go WebView2 Demo")
	w.SetSize(980, 700, webview.HintNone)

	if err := w.Bind("greet", greet); err != nil {
		panic(err)
	}

	if err := w.Bind("getRuntime", getRuntime); err != nil {
		panic(err)
	}

	server, baseURL, err := startStaticServer()
	if err != nil {
		panic(err)
	}
	defer shutdownServer(server)

	w.Navigate(baseURL)
	w.Run()
}

func startStaticServer() (*http.Server, string, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, "", err
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("web")))

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

func greet(name string) string {
	return fmt.Sprintf("Hello %s. Native Go says hi at %s", name, time.Now().Format(time.RFC1123))
}

func getRuntime() string {
	return fmt.Sprintf("Running on %s/%s", runtime.GOOS, runtime.GOARCH)
}
