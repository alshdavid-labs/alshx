package paths

import (
	"fmt"
	"time"
)

func ArchiveNamePath() string {
	t := time.Now().UTC()

	var year = fmt.Sprintf("%d", t.Year())
	var month = fmt.Sprintf("%d", t.Month())
	var day = fmt.Sprintf("%d", t.Day())

	year = year[2:]

	if len(month) == 1 {
		month = fmt.Sprintf("0%s", month)
	}

	if len(day) == 1 {
		day = fmt.Sprintf("0%s", day)
	}

	return fmt.Sprintf("%s-%s-%s", year, month, day)
}
