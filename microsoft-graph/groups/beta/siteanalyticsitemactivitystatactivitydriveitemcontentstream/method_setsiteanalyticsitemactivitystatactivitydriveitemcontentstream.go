package siteanalyticsitemactivitystatactivitydriveitemcontentstream

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

type SetSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions() SetSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions {
	return SetSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions{}
}

func (o SetSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetSiteAnalyticsItemActivityStatActivityDriveItemContentStream - Update contentStream for the navigation property
// driveItem in groups. The content stream, if the item represents a file.
func (c SiteAnalyticsItemActivityStatActivityDriveItemContentStreamClient) SetSiteAnalyticsItemActivityStatActivityDriveItemContentStream(ctx context.Context, id beta.GroupIdSiteIdAnalyticsItemActivityStatIdActivityId, input []byte, options SetSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) (result SetSiteAnalyticsItemActivityStatActivityDriveItemContentStreamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: options.ContentType,
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/driveItem/contentStream", id.ID()),
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
