// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Makes a copy of an existing snapshot. This operation is valid for Redis only.
// Users or groups that have permissions to use the CopySnapshot operation can
// create their own Amazon S3 buckets and copy snapshots to it. To control access
// to your snapshots, use an IAM policy to control who has the ability to use the
// CopySnapshot operation. For more information about using IAM to control the use
// of ElastiCache operations, see Exporting Snapshots (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html)
// and Authentication & Access Control (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/IAM.html)
// . You could receive the following error messages. Error Messages
//   - Error Message: The S3 bucket %s is outside of the region. Solution: Create
//     an Amazon S3 bucket in the same region as your snapshot. For more information,
//     see Step 1: Create an Amazon S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-create-s3-bucket)
//     in the ElastiCache User Guide.
//   - Error Message: The S3 bucket %s does not exist. Solution: Create an Amazon
//     S3 bucket in the same region as your snapshot. For more information, see Step
//     1: Create an Amazon S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-create-s3-bucket)
//     in the ElastiCache User Guide.
//   - Error Message: The S3 bucket %s is not owned by the authenticated user.
//     Solution: Create an Amazon S3 bucket in the same region as your snapshot. For
//     more information, see Step 1: Create an Amazon S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-create-s3-bucket)
//     in the ElastiCache User Guide.
//   - Error Message: The authenticated user does not have sufficient permissions
//     to perform the desired activity. Solution: Contact your system administrator to
//     get the needed permissions.
//   - Error Message: The S3 bucket %s already contains an object with key %s.
//     Solution: Give the TargetSnapshotName a new and unique value. If exporting a
//     snapshot, you could alternatively create a new Amazon S3 bucket and use this
//     same value for TargetSnapshotName .
//   - Error Message: ElastiCache has not been granted READ permissions %s on the
//     S3 Bucket. Solution: Add List and Read permissions on the bucket. For more
//     information, see Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-grant-access)
//     in the ElastiCache User Guide.
//   - Error Message: ElastiCache has not been granted WRITE permissions %s on the
//     S3 Bucket. Solution: Add Upload/Delete permissions on the bucket. For more
//     information, see Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-grant-access)
//     in the ElastiCache User Guide.
//   - Error Message: ElastiCache has not been granted READ_ACP permissions %s on
//     the S3 Bucket. Solution: Add View Permissions on the bucket. For more
//     information, see Step 2: Grant ElastiCache Access to Your Amazon S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-grant-access)
//     in the ElastiCache User Guide.
func (c *Client) CopySnapshot(ctx context.Context, params *CopySnapshotInput, optFns ...func(*Options)) (*CopySnapshotOutput, error) {
	if params == nil {
		params = &CopySnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopySnapshot", params, optFns, c.addOperationCopySnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopySnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CopySnapshotMessage operation.
type CopySnapshotInput struct {

	// The name of an existing snapshot from which to make a copy.
	//
	// This member is required.
	SourceSnapshotName *string

	// A name for the snapshot copy. ElastiCache does not permit overwriting a
	// snapshot, therefore this name must be unique within its context - ElastiCache or
	// an Amazon S3 bucket if exporting.
	//
	// This member is required.
	TargetSnapshotName *string

	// The ID of the KMS key used to encrypt the target snapshot.
	KmsKeyId *string

	// A list of tags to be added to this resource. A tag is a key-value pair. A tag
	// key must be accompanied by a tag value, although null is accepted.
	Tags []types.Tag

	// The Amazon S3 bucket to which the snapshot is exported. This parameter is used
	// only when exporting a snapshot for external access. When using this parameter to
	// export a snapshot, be sure Amazon ElastiCache has the needed permissions to this
	// S3 bucket. For more information, see Step 2: Grant ElastiCache Access to Your
	// Amazon S3 Bucket (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html#backups-exporting-grant-access)
	// in the Amazon ElastiCache User Guide. For more information, see Exporting a
	// Snapshot (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/backups-exporting.html)
	// in the Amazon ElastiCache User Guide.
	TargetBucket *string

	noSmithyDocumentSerde
}

type CopySnapshotOutput struct {

	// Represents a copy of an entire Redis cluster as of the time when the snapshot
	// was taken.
	Snapshot *types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCopySnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCopySnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCopySnapshot{}, middleware.After)
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
	if err = addOpCopySnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCopySnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCopySnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "CopySnapshot",
	}
}
