package pendingaccessreviewinstancedecisioninsight

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPendingAccessReviewInstanceDecisionInsightOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.GovernanceInsight
}

type GetPendingAccessReviewInstanceDecisionInsightOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetPendingAccessReviewInstanceDecisionInsightOperationOptions() GetPendingAccessReviewInstanceDecisionInsightOperationOptions {
	return GetPendingAccessReviewInstanceDecisionInsightOperationOptions{}
}

func (o GetPendingAccessReviewInstanceDecisionInsightOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPendingAccessReviewInstanceDecisionInsightOperationOptions) ToOData() *odata.Query {
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

func (o GetPendingAccessReviewInstanceDecisionInsightOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPendingAccessReviewInstanceDecisionInsight - Get insights from me. Insights are recommendations to reviewers on
// whether to approve or deny a decision. There can be multiple insights associated with an
// accessReviewInstanceDecisionItem.
func (c PendingAccessReviewInstanceDecisionInsightClient) GetPendingAccessReviewInstanceDecisionInsight(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionIdInsightId, options GetPendingAccessReviewInstanceDecisionInsightOperationOptions) (result GetPendingAccessReviewInstanceDecisionInsightOperationResponse, err error) {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalGovernanceInsightImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
