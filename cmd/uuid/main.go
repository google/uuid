// Package main is the main package for the UUID generator application.
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

// parse is a command-line flag to parse the given UUID. ('-' for STDIN)
var (
	parse = flag.String("d", "", "Parse given UUID. ('-' for STDIN)")
)

func main() {
	flag.Parse()
	var lreader io.Reader

	// If no UUID is given as input, generate a new UUID and print it.
	if *parse == "" {
		u, _ := uuid.NewUUID()
		fmt.Printf("%s\n", u)
	} else if *parse != "" { // If a UUID is provided as input, parse and process it.
		if *parse == "-" { // If input is "-", read from STDIN.
			lreader = os.Stdin
		} else if *parse != "-" { // Otherwise, read from the provided input string.
			lreader = strings.NewReader(*parse)
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(lreader)
		s := buf.String()
		s = strings.TrimSuffix(s, "\n")

		// Parse the UUID string to a UUID object.
		uuid, _ := uuid.Parse(s)
		fmt.Printf("UUID= %s\n", uuid)

		// Get the timestamp information from the parsed UUID.
		id := uuid
		t := id.Time()
		sec, nsec := t.UnixTime()
		timeStamp := time.Unix(sec, nsec)

		// Print the timestamp in a formatted date string.
		fmt.Printf("DATE= %v \n", timeStamp.Format("2006-01-02 Mon 15:04:05.00000Z -0700"))
	}
}
