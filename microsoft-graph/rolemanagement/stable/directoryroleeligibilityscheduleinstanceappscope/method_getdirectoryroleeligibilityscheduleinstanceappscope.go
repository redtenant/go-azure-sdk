package directoryroleeligibilityscheduleinstanceappscope

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDirectoryRoleEligibilityScheduleInstanceAppScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AppScope
}

type GetDirectoryRoleEligibilityScheduleInstanceAppScopeOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetDirectoryRoleEligibilityScheduleInstanceAppScopeOperationOptions() GetDirectoryRoleEligibilityScheduleInstanceAppScopeOperationOptions {
	return GetDirectoryRoleEligibilityScheduleInstanceAppScopeOperationOptions{}
}

func (o GetDirectoryRoleEligibilityScheduleInstanceAppScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDirectoryRoleEligibilityScheduleInstanceAppScopeOperationOptions) ToOData() *odata.Query {
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

func (o GetDirectoryRoleEligibilityScheduleInstanceAppScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDirectoryRoleEligibilityScheduleInstanceAppScope - Get appScope from roleManagement. Read-only property with
// details of the app-specific scope when the assignment or role eligibility is scoped to an app. Nullable.
func (c DirectoryRoleEligibilityScheduleInstanceAppScopeClient) GetDirectoryRoleEligibilityScheduleInstanceAppScope(ctx context.Context, id stable.RoleManagementDirectoryRoleEligibilityScheduleInstanceId, options GetDirectoryRoleEligibilityScheduleInstanceAppScopeOperationOptions) (result GetDirectoryRoleEligibilityScheduleInstanceAppScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/appScope", id.ID()),
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

	var model stable.AppScope
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
