package lifecycleworkflowdeleteditemworkflowuserprocessingresultsubject

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

type GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.User
}

type GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectOperationOptions() GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectOperationOptions {
	return GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectOperationOptions{}
}

func (o GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectOperationOptions) ToOData() *odata.Query {
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

func (o GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubject - Get subject from identityGovernance. The unique
// identifier of the AAD user targeted for the taskProcessingResult.Supports $filter(eq, ne) and $expand.
func (c LifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectClient) GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubject(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultId, options GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectOperationOptions) (result GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultSubjectOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/subject", id.ID()),
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

	var model stable.User
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
