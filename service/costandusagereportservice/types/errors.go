// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// A report with the specified name already exists in the account. Specify a
// different report name.
type DuplicateReportNameException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *DuplicateReportNameException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DuplicateReportNameException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DuplicateReportNameException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "DuplicateReportNameException"
	}
	return *e.ErrorCodeOverride
}
func (e *DuplicateReportNameException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An error on the server occurred during the processing of your request. Try
// again later.
type InternalErrorException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "InternalErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// This account already has five reports defined. To define a new report, you must
// delete an existing report.
type ReportLimitReachedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ReportLimitReachedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReportLimitReachedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReportLimitReachedException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ReportLimitReachedException"
	}
	return *e.ErrorCodeOverride
}
func (e *ReportLimitReachedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The input fails to satisfy the constraints specified by an AWS service.
type ValidationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ValidationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ValidationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "ValidationException"
	}
	return *e.ErrorCodeOverride
}
func (e *ValidationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
