package tokenlifetimepolicy

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

type AddTokenLifetimePolicyRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddTokenLifetimePolicyRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddTokenLifetimePolicyRefOperationOptions() AddTokenLifetimePolicyRefOperationOptions {
	return AddTokenLifetimePolicyRefOperationOptions{}
}

func (o AddTokenLifetimePolicyRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddTokenLifetimePolicyRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddTokenLifetimePolicyRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddTokenLifetimePolicyRef - Assign tokenLifetimePolicy. Assign a tokenLifetimePolicy to an application or
// servicePrincipal. You can have multiple tokenLifetimePolicy policies in a tenant but can assign only one
// tokenLifetimePolicy per application.
func (c TokenLifetimePolicyClient) AddTokenLifetimePolicyRef(ctx context.Context, id beta.ApplicationId, input beta.ReferenceCreate, options AddTokenLifetimePolicyRefOperationOptions) (result AddTokenLifetimePolicyRefOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/tokenLifetimePolicies/$ref", id.ID()),
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

	return
}
