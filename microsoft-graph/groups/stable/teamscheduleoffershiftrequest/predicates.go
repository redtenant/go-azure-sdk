package teamscheduleoffershiftrequest

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"

type OfferShiftRequestOperationPredicate struct {
}

func (p OfferShiftRequestOperationPredicate) Matches(input stable.OfferShiftRequest) bool {

	return true
}
