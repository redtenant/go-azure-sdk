package entitlementmanagementroleeligibilityschedulerequestdirectoryscope

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.DirectoryObject
}

type GetEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeOperationOptions() GetEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeOperationOptions {
	return GetEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeOperationOptions{}
}

func (o GetEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeOperationOptions) ToOData() *odata.Query {
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

func (o GetEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementRoleEligibilityScheduleRequestDirectoryScope - Get directoryScope from roleManagement. The
// directory object that is the scope of the role eligibility. Read-only. Supports $expand.
func (c EntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeClient) GetEntitlementManagementRoleEligibilityScheduleRequestDirectoryScope(ctx context.Context, id beta.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId, options GetEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeOperationOptions) (result GetEntitlementManagementRoleEligibilityScheduleRequestDirectoryScopeOperationResponse, err error) {
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
	model, err := beta.UnmarshalDirectoryObjectImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
