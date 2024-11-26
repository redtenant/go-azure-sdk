package pendingaccessreviewinstancedecisioninstance

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

type SendPendingAccessReviewInstanceDecisionInstanceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SendPendingAccessReviewInstanceDecisionInstanceReminderOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSendPendingAccessReviewInstanceDecisionInstanceReminderOperationOptions() SendPendingAccessReviewInstanceDecisionInstanceReminderOperationOptions {
	return SendPendingAccessReviewInstanceDecisionInstanceReminderOperationOptions{}
}

func (o SendPendingAccessReviewInstanceDecisionInstanceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SendPendingAccessReviewInstanceDecisionInstanceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SendPendingAccessReviewInstanceDecisionInstanceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SendPendingAccessReviewInstanceDecisionInstanceReminder - Invoke action sendReminder. Send a reminder to the
// reviewers of a currently active accessReviewInstance.
func (c PendingAccessReviewInstanceDecisionInstanceClient) SendPendingAccessReviewInstanceDecisionInstanceReminder(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionId, options SendPendingAccessReviewInstanceDecisionInstanceReminderOperationOptions) (result SendPendingAccessReviewInstanceDecisionInstanceReminderOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/instance/sendReminder", id.ID()),
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
