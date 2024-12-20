package devops

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/client/pollers"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureDevOpsReposUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *AzureDevOpsRepository
}

// AzureDevOpsReposUpdate ...
func (c DevOpsClient) AzureDevOpsReposUpdate(ctx context.Context, id ProjectRepoId, input AzureDevOpsRepository) (result AzureDevOpsReposUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPatch,
		Path:       id.ID(),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// AzureDevOpsReposUpdateThenPoll performs AzureDevOpsReposUpdate then polls until it's completed
func (c DevOpsClient) AzureDevOpsReposUpdateThenPoll(ctx context.Context, id ProjectRepoId, input AzureDevOpsRepository) error {
	result, err := c.AzureDevOpsReposUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing AzureDevOpsReposUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after AzureDevOpsReposUpdate: %+v", err)
	}

	return nil
}
