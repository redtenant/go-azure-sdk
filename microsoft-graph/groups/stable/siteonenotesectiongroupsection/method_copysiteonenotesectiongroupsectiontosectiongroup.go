package siteonenotesectiongroupsection

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

type CopySiteOnenoteSectionGroupSectionToSectionGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OnenoteOperation
}

type CopySiteOnenoteSectionGroupSectionToSectionGroupOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCopySiteOnenoteSectionGroupSectionToSectionGroupOperationOptions() CopySiteOnenoteSectionGroupSectionToSectionGroupOperationOptions {
	return CopySiteOnenoteSectionGroupSectionToSectionGroupOperationOptions{}
}

func (o CopySiteOnenoteSectionGroupSectionToSectionGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CopySiteOnenoteSectionGroupSectionToSectionGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CopySiteOnenoteSectionGroupSectionToSectionGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CopySiteOnenoteSectionGroupSectionToSectionGroup - Invoke action copyToSectionGroup. For Copy operations, you follow
// an asynchronous calling pattern: First call the Copy action, and then poll the operation endpoint for the result.
func (c SiteOnenoteSectionGroupSectionClient) CopySiteOnenoteSectionGroupSectionToSectionGroup(ctx context.Context, id stable.GroupIdSiteIdOnenoteSectionGroupIdSectionId, input CopySiteOnenoteSectionGroupSectionToSectionGroupRequest, options CopySiteOnenoteSectionGroupSectionToSectionGroupOperationOptions) (result CopySiteOnenoteSectionGroupSectionToSectionGroupOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/copyToSectionGroup", id.ID()),
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

	var model stable.OnenoteOperation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
