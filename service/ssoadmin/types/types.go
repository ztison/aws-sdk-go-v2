// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// These are IAM Identity Center identity store attributes that you can configure
// for use in attributes-based access control (ABAC). You can create permissions
// policies that determine who can access your AWS resources based upon the
// configured attribute values. When you enable ABAC and specify
// AccessControlAttributes , IAM Identity Center passes the attribute values of the
// authenticated user into IAM for use in policy evaluation.
type AccessControlAttribute struct {

	// The name of the attribute associated with your identities in your identity
	// source. This is used to map a specified attribute in your identity source with
	// an attribute in IAM Identity Center.
	//
	// This member is required.
	Key *string

	// The value used for mapping a specified attribute to an identity source.
	//
	// This member is required.
	Value *AccessControlAttributeValue

	noSmithyDocumentSerde
}

// The value used for mapping a specified attribute to an identity source. For
// more information, see Attribute mappings (https://docs.aws.amazon.com/singlesignon/latest/userguide/attributemappingsconcept.html)
// in the IAM Identity Center User Guide.
type AccessControlAttributeValue struct {

	// The identity source to use when mapping a specified attribute to IAM Identity
	// Center.
	//
	// This member is required.
	Source []string

	noSmithyDocumentSerde
}

// The assignment that indicates a principal's limited access to a specified AWS
// account with a specified permission set. The term principal here refers to a
// user or group that is defined in IAM Identity Center.
type AccountAssignment struct {

	// The identifier of the AWS account.
	AccountId *string

	// The ARN of the permission set. For more information about ARNs, see Amazon
	// Resource Names (ARNs) and AWS Service Namespaces in the AWS General Reference.
	PermissionSetArn *string

	// An identifier for an object in IAM Identity Center, such as a user or group.
	// PrincipalIds are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6). For
	// more information about PrincipalIds in IAM Identity Center, see the IAM
	// Identity Center Identity Store API Reference .
	PrincipalId *string

	// The entity type for which the assignment will be created.
	PrincipalType PrincipalType

	noSmithyDocumentSerde
}

// The status of the creation or deletion operation of an assignment that a
// principal needs to access an account.
type AccountAssignmentOperationStatus struct {

	// The date that the permission set was created.
	CreatedDate *time.Time

	// The message that contains an error or exception in case of an operation failure.
	FailureReason *string

	// The ARN of the permission set. For more information about ARNs, see Amazon
	// Resource Names (ARNs) and AWS Service Namespaces in the AWS General Reference.
	PermissionSetArn *string

	// An identifier for an object in IAM Identity Center, such as a user or group.
	// PrincipalIds are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6). For
	// more information about PrincipalIds in IAM Identity Center, see the IAM
	// Identity Center Identity Store API Reference .
	PrincipalId *string

	// The entity type for which the assignment will be created.
	PrincipalType PrincipalType

	// The identifier for tracking the request operation that is generated by the
	// universally unique identifier (UUID) workflow.
	RequestId *string

	// The status of the permission set provisioning process.
	Status StatusValues

	// TargetID is an AWS account identifier, typically a 10-12 digit string (For
	// example, 123456789012).
	TargetId *string

	// The entity type for which the assignment will be created.
	TargetType TargetType

	noSmithyDocumentSerde
}

// Provides information about the AccountAssignment creation request.
type AccountAssignmentOperationStatusMetadata struct {

	// The date that the permission set was created.
	CreatedDate *time.Time

	// The identifier for tracking the request operation that is generated by the
	// universally unique identifier (UUID) workflow.
	RequestId *string

	// The status of the permission set provisioning process.
	Status StatusValues

	noSmithyDocumentSerde
}

// A structure that stores the details of the AWS managed policy.
type AttachedManagedPolicy struct {

	// The ARN of the AWS managed policy. For more information about ARNs, see Amazon
	// Resource Names (ARNs) and AWS Service Namespaces in the AWS General Reference.
	Arn *string

	// The name of the AWS managed policy.
	Name *string

	noSmithyDocumentSerde
}

// Specifies the name and path of a customer managed policy. You must have an IAM
// policy that matches the name and path in each AWS account where you want to
// deploy your permission set.
type CustomerManagedPolicyReference struct {

	// The name of the IAM policy that you have configured in each account where you
	// want to deploy your permission set.
	//
	// This member is required.
	Name *string

	// The path to the IAM policy that you have configured in each account where you
	// want to deploy your permission set. The default is / . For more information, see
	// Friendly names and paths (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names)
	// in the IAM User Guide.
	Path *string

	noSmithyDocumentSerde
}

