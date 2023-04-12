// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the field-level encryption configuration information.
func (c *Client) GetFieldLevelEncryptionConfig(ctx context.Context, params *GetFieldLevelEncryptionConfigInput, optFns ...func(*Options)) (*GetFieldLevelEncryptionConfigOutput, error) {
	if params == nil {
		params = &GetFieldLevelEncryptionConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFieldLevelEncryptionConfig", params, optFns, c.addOperationGetFieldLevelEncryptionConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFieldLevelEncryptionConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFieldLevelEncryptionConfigInput struct {

	// Request the ID for the field-level encryption configuration information.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetFieldLevelEncryptionConfigOutput struct {

	// The current version of the field level encryption configuration. For example:
	// E2QWRUHAPOMQZL .
	ETag *string

	// Return the field-level encryption configuration information.
	FieldLevelEncryptionConfig *types.FieldLevelEncryptionConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFieldLevelEncryptionConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetFieldLevelEncryptionConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetFieldLevelEncryptionConfig{}, middleware.After)
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
	if err = addOpGetFieldLevelEncryptionConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFieldLevelEncryptionConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFieldLevelEncryptionConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetFieldLevelEncryptionConfig",
	}
}
