package owneddevice

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

type GetOwnedDevicesCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetOwnedDevicesCountOperationOptions struct {
	ConsistencyLevel *odata.ConsistencyLevel
	Filter           *string
	Metadata         *odata.Metadata
	RetryFunc        client.RequestRetryFunc
	Search           *string
}

func DefaultGetOwnedDevicesCountOperationOptions() GetOwnedDevicesCountOperationOptions {
	return GetOwnedDevicesCountOperationOptions{}
}

func (o GetOwnedDevicesCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetOwnedDevicesCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.ConsistencyLevel != nil {
		out.ConsistencyLevel = *o.ConsistencyLevel
	}
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

func (o GetOwnedDevicesCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetOwnedDevicesCount - Get the number of the resource
func (c OwnedDeviceClient) GetOwnedDevicesCount(ctx context.Context, id stable.UserId, options GetOwnedDevicesCountOperationOptions) (result GetOwnedDevicesCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/ownedDevices/$count", id.ID()),
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
