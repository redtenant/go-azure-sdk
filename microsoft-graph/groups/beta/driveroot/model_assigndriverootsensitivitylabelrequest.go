package driveroot

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDriveRootSensitivityLabelRequest struct {
	AssignmentMethod   *beta.SensitivityLabelAssignmentMethod `json:"assignmentMethod,omitempty"`
	JustificationText  nullable.Type[string]                  `json:"justificationText,omitempty"`
	SensitivityLabelId nullable.Type[string]                  `json:"sensitivityLabelId,omitempty"`
}
