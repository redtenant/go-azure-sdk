package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2cUserFlowIdIdentityProviderId{}

// IdentityB2cUserFlowIdIdentityProviderId is a struct representing the Resource ID for a Identity B 2c User Flow Id Identity Provider
type IdentityB2cUserFlowIdIdentityProviderId struct {
	B2cIdentityUserFlowId string
	IdentityProviderId    string
}

// NewIdentityB2cUserFlowIdIdentityProviderID returns a new IdentityB2cUserFlowIdIdentityProviderId struct
func NewIdentityB2cUserFlowIdIdentityProviderID(b2cIdentityUserFlowId string, identityProviderId string) IdentityB2cUserFlowIdIdentityProviderId {
	return IdentityB2cUserFlowIdIdentityProviderId{
		B2cIdentityUserFlowId: b2cIdentityUserFlowId,
		IdentityProviderId:    identityProviderId,
	}
}

// ParseIdentityB2cUserFlowIdIdentityProviderID parses 'input' into a IdentityB2cUserFlowIdIdentityProviderId
func ParseIdentityB2cUserFlowIdIdentityProviderID(input string) (*IdentityB2cUserFlowIdIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowIdIdentityProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowIdIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2cUserFlowIdIdentityProviderIDInsensitively parses 'input' case-insensitively into a IdentityB2cUserFlowIdIdentityProviderId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2cUserFlowIdIdentityProviderIDInsensitively(input string) (*IdentityB2cUserFlowIdIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowIdIdentityProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowIdIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2cUserFlowIdIdentityProviderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.B2cIdentityUserFlowId, ok = input.Parsed["b2cIdentityUserFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "b2cIdentityUserFlowId", input)
	}

	if id.IdentityProviderId, ok = input.Parsed["identityProviderId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityProviderId", input)
	}

	return nil
}

// ValidateIdentityB2cUserFlowIdIdentityProviderID checks that 'input' can be parsed as a Identity B 2c User Flow Id Identity Provider ID
func ValidateIdentityB2cUserFlowIdIdentityProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2cUserFlowIdIdentityProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2c User Flow Id Identity Provider ID
func (id IdentityB2cUserFlowIdIdentityProviderId) ID() string {
	fmtString := "/identity/b2cUserFlows/%s/identityProviders/%s"
	return fmt.Sprintf(fmtString, id.B2cIdentityUserFlowId, id.IdentityProviderId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2c User Flow Id Identity Provider ID
func (id IdentityB2cUserFlowIdIdentityProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2cUserFlows", "b2cUserFlows", "b2cUserFlows"),
		resourceids.UserSpecifiedSegment("b2cIdentityUserFlowId", "b2cIdentityUserFlowId"),
		resourceids.StaticSegment("identityProviders", "identityProviders", "identityProviders"),
		resourceids.UserSpecifiedSegment("identityProviderId", "identityProviderId"),
	}
}

// String returns a human-readable description of this Identity B 2c User Flow Id Identity Provider ID
func (id IdentityB2cUserFlowIdIdentityProviderId) String() string {
	components := []string{
		fmt.Sprintf("B 2c Identity User Flow: %q", id.B2cIdentityUserFlowId),
		fmt.Sprintf("Identity Provider: %q", id.IdentityProviderId),
	}
	return fmt.Sprintf("Identity B 2c User Flow Id Identity Provider (%s)", strings.Join(components, "\n"))
}
