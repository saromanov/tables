package tables

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// App is provides create of the new app
type App struct {
	writer *tabwriter.Writer
}

// New provides create of the new app
func New() *App {
	return &App{
		writer: tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0),
	}
}

// AddLine will write a new table line
func (t *App) AddLine(args ...interface{}) {
	formatString := t.buildFormatString(args)
	fmt.Fprintf(t.writer, formatString, args...)
}

// AddHeader will write a new table line followed by a seperator
func (t *App) AddHeader(args ...interface{}) {
	t.AddLine(args...)
	t.addSeperator(args)
}

// Build will write the table to the terminal
func (t *App) Build() {
	t.writer.Flush()
}

// addSeperator will write a new dash seperator line based on the args length
func (t *App) addSeperator(args []interface{}) {
	var b bytes.Buffer
	l := len(args)
	for idx, arg := range args {
		length := len(fmt.Sprintf("%v", arg))
		b.WriteString(strings.Repeat("-", length))
		if idx+1 != l {
			b.WriteString("\t")
		}
	}
	fmt.Fprintln(t.writer, b.String())
}

// buildFormatString will build up the formatting string used by the *tabwriter.Writer
func (t *App) buildFormatString(args []interface{}) string {
	var b bytes.Buffer
	l := len(args)
	for idx := range args {
		b.WriteString("%v")
		if idx+1 != l {
			b.WriteString("\t")
		}
	}
	b.WriteString("\n")
	return b.String()
}
