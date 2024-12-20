package eventexceptionoccurrenceinstance

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AcceptEventExceptionOccurrenceInstanceRequest struct {
	Comment      nullable.Type[string] `json:"Comment,omitempty"`
	SendResponse nullable.Type[bool]   `json:"SendResponse,omitempty"`
}
