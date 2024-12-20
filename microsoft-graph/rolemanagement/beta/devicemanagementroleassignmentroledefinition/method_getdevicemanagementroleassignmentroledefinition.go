package devicemanagementroleassignmentroledefinition

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceManagementRoleAssignmentRoleDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleDefinition
}

type GetDeviceManagementRoleAssignmentRoleDefinitionOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetDeviceManagementRoleAssignmentRoleDefinitionOperationOptions() GetDeviceManagementRoleAssignmentRoleDefinitionOperationOptions {
	return GetDeviceManagementRoleAssignmentRoleDefinitionOperationOptions{}
}

func (o GetDeviceManagementRoleAssignmentRoleDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceManagementRoleAssignmentRoleDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDeviceManagementRoleAssignmentRoleDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceManagementRoleAssignmentRoleDefinition - Get roleDefinition from roleManagement. Specifies the
// roleDefinition that the assignment is for. Provided so that callers can get the role definition using $expand at the
// same time as getting the role assignment. Supports $filter (eq operator on id, isBuiltIn, and displayName, and
// startsWith operator on displayName) and $expand.
func (c DeviceManagementRoleAssignmentRoleDefinitionClient) GetDeviceManagementRoleAssignmentRoleDefinition(ctx context.Context, id beta.RoleManagementDeviceManagementRoleAssignmentId, options GetDeviceManagementRoleAssignmentRoleDefinitionOperationOptions) (result GetDeviceManagementRoleAssignmentRoleDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/roleDefinition", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.UnifiedRoleDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
