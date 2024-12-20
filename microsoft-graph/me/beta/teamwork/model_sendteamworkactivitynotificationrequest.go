package teamwork

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendTeamworkActivityNotificationRequest struct {
	ActivityType       nullable.Type[string]       `json:"activityType,omitempty"`
	ChainId            nullable.Type[int64]        `json:"chainId,omitempty"`
	PreviewText        *beta.ItemBody              `json:"previewText,omitempty"`
	TeamsAppId         nullable.Type[string]       `json:"teamsAppId,omitempty"`
	TemplateParameters *[]beta.KeyValuePair        `json:"templateParameters,omitempty"`
	Topic              *beta.TeamworkActivityTopic `json:"topic,omitempty"`
}
