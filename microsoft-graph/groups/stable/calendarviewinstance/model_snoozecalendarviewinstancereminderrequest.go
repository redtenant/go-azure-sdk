package calendarviewinstance

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SnoozeCalendarViewInstanceReminderRequest struct {
	NewReminderTime *stable.DateTimeTimeZone `json:"NewReminderTime,omitempty"`
}
