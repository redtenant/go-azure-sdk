package accessreviewdefinitioninstancedecisioninstancecontactedreviewer

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAccessReviewDefinitionInstanceDecisionInstanceContactedReviewerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessReviewReviewer
}

type GetAccessReviewDefinitionInstanceDecisionInstanceContactedReviewerOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAccessReviewDefinitionInstanceDecisionInstanceContactedReviewerOperationOptions() GetAccessReviewDefinitionInstanceDecisionInstanceContactedReviewerOperationOptions {
	return GetAccessReviewDefinitionInstanceDecisionInstanceContactedReviewerOperationOptions{}
}

func (o GetAccessReviewDefinitionInstanceDecisionInstanceContactedReviewerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAccessReviewDefinitionInstanceDecisionInstanceContactedReviewerOperationOptions) ToOData() *odata.Query {
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

func (o GetAccessReviewDefinitionInstanceDecisionInstanceContactedReviewerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAccessReviewDefinitionInstanceDecisionInstanceContactedReviewer - Get contactedReviewers from identityGovernance.
// Returns the collection of reviewers who were contacted to complete this review. While the reviewers and
// fallbackReviewers properties of the accessReviewScheduleDefinition might specify group owners or managers as
// reviewers, contactedReviewers returns their individual identities. Supports $select. Read-only.
func (c AccessReviewDefinitionInstanceDecisionInstanceContactedReviewerClient) GetAccessReviewDefinitionInstanceDecisionInstanceContactedReviewer(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId, options GetAccessReviewDefinitionInstanceDecisionInstanceContactedReviewerOperationOptions) (result GetAccessReviewDefinitionInstanceDecisionInstanceContactedReviewerOperationResponse, err error) {
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

	var model beta.AccessReviewReviewer
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
