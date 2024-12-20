package applicationtemplate

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

type InstantiateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ApplicationServicePrincipal
}

type InstantiateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultInstantiateOperationOptions() InstantiateOperationOptions {
	return InstantiateOperationOptions{}
}

func (o InstantiateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o InstantiateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o InstantiateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// Instantiate - Invoke action instantiate. Add an instance of an application from the Microsoft Entra application
// gallery into your directory. The application template with ID 8adf8e6e-67b2-4cf2-a259-e3dc5476c621 can be used to add
// a non-gallery app that you can configure different single-sign on (SSO) modes like SAML SSO and password-based SSO.
func (c ApplicationTemplateClient) Instantiate(ctx context.Context, id stable.ApplicationTemplateId, input InstantiateRequest, options InstantiateOperationOptions) (result InstantiateOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/instantiate", id.ID()),
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

	var model stable.ApplicationServicePrincipal
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
