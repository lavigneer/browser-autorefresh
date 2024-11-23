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
	t.Parallel()
	testTemplate := template.New("main")

	_, err := autorefresh.New(testTemplate, "__test_path__", 250)
	if err != nil {
		t.Fatalf("Could not create template. %v", err)
	}
	var b bytes.Buffer
	err = testTemplate.Execute(&b, nil)
	if err != nil {
		t.Fatalf("Could not render template. %v", err)
	}
	if !strings.Contains(b.String(), "new WebSocket(\"__test_path__\")") {
		t.Fatalf("Did not insert path correctly for the websocket. Rendered %s", b.String())
	}
	if !regexp.MustCompile("setTimeout.*250").MatchString(b.String()) {
		t.Fatalf("Did not insert timeout correctly for the websocket. Rendered %s", b.String())
	}
}
