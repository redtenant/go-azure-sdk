package notebookworkspacesresource

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

type NotebookWorkspacesCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *NotebookWorkspace
}

// NotebookWorkspacesCreateOrUpdate ...
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesCreateOrUpdate(ctx context.Context, id DatabaseAccountId, input ARMProxyResource) (result NotebookWorkspacesCreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/notebookWorkspaces/default", id.ID()),
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

// NotebookWorkspacesCreateOrUpdateThenPoll performs NotebookWorkspacesCreateOrUpdate then polls until it's completed
func (c NotebookWorkspacesResourceClient) NotebookWorkspacesCreateOrUpdateThenPoll(ctx context.Context, id DatabaseAccountId, input ARMProxyResource) error {
	result, err := c.NotebookWorkspacesCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing NotebookWorkspacesCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after NotebookWorkspacesCreateOrUpdate: %+v", err)
	}

	return nil
}
