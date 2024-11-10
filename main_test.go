package autorefresh_test

import (
	"bytes"
	"html/template"
	"regexp"
	"strings"
	"testing"

	autorefresh "github.com/lavigneer/browser-autorefresh"
)

func TestTemplate(t *testing.T) {
	testTemplate := template.New("main")

	autorefresh.New(testTemplate, "__test_path__", 250)
	var b bytes.Buffer
	testTemplate.Execute(&b, nil)
	if !strings.Contains(b.String(), "new WebSocket(\"__test_path__\")") {
		t.Fatalf("Did not insert path correctly for the websocket. Rendered %s", b.String())
	}
	if !regexp.MustCompile("setTimeout.*250").MatchString(b.String()) {
		t.Fatalf("Did not insert timeout correctly for the websocket. Rendered %s", b.String())
	}
}
