// Code generated by smithy-go-codegen DO NOT EDIT.

package ebs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ebs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Writes a block of data to a snapshot. If the specified block contains data, the
// existing data is overwritten. The target snapshot must be in the pending state.
// Data written to a snapshot must be aligned with 512-KiB sectors.
func (c *Client) PutSnapshotBlock(ctx context.Context, params *PutSnapshotBlockInput, optFns ...func(*Options)) (*PutSnapshotBlockOutput, error) {
	if params == nil {
		params = &PutSnapshotBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSnapshotBlock", params, optFns, c.addOperationPutSnapshotBlockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSnapshotBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSnapshotBlockInput struct {

	// The data to write to the block. The block data is not signed as part of the
	// Signature Version 4 signing process. As a result, you must generate and provide
	// a Base64-encoded SHA256 checksum for the block data using the x-amz-Checksum
	// header. Also, you must specify the checksum algorithm using the
	// x-amz-Checksum-Algorithm header. The checksum that you provide is part of the
	// Signature Version 4 signing process. It is validated against a checksum
	// generated by Amazon EBS to ensure the validity and authenticity of the data. If
	// the checksums do not correspond, the request fails. For more information, see
	// Using checksums with the EBS direct APIs (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-accessing-snapshot.html#ebsapis-using-checksums)
	// in the Amazon Elastic Compute Cloud User Guide.
	//
	// This member is required.
	BlockData io.Reader

	// The block index of the block in which to write the data. A block index is a
	// logical index in units of 512 KiB blocks. To identify the block index, divide
	// the logical offset of the data in the logical volume by the block size (logical
	// offset of data/ 524288 ). The logical offset of the data must be 512 KiB
	// aligned.
	//
	// This member is required.
	BlockIndex *int32

	// A Base64-encoded SHA256 checksum of the data. Only SHA256 checksums are
	// supported.
	//
	// This member is required.
	Checksum *string

	// The algorithm used to generate the checksum. Currently, the only supported
	// algorithm is SHA256 .
	//
	// This member is required.
	ChecksumAlgorithm types.ChecksumAlgorithm

	// The size of the data to write to the block, in bytes. Currently, the only
	// supported size is 524288 bytes. Valid values: 524288
	//
	// This member is required.
	DataLength *int32

	// The ID of the snapshot. If the specified snapshot is encrypted, you must have
	// permission to use the KMS key that was used to encrypt the snapshot. For more
	// information, see Using encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html)
	// in the Amazon Elastic Compute Cloud User Guide..
	//
	// This member is required.
	SnapshotId *string

	// The progress of the write process, as a percentage.
	Progress *int32

	noSmithyDocumentSerde
}

type PutSnapshotBlockOutput struct {

	// The SHA256 checksum generated for the block data by Amazon EBS.
	Checksum *string

	// The algorithm used by Amazon EBS to generate the checksum.
	ChecksumAlgorithm types.ChecksumAlgorithm

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSnapshotBlockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutSnapshotBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSnapshotBlock{}, middleware.After)
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
	if err = v4.AddUnsignedPayloadMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
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
	if err = addOpPutSnapshotBlockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSnapshotBlock(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSnapshotBlock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ebs",
		OperationName: "PutSnapshotBlock",
	}
}
