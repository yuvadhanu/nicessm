package shared

import (
	"errors"
	"fmt"
	"time"
)

func (s *Shared) UniqueDateStr(date *time.Time) (string, error) {
	if date == nil {
		return "", errors.New("Date is nil")
	}
	str := fmt.Sprintf("%v_%v_%v", date.Day(), date.Month().String(), date.Year())
	return str, nil
}
