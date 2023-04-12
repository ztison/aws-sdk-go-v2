// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches a block storage disk to a running or stopped Lightsail instance and
// exposes it to the instance with the specified disk name. The attach disk
// operation supports tag-based access control via resource tags applied to the
// resource identified by disk name . For more information, see the Amazon
// Lightsail Developer Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-controlling-access-using-tags)
// .
func (c *Client) AttachDisk(ctx context.Context, params *AttachDiskInput, optFns ...func(*Options)) (*AttachDiskOutput, error) {
	if params == nil {
		params = &AttachDiskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachDisk", params, optFns, c.addOperationAttachDiskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachDiskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachDiskInput struct {

	// The unique Lightsail disk name (e.g., my-disk ).
	//
	// This member is required.
	DiskName *string

	// The disk path to expose to the instance (e.g., /dev/xvdf ).
	//
	// This member is required.
	DiskPath *string

	// The name of the Lightsail instance where you want to utilize the storage disk.
	//
	// This member is required.
	InstanceName *string

	// A Boolean value used to determine the automatic mounting of a storage volume to
	// a virtual computer. The default value is False . This value only applies to
	// Lightsail for Research resources.
	AutoMounting *bool

	noSmithyDocumentSerde
}

type AttachDiskOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachDiskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAttachDisk{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachDisk{}, middleware.After)
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
	if err = addOpAttachDiskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachDisk(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachDisk(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "AttachDisk",
	}
}
