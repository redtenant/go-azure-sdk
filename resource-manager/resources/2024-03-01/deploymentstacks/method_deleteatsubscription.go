package deploymentstacks

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

type DeleteAtSubscriptionOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteAtSubscriptionOperationOptions struct {
	BypassStackOutOfSyncError      *bool
	UnmanageActionManagementGroups *UnmanageActionManagementGroupMode
	UnmanageActionResourceGroups   *UnmanageActionResourceGroupMode
	UnmanageActionResources        *UnmanageActionResourceMode
}

func DefaultDeleteAtSubscriptionOperationOptions() DeleteAtSubscriptionOperationOptions {
	return DeleteAtSubscriptionOperationOptions{}
}

func (o DeleteAtSubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DeleteAtSubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteAtSubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.BypassStackOutOfSyncError != nil {
		out.Append("bypassStackOutOfSyncError", fmt.Sprintf("%v", *o.BypassStackOutOfSyncError))
	}
	if o.UnmanageActionManagementGroups != nil {
		out.Append("unmanageAction.ManagementGroups", fmt.Sprintf("%v", *o.UnmanageActionManagementGroups))
	}
	if o.UnmanageActionResourceGroups != nil {
		out.Append("unmanageAction.ResourceGroups", fmt.Sprintf("%v", *o.UnmanageActionResourceGroups))
	}
	if o.UnmanageActionResources != nil {
		out.Append("unmanageAction.Resources", fmt.Sprintf("%v", *o.UnmanageActionResources))
	}
	return &out
}

// DeleteAtSubscription ...
func (c DeploymentStacksClient) DeleteAtSubscription(ctx context.Context, id DeploymentStackId, options DeleteAtSubscriptionOperationOptions) (result DeleteAtSubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// DeleteAtSubscriptionThenPoll performs DeleteAtSubscription then polls until it's completed
func (c DeploymentStacksClient) DeleteAtSubscriptionThenPoll(ctx context.Context, id DeploymentStackId, options DeleteAtSubscriptionOperationOptions) error {
	result, err := c.DeleteAtSubscription(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing DeleteAtSubscription: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DeleteAtSubscription: %+v", err)
	}

	return nil
}
