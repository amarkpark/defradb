// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package one_to_one_to_one

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

var bookAuthorGQLSchema = `
	type book {
		name: String
		rating: Float
		author: author
		publisher: publisher @primary
	}

	type author {
		name: String
		age: Int
		verified: Boolean
		published: book @primary
	}

	type publisher {
		name: String
		address: String
		yearOpened: Int
		printed: book
	}
`

func executeTestCase(t *testing.T, test testUtils.QueryTestCase) {
	testUtils.ExecuteQueryTestCase(t, bookAuthorGQLSchema, []string{"book", "author", "publisher"}, test)
}