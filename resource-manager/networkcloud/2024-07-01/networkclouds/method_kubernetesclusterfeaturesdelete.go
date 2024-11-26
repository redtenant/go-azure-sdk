package networkclouds

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

type KubernetesClusterFeaturesDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// KubernetesClusterFeaturesDelete ...
func (c NetworkcloudsClient) KubernetesClusterFeaturesDelete(ctx context.Context, id FeatureId) (result KubernetesClusterFeaturesDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
		},
		HttpMethod: http.MethodDelete,
		Path:       id.ID(),
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

// KubernetesClusterFeaturesDeleteThenPoll performs KubernetesClusterFeaturesDelete then polls until it's completed
func (c NetworkcloudsClient) KubernetesClusterFeaturesDeleteThenPoll(ctx context.Context, id FeatureId) error {
	result, err := c.KubernetesClusterFeaturesDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing KubernetesClusterFeaturesDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after KubernetesClusterFeaturesDelete: %+v", err)
	}

	return nil
}
