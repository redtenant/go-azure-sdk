package accessreviewdefinitioninstancestagedecisioninstancedecision

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

type CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecisionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecisionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecisionOperationOptions() CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecisionOperationOptions {
	return CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecisionOperationOptions{}
}

func (o CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecisionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecisionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecisionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecision - Invoke action
// recordAllDecisions. As a reviewer of an access review, record a decision for an accessReviewInstanceDecisionItem that
// is assigned to you and that matches the principal or resource IDs specified. If no IDs are specified, the decisions
// will apply to every accessReviewInstanceDecisionItem for which you are the reviewer.
func (c AccessReviewDefinitionInstanceStageDecisionInstanceDecisionClient) CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecision(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId, input CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecisionRequest, options CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecisionOperationOptions) (result CreateAccessReviewDefinitionInstanceStageDecisionInstanceDecisionRecordAllDecisionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/instance/decisions/recordAllDecisions", id.ID()),
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

	return
}
