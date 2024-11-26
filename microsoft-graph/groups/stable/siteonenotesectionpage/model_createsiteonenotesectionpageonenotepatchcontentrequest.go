package siteonenotesectionpage

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSiteOnenoteSectionPageOnenotePatchContentRequest struct {
	Commands *[]stable.OnenotePatchContentCommand `json:"commands,omitempty"`
}
