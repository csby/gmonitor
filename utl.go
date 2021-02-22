package gmonitor

import (
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t DateTime) String() string {
	return time.Time(t).Format(timeFormat)
}

func toText(v float64) string {
	kb := float64(1024)
	mb := 1024 * kb
	gb := 1024 * mb

	if v >= gb {
		return fmt.Sprintf("%.1fGB", v/gb)
	} else if v >= mb {
		return fmt.Sprintf("%.1fMB", v/mb)
	} else if v >= kb {
		return fmt.Sprintf("%.1fKB", v/kb)
	} else {
		return fmt.Sprintf("%.0fB", v)
	}
}
