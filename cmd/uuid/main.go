package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	parse  = flag.String("d", "", "Parse given UUID.")
	random = flag.Bool("r", false, "Generate Random UUID.")
)

func main() {
	flag.Parse()
	var lreader io.Reader

	if *parse == "" && *random == false {
		u, _ := uuid.NewUUID()
		fmt.Printf("%s\n", u)
	} else if *parse == "" && *random == true {
		u := uuid.New()
		fmt.Printf("%s\n", u)
	} else if *parse != "" {
		if *parse == "-" {
			lreader = os.Stdin
		} else if *parse != "-" {
			lreader = strings.NewReader(*parse)
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(lreader)
		s := buf.String()

		uuid, _ := uuid.Parse(s)
		fmt.Printf("uuid : %s\n", uuid)

		id := uuid
		t := id.Time()
		sec, nsec := t.UnixTime()
		timeStamp := time.Unix(sec, nsec)
		fmt.Printf("date : %v \n", timeStamp.Format("2006-01-02 Mon 15:04:05.000Z -0700"))
	}
}
