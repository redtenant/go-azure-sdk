package application

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddKeyRequest struct {
	KeyCredential      *stable.KeyCredential      `json:"keyCredential,omitempty"`
	PasswordCredential *stable.PasswordCredential `json:"passwordCredential,omitempty"`
	Proof              *string                    `json:"proof,omitempty"`
}
