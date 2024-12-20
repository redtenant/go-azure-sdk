package windowsdriverupdateprofile

import (
	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateWindowsDriverUpdateProfileExecuteActionRequest struct {
	// An enum type to represent approval actions of single or list of drivers.
	ActionName *beta.DriverApprovalAction `json:"actionName,omitempty"`

	DeploymentDate nullable.Type[string] `json:"deploymentDate,omitempty"`
	DriverIds      *[]string             `json:"driverIds,omitempty"`
}
