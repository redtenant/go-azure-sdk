package calendareventattachment

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCalendarEventAttachmentsUploadSessionRequest struct {
	AttachmentItem *beta.AttachmentItem `json:"AttachmentItem,omitempty"`
}
