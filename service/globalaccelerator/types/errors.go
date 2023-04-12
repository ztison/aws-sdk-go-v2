// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The accelerator that you specified could not be disabled.
type AcceleratorNotDisabledException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AcceleratorNotDisabledException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AcceleratorNotDisabledException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AcceleratorNotDisabledException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AcceleratorNotDisabledException"
	}
	return *e.ErrorCodeOverride
}
func (e *AcceleratorNotDisabledException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The accelerator that you specified doesn't exist.
type AcceleratorNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AcceleratorNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AcceleratorNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AcceleratorNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AcceleratorNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *AcceleratorNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You don't have access permission.
type AccessDeniedException struct {
	Message *string

	ErrorCodeOverride *string

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

// The listener that you specified has an endpoint group associated with it. You
// must remove all dependent resources from a listener before you can delete it.
type AssociatedEndpointGroupFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AssociatedEndpointGroupFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AssociatedEndpointGroupFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AssociatedEndpointGroupFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AssociatedEndpointGroupFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *AssociatedEndpointGroupFoundException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The accelerator that you specified has a listener associated with it. You must
// remove all dependent resources from an accelerator before you can delete it.
type AssociatedListenerFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *AssociatedListenerFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AssociatedListenerFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AssociatedListenerFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "AssociatedListenerFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *AssociatedListenerFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The CIDR that you specified was not found or is incorrect.
type ByoipCidrNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ByoipCidrNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ByoipCidrNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ByoipCidrNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ByoipCidrNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ByoipCidrNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You can't use both of those options.
type ConflictException struct {
	Message *string

	ErrorCodeOverride *string

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

// The endpoint that you specified doesn't exist.
type EndpointAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *EndpointAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EndpointAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EndpointAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EndpointAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *EndpointAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The endpoint group that you specified already exists.
type EndpointGroupAlreadyExistsException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *EndpointGroupAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EndpointGroupAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EndpointGroupAlreadyExistsException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EndpointGroupAlreadyExistsException"
	}
	return *e.ErrorCodeOverride
}
func (e *EndpointGroupAlreadyExistsException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The endpoint group that you specified doesn't exist.
type EndpointGroupNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *EndpointGroupNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EndpointGroupNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EndpointGroupNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EndpointGroupNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *EndpointGroupNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The endpoint that you specified doesn't exist.
type EndpointNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *EndpointNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EndpointNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EndpointNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "EndpointNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *EndpointNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The CIDR that you specified is not valid for this action. For example, the
// state of the CIDR might be incorrect for this action.
type IncorrectCidrStateException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *IncorrectCidrStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IncorrectCidrStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IncorrectCidrStateException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "IncorrectCidrStateException"
	}
	return *e.ErrorCodeOverride
}
func (e *IncorrectCidrStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There was an internal error for Global Accelerator.
type InternalServiceErrorException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServiceErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServiceErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServiceErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalServiceErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServiceErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// An argument that you specified is invalid.
type InvalidArgumentException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidArgumentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArgumentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArgumentException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidArgumentException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidArgumentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There isn't another item to return.
type InvalidNextTokenException struct {
	Message *string

	ErrorCodeOverride *string

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

// The port numbers that you specified are not valid numbers or are not unique for
// this accelerator.
type InvalidPortRangeException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidPortRangeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPortRangeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPortRangeException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InvalidPortRangeException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidPortRangeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Processing your request would cause you to exceed an Global Accelerator limit.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

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

// The listener that you specified doesn't exist.
type ListenerNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ListenerNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ListenerNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ListenerNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ListenerNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ListenerNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There's already a transaction in progress. Another transaction can't be
// processed.
type TransactionInProgressException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *TransactionInProgressException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TransactionInProgressException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TransactionInProgressException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "TransactionInProgressException"
	}
	return *e.ErrorCodeOverride
}
func (e *TransactionInProgressException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
