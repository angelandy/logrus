package logrus

/**
add seaslog by angelandy for golang
 */

import (
	"bytes"
	"sync"
)

// TextFormatter formats logs into text
type SeasFormatter struct {
	TimestampFormat string
	sync.Once
}



// Format renders a single log entry
func (f *SeasFormatter) Format(entry *Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	//f.Do(func() { f.init(entry) })

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	b.WriteString(entry.Time.Format(timestampFormat))
	b.WriteString(" | ")
	b.WriteString(entry.Level.String())
	b.WriteString(" | ")
	b.WriteString(entry.Message)
	b.WriteByte('\n')
	return b.Bytes(), nil
}