package dailyprintusagebyprinter

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/redtenant/go-azure-sdk/microsoft-graph/common-types/stable"

type PrintUsageByPrinterOperationPredicate struct {
}

func (p PrintUsageByPrinterOperationPredicate) Matches(input stable.PrintUsageByPrinter) bool {

	return true
}
