package devicehealthscript

import (
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceHealthScriptGlobalScriptResult struct {
	Value nullable.Type[string] `json:"value,omitempty"`
}
