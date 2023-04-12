// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// You don't have access to this item. The provided credentials couldn't be
// validated. You might not be authorized to carry out the request. Make sure that
// your account is authorized to use the Amazon QuickSight service, that your
// policies have the correct permissions, and that you are using the correct
// credentials.
type AccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AccessDeniedException"
	}
	return *e.ErrorCodeOverride
}
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource is already in a state that indicates an operation is happening that
// must complete before a new update can be applied.
type ConcurrentUpdatingException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *ConcurrentUpdatingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentUpdatingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentUpdatingException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConcurrentUpdatingException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConcurrentUpdatingException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Updating or deleting a resource can cause an inconsistent state.
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *ConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ConflictException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The domain specified isn't on the allow list. All domains for embedded
// dashboards must be added to the approved list by an Amazon QuickSight admin.
type DomainNotWhitelistedException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *DomainNotWhitelistedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DomainNotWhitelistedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DomainNotWhitelistedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DomainNotWhitelistedException"
	}
	return *e.ErrorCodeOverride
}
func (e *DomainNotWhitelistedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The identity type specified isn't supported. Supported identity types include
// IAM and QUICKSIGHT .
type IdentityTypeNotSupportedException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *IdentityTypeNotSupportedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdentityTypeNotSupportedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdentityTypeNotSupportedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "IdentityTypeNotSupportedException"
	}
	return *e.ErrorCodeOverride
}
func (e *IdentityTypeNotSupportedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An internal failure occurred.
type InternalFailureException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *InternalFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalFailureException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalFailureException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The NextToken value isn't valid.
type InvalidNextTokenException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidNextTokenException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more parameters has a value that isn't valid.
type InvalidParameterValueException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *InvalidParameterValueException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterValueException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterValueException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidParameterValueException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidParameterValueException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You don't have this feature activated for your account. To fix this issue,
// contact Amazon Web Services support.
type InvalidRequestException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A limit is exceeded.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceType ExceptionResourceType
	RequestId    *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more preconditions aren't met.
type PreconditionNotMetException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *PreconditionNotMetException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PreconditionNotMetException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PreconditionNotMetException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "PreconditionNotMetException"
	}
	return *e.ErrorCodeOverride
}
func (e *PreconditionNotMetException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The user with the provided name isn't found. This error can happen in any
// operation that requires finding a user based on a provided user name, such as
// DeleteUser , DescribeUser , and so on.
type QuickSightUserNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *QuickSightUserNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *QuickSightUserNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *QuickSightUserNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "QuickSightUserNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *QuickSightUserNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource specified already exists.
type ResourceExistsException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceType ExceptionResourceType
	RequestId    *string

	noSmithyDocumentSerde
}

func (e *ResourceExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more resources can't be found.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceType ExceptionResourceType
	RequestId    *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This resource is currently unavailable.
type ResourceUnavailableException struct {
	Message *string

	ErrorCodeOverride *string

	ResourceType ExceptionResourceType
	RequestId    *string

	noSmithyDocumentSerde
}

func (e *ResourceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceUnavailableException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ResourceUnavailableException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The number of minutes specified for the lifetime of a session isn't valid. The
// session lifetime must be 15-600 minutes.
type SessionLifetimeInMinutesInvalidException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *SessionLifetimeInMinutesInvalidException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SessionLifetimeInMinutesInvalidException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SessionLifetimeInMinutesInvalidException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "SessionLifetimeInMinutesInvalidException"
	}
	return *e.ErrorCodeOverride
}
func (e *SessionLifetimeInMinutesInvalidException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Access is throttled.
type ThrottlingException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ThrottlingException"
	}
	return *e.ErrorCodeOverride
}
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This error indicates that you are calling an embedding operation in Amazon
// QuickSight without the required pricing plan on your Amazon Web Services
// account. Before you can use embedding for anonymous users, a QuickSight
// administrator needs to add capacity pricing to Amazon QuickSight. You can do
// this on the Manage Amazon QuickSight page. After capacity pricing is added, you
// can use the GetDashboardEmbedUrl (https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GetDashboardEmbedUrl.html)
// API operation with the --identity-type ANONYMOUS option.
type UnsupportedPricingPlanException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *UnsupportedPricingPlanException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedPricingPlanException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedPricingPlanException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedPricingPlanException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedPricingPlanException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This error indicates that you are calling an operation on an Amazon QuickSight
// subscription where the edition doesn't include support for that operation.
// Amazon Amazon QuickSight currently has Standard Edition and Enterprise Edition.
// Not every operation and capability is available in every edition.
type UnsupportedUserEditionException struct {
	Message *string

	ErrorCodeOverride *string

	RequestId *string

	noSmithyDocumentSerde
}

func (e *UnsupportedUserEditionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedUserEditionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedUserEditionException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "UnsupportedUserEditionException"
	}
	return *e.ErrorCodeOverride
}
func (e *UnsupportedUserEditionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
