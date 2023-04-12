// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes one or more attributes, of the same attribute type, from all the
// endpoints that are associated with an application.
func (c *Client) RemoveAttributes(ctx context.Context, params *RemoveAttributesInput, optFns ...func(*Options)) (*RemoveAttributesOutput, error) {
	if params == nil {
		params = &RemoveAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveAttributes", params, optFns, c.addOperationRemoveAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveAttributesInput struct {

	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	//
	// This member is required.
	ApplicationId *string

	// The type of attribute or attributes to remove. Valid values are:
	//   - endpoint-custom-attributes - Custom attributes that describe endpoints,
	//   such as the date when an associated user opted in or out of receiving
	//   communications from you through a specific type of channel.
	//   - endpoint-metric-attributes - Custom metrics that your app reports to Amazon
	//   Pinpoint for endpoints, such as the number of app sessions or the number of
	//   items left in a cart.
	//   - endpoint-user-attributes - Custom attributes that describe users, such as
	//   first name, last name, and age.
	//
	// This member is required.
	AttributeType *string

	// Specifies one or more attributes to remove from all the endpoints that are
	// associated with an application.
	//
	// This member is required.
	UpdateAttributesRequest *types.UpdateAttributesRequest

	noSmithyDocumentSerde
}

type RemoveAttributesOutput struct {

	// Provides information about the type and the names of attributes that were
	// removed from all the endpoints that are associated with an application.
	//
	// This member is required.
	AttributesResource *types.AttributesResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRemoveAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveAttributes{}, middleware.After)
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
	if err = addOpRemoveAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "RemoveAttributes",
	}
}
