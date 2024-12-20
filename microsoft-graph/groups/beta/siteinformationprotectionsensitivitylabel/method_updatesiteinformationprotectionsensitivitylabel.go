package siteinformationprotectionsensitivitylabel

import (
	"context"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSiteInformationProtectionSensitivityLabelOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSiteInformationProtectionSensitivityLabelOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateSiteInformationProtectionSensitivityLabelOperationOptions() UpdateSiteInformationProtectionSensitivityLabelOperationOptions {
	return UpdateSiteInformationProtectionSensitivityLabelOperationOptions{}
}

func (o UpdateSiteInformationProtectionSensitivityLabelOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSiteInformationProtectionSensitivityLabelOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSiteInformationProtectionSensitivityLabelOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSiteInformationProtectionSensitivityLabel - Update the navigation property sensitivityLabels in groups
func (c SiteInformationProtectionSensitivityLabelClient) UpdateSiteInformationProtectionSensitivityLabel(ctx context.Context, id beta.GroupIdSiteIdInformationProtectionSensitivityLabelId, input beta.SensitivityLabel, options UpdateSiteInformationProtectionSensitivityLabelOperationOptions) (result UpdateSiteInformationProtectionSensitivityLabelOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
