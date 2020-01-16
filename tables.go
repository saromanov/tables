package tables

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type Hook func(string) string

// App is provides create of the new app
type App struct {
	writer    *tabwriter.Writer
	lines     string
	lineHooks []Hook
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
	if len(t.lineHooks) > 0 {
		for _, h := range t.lineHooks {
			formatString = h(formatString)
		}
	}
	t.lines += fmt.Sprintf(formatString, args...)
	fmt.Fprintf(t.writer, formatString, args...)
}

// AddLineHooks provides adding of line hooks
// Its call at the stage of formatting. So, its doesen't change
// content of the lines. For example, this is possible
// to add color of the lines
func (t *App) AddLineHooks(hooks ...Hook) {
	t.lineHooks = hooks
}

// AddHeader will write a new table line followed by a separator
func (t *App) AddHeader(args ...interface{}) {
	t.AddLine(args...)
	t.addSeparator(args)
}

// Build will write the table to the terminal
func (t *App) Build() {
	t.writer.Flush()
}

// String returns formated string
func (t *App) String() string {
	return t.lines
}

// addSeparator will write a new dash separator line based on the args length
func (t *App) addSeparator(args []interface{}) {
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

// buildFormatString will build up the formatting string
func (t *App) buildFormatString(args []interface{}) string {
	var b bytes.Buffer
	l := len(args)
	for idx := range args {
		b.WriteString("%v")
		if idx+1 == l {
			break
		}
		b.WriteString("\t")
	}
	b.WriteString("\n")
	return b.String()
}
