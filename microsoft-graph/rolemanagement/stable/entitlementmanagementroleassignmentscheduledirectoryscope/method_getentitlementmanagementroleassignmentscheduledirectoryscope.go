package entitlementmanagementroleassignmentscheduledirectoryscope

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntitlementManagementRoleAssignmentScheduleDirectoryScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.DirectoryObject
}

type GetEntitlementManagementRoleAssignmentScheduleDirectoryScopeOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementRoleAssignmentScheduleDirectoryScopeOperationOptions() GetEntitlementManagementRoleAssignmentScheduleDirectoryScopeOperationOptions {
	return GetEntitlementManagementRoleAssignmentScheduleDirectoryScopeOperationOptions{}
}

func (o GetEntitlementManagementRoleAssignmentScheduleDirectoryScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementRoleAssignmentScheduleDirectoryScopeOperationOptions) ToOData() *odata.Query {
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

func (o GetEntitlementManagementRoleAssignmentScheduleDirectoryScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementRoleAssignmentScheduleDirectoryScope - Get directoryScope from roleManagement. The directory
// object that is the scope of the role eligibility or assignment. Read-only.
func (c EntitlementManagementRoleAssignmentScheduleDirectoryScopeClient) GetEntitlementManagementRoleAssignmentScheduleDirectoryScope(ctx context.Context, id stable.RoleManagementEntitlementManagementRoleAssignmentScheduleId, options GetEntitlementManagementRoleAssignmentScheduleDirectoryScopeOperationOptions) (result GetEntitlementManagementRoleAssignmentScheduleDirectoryScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/directoryScope", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalDirectoryObjectImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
