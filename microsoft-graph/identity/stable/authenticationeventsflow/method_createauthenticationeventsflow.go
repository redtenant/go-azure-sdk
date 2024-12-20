package authenticationeventsflow

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAuthenticationEventsFlowOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.AuthenticationEventsFlow
}

type CreateAuthenticationEventsFlowOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAuthenticationEventsFlowOperationOptions() CreateAuthenticationEventsFlowOperationOptions {
	return CreateAuthenticationEventsFlowOperationOptions{}
}

func (o CreateAuthenticationEventsFlowOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAuthenticationEventsFlowOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAuthenticationEventsFlowOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAuthenticationEventsFlow - Create authenticationEventsFlow. Create a new authenticationEventsFlow object that
// is of the type specified in the request body. The following derived subtypes are supported: -
// externalUsersSelfServiceSignupEventsFlow object type.
func (c AuthenticationEventsFlowClient) CreateAuthenticationEventsFlow(ctx context.Context, input stable.AuthenticationEventsFlow, options CreateAuthenticationEventsFlowOperationOptions) (result CreateAuthenticationEventsFlowOperationResponse, err error) {
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
		Path:          "/identity/authenticationEventsFlows",
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
	model, err := stable.UnmarshalAuthenticationEventsFlowImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
