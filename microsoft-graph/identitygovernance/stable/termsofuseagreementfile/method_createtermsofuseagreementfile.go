package termsofuseagreementfile

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

type CreateTermsOfUseAgreementFileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AgreementFileLocalization
}

type CreateTermsOfUseAgreementFileOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTermsOfUseAgreementFileOperationOptions() CreateTermsOfUseAgreementFileOperationOptions {
	return CreateTermsOfUseAgreementFileOperationOptions{}
}

func (o CreateTermsOfUseAgreementFileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTermsOfUseAgreementFileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTermsOfUseAgreementFileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTermsOfUseAgreementFile - Create agreementFileLocalization. Create a new localized agreement file.
func (c TermsOfUseAgreementFileClient) CreateTermsOfUseAgreementFile(ctx context.Context, id stable.IdentityGovernanceTermsOfUseAgreementId, input stable.AgreementFileLocalization, options CreateTermsOfUseAgreementFileOperationOptions) (result CreateTermsOfUseAgreementFileOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/files", id.ID()),
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

	var model stable.AgreementFileLocalization
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
