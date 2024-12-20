package privilegedaccessgroupassignmentscheduleactivatedusing

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

type GetPrivilegedAccessGroupAssignmentScheduleActivatedUsingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PrivilegedAccessGroupEligibilitySchedule
}

type GetPrivilegedAccessGroupAssignmentScheduleActivatedUsingOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetPrivilegedAccessGroupAssignmentScheduleActivatedUsingOperationOptions() GetPrivilegedAccessGroupAssignmentScheduleActivatedUsingOperationOptions {
	return GetPrivilegedAccessGroupAssignmentScheduleActivatedUsingOperationOptions{}
}

func (o GetPrivilegedAccessGroupAssignmentScheduleActivatedUsingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPrivilegedAccessGroupAssignmentScheduleActivatedUsingOperationOptions) ToOData() *odata.Query {
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

func (o GetPrivilegedAccessGroupAssignmentScheduleActivatedUsingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPrivilegedAccessGroupAssignmentScheduleActivatedUsing - Get activatedUsing from identityGovernance. When the
// request activates an ownership or membership assignment in PIM for groups, this object represents the eligibility
// relationship. Otherwise, it's null. Supports $expand.
func (c PrivilegedAccessGroupAssignmentScheduleActivatedUsingClient) GetPrivilegedAccessGroupAssignmentScheduleActivatedUsing(ctx context.Context, id stable.IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId, options GetPrivilegedAccessGroupAssignmentScheduleActivatedUsingOperationOptions) (result GetPrivilegedAccessGroupAssignmentScheduleActivatedUsingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/activatedUsing", id.ID()),
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

	var model stable.PrivilegedAccessGroupEligibilitySchedule
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
