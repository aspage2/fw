package main

import (
	"flag"
	"fmt"
	"github.com/aspage2/fw/fw"
	"io/ioutil"
	"os"
	"strings"
)

// GetQuery gets the candidate querystring from either
// the commandline args or stdin. Commandline args are
// parsed with flag.Args and are joined with a space character.
// If flag.Args is empty, stdin is used instead.
func GetQuery() (string, error) {
	if !flag.Parsed() {
		flag.Parse()
	}
	args := flag.Args()
	if len(args) == 0 {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return "", err
		}
		return string(data), nil
	} else {
		return strings.Join(args, " "), nil
	}
}

func main() {
	query, err := GetQuery()
	if err != nil {
		panic(err.Error())
	}

	var resp strings.Builder
	for _, r := range query {
		resp.WriteRune(fw.Convert(r))
	}

	fmt.Print(resp.String())
}