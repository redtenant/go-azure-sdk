package conditionalaccessauthenticationstrengthauthenticationmethodmode

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetConditionalAccessAuthenticationStrengthAuthenticationMethodModesCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetConditionalAccessAuthenticationStrengthAuthenticationMethodModesCountOperationOptions struct {
	Filter    *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Search    *string
}

func DefaultGetConditionalAccessAuthenticationStrengthAuthenticationMethodModesCountOperationOptions() GetConditionalAccessAuthenticationStrengthAuthenticationMethodModesCountOperationOptions {
	return GetConditionalAccessAuthenticationStrengthAuthenticationMethodModesCountOperationOptions{}
}

func (o GetConditionalAccessAuthenticationStrengthAuthenticationMethodModesCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetConditionalAccessAuthenticationStrengthAuthenticationMethodModesCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetConditionalAccessAuthenticationStrengthAuthenticationMethodModesCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetConditionalAccessAuthenticationStrengthAuthenticationMethodModesCount - Get the number of the resource
func (c ConditionalAccessAuthenticationStrengthAuthenticationMethodModeClient) GetConditionalAccessAuthenticationStrengthAuthenticationMethodModesCount(ctx context.Context, options GetConditionalAccessAuthenticationStrengthAuthenticationMethodModesCountOperationOptions) (result GetConditionalAccessAuthenticationStrengthAuthenticationMethodModesCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/identity/conditionalAccess/authenticationStrength/authenticationMethodModes/$count",
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
