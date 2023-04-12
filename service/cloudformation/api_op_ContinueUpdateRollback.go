// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For a specified stack that's in the UPDATE_ROLLBACK_FAILED state, continues
// rolling it back to the UPDATE_ROLLBACK_COMPLETE state. Depending on the cause
// of the failure, you can manually fix the error (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/troubleshooting.html#troubleshooting-errors-update-rollback-failed)
// and continue the rollback. By continuing the rollback, you can return your stack
// to a working state (the UPDATE_ROLLBACK_COMPLETE state), and then try to update
// the stack again. A stack goes into the UPDATE_ROLLBACK_FAILED state when
// CloudFormation can't roll back all changes after a failed stack update. For
// example, you might have a stack that's rolling back to an old database instance
// that was deleted outside of CloudFormation. Because CloudFormation doesn't know
// the database was deleted, it assumes that the database instance still exists and
// attempts to roll back to it, causing the update rollback to fail.
func (c *Client) ContinueUpdateRollback(ctx context.Context, params *ContinueUpdateRollbackInput, optFns ...func(*Options)) (*ContinueUpdateRollbackOutput, error) {
	if params == nil {
		params = &ContinueUpdateRollbackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ContinueUpdateRollback", params, optFns, c.addOperationContinueUpdateRollbackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ContinueUpdateRollbackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ContinueUpdateRollback action.
type ContinueUpdateRollbackInput struct {

	// The name or the unique ID of the stack that you want to continue rolling back.
	// Don't specify the name of a nested stack (a stack that was created by using the
	// AWS::CloudFormation::Stack resource). Instead, use this operation on the parent
	// stack (the stack that contains the AWS::CloudFormation::Stack resource).
	//
	// This member is required.
	StackName *string

	// A unique identifier for this ContinueUpdateRollback request. Specify this token
	// if you plan to retry requests so that CloudFormationknows that you're not
	// attempting to continue the rollback to a stack with the same name. You might
	// retry ContinueUpdateRollback requests to ensure that CloudFormation
	// successfully received them.
	ClientRequestToken *string

	// A list of the logical IDs of the resources that CloudFormation skips during the
	// continue update rollback operation. You can specify only resources that are in
	// the UPDATE_FAILED state because a rollback failed. You can't specify resources
	// that are in the UPDATE_FAILED state for other reasons, for example, because an
	// update was canceled. To check why a resource update failed, use the
	// DescribeStackResources action, and view the resource status reason. Specify this
	// property to skip rolling back resources that CloudFormation can't successfully
	// roll back. We recommend that you troubleshoot (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/troubleshooting.html#troubleshooting-errors-update-rollback-failed)
	// resources before skipping them. CloudFormation sets the status of the specified
	// resources to UPDATE_COMPLETE and continues to roll back the stack. After the
	// rollback is complete, the state of the skipped resources will be inconsistent
	// with the state of the resources in the stack template. Before performing another
	// stack update, you must update the stack or resources to be consistent with each
	// other. If you don't, subsequent stack updates might fail, and the stack will
	// become unrecoverable. Specify the minimum number of resources required to
	// successfully roll back your stack. For example, a failed resource update might
	// cause dependent resources to fail. In this case, it might not be necessary to
	// skip the dependent resources. To skip resources that are part of nested stacks,
	// use the following format: NestedStackName.ResourceLogicalID . If you want to
	// specify the logical ID of a stack resource ( Type: AWS::CloudFormation::Stack )
	// in the ResourcesToSkip list, then its corresponding embedded stack must be in
	// one of the following states: DELETE_IN_PROGRESS , DELETE_COMPLETE , or
	// DELETE_FAILED . Don't confuse a child stack's name with its corresponding
	// logical ID defined in the parent stack. For an example of a continue update
	// rollback operation with nested stacks, see Using ResourcesToSkip to recover a
	// nested stacks hierarchy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-continueupdaterollback.html#nested-stacks)
	// .
	ResourcesToSkip []string

	// The Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role
	// that CloudFormation assumes to roll back the stack. CloudFormation uses the
	// role's credentials to make calls on your behalf. CloudFormation always uses this
	// role for all future operations on the stack. Provided that users have permission
	// to operate on the stack, CloudFormation uses this role even if the users don't
	// have permission to pass it. Ensure that the role grants least permission. If you
	// don't specify a value, CloudFormation uses the role that was previously
	// associated with the stack. If no role is available, CloudFormation uses a
	// temporary session that's generated from your user credentials.
	RoleARN *string

	noSmithyDocumentSerde
}

// The output for a ContinueUpdateRollback operation.
type ContinueUpdateRollbackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationContinueUpdateRollbackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpContinueUpdateRollback{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpContinueUpdateRollback{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpContinueUpdateRollbackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opContinueUpdateRollback(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opContinueUpdateRollback(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ContinueUpdateRollback",
	}
}
