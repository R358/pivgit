// Copyright (c) 2014, R358.org
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
// list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// * Neither the name of pivgit nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

//
// PivotalTracker is a Registered Trademark of http://pivotallabs.com
//

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/R358/pivgit/pivapi5"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	FETCH_PROJECT_LIST       = "https://www.pivotaltracker.com/services/v5/projects"
	FETCH_STORIES_ON_PROJECT = "https://www.pivotaltracker.com/services/v5/projects/%v/stories"
	PIVOTAL_TOKEN_HEADER     = "X-TrackerToken"
)

func main() {

	var token string = os.Getenv("PIVGIT_TOKEN")

	var prjOffset = 0
	if token == "" {
		if len(os.Args) < 2 {
			log.Fatalln("Token not set, either supply the token as the first parameter or set PIVGET_TOKEN")
		} else {
			prjOffset = 1
			token = os.Args[1]
		}
	}

	//
	// Make a map of projects to easily check against during story scan.
	//

	checkPrj := make(map[string]string)

	if len(os.Args) > 1 {
		for i, v := range os.Args {
			if i > prjOffset {
				checkPrj[v] = v
			}
		}
	}

	m := track(token, checkPrj)

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("No Pipe.  For example 'git branch | pivgit' ")
	} else {

		reg, err := regexp.Compile("\\D+")
		if err != nil {
			log.Fatal(err)
		}

		bio := bufio.NewScanner(os.Stdin)
		var ls bool = false
		for bio.Scan() {
			tokens := strings.Fields(bio.Text())
			ls = false
			for _, v := range tokens {
				if ls {
					print(" ")
				} else {
					ls = true
				}
				print(v)
				token = reg.ReplaceAllString(v, "")
				if token != "" {
					if v, ok := m[token]; ok {
						print(" " + v.Name)
					}
				}
			}

			println()

		}

	}

}

// Scan all or specific projects accessible for a user's token.
//
func track(token string, checkPrj map[string]string) map[string]pivapi5.Story {

	client := http.Client{}
	req, _ := http.NewRequest("GET", FETCH_PROJECT_LIST, nil)

	req.Header.Add(PIVOTAL_TOKEN_HEADER, token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var projects []pivapi5.Project
	err = json.Unmarshal(body, &projects)

	if err != nil {
		log.Fatal(err.Error())
	}

	m := make(map[string]pivapi5.Story)

	for _, v := range projects {

		//
		// Empty map or in the map.
		//
		if _, ok := checkPrj[v.Name]; len(checkPrj) == 0 || ok {
			stories(v, token, func(s pivapi5.Story) {
				m[strconv.FormatInt(s.Id, 10)] = s
			})
		}

	}

	return m
}

// Load the stories for a given project.
//
func stories(project pivapi5.Project, token string, c func(pivapi5.Story)) {
	client := http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf(FETCH_STORIES_ON_PROJECT, project.Id), nil)

	req.Header.Add(PIVOTAL_TOKEN_HEADER, token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Fetching Story: " + err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var stories []pivapi5.Story
	err = json.Unmarshal(body, &stories)

	if err != nil {
		log.Fatal("Unmarshal story: " + err.Error())
	}

	for _, v := range stories {
		c(v)
	}

}
