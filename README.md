# Go WebView2 Demo

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)

A tiny Go desktop app with Go templates, HTMX, and a static Tailwind CSS file using [`github.com/jchv/go-webview2`](https://github.com/jchv/go-webview2).

## What this demonstrates

- Native app window launched from Go
- Browser UI rendered in the embedded platform web engine
- Local HTTP server rendering explicit Go templates from `web/templates/`
- HTMX requests calling Go HTTP handlers for dynamic partial updates
- Monochrome light theme with semantic template classes and compiled external Tailwind CSS in `web/static/app.css`

## Platform support

This demo is Windows-only and uses Microsoft Edge WebView2.

## Run

```bash
go run ./cmd/demo
```

The app starts a localhost HTTP server and loads the template-rendered page in WebView2.

## Windows prerequisites

- Microsoft Edge WebView2 Runtime installed

No cgo toolchain is required for this version.

## License

MIT. See [LICENSE](LICENSE).
