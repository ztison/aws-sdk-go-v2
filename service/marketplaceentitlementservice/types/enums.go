// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type GetEntitlementFilterName string

// Enum values for GetEntitlementFilterName
const (
	GetEntitlementFilterNameCustomerIdentifier GetEntitlementFilterName = "CUSTOMER_IDENTIFIER"
	GetEntitlementFilterNameDimension          GetEntitlementFilterName = "DIMENSION"
)

// Values returns all known values for GetEntitlementFilterName. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (GetEntitlementFilterName) Values() []GetEntitlementFilterName {
	return []GetEntitlementFilterName{
		"CUSTOMER_IDENTIFIER",
		"DIMENSION",
	}
}
