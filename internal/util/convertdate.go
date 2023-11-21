// internal/util/convertdate.go

package util

import (
	"fmt"
	"time"
)

func ConvertDate(date string) string {
	dateParse, err := time.Parse("2006-01-02T15:04:05.999999999Z", date)
	if err != nil {
		fmt.Println("Error converting date: ", err)
		dateParse = time.Now()
	}

	dateConv := dateParse.Format("Mon 2 Jan 2006 15:04:05")

	return dateConv
}
