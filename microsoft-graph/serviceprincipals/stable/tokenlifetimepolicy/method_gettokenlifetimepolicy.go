package tokenlifetimepolicy

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTokenLifetimePolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.TokenLifetimePolicy
}

type GetTokenLifetimePolicyOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetTokenLifetimePolicyOperationOptions() GetTokenLifetimePolicyOperationOptions {
	return GetTokenLifetimePolicyOperationOptions{}
}

func (o GetTokenLifetimePolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTokenLifetimePolicyOperationOptions) ToOData() *odata.Query {
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

func (o GetTokenLifetimePolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTokenLifetimePolicy - Get tokenLifetimePolicies from servicePrincipals. The tokenLifetimePolicies assigned to this
// service principal.
func (c TokenLifetimePolicyClient) GetTokenLifetimePolicy(ctx context.Context, id stable.ServicePrincipalIdTokenLifetimePolicyId, options GetTokenLifetimePolicyOperationOptions) (result GetTokenLifetimePolicyOperationResponse, err error) {
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

	var model stable.TokenLifetimePolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
