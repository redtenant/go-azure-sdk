package resourceaccessprofile

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignResourceAccessProfilesRequest struct {
	Assignments *[]beta.DeviceManagementResourceAccessProfileAssignment `json:"assignments,omitempty"`
}
