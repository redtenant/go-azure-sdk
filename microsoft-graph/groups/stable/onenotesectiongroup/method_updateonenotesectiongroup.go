package onenotesectiongroup

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOnenoteSectionGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateOnenoteSectionGroupOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateOnenoteSectionGroupOperationOptions() UpdateOnenoteSectionGroupOperationOptions {
	return UpdateOnenoteSectionGroupOperationOptions{}
}

func (o UpdateOnenoteSectionGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateOnenoteSectionGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateOnenoteSectionGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateOnenoteSectionGroup - Update the navigation property sectionGroups in groups
func (c OnenoteSectionGroupClient) UpdateOnenoteSectionGroup(ctx context.Context, id stable.GroupIdOnenoteSectionGroupId, input stable.SectionGroup, options UpdateOnenoteSectionGroupOperationOptions) (result UpdateOnenoteSectionGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
