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
	person = flag.Bool("p", false, "Generate DCE Person UUID.")
	random = flag.Bool("r", false, "Generate Random UUID.")
)

func main() {
	flag.Parse()
	var lreader io.Reader

	if *parse == "" && *person == false {
		u := uuid.New()
		fmt.Printf("%s", u)
	} else if *parse == "" && *person == true {
		u, _ := uuid.NewDCEPerson()
		fmt.Printf("%s", u)
	} else if *parse == "" && *random == true {
		u, _ := uuid.NewRandom()
		fmt.Printf("%s", u)
	} else if *parse == "-" {
		lreader = os.Stdin
	} else if *parse != "" && *parse != "-" {
		lreader = strings.NewReader(*parse)

		buf := new(bytes.Buffer)
		buf.ReadFrom(lreader)
		s := buf.String()

		uuid, _ := uuid.Parse(s)
		fmt.Printf("Successfully parsed UID : %s\n", uuid)

		id := uuid
		t := id.Time()
		sec, nsec := t.UnixTime()
		timeStamp := time.Unix(sec, nsec)
		fmt.Printf("The id was generated at : %v \n", timeStamp.Format("2006-01-02 Mon 15:04:05.000Z -0700"))
	}
}
