package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const PREFIX string = "Date:"

const DT_FORMAT = "Mon, 02 Jan 2006 15:04:05 -0700 (MST)"

var DTFormats = [6]string{
	time.RFC1123Z,
	time.RFC1123,
	"Mon, 2 Jan 2006 15:04:05 -0700", // RFC1123Z but use 2 instead of 02
	"Mon, 2 Jan 2006 15:04:05 MST",   // RFC1123 but use 2 instead of 02
	DT_FORMAT,
	"Mon, 2 Jan 2006 15:04:05 -0700 (MST)",
}

func LocalizeTime(dt string) (string, error) {
	for _, format := range DTFormats {
		t, err := time.Parse(format, dt)
		if err == nil {
			localTime := t.Local()
			return localTime.Format(DT_FORMAT), nil
		}
	}
	return "", errors.New("Unknown date format")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if strings.HasPrefix(line, PREFIX) {
			dt := strings.TrimPrefix(line, PREFIX)
			dt = strings.TrimSpace(dt)
			local_dt, err := LocalizeTime(dt)
			if err != nil {
				fmt.Print(line)
			} else {
				fmt.Println(PREFIX, local_dt)
			}
		} else {
			fmt.Print(line)
		}
	}
}
