package sitelistcontenttype

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopySiteListContentTypeToDefaultContentLocationRequest struct {
	DestinationFileName nullable.Type[string] `json:"destinationFileName,omitempty"`
	SourceFile          *beta.ItemReference   `json:"sourceFile,omitempty"`
}
