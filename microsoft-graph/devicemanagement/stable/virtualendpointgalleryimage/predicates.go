package virtualendpointgalleryimage

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"

type CloudPCGalleryImageOperationPredicate struct {
}

func (p CloudPCGalleryImageOperationPredicate) Matches(input stable.CloudPCGalleryImage) bool {

	return true
}
