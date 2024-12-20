package accessreviewhistorydefinitioninstance

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

type CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUriOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessReviewHistoryInstance
}

type CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUriOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUriOperationOptions() CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUriOperationOptions {
	return CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUriOperationOptions{}
}

func (o CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUriOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUriOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUriOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUri - Invoke action generateDownloadUri. Generates a URI
// for an accessReviewHistoryInstance object the status for which is done. Each URI can be used to retrieve the
// instance's review history data. Each URI is valid for 24 hours and can be retrieved by fetching the downloadUri
// property from the accessReviewHistoryInstance object.
func (c AccessReviewHistoryDefinitionInstanceClient) CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUri(ctx context.Context, id beta.IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId, options CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUriOperationOptions) (result CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUriOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/generateDownloadUri", id.ID()),
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
