package mailfolderchildfoldermessage

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMailFolderChildFolderMessageReplyAllRequest struct {
	Comment nullable.Type[string] `json:"Comment,omitempty"`
	Message *beta.Message         `json:"Message,omitempty"`
}
