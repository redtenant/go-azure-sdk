package group

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidatePropertiesRequest struct {
	DisplayName      nullable.Type[string] `json:"displayName,omitempty"`
	MailNickname     nullable.Type[string] `json:"mailNickname,omitempty"`
	OnBehalfOfUserId nullable.Type[string] `json:"onBehalfOfUserId,omitempty"`
}
