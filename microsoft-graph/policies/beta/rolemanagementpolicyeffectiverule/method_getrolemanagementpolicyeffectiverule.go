package rolemanagementpolicyeffectiverule

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRoleManagementPolicyEffectiveRuleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.UnifiedRoleManagementPolicyRule
}

type GetRoleManagementPolicyEffectiveRuleOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetRoleManagementPolicyEffectiveRuleOperationOptions() GetRoleManagementPolicyEffectiveRuleOperationOptions {
	return GetRoleManagementPolicyEffectiveRuleOperationOptions{}
}

func (o GetRoleManagementPolicyEffectiveRuleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetRoleManagementPolicyEffectiveRuleOperationOptions) ToOData() *odata.Query {
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

func (o GetRoleManagementPolicyEffectiveRuleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetRoleManagementPolicyEffectiveRule - Get effectiveRules from policies. The list of effective rules like approval
// rules and expiration rules evaluated based on inherited referenced rules. For example, if there is a tenant-wide
// policy to enforce enabling an approval rule, the effective rule will be to enable approval even if the policy has a
// rule to disable approval. Supports $expand.
func (c RoleManagementPolicyEffectiveRuleClient) GetRoleManagementPolicyEffectiveRule(ctx context.Context, id beta.PolicyRoleManagementPolicyIdEffectiveRuleId, options GetRoleManagementPolicyEffectiveRuleOperationOptions) (result GetRoleManagementPolicyEffectiveRuleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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
	model, err := beta.UnmarshalUnifiedRoleManagementPolicyRuleImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
