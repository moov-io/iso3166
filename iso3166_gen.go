// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

//go:build ignore
// +build ignore

// Generates iso3166.go.
//
// This file grabs the ISO 3166-1-alpha2 codes and writes them
// into source code so we don't rely on any external files (zip,
// json, etc).
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"slices"
	"strings"
	"time"
)

var (
	// From https://datahub.io/core/country-list#data
	downloadUrl = "https://raw.githubusercontent.com/fannarsh/country-list/master/data.json"
)

// [{"Code": "AF", "Name": "Afghanistan"}, ...]
type country struct {
	Code string `json:"Code"`
	Name string `json:"Name"`
}

func main() {
	when := time.Now().Format("2006-01-02T03:04:05Z")
	who, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to get user on %s", runtime.GOOS)
	}

	// Write copyright header
	var buf bytes.Buffer
	fmt.Fprintf(&buf, `// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Generated on %s by %s, any modifications will be overwritten
package iso3166
`, when, who.Username)

	// Download certs
	resp, err := http.Get(downloadUrl)
	if err != nil {
		log.Fatalf("error while downloading %s: %v", downloadUrl, err)
	}
	defer resp.Body.Close()

	var countries []country
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil {
		log.Fatalf("error while parsing country response: %v", err)
	}
	slices.SortFunc(countries, func(a, b country) int {
		if a.Name < b.Name {
			return -1
		}
		return 1
	})

	// Write countries to source code
	fmt.Fprintln(&buf, "var countryCodes = map[string][]string{")
	for i := range countries {
		countryCode := countries[i].Code

		// Add some extra names to certain countries
		countryNames := []string{countries[i].Name}
		switch countryNames[0] {
		case "United Kingdom of Great Britain and Northern Ireland":
			countryNames = append(countryNames, "United Kingdom", "England", "Great Britain")
		case "United States of America":
			countryNames = append(countryNames, "USA", "United States")
		}
		countryNamesValue := strings.Join(countryNames, `","`)

		fmt.Fprintf(&buf, fmt.Sprintf(`"%s": []string{"%s"},`+"\n", countryCode, countryNamesValue))
	}
	fmt.Fprintln(&buf, "}")

	// format source code and write file
	out, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
		log.Fatalf("error formatting output code, err=%v", err)
	}

	err = os.WriteFile("data.go", out, 0644)
	if err != nil {
		log.Fatalf("error writing file, err=%v", err)
	}
}
