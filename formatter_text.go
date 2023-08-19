// to implement a text formatter
// fundemental text output format
package naivelog

import (
	"fmt"
	"time"
)

type TextFormatter struct {
	// to determine if the basic info should be printed
	IgnoreBasicFields bool
}

func (f *TextFormatter) Format(e *Entry) error {
	if !f.IgnoreBasicFields {
		e.Buffer.WriteString(fmt.Sprintf("%s %s", e.Time.Format(time.Layout), LevelNameMap[e.Level]))
		if e.File != "" {
			// to eliminate the file path
			shortName := e.File
			fileNameEndIdx := len(e.File) - 1
			for i := fileNameEndIdx; i > 0; i-- {
				if e.File[i] == '/' {
					shortName = e.File[i+1:]
					break
				}
			}
			e.Buffer.WriteString(fmt.Sprintf(" %s:%d", shortName, e.Line))
		}
		e.Buffer.WriteString(" ")
	}

	if e.Format == FmtEmptySeparate {
		e.Buffer.WriteString(fmt.Sprint(e.Args...))
	} else {
		e.Buffer.WriteString(fmt.Sprintf(e.Format, e.Args...))
	}
	e.Buffer.WriteString("\n")

	return nil
}
