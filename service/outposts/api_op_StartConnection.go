// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Amazon Web Services uses this action to install Outpost servers. Starts the
// connection required for Outpost server installation. Use CloudTrail to monitor
// this action or Amazon Web Services managed policy for Amazon Web Services
// Outposts to secure it. For more information, see Amazon Web Services managed
// policies for Amazon Web Services Outposts (https://docs.aws.amazon.com/outposts/latest/userguide/security-iam-awsmanpol.html)
// and Logging Amazon Web Services Outposts API calls with Amazon Web Services
// CloudTrail (https://docs.aws.amazon.com/outposts/latest/userguide/logging-using-cloudtrail.html)
// in the Amazon Web Services Outposts User Guide.
func (c *Client) StartConnection(ctx context.Context, params *StartConnectionInput, optFns ...func(*Options)) (*StartConnectionOutput, error) {
	if params == nil {
		params = &StartConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartConnection", params, optFns, c.addOperationStartConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartConnectionInput struct {

	// The ID of the Outpost server.
	//
	// This member is required.
	AssetId *string

	// The public key of the client.
	//
	// This member is required.
	ClientPublicKey *string

	// The serial number of the dongle.
	//
	// This member is required.
	DeviceSerialNumber *string

	// The device index of the network interface on the Outpost server.
	//
	// This member is required.
	NetworkInterfaceDeviceIndex int32

	noSmithyDocumentSerde
}

type StartConnectionOutput struct {

	// The ID of the connection.
	ConnectionId *string

	// The underlay IP address.
	UnderlayIpAddress *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartConnection{}, middleware.After)
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
	if err = addOpStartConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "outposts",
		OperationName: "StartConnection",
	}
}
