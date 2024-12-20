package beta

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAnswerKeyword struct {
	// A collection of keywords used to trigger the search answer.
	Keywords *[]string `json:"keywords,omitempty"`

	// If true, indicates that the search term contains similar words to the keywords that should trigger the search answer.
	MatchSimilarKeywords nullable.Type[bool] `json:"matchSimilarKeywords,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Unique keywords that guarantee the search answer is triggered.
	ReservedKeywords *[]string `json:"reservedKeywords,omitempty"`
}
