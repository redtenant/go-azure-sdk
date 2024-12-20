package accessreviewhistorydefinitioninstance

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAccessReviewHistoryDefinitionInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessReviewHistoryInstance
}

type GetAccessReviewHistoryDefinitionInstanceOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAccessReviewHistoryDefinitionInstanceOperationOptions() GetAccessReviewHistoryDefinitionInstanceOperationOptions {
	return GetAccessReviewHistoryDefinitionInstanceOperationOptions{}
}

func (o GetAccessReviewHistoryDefinitionInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAccessReviewHistoryDefinitionInstanceOperationOptions) ToOData() *odata.Query {
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

func (o GetAccessReviewHistoryDefinitionInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAccessReviewHistoryDefinitionInstance - Get instances from identityGovernance. If the
// accessReviewHistoryDefinition is a recurring definition, instances represent each recurrence. A definition that
// doesn't recur will have exactly one instance.
func (c AccessReviewHistoryDefinitionInstanceClient) GetAccessReviewHistoryDefinitionInstance(ctx context.Context, id beta.IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId, options GetAccessReviewHistoryDefinitionInstanceOperationOptions) (result GetAccessReviewHistoryDefinitionInstanceOperationResponse, err error) {
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

	var model beta.AccessReviewHistoryInstance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
