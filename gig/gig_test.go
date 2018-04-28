package gig

import (
	"testing"
	"bytes"
	"os"
	"io"
	"fmt"
	"strings"
)

func Test_Run(t *testing.T)  {
	tests := []struct {
		args         []string
		containing   []string
		expectedCode int
	}{
		{[]string{"output gitignore", "Go", "-q"}, []string{"# Test binary, build with `go test -c`"}, 0},
		{[]string{"shows list", "-l"}, []string{"Go", "Rails","Kotlin"}, 0},
		{[]string{"shows version", "-v"}, []string{"gig version"}, 1},
		{[]string{"show usage", "-q"}, []string{"usage: "}, 1},
	}

	for _, te := range tests {
		stream := new(bytes.Buffer)
		cli := &Gig{OutStream: stream, ErrStream: stream, Output: []io.Writer{stream}}
		os.Args = te.args
		status := cli.Run()

		if status != te.expectedCode {
			t.Errorf("ExitStatus=%d, want %d", status, te.expectedCode)
		}

		for _, v := range te.containing {
			containing := fmt.Sprintf(v)
			if !strings.Contains(stream.String(), containing) {
				t.Errorf("[%s] actual: %s, want: %s", te.args[0], stream.String(), containing)
			}
		}

	}
}