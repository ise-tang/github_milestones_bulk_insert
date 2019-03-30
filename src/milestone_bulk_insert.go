package main

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
  "gopkg.in/yaml.v2"
  "io/ioutil"
	"context"
	"os"
	"io"
	"encoding/csv"
	"time"
	"fmt"
)

func main() {
	// Load setting.yml
	buf, err := ioutil.ReadFile("settings.yml")
	if err != nil {
    return
	}

  m := make(map[interface{}]interface{})
  err = yaml.Unmarshal(buf, &m)
  if err != nil {
    panic(err)
  }

	// create GitHub client
	ctx := context.Background()
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: m["access_token"].(string)},
  )
  tc := oauth2.NewClient(oauth2.NoContext, ts)

  client := github.NewClient(tc)

  // read CSV
  var fp *os.File
  if len(os.Args) < 2 {
    fp = os.Stdin
  } else {
    var err error
    fp, err = os.Open(os.Args[1])
    if err != nil {
    				panic(err)
    }
    defer fp.Close()
  }

	reader := csv.NewReader(fp)
	reader.Comma = ','
	reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない！
	var header []string
	layout := "2006/01/02"

	for {
		record, err := reader.Read()
		if err == io.EOF {
						break
		} else if err != nil {
						panic(err)
		} else if header == nil {
			header = record
			continue
		}
		due_on, err := time.Parse(layout, record[2])
		if err != nil {
		  panic(err)
		}

		milestone :=  github.Milestone {
			Title: github.String(record[0]),
			Description: github.String(record[1]),
			DueOn: &due_on,
    }
		// create milestone
		repos, _, err := client.Issues.CreateMilestone(ctx, m["owner"].(string), m["repo"].(string), &milestone)

		if err != nil {
			panic(err)
		}

		if repos != nil {
			fmt.Println("created.")
		}
	}

}