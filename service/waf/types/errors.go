// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

type WAFBadRequestException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFBadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFBadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFBadRequestException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFBadRequestException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFBadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The name specified is invalid.
type WAFDisallowedNameException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFDisallowedNameException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFDisallowedNameException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFDisallowedNameException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFDisallowedNameException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFDisallowedNameException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation failed due to a problem with the migration. The failure cause is
// provided in the exception, in the MigrationErrorType :
//   - ENTITY_NOT_SUPPORTED - The web ACL has an unsupported entity but the
//     IgnoreUnsupportedType is not set to true.
//   - ENTITY_NOT_FOUND - The web ACL doesn't exist.
//   - S3_BUCKET_NO_PERMISSION - You don't have permission to perform the PutObject
//     action to the specified Amazon S3 bucket.
//   - S3_BUCKET_NOT_ACCESSIBLE - The bucket policy doesn't allow AWS WAF to
//     perform the PutObject action in the bucket.
//   - S3_BUCKET_NOT_FOUND - The S3 bucket doesn't exist.
//   - S3_BUCKET_INVALID_REGION - The S3 bucket is not in the same Region as the
//     web ACL.
//   - S3_INTERNAL_ERROR - AWS WAF failed to create the template in the S3 bucket
//     for another reason.
type WAFEntityMigrationException struct {
	Message *string

	ErrorCodeOverride *string

	MigrationErrorType   MigrationErrorType
	MigrationErrorReason *string

	noSmithyDocumentSerde
}

