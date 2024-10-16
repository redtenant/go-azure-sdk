package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2cUserFlowId{}

// IdentityB2cUserFlowId is a struct representing the Resource ID for a Identity B 2c User Flow
type IdentityB2cUserFlowId struct {
	B2cIdentityUserFlowId string
}

// NewIdentityB2cUserFlowID returns a new IdentityB2cUserFlowId struct
func NewIdentityB2cUserFlowID(b2cIdentityUserFlowId string) IdentityB2cUserFlowId {
	return IdentityB2cUserFlowId{
		B2cIdentityUserFlowId: b2cIdentityUserFlowId,
	}
}

// ParseIdentityB2cUserFlowID parses 'input' into a IdentityB2cUserFlowId
func ParseIdentityB2cUserFlowID(input string) (*IdentityB2cUserFlowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2cUserFlowIDInsensitively parses 'input' case-insensitively into a IdentityB2cUserFlowId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2cUserFlowIDInsensitively(input string) (*IdentityB2cUserFlowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2cUserFlowId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.B2cIdentityUserFlowId, ok = input.Parsed["b2cIdentityUserFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "b2cIdentityUserFlowId", input)
	}

	return nil
}

// ValidateIdentityB2cUserFlowID checks that 'input' can be parsed as a Identity B 2c User Flow ID
func ValidateIdentityB2cUserFlowID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2cUserFlowID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2c User Flow ID
func (id IdentityB2cUserFlowId) ID() string {
	fmtString := "/identity/b2cUserFlows/%s"
	return fmt.Sprintf(fmtString, id.B2cIdentityUserFlowId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2c User Flow ID
func (id IdentityB2cUserFlowId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2cUserFlows", "b2cUserFlows", "b2cUserFlows"),
		resourceids.UserSpecifiedSegment("b2cIdentityUserFlowId", "b2cIdentityUserFlowId"),
	}
}

// String returns a human-readable description of this Identity B 2c User Flow ID
func (id IdentityB2cUserFlowId) String() string {
	components := []string{
		fmt.Sprintf("B 2c Identity User Flow: %q", id.B2cIdentityUserFlowId),
	}
	return fmt.Sprintf("Identity B 2c User Flow (%s)", strings.Join(components, "\n"))
}
