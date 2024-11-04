// Package autorefresh provides a mechanism for attaching browser refreshing to your templates during development.
// When it is detected that your program has restarted (e.g., by using a live-reload tool like "air"), it will
// trigger the browser page to refresh itself automatically.
package autorefresh

import (
	_ "embed"
	"errors"
	"html/template"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

//go:embed "reload.go.html"
var Script string

type PageReloader struct {
	Template    *template.Template
	Path        string
	RefreshRate uint
}

func New(t *template.Template, path string, refreshRate uint) (*PageReloader, error) {
	// If there was no template passed, create our own and let it get used in some other way
	if t == nil {
		t = template.New("autorefresh")
	}
	if refreshRate < 100 {
		return nil, errors.New("refreshRate must be at least 100ms")
	}
	t, err := t.Funcs(template.FuncMap{
		"path":        func() string { return path },
		"refreshRate": func() uint { return refreshRate },
	}).Parse(Script)
	if err != nil {
		return nil, err
	}
	return &PageReloader{Path: path, Template: t}, nil
}

func (p *PageReloader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)
	if err != nil {
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")
	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)
	for {
		socket.Ping(socketCtx)
		time.Sleep(time.Second * 2)
	}
}
