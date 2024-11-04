[![Go Reference](https://pkg.go.dev/badge/github.com/lavigneer/browser-autorefresh.svg)](https://pkg.go.dev/github.com/lavigneer/browser-autorefresh)

# browser-autorefresh
This is a small Golang utility that can be installed as part of a web server that facilitates browser page auto-refresh when the service is restarted via some live-relaod function

There are two components this utility provides:

1. An http handler for a websocket endpoint
2. A template that provides a JS script to include in a page's html that uses the websocket endpoint to detect when the server is operational.

When the websocket connection is lost, a retry on the client-side occurs until it is able to reconnect, at which point it reloads the page to pull in the latest changes.

This is most useful when combined with a server live reload tool (e.g., (air)[https://github.com/air-verse/air]).

# Installation

```bash
go get github.com/lavigneer/browser-autorefresh
```

# Example

See [examples/std](examples/std)
