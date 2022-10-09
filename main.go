package pagedb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// main
func main() {
	// start a http server
	http.HandleFunc("/query", handler)

}

// handler
func handler(w http.ResponseWriter, r *http.Request) {
	// request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// log error
		return
	}
	// json unmarshal
	type Query struct {
		Query string `json:"query"`
	}
	var query Query
	err = json.Unmarshal(body, &query)
	if err != nil {
		// log error
		return
	}

	// parse query
	/*
		Query Example:
			select
			id,name
			user.meta
			id > 1 //must be space separated
	*/
	// split query
	querySplit := strings.Split(strings.TrimSpace(query.Query), "\n")
	// query type
	queryType := strings.TrimSpace(querySplit[0])
	// query fields
	queryFields := strings.Split(strings.TrimSpace(querySplit[1]), ",")
	// query table
	queryTable := strings.TrimSpace(querySplit[2])
	// query condition
	queryCondition := strings.TrimSpace(querySplit[3])

	// if query type is not select , add, or delete, return 400 error
	if queryType != "select" && queryType != "add" && queryType != "delete" {
		w.WriteHeader(400)
		return
	}

	// if query type is select
	if queryType == "select" {
		// fopen queryTable.json

		fileHandler, err := os.Open(queryTable + ".json")
		if err != nil {
			// 404
			w.WriteHeader(404)
			return
		}
		// fileHandler read byte by byte

		b := make([]byte, 1)
		for {
			_, e := fileHandler.Read(b)
			if e != nil && !errors.Is(e, io.EOF) {
				fmt.Println(err)
				break
			}
			if errors.Is(e, io.EOF) {
				break
			}

			//check byte is 0x82 (start of object)
			if b[0] == 0x82 {
			}

		}
	}

}
