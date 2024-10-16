package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityUserFlowId{}

// IdentityUserFlowId is a struct representing the Resource ID for a Identity User Flow
type IdentityUserFlowId struct {
	IdentityUserFlowId string
}

// NewIdentityUserFlowID returns a new IdentityUserFlowId struct
func NewIdentityUserFlowID(identityUserFlowId string) IdentityUserFlowId {
	return IdentityUserFlowId{
		IdentityUserFlowId: identityUserFlowId,
	}
}

// ParseIdentityUserFlowID parses 'input' into a IdentityUserFlowId
func ParseIdentityUserFlowID(input string) (*IdentityUserFlowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityUserFlowId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityUserFlowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityUserFlowIDInsensitively parses 'input' case-insensitively into a IdentityUserFlowId
// note: this method should only be used for API response data and not user input
func ParseIdentityUserFlowIDInsensitively(input string) (*IdentityUserFlowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityUserFlowId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityUserFlowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityUserFlowId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.IdentityUserFlowId, ok = input.Parsed["identityUserFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityUserFlowId", input)
	}

	return nil
}

// ValidateIdentityUserFlowID checks that 'input' can be parsed as a Identity User Flow ID
func ValidateIdentityUserFlowID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityUserFlowID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity User Flow ID
func (id IdentityUserFlowId) ID() string {
	fmtString := "/identity/userFlows/%s"
	return fmt.Sprintf(fmtString, id.IdentityUserFlowId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity User Flow ID
func (id IdentityUserFlowId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("userFlows", "userFlows", "userFlows"),
		resourceids.UserSpecifiedSegment("identityUserFlowId", "identityUserFlowId"),
	}
}

// String returns a human-readable description of this Identity User Flow ID
func (id IdentityUserFlowId) String() string {
	components := []string{
		fmt.Sprintf("Identity User Flow: %q", id.IdentityUserFlowId),
	}
	return fmt.Sprintf("Identity User Flow (%s)", strings.Join(components, "\n"))
}
