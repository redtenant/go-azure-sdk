package userexperienceanalyticsapphealthoverview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteUserExperienceAnalyticsAppHealthOverviewOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteUserExperienceAnalyticsAppHealthOverviewOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteUserExperienceAnalyticsAppHealthOverviewOperationOptions() DeleteUserExperienceAnalyticsAppHealthOverviewOperationOptions {
	return DeleteUserExperienceAnalyticsAppHealthOverviewOperationOptions{}
}

func (o DeleteUserExperienceAnalyticsAppHealthOverviewOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteUserExperienceAnalyticsAppHealthOverviewOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteUserExperienceAnalyticsAppHealthOverviewOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteUserExperienceAnalyticsAppHealthOverview - Delete navigation property userExperienceAnalyticsAppHealthOverview
// for deviceManagement
func (c UserExperienceAnalyticsAppHealthOverviewClient) DeleteUserExperienceAnalyticsAppHealthOverview(ctx context.Context, options DeleteUserExperienceAnalyticsAppHealthOverviewOperationOptions) (result DeleteUserExperienceAnalyticsAppHealthOverviewOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsAppHealthOverview",
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

	return
}
