package rolemanagementpolicyeffectiverule

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

type CreateRoleManagementPolicyEffectiveRuleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.UnifiedRoleManagementPolicyRule
}

type CreateRoleManagementPolicyEffectiveRuleOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateRoleManagementPolicyEffectiveRuleOperationOptions() CreateRoleManagementPolicyEffectiveRuleOperationOptions {
	return CreateRoleManagementPolicyEffectiveRuleOperationOptions{}
}

func (o CreateRoleManagementPolicyEffectiveRuleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateRoleManagementPolicyEffectiveRuleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateRoleManagementPolicyEffectiveRuleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateRoleManagementPolicyEffectiveRule - Create new navigation property to effectiveRules for policies
func (c RoleManagementPolicyEffectiveRuleClient) CreateRoleManagementPolicyEffectiveRule(ctx context.Context, id beta.PolicyRoleManagementPolicyId, input beta.UnifiedRoleManagementPolicyRule, options CreateRoleManagementPolicyEffectiveRuleOperationOptions) (result CreateRoleManagementPolicyEffectiveRuleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/effectiveRules", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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
