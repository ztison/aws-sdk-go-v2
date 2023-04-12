// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified local gateway route.
func (c *Client) ModifyLocalGatewayRoute(ctx context.Context, params *ModifyLocalGatewayRouteInput, optFns ...func(*Options)) (*ModifyLocalGatewayRouteOutput, error) {
	if params == nil {
		params = &ModifyLocalGatewayRouteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyLocalGatewayRoute", params, optFns, c.addOperationModifyLocalGatewayRouteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyLocalGatewayRouteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyLocalGatewayRouteInput struct {

	// The ID of the local gateway route table.
	//
	// This member is required.
	LocalGatewayRouteTableId *string

	// The CIDR block used for destination matches. The value that you provide must
	// match the CIDR of an existing route in the table.
	DestinationCidrBlock *string

	// The ID of the prefix list. Use a prefix list in place of DestinationCidrBlock .
	// You cannot use DestinationPrefixListId and DestinationCidrBlock in the same
	// request.
	DestinationPrefixListId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The ID of the virtual interface group.
	LocalGatewayVirtualInterfaceGroupId *string

	// The ID of the network interface.
	NetworkInterfaceId *string

	noSmithyDocumentSerde
}

type ModifyLocalGatewayRouteOutput struct {

	// Information about the local gateway route table.
	Route *types.LocalGatewayRoute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyLocalGatewayRouteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyLocalGatewayRoute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyLocalGatewayRoute{}, middleware.After)
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
	if err = addOpModifyLocalGatewayRouteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyLocalGatewayRoute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyLocalGatewayRoute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyLocalGatewayRoute",
	}
}
