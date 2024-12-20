package onlinemeetingregistrationregistrant

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/redtenant/go-azure-sdk/sdk/client"
	"github.com/redtenant/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOnlineMeetingRegistrationRegistrantOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.MeetingRegistrantBase
}

type CreateOnlineMeetingRegistrationRegistrantOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateOnlineMeetingRegistrationRegistrantOperationOptions() CreateOnlineMeetingRegistrationRegistrantOperationOptions {
	return CreateOnlineMeetingRegistrationRegistrantOperationOptions{}
}

func (o CreateOnlineMeetingRegistrationRegistrantOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOnlineMeetingRegistrationRegistrantOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOnlineMeetingRegistrationRegistrantOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOnlineMeetingRegistrationRegistrant - Create externalMeetingRegistrant (deprecated). Enroll an
// externalMeetingRegistrant in an online meeting which has externalMeetingRegistration enabled. The meeting organizer
// enrolls someone by providing a unique id in the external registration system and gets the unique joinWebUrl of this
// registrant.
func (c OnlineMeetingRegistrationRegistrantClient) CreateOnlineMeetingRegistrationRegistrant(ctx context.Context, id beta.MeOnlineMeetingId, input beta.MeetingRegistrantBase, options CreateOnlineMeetingRegistrationRegistrantOperationOptions) (result CreateOnlineMeetingRegistrationRegistrantOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/registration/registrants", id.ID()),
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
	model, err := beta.UnmarshalMeetingRegistrantBaseImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
