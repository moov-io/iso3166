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

package iso3166

import (
	"strings"
)

// Valid returns successful if code is a valid ISO 3166-1-alpha-2
// country code.
//
// Example: US
func Valid(code string) bool {
	_, ok := countryCodes[strings.ToUpper(code)]
	return ok
}

var (
	punctuationRemover = strings.NewReplacer(".", "", ",", "")
)

// LookupCode will attempt to find a valid ISO 3166-1-alpha-2 code
// for the given country name.
//
// Example: "United States" = "US"
func LookupCode(input string) string {
	input = punctuationRemover.Replace(input)
	input = strings.ReplaceAll(input, "  ", " ")

	for code, names := range countryCodes {
		for _, name := range names {
			if strings.EqualFold(input, name) {
				return code
			}
		}
	}
	return ""
}

// GetName returns the ISO 3166 name for a given ISO 3166-1-alpha-2 code
func GetName(code string) string {
	names := countryCodes[strings.ToUpper(code)]
	if len(names) > 0 {
		return names[0]
	}
	return ""
}
