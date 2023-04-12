// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation completes the vault locking process by transitioning the vault
// lock from the InProgress state to the Locked state, which causes the vault lock
// policy to become unchangeable. A vault lock is put into the InProgress state by
// calling InitiateVaultLock . You can obtain the state of the vault lock by
// calling GetVaultLock . For more information about the vault locking process,
// Amazon Glacier Vault Lock (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html)
// . This operation is idempotent. This request is always successful if the vault
// lock is in the Locked state and the provided lock ID matches the lock ID
// originally used to lock the vault. If an invalid lock ID is passed in the
// request when the vault lock is in the Locked state, the operation returns an
// AccessDeniedException error. If an invalid lock ID is passed in the request when
// the vault lock is in the InProgress state, the operation throws an
// InvalidParameter error.
func (c *Client) CompleteVaultLock(ctx context.Context, params *CompleteVaultLockInput, optFns ...func(*Options)) (*CompleteVaultLockOutput, error) {
	if params == nil {
		params = &CompleteVaultLockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CompleteVaultLock", params, optFns, c.addOperationCompleteVaultLockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CompleteVaultLockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input values for CompleteVaultLock .
type CompleteVaultLockInput struct {

	// The AccountId value is the AWS account ID. This value must match the AWS
	// account ID associated with the credentials used to sign the request. You can
	// either specify an AWS account ID or optionally a single ' - ' (hyphen), in which
	// case Amazon Glacier uses the AWS account ID associated with the credentials used
	// to sign the request. If you specify your account ID, do not include any hyphens
	// ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The lockId value is the lock ID obtained from a InitiateVaultLock request.
	//
	// This member is required.
	LockId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	noSmithyDocumentSerde
}

type CompleteVaultLockOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCompleteVaultLockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCompleteVaultLock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCompleteVaultLock{}, middleware.After)
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
	if err = addOpCompleteVaultLockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCompleteVaultLock(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCompleteVaultLock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "CompleteVaultLock",
	}
}
