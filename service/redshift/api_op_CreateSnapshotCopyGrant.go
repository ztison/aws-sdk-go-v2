// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a snapshot copy grant that permits Amazon Redshift to use an encrypted
// symmetric key from Key Management Service (KMS) to encrypt copied snapshots in a
// destination region. For more information about managing snapshot copy grants, go
// to Amazon Redshift Database Encryption (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) CreateSnapshotCopyGrant(ctx context.Context, params *CreateSnapshotCopyGrantInput, optFns ...func(*Options)) (*CreateSnapshotCopyGrantOutput, error) {
	if params == nil {
		params = &CreateSnapshotCopyGrantInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSnapshotCopyGrant", params, optFns, c.addOperationCreateSnapshotCopyGrantMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSnapshotCopyGrantOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The result of the CreateSnapshotCopyGrant action.
type CreateSnapshotCopyGrantInput struct {

	// The name of the snapshot copy grant. This name must be unique in the region for
	// the Amazon Web Services account. Constraints:
	//   - Must contain from 1 to 63 alphanumeric characters or hyphens.
	//   - Alphabetic characters must be lowercase.
	//   - First character must be a letter.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	//   - Must be unique for all clusters within an Amazon Web Services account.
	//
	// This member is required.
	SnapshotCopyGrantName *string

	// The unique identifier of the encrypted symmetric key to which to grant Amazon
	// Redshift permission. If no key is specified, the default key is used.
	KmsKeyId *string

	// A list of tag instances.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateSnapshotCopyGrantOutput struct {

	// The snapshot copy grant that grants Amazon Redshift permission to encrypt
	// copied snapshots with the specified encrypted symmetric key from Amazon Web
	// Services KMS in the destination region. For more information about managing
	// snapshot copy grants, go to Amazon Redshift Database Encryption (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html)
	// in the Amazon Redshift Cluster Management Guide.
	SnapshotCopyGrant *types.SnapshotCopyGrant

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSnapshotCopyGrantMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateSnapshotCopyGrant{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateSnapshotCopyGrant{}, middleware.After)
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
	if err = addOpCreateSnapshotCopyGrantValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSnapshotCopyGrant(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSnapshotCopyGrant(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateSnapshotCopyGrant",
	}
}
