package privilegedaccessgroupeligibilityscheduleinstancegroup

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

type GetPrivilegedAccessGroupEligibilityScheduleInstanceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Group
}

type GetPrivilegedAccessGroupEligibilityScheduleInstanceGroupOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetPrivilegedAccessGroupEligibilityScheduleInstanceGroupOperationOptions() GetPrivilegedAccessGroupEligibilityScheduleInstanceGroupOperationOptions {
	return GetPrivilegedAccessGroupEligibilityScheduleInstanceGroupOperationOptions{}
}

func (o GetPrivilegedAccessGroupEligibilityScheduleInstanceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPrivilegedAccessGroupEligibilityScheduleInstanceGroupOperationOptions) ToOData() *odata.Query {
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

func (o GetPrivilegedAccessGroupEligibilityScheduleInstanceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPrivilegedAccessGroupEligibilityScheduleInstanceGroup - Get group from identityGovernance. References the group
// that is the scope of the membership or ownership eligibility through PIM for groups. Supports $expand.
func (c PrivilegedAccessGroupEligibilityScheduleInstanceGroupClient) GetPrivilegedAccessGroupEligibilityScheduleInstanceGroup(ctx context.Context, id beta.IdentityGovernancePrivilegedAccessGroupEligibilityScheduleInstanceId, options GetPrivilegedAccessGroupEligibilityScheduleInstanceGroupOperationOptions) (result GetPrivilegedAccessGroupEligibilityScheduleInstanceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/group", id.ID()),
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

	var model beta.Group
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
