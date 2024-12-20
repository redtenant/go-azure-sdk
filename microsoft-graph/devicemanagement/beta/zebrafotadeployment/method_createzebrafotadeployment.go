package zebrafotadeployment

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateZebraFotaDeploymentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ZebraFotaDeployment
}

type CreateZebraFotaDeploymentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateZebraFotaDeploymentOperationOptions() CreateZebraFotaDeploymentOperationOptions {
	return CreateZebraFotaDeploymentOperationOptions{}
}

func (o CreateZebraFotaDeploymentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateZebraFotaDeploymentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateZebraFotaDeploymentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateZebraFotaDeployment - Create new navigation property to zebraFotaDeployments for deviceManagement
func (c ZebraFotaDeploymentClient) CreateZebraFotaDeployment(ctx context.Context, input beta.ZebraFotaDeployment, options CreateZebraFotaDeploymentOperationOptions) (result CreateZebraFotaDeploymentOperationResponse, err error) {
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
		Path:          "/deviceManagement/zebraFotaDeployments",
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

	var model beta.ZebraFotaDeployment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
