package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const DT_FORMAT = "Mon, 2 Jan 2006 15:04:05 -0700"
const PREFIX = "Date:"

func LocalizeTime(dt string) (string, error) {
	t, err := time.Parse(DT_FORMAT, dt)
	if err != nil {
		return "", err
	}
	local_t := t.Local()
	return local_t.Format(DT_FORMAT), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if strings.HasPrefix(line, PREFIX) {
			// Clean up date. Sometimes it has (PST) in the end
			dt := strings.TrimPrefix(line, PREFIX)
			dt = strings.Split(dt, "(")[0]
			dt = strings.TrimSpace(dt)
			local_dt, err := LocalizeTime(dt)
			if err != nil {
				fmt.Print(line)
			}
			fmt.Println(PREFIX, local_dt)
		} else {
			fmt.Print(line)
		}
	}
}
