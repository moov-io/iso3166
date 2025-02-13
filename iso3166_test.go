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

package iso3166_test

import (
	"testing"

	"github.com/moov-io/iso3166"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	require.True(t, iso3166.Valid("US"))
	require.True(t, iso3166.Valid("SS"))
	require.False(t, iso3166.Valid(""))
	require.False(t, iso3166.Valid("QZ"))
}

func TestLookupCode(t *testing.T) {
	require.Equal(t, "US", iso3166.LookupCode("U.S.A"))
	require.Equal(t, "US", iso3166.LookupCode("united states"))
	require.Equal(t, "GB", iso3166.LookupCode("ENGLAND"))
	require.Equal(t, "VI", iso3166.LookupCode("U.S.,  Virgin Islands"))
}

func TestGetName(t *testing.T) {
	require.Equal(t, "", iso3166.GetName(""))
	require.Equal(t, "", iso3166.GetName("not-found"))

	require.Equal(t, "United States of America", iso3166.GetName("us"))
}