// Specifies the attributes to add to your attribute-based access control (ABAC)
// configuration.
type InstanceAccessControlAttributeConfiguration struct {

	// Lists the attributes that are configured for ABAC in the specified IAM Identity
	// Center instance.
	//
	// This member is required.
	AccessControlAttributes []AccessControlAttribute

	noSmithyDocumentSerde
}

// Provides information about the IAM Identity Center instance.
type InstanceMetadata struct {

	// The identifier of the identity store that is connected to the IAM Identity
	// Center instance.
	IdentityStoreId *string

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed. For more information about ARNs, see Amazon Resource Names (ARNs) and
	// AWS Service Namespaces in the AWS General Reference.
	InstanceArn *string

	noSmithyDocumentSerde
}

// Filters he operation status list based on the passed attribute value.
type OperationStatusFilter struct {

	// Filters the list operations result based on the status attribute.
	Status StatusValues

	noSmithyDocumentSerde
}

// Specifies the configuration of the AWS managed or customer managed policy that
// you want to set as a permissions boundary. Specify either
// CustomerManagedPolicyReference to use the name and path of a customer managed
// policy, or ManagedPolicyArn to use the ARN of an AWS managed policy. A
// permissions boundary represents the maximum permissions that any policy can
// grant your role. For more information, see Permissions boundaries for IAM
// entities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
// in the IAM User Guide. Policies used as permissions boundaries don't provide
// permissions. You must also attach an IAM policy to the role. To learn how the
// effective permissions for a role are evaluated, see IAM JSON policy evaluation
// logic (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html)
// in the IAM User Guide.
type PermissionsBoundary struct {

	// Specifies the name and path of a customer managed policy. You must have an IAM
	// policy that matches the name and path in each AWS account where you want to
	// deploy your permission set.
	CustomerManagedPolicyReference *CustomerManagedPolicyReference

	// The AWS managed policy ARN that you want to attach to a permission set as a
	// permissions boundary.
	ManagedPolicyArn *string

	noSmithyDocumentSerde
}

// An entity that contains IAM policies.
type PermissionSet struct {

	// The date that the permission set was created.
	CreatedDate *time.Time

	// The description of the PermissionSet .
	Description *string

	// The name of the permission set.
	Name *string

	// The ARN of the permission set. For more information about ARNs, see Amazon
	// Resource Names (ARNs) and AWS Service Namespaces in the AWS General Reference.
	PermissionSetArn *string

	// Used to redirect users within the application during the federation
	// authentication process.
	RelayState *string

	// The length of time that the application user sessions are valid for in the
	// ISO-8601 standard.
	SessionDuration *string

	noSmithyDocumentSerde
}

// A structure that is used to provide the status of the provisioning operation
// for a specified permission set.
type PermissionSetProvisioningStatus struct {

	// The identifier of the AWS account from which to list the assignments.
	AccountId *string

	// The date that the permission set was created.
	CreatedDate *time.Time

	// The message that contains an error or exception in case of an operation failure.
	FailureReason *string

	// The ARN of the permission set that is being provisioned. For more information
	// about ARNs, see Amazon Resource Names (ARNs) and AWS Service Namespaces in the
	// AWS General Reference.
	PermissionSetArn *string

	// The identifier for tracking the request operation that is generated by the
	// universally unique identifier (UUID) workflow.
	RequestId *string

	// The status of the permission set provisioning process.
	Status StatusValues

	noSmithyDocumentSerde
}

// Provides information about the permission set provisioning status.
type PermissionSetProvisioningStatusMetadata struct {

	// The date that the permission set was created.
	CreatedDate *time.Time

	// The identifier for tracking the request operation that is generated by the
	// universally unique identifier (UUID) workflow.
	RequestId *string

	// The status of the permission set provisioning process.
	Status StatusValues

	noSmithyDocumentSerde
}

// A set of key-value pairs that are used to manage the resource. Tags can only be
// applied to permission sets and cannot be applied to corresponding roles that IAM
// Identity Center creates in AWS accounts.
type Tag struct {

	// The key for the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
