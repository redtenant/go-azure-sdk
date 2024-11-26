package policystates

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/client/pollers"
	"github.com/redtenant/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerResourceGroupEvaluationOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// TriggerResourceGroupEvaluation ...
func (c PolicyStatesClient) TriggerResourceGroupEvaluation(ctx context.Context, id commonids.ResourceGroupId) (result TriggerResourceGroupEvaluationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/providers/Microsoft.PolicyInsights/policyStates/latest/triggerEvaluation", id.ID()),
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

// TriggerResourceGroupEvaluationThenPoll performs TriggerResourceGroupEvaluation then polls until it's completed
func (c PolicyStatesClient) TriggerResourceGroupEvaluationThenPoll(ctx context.Context, id commonids.ResourceGroupId) error {
	result, err := c.TriggerResourceGroupEvaluation(ctx, id)
	if err != nil {
		return fmt.Errorf("performing TriggerResourceGroupEvaluation: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after TriggerResourceGroupEvaluation: %+v", err)
	}

	return nil
}
