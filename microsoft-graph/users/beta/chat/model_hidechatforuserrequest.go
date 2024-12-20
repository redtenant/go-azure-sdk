package chat

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HideChatForUserRequest struct {
	TenantId nullable.Type[string]      `json:"tenantId,omitempty"`
	User     *beta.TeamworkUserIdentity `json:"user,omitempty"`
}
