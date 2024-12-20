package domain

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateForceDeleteRequest struct {
	DisableUserAccounts nullable.Type[bool] `json:"disableUserAccounts,omitempty"`
}
