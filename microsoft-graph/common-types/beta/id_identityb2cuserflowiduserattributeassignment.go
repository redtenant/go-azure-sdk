package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2cUserFlowIdUserAttributeAssignmentId{}

// IdentityB2cUserFlowIdUserAttributeAssignmentId is a struct representing the Resource ID for a Identity B 2c User Flow Id User Attribute Assignment
type IdentityB2cUserFlowIdUserAttributeAssignmentId struct {
	B2cIdentityUserFlowId                 string
	IdentityUserFlowAttributeAssignmentId string
}

// NewIdentityB2cUserFlowIdUserAttributeAssignmentID returns a new IdentityB2cUserFlowIdUserAttributeAssignmentId struct
func NewIdentityB2cUserFlowIdUserAttributeAssignmentID(b2cIdentityUserFlowId string, identityUserFlowAttributeAssignmentId string) IdentityB2cUserFlowIdUserAttributeAssignmentId {
	return IdentityB2cUserFlowIdUserAttributeAssignmentId{
		B2cIdentityUserFlowId:                 b2cIdentityUserFlowId,
		IdentityUserFlowAttributeAssignmentId: identityUserFlowAttributeAssignmentId,
	}
}

// ParseIdentityB2cUserFlowIdUserAttributeAssignmentID parses 'input' into a IdentityB2cUserFlowIdUserAttributeAssignmentId
func ParseIdentityB2cUserFlowIdUserAttributeAssignmentID(input string) (*IdentityB2cUserFlowIdUserAttributeAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowIdUserAttributeAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowIdUserAttributeAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2cUserFlowIdUserAttributeAssignmentIDInsensitively parses 'input' case-insensitively into a IdentityB2cUserFlowIdUserAttributeAssignmentId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2cUserFlowIdUserAttributeAssignmentIDInsensitively(input string) (*IdentityB2cUserFlowIdUserAttributeAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowIdUserAttributeAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowIdUserAttributeAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2cUserFlowIdUserAttributeAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.B2cIdentityUserFlowId, ok = input.Parsed["b2cIdentityUserFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "b2cIdentityUserFlowId", input)
	}

	if id.IdentityUserFlowAttributeAssignmentId, ok = input.Parsed["identityUserFlowAttributeAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityUserFlowAttributeAssignmentId", input)
	}

	return nil
}

// ValidateIdentityB2cUserFlowIdUserAttributeAssignmentID checks that 'input' can be parsed as a Identity B 2c User Flow Id User Attribute Assignment ID
func ValidateIdentityB2cUserFlowIdUserAttributeAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2cUserFlowIdUserAttributeAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2c User Flow Id User Attribute Assignment ID
func (id IdentityB2cUserFlowIdUserAttributeAssignmentId) ID() string {
	fmtString := "/identity/b2cUserFlows/%s/userAttributeAssignments/%s"
	return fmt.Sprintf(fmtString, id.B2cIdentityUserFlowId, id.IdentityUserFlowAttributeAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2c User Flow Id User Attribute Assignment ID
func (id IdentityB2cUserFlowIdUserAttributeAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2cUserFlows", "b2cUserFlows", "b2cUserFlows"),
		resourceids.UserSpecifiedSegment("b2cIdentityUserFlowId", "b2cIdentityUserFlowId"),
		resourceids.StaticSegment("userAttributeAssignments", "userAttributeAssignments", "userAttributeAssignments"),
		resourceids.UserSpecifiedSegment("identityUserFlowAttributeAssignmentId", "identityUserFlowAttributeAssignmentId"),
	}
}

// String returns a human-readable description of this Identity B 2c User Flow Id User Attribute Assignment ID
func (id IdentityB2cUserFlowIdUserAttributeAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("B 2c Identity User Flow: %q", id.B2cIdentityUserFlowId),
		fmt.Sprintf("Identity User Flow Attribute Assignment: %q", id.IdentityUserFlowAttributeAssignmentId),
	}
	return fmt.Sprintf("Identity B 2c User Flow Id User Attribute Assignment (%s)", strings.Join(components, "\n"))
}
