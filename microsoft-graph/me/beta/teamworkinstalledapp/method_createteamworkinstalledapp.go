package teamworkinstalledapp

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTeamworkInstalledAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UserScopeTeamsAppInstallation
}

type CreateTeamworkInstalledAppOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamworkInstalledAppOperationOptions() CreateTeamworkInstalledAppOperationOptions {
	return CreateTeamworkInstalledAppOperationOptions{}
}

func (o CreateTeamworkInstalledAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamworkInstalledAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamworkInstalledAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamworkInstalledApp - Create new navigation property to installedApps for me
func (c TeamworkInstalledAppClient) CreateTeamworkInstalledApp(ctx context.Context, input beta.UserScopeTeamsAppInstallation, options CreateTeamworkInstalledAppOperationOptions) (result CreateTeamworkInstalledAppOperationResponse, err error) {
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
		Path:          "/me/teamwork/installedApps",
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

	var model beta.UserScopeTeamsAppInstallation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