func (e *WAFEntityMigrationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFEntityMigrationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFEntityMigrationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFEntityMigrationException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFEntityMigrationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation failed because of a system problem, even though the request was
// valid. Retry your request.
type WAFInternalErrorException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFInternalErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFInternalErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFInternalErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFInternalErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFInternalErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The operation failed because you tried to create, update, or delete an object
// by using an invalid account identifier.
type WAFInvalidAccountException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFInvalidAccountException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFInvalidAccountException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFInvalidAccountException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFInvalidAccountException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFInvalidAccountException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation failed because there was nothing to do. For example:
//   - You tried to remove a Rule from a WebACL , but the Rule isn't in the
//     specified WebACL .
//   - You tried to remove an IP address from an IPSet , but the IP address isn't
//     in the specified IPSet .
//   - You tried to remove a ByteMatchTuple from a ByteMatchSet , but the
//     ByteMatchTuple isn't in the specified WebACL .
//   - You tried to add a Rule to a WebACL , but the Rule already exists in the
//     specified WebACL .
//   - You tried to add a ByteMatchTuple to a ByteMatchSet , but the ByteMatchTuple
//     already exists in the specified WebACL .
type WAFInvalidOperationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFInvalidOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFInvalidOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFInvalidOperationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFInvalidOperationException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFInvalidOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation failed because AWS WAF didn't recognize a parameter in the
// request. For example:
//   - You specified an invalid parameter name.
//   - You specified an invalid value.
//   - You tried to update an object ( ByteMatchSet , IPSet , Rule , or WebACL )
//     using an action other than INSERT or DELETE .
//   - You tried to create a WebACL with a DefaultAction Type other than ALLOW ,
//     BLOCK , or COUNT .
//   - You tried to create a RateBasedRule with a RateKey value other than IP .
//   - You tried to update a WebACL with a WafAction Type other than ALLOW , BLOCK
//     , or COUNT .
//   - You tried to update a ByteMatchSet with a FieldToMatch Type other than
//     HEADER, METHOD, QUERY_STRING, URI, or BODY.
//   - You tried to update a ByteMatchSet with a Field of HEADER but no value for
//     Data .
//   - Your request references an ARN that is malformed, or corresponds to a
//     resource with which a web ACL cannot be associated.
type WAFInvalidParameterException struct {
	Message *string

	ErrorCodeOverride *string

	Field     ParameterExceptionField
	Parameter *string
	Reason    ParameterExceptionReason

	noSmithyDocumentSerde
}

func (e *WAFInvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFInvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFInvalidParameterException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFInvalidParameterException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFInvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation failed because the specified policy is not in the proper format.
// The policy is subject to the following restrictions:
//   - You can attach only one policy with each PutPermissionPolicy request.
//   - The policy must include an Effect , Action and Principal .
//   - Effect must specify Allow .
//   - The Action in the policy must be waf:UpdateWebACL ,
//     waf-regional:UpdateWebACL , waf:GetRuleGroup and waf-regional:GetRuleGroup .
//     Any extra or wildcard actions in the policy will be rejected.
//   - The policy cannot include a Resource parameter.
//   - The ARN in the request must be a valid WAF RuleGroup ARN and the RuleGroup
//     must exist in the same region.
//   - The user making the request must be the owner of the RuleGroup.
//   - Your policy must be composed using IAM Policy version 2012-10-17.
type WAFInvalidPermissionPolicyException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFInvalidPermissionPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFInvalidPermissionPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFInvalidPermissionPolicyException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFInvalidPermissionPolicyException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFInvalidPermissionPolicyException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The regular expression (regex) you specified in RegexPatternString is invalid.
type WAFInvalidRegexPatternException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFInvalidRegexPatternException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFInvalidRegexPatternException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFInvalidRegexPatternException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFInvalidRegexPatternException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFInvalidRegexPatternException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation exceeds a resource limit, for example, the maximum number of
// WebACL objects that you can create for an AWS account. For more information, see
// Limits (https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in
// the AWS WAF Developer Guide.
type WAFLimitsExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFLimitsExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFLimitsExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFLimitsExceededException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFLimitsExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFLimitsExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation failed because you tried to delete an object that isn't empty.
// For example:
//   - You tried to delete a WebACL that still contains one or more Rule objects.
//   - You tried to delete a Rule that still contains one or more ByteMatchSet
//     objects or other predicates.
//   - You tried to delete a ByteMatchSet that contains one or more ByteMatchTuple
//     objects.
//   - You tried to delete an IPSet that references one or more IP addresses.
type WAFNonEmptyEntityException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFNonEmptyEntityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFNonEmptyEntityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFNonEmptyEntityException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFNonEmptyEntityException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFNonEmptyEntityException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation failed because you tried to add an object to or delete an object
// from another object that doesn't exist. For example:
//   - You tried to add a Rule to or delete a Rule from a WebACL that doesn't
//     exist.
//   - You tried to add a ByteMatchSet to or delete a ByteMatchSet from a Rule that
//     doesn't exist.
//   - You tried to add an IP address to or delete an IP address from an IPSet that
//     doesn't exist.
//   - You tried to add a ByteMatchTuple to or delete a ByteMatchTuple from a
//     ByteMatchSet that doesn't exist.
type WAFNonexistentContainerException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFNonexistentContainerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFNonexistentContainerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFNonexistentContainerException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFNonexistentContainerException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFNonexistentContainerException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation failed because the referenced object doesn't exist.
type WAFNonexistentItemException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFNonexistentItemException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFNonexistentItemException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFNonexistentItemException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFNonexistentItemException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFNonexistentItemException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The operation failed because you tried to delete an object that is still in
// use. For example:
//   - You tried to delete a ByteMatchSet that is still referenced by a Rule .
//   - You tried to delete a Rule that is still referenced by a WebACL .
type WAFReferencedItemException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFReferencedItemException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFReferencedItemException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFReferencedItemException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFReferencedItemException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFReferencedItemException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// AWS WAF is not able to access the service linked role. This can be caused by a
// previous PutLoggingConfiguration request, which can lock the service linked
// role for about 20 seconds. Please try your request again. The service linked
// role can also be locked by a previous DeleteServiceLinkedRole request, which
// can lock the role for 15 minutes or more. If you recently made a
// DeleteServiceLinkedRole , wait at least 15 minutes and try the request again. If
// you receive this same exception again, you will have to wait additional time
// until the role is unlocked.
type WAFServiceLinkedRoleErrorException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFServiceLinkedRoleErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFServiceLinkedRoleErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFServiceLinkedRoleErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFServiceLinkedRoleErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFServiceLinkedRoleErrorException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The operation failed because you tried to create, update, or delete an object
// by using a change token that has already been used.
type WAFStaleDataException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFStaleDataException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFStaleDataException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFStaleDataException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFStaleDataException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFStaleDataException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified subscription does not exist.
type WAFSubscriptionNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFSubscriptionNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFSubscriptionNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFSubscriptionNotFoundException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFSubscriptionNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFSubscriptionNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type WAFTagOperationException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFTagOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFTagOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFTagOperationException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFTagOperationException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFTagOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

type WAFTagOperationInternalErrorException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *WAFTagOperationInternalErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *WAFTagOperationInternalErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *WAFTagOperationInternalErrorException) ErrorCode() string {
	if e == nil || e.ErrorCodeOverride == nil {
		return "WAFTagOperationInternalErrorException"
	}
	return *e.ErrorCodeOverride
}
func (e *WAFTagOperationInternalErrorException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultServer
}
