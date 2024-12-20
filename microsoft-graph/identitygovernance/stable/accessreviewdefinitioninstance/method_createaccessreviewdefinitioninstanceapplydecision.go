package accessreviewdefinitioninstance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAccessReviewDefinitionInstanceApplyDecisionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateAccessReviewDefinitionInstanceApplyDecisionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAccessReviewDefinitionInstanceApplyDecisionOperationOptions() CreateAccessReviewDefinitionInstanceApplyDecisionOperationOptions {
	return CreateAccessReviewDefinitionInstanceApplyDecisionOperationOptions{}
}

func (o CreateAccessReviewDefinitionInstanceApplyDecisionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAccessReviewDefinitionInstanceApplyDecisionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAccessReviewDefinitionInstanceApplyDecisionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAccessReviewDefinitionInstanceApplyDecision - Invoke action applyDecisions. Apply review decisions on an
// accessReviewInstance if the decisions were not applied automatically because the autoApplyDecisionsEnabled property
// is false in the review's accessReviewScheduleSettings. The status of the accessReviewInstance must be Completed to
// call this method.
func (c AccessReviewDefinitionInstanceClient) CreateAccessReviewDefinitionInstanceApplyDecision(ctx context.Context, id stable.IdentityGovernanceAccessReviewDefinitionIdInstanceId, options CreateAccessReviewDefinitionInstanceApplyDecisionOperationOptions) (result CreateAccessReviewDefinitionInstanceApplyDecisionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/applyDecisions", id.ID()),
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
