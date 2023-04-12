// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns metadata about your DataSync location for an object storage system.
func (c *Client) DescribeLocationObjectStorage(ctx context.Context, params *DescribeLocationObjectStorageInput, optFns ...func(*Options)) (*DescribeLocationObjectStorageOutput, error) {
	if params == nil {
		params = &DescribeLocationObjectStorageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocationObjectStorage", params, optFns, c.addOperationDescribeLocationObjectStorageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocationObjectStorageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeLocationObjectStorageRequest
type DescribeLocationObjectStorageInput struct {

	// The Amazon Resource Name (ARN) of the object storage system location that you
	// want information about.
	//
	// This member is required.
	LocationArn *string

	noSmithyDocumentSerde
}

// DescribeLocationObjectStorageResponse
type DescribeLocationObjectStorageOutput struct {

	// The access key (for example, a user name) required to authenticate with the
	// object storage system.
	AccessKey *string

	// The ARNs of the DataSync agents that can securely connect with your location.
	AgentArns []string

	// The time that the location was created.
	CreationTime *time.Time

	// The ARN of the object storage system location.
	LocationArn *string

	// The URL of the object storage system location.
	LocationUri *string

	// The self-signed certificate that DataSync uses to securely authenticate with
	// your object storage system.
	ServerCertificate []byte

	// The port that your object storage server accepts inbound network traffic on
	// (for example, port 443).
	ServerPort *int32

	// The protocol that your object storage system uses to communicate.
	ServerProtocol types.ObjectStorageServerProtocol

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLocationObjectStorageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLocationObjectStorage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLocationObjectStorage{}, middleware.After)
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
	if err = addOpDescribeLocationObjectStorageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocationObjectStorage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLocationObjectStorage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeLocationObjectStorage",
	}
}
