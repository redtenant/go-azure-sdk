package signalr

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomDomainsGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *CustomDomain
}

// CustomDomainsGet ...
func (c SignalRClient) CustomDomainsGet(ctx context.Context, id CustomDomainId) (result CustomDomainsGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       id.ID(),
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

	var model CustomDomain
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
