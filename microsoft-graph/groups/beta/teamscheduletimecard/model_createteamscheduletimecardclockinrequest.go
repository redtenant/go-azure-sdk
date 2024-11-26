package teamscheduletimecard

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTeamScheduleTimeCardClockInRequest struct {
	AtApprovedLocation nullable.Type[bool]   `json:"atApprovedLocation,omitempty"`
	Notes              *beta.ItemBody        `json:"notes,omitempty"`
	OnBehalfOfUserId   nullable.Type[string] `json:"onBehalfOfUserId,omitempty"`
}
