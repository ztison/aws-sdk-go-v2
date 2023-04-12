// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an API key resource in your Amazon Web Services account, which lets you
// grant geo:GetMap* actions for Amazon Location Map resources to the API key
// bearer. The API keys feature is in preview. We may add, change, or remove
// features before announcing general availability. For more information, see
// Using API keys (https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html)
// .
func (c *Client) CreateKey(ctx context.Context, params *CreateKeyInput, optFns ...func(*Options)) (*CreateKeyOutput, error) {
	if params == nil {
		params = &CreateKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKey", params, optFns, c.addOperationCreateKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKeyInput struct {

	// A custom name for the API key resource. Requirements:
	//   - Contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods
	//   (.), and underscores (_).
	//   - Must be a unique API key name.
	//   - No spaces allowed. For example, ExampleAPIKey .
	//
	// This member is required.
	KeyName *string

	// The API key restrictions for the API key resource.
	//
	// This member is required.
	Restrictions *types.ApiKeyRestrictions

	// An optional description for the API key resource.
	Description *string

	// The optional timestamp for when the API key resource will expire in  ISO 8601 (https://www.iso.org/iso-8601-date-and-time-format.html)
	// format: YYYY-MM-DDThh:mm:ss.sssZ . One of NoExpiry or ExpireTime must be set.
	ExpireTime *time.Time

	// Optionally set to true to set no expiration time for the API key. One of
	// NoExpiry or ExpireTime must be set.
	NoExpiry *bool

	// Applies one or more tags to the map resource. A tag is a key-value pair that
	// helps manage, identify, search, and filter your resources by labelling them.
	// Format: "key" : "value" Restrictions:
	//   - Maximum 50 tags per resource
	//   - Each resource tag must be unique with a maximum of one value.
	//   - Maximum key length: 128 Unicode characters in UTF-8
	//   - Maximum value length: 256 Unicode characters in UTF-8
	//   - Can use alphanumeric characters (A–Z, a–z, 0–9), and the following
	//   characters: + - = . _ : / @.
	//   - Cannot use "aws:" as a prefix for a key.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateKeyOutput struct {

	// The timestamp for when the API key resource was created in  ISO 8601 (https://www.iso.org/iso-8601-date-and-time-format.html)
	// format: YYYY-MM-DDThh:mm:ss.sssZ .
	//
	// This member is required.
	CreateTime *time.Time

	// The key value/string of an API key. This value is used when making API calls to
	// authorize the call. For example, see GetMapGlyphs (https://docs.aws.amazon.com/location/latest/APIReference/API_GetMapGlyphs.html)
	// .
	//
	// This member is required.
	Key *string

	// The Amazon Resource Name (ARN) for the API key resource. Used when you need to
	// specify a resource across all Amazon Web Services.
	//   - Format example: arn:aws:geo:region:account-id:key/ExampleKey
	//
	// This member is required.
	KeyArn *string

	// The name of the API key resource.
	//
	// This member is required.
	KeyName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateKey{}, middleware.After)
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
	if err = addEndpointPrefix_opCreateKeyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKey(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateKeyMiddleware struct {
}

func (*endpointPrefix_opCreateKeyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateKeyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "metadata." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreateKeyMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreateKeyMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opCreateKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "CreateKey",
	}
}
