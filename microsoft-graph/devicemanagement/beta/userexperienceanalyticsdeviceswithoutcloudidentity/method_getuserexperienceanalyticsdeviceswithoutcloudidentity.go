package userexperienceanalyticsdeviceswithoutcloudidentity

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserExperienceAnalyticsDeviceWithoutCloudIdentity
}

type GetUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions() GetUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions {
	return GetUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions{}
}

func (o GetUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetUserExperienceAnalyticsDevicesWithoutCloudIdentity - Get userExperienceAnalyticsDevicesWithoutCloudIdentity from
// deviceManagement. User experience analytics devices without cloud identity.
func (c UserExperienceAnalyticsDevicesWithoutCloudIdentityClient) GetUserExperienceAnalyticsDevicesWithoutCloudIdentity(ctx context.Context, id beta.DeviceManagementUserExperienceAnalyticsDevicesWithoutCloudIdentityId, options GetUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationOptions) (result GetUserExperienceAnalyticsDevicesWithoutCloudIdentityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.UserExperienceAnalyticsDeviceWithoutCloudIdentity
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
