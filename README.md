# Go WebView2 Demo

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)

A tiny Go desktop app with HTML/CSS/JS UI using [`github.com/jchv/go-webview2`](https://github.com/jchv/go-webview2).

## What this demonstrates

- Native app window launched from Go
- Browser UI rendered in the embedded platform web engine
- Local HTTP server for hot-editable UI files in `web/`
- JS calling Go (`window.greet` and `window.getRuntime`) with returned values

## Platform support

This demo is Windows-only and uses Microsoft Edge WebView2.

## Run

```bash
go run .
```

The app starts a localhost HTTP server and loads `web/index.html` in WebView2.

## Windows prerequisites

- Microsoft Edge WebView2 Runtime installed

No cgo toolchain is required for this version.

## License

MIT. See [LICENSE](LICENSE).
