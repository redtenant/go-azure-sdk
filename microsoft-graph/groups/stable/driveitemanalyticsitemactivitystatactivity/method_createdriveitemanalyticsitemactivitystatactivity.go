package driveitemanalyticsitemactivitystatactivity

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

type CreateDriveItemAnalyticsItemActivityStatActivityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ItemActivity
}

type CreateDriveItemAnalyticsItemActivityStatActivityOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDriveItemAnalyticsItemActivityStatActivityOperationOptions() CreateDriveItemAnalyticsItemActivityStatActivityOperationOptions {
	return CreateDriveItemAnalyticsItemActivityStatActivityOperationOptions{}
}

func (o CreateDriveItemAnalyticsItemActivityStatActivityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDriveItemAnalyticsItemActivityStatActivityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDriveItemAnalyticsItemActivityStatActivityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDriveItemAnalyticsItemActivityStatActivity - Create new navigation property to activities for groups
func (c DriveItemAnalyticsItemActivityStatActivityClient) CreateDriveItemAnalyticsItemActivityStatActivity(ctx context.Context, id stable.GroupIdDriveIdItemIdAnalyticsItemActivityStatId, input stable.ItemActivity, options CreateDriveItemAnalyticsItemActivityStatActivityOperationOptions) (result CreateDriveItemAnalyticsItemActivityStatActivityOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/activities", id.ID()),
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

	var model stable.ItemActivity
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
