// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a SignalingChannel to a stream to store the media. There are two
// signaling modes that can specified :
//   - If the StorageStatus is disabled, no data will be stored, and the StreamARN
//     parameter will not be needed.
//   - If the StorageStatus is enabled, the data will be stored in the StreamARN
//     provided.
func (c *Client) UpdateMediaStorageConfiguration(ctx context.Context, params *UpdateMediaStorageConfigurationInput, optFns ...func(*Options)) (*UpdateMediaStorageConfigurationOutput, error) {
	if params == nil {
		params = &UpdateMediaStorageConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMediaStorageConfiguration", params, optFns, c.addOperationUpdateMediaStorageConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMediaStorageConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMediaStorageConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the channel.
	//
	// This member is required.
	ChannelARN *string

	// A structure that encapsulates, or contains, the media storage configuration
	// properties.
	//
	// This member is required.
	MediaStorageConfiguration *types.MediaStorageConfiguration

	noSmithyDocumentSerde
}

type UpdateMediaStorageConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMediaStorageConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMediaStorageConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMediaStorageConfiguration{}, middleware.After)
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
	if err = addOpUpdateMediaStorageConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMediaStorageConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateMediaStorageConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "UpdateMediaStorageConfiguration",
	}
}
