package profilepublication

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

type GetProfilePublicationsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetProfilePublicationsCountOperationOptions struct {
	Filter    *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Search    *string
}

func DefaultGetProfilePublicationsCountOperationOptions() GetProfilePublicationsCountOperationOptions {
	return GetProfilePublicationsCountOperationOptions{}
}

func (o GetProfilePublicationsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetProfilePublicationsCountOperationOptions) ToOData() *odata.Query {
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

func (o GetProfilePublicationsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetProfilePublicationsCount - Get the number of the resource
func (c ProfilePublicationClient) GetProfilePublicationsCount(ctx context.Context, id beta.UserId, options GetProfilePublicationsCountOperationOptions) (result GetProfilePublicationsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/profile/publications/$count", id.ID()),
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
