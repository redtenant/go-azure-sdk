package sitecontentmodel

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveSiteContentModelFromDriveRequest struct {
	DriveId nullable.Type[string] `json:"driveId,omitempty"`
}
