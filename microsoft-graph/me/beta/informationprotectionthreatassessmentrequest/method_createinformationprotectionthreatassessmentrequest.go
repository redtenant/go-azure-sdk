package informationprotectionthreatassessmentrequest

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

type CreateInformationProtectionThreatAssessmentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.ThreatAssessmentRequest
}

type CreateInformationProtectionThreatAssessmentRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateInformationProtectionThreatAssessmentRequestOperationOptions() CreateInformationProtectionThreatAssessmentRequestOperationOptions {
	return CreateInformationProtectionThreatAssessmentRequestOperationOptions{}
}

func (o CreateInformationProtectionThreatAssessmentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateInformationProtectionThreatAssessmentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateInformationProtectionThreatAssessmentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateInformationProtectionThreatAssessmentRequest - Create new navigation property to threatAssessmentRequests for
// me
func (c InformationProtectionThreatAssessmentRequestClient) CreateInformationProtectionThreatAssessmentRequest(ctx context.Context, input beta.ThreatAssessmentRequest, options CreateInformationProtectionThreatAssessmentRequestOperationOptions) (result CreateInformationProtectionThreatAssessmentRequestOperationResponse, err error) {
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
		Path:          "/me/informationProtection/threatAssessmentRequests",
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalThreatAssessmentRequestImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
