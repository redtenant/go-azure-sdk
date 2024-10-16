package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityB2cUserFlowIdUserFlowIdentityProviderId{}

// IdentityB2cUserFlowIdUserFlowIdentityProviderId is a struct representing the Resource ID for a Identity B 2c User Flow Id User Flow Identity Provider
type IdentityB2cUserFlowIdUserFlowIdentityProviderId struct {
	B2cIdentityUserFlowId  string
	IdentityProviderBaseId string
}

// NewIdentityB2cUserFlowIdUserFlowIdentityProviderID returns a new IdentityB2cUserFlowIdUserFlowIdentityProviderId struct
func NewIdentityB2cUserFlowIdUserFlowIdentityProviderID(b2cIdentityUserFlowId string, identityProviderBaseId string) IdentityB2cUserFlowIdUserFlowIdentityProviderId {
	return IdentityB2cUserFlowIdUserFlowIdentityProviderId{
		B2cIdentityUserFlowId:  b2cIdentityUserFlowId,
		IdentityProviderBaseId: identityProviderBaseId,
	}
}

// ParseIdentityB2cUserFlowIdUserFlowIdentityProviderID parses 'input' into a IdentityB2cUserFlowIdUserFlowIdentityProviderId
func ParseIdentityB2cUserFlowIdUserFlowIdentityProviderID(input string) (*IdentityB2cUserFlowIdUserFlowIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowIdUserFlowIdentityProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowIdUserFlowIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityB2cUserFlowIdUserFlowIdentityProviderIDInsensitively parses 'input' case-insensitively into a IdentityB2cUserFlowIdUserFlowIdentityProviderId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2cUserFlowIdUserFlowIdentityProviderIDInsensitively(input string) (*IdentityB2cUserFlowIdUserFlowIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityB2cUserFlowIdUserFlowIdentityProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityB2cUserFlowIdUserFlowIdentityProviderId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityB2cUserFlowIdUserFlowIdentityProviderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.B2cIdentityUserFlowId, ok = input.Parsed["b2cIdentityUserFlowId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "b2cIdentityUserFlowId", input)
	}

	if id.IdentityProviderBaseId, ok = input.Parsed["identityProviderBaseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "identityProviderBaseId", input)
	}

	return nil
}

// ValidateIdentityB2cUserFlowIdUserFlowIdentityProviderID checks that 'input' can be parsed as a Identity B 2c User Flow Id User Flow Identity Provider ID
func ValidateIdentityB2cUserFlowIdUserFlowIdentityProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2cUserFlowIdUserFlowIdentityProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2c User Flow Id User Flow Identity Provider ID
func (id IdentityB2cUserFlowIdUserFlowIdentityProviderId) ID() string {
	fmtString := "/identity/b2cUserFlows/%s/userFlowIdentityProviders/%s"
	return fmt.Sprintf(fmtString, id.B2cIdentityUserFlowId, id.IdentityProviderBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2c User Flow Id User Flow Identity Provider ID
func (id IdentityB2cUserFlowIdUserFlowIdentityProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("b2cUserFlows", "b2cUserFlows", "b2cUserFlows"),
		resourceids.UserSpecifiedSegment("b2cIdentityUserFlowId", "b2cIdentityUserFlowId"),
		resourceids.StaticSegment("userFlowIdentityProviders", "userFlowIdentityProviders", "userFlowIdentityProviders"),
		resourceids.UserSpecifiedSegment("identityProviderBaseId", "identityProviderBaseId"),
	}
}

// String returns a human-readable description of this Identity B 2c User Flow Id User Flow Identity Provider ID
func (id IdentityB2cUserFlowIdUserFlowIdentityProviderId) String() string {
	components := []string{
		fmt.Sprintf("B 2c Identity User Flow: %q", id.B2cIdentityUserFlowId),
		fmt.Sprintf("Identity Provider Base: %q", id.IdentityProviderBaseId),
	}
	return fmt.Sprintf("Identity B 2c User Flow Id User Flow Identity Provider (%s)", strings.Join(components, "\n"))
}
