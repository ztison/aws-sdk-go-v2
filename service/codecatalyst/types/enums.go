// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CatalogActionVersionFileRecordType string

// Enum values for CatalogActionVersionFileRecordType
const (
	CatalogActionVersionFileRecordTypeLicense CatalogActionVersionFileRecordType = "LICENSE"
	CatalogActionVersionFileRecordTypeReadme  CatalogActionVersionFileRecordType = "README"
)

// Values returns all known values for CatalogActionVersionFileRecordType. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (CatalogActionVersionFileRecordType) Values() []CatalogActionVersionFileRecordType {
	return []CatalogActionVersionFileRecordType{
		"LICENSE",
		"README",
	}
}

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorEquals              ComparisonOperator = "EQ"
	ComparisonOperatorGreaterThan         ComparisonOperator = "GT"
	ComparisonOperatorGreaterThanOrEquals ComparisonOperator = "GE"
	ComparisonOperatorLessThan            ComparisonOperator = "LT"
	ComparisonOperatorLessThanOrEquals    ComparisonOperator = "LE"
)

// Values returns all known values for ComparisonOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComparisonOperator) Values() []ComparisonOperator {
	return []ComparisonOperator{
		"EQ",
		"GT",
		"GE",
		"LT",
		"LE",
	}
}

type DevEnvironmentSessionType string

// Enum values for DevEnvironmentSessionType
const (
	DevEnvironmentSessionTypeSsm DevEnvironmentSessionType = "SSM"
	DevEnvironmentSessionTypeSsh DevEnvironmentSessionType = "SSH"
)

// Values returns all known values for DevEnvironmentSessionType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DevEnvironmentSessionType) Values() []DevEnvironmentSessionType {
	return []DevEnvironmentSessionType{
		"SSM",
		"SSH",
	}
}

type DevEnvironmentStatus string

// Enum values for DevEnvironmentStatus
const (
	DevEnvironmentStatusPending  DevEnvironmentStatus = "PENDING"
	DevEnvironmentStatusRunning  DevEnvironmentStatus = "RUNNING"
	DevEnvironmentStatusStarting DevEnvironmentStatus = "STARTING"
	DevEnvironmentStatusStopping DevEnvironmentStatus = "STOPPING"
	DevEnvironmentStatusStopped  DevEnvironmentStatus = "STOPPED"
	DevEnvironmentStatusFailed   DevEnvironmentStatus = "FAILED"
	DevEnvironmentStatusDeleting DevEnvironmentStatus = "DELETING"
	DevEnvironmentStatusDeleted  DevEnvironmentStatus = "DELETED"
)

// Values returns all known values for DevEnvironmentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DevEnvironmentStatus) Values() []DevEnvironmentStatus {
	return []DevEnvironmentStatus{
		"PENDING",
		"RUNNING",
		"STARTING",
		"STOPPING",
		"STOPPED",
		"FAILED",
		"DELETING",
		"DELETED",
	}
}

type FilterKey string

// Enum values for FilterKey
const (
	FilterKeyHasAccessTo FilterKey = "hasAccessTo"
)

// Values returns all known values for FilterKey. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FilterKey) Values() []FilterKey {
	return []FilterKey{
		"hasAccessTo",
	}
}

type InstanceType string

// Enum values for InstanceType
const (
	InstanceTypeDevStandard1Small  InstanceType = "dev.standard1.small"
	InstanceTypeDevStandard1Medium InstanceType = "dev.standard1.medium"
	InstanceTypeDevStandard1Large  InstanceType = "dev.standard1.large"
	InstanceTypeDevStandard1Xlarge InstanceType = "dev.standard1.xlarge"
)

// Values returns all known values for InstanceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InstanceType) Values() []InstanceType {
	return []InstanceType{
		"dev.standard1.small",
		"dev.standard1.medium",
		"dev.standard1.large",
		"dev.standard1.xlarge",
	}
}

type OperationType string

// Enum values for OperationType
const (
	OperationTypeReadonly OperationType = "READONLY"
	OperationTypeMutation OperationType = "MUTATION"
)

// Values returns all known values for OperationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OperationType) Values() []OperationType {
	return []OperationType{
		"READONLY",
		"MUTATION",
	}
}

type UserType string

// Enum values for UserType
const (
	UserTypeUser       UserType = "USER"
	UserTypeAwsAccount UserType = "AWS_ACCOUNT"
	UserTypeUnknown    UserType = "UNKNOWN"
)

// Values returns all known values for UserType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (UserType) Values() []UserType {
	return []UserType{
		"USER",
		"AWS_ACCOUNT",
		"UNKNOWN",
	}
}
