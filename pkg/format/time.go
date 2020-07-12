package format

import "time"

type TimeFormat string

const (
	Time TimeFormat = "2006-01-02T15:04:05Z"
)

func (tf TimeFormat) String() string {
	return string(tf)
}

func (tf TimeFormat) Build(t time.Time) string {
	return t.Format(tf.String())
}
