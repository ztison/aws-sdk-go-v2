// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a specified traffic policy instance. After you submit a
// CreateTrafficPolicyInstance or an UpdateTrafficPolicyInstance request, there's
// a brief delay while Amazon Route 53 creates the resource record sets that are
// specified in the traffic policy definition. For more information, see the State
// response element. In the Route 53 console, traffic policy instances are known as
// policy records.
func (c *Client) GetTrafficPolicyInstance(ctx context.Context, params *GetTrafficPolicyInstanceInput, optFns ...func(*Options)) (*GetTrafficPolicyInstanceOutput, error) {
	if params == nil {
		params = &GetTrafficPolicyInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTrafficPolicyInstance", params, optFns, c.addOperationGetTrafficPolicyInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTrafficPolicyInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets information about a specified traffic policy instance.
type GetTrafficPolicyInstanceInput struct {

	// The ID of the traffic policy instance that you want to get information about.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

// A complex type that contains information about the resource record sets that
// Amazon Route 53 created based on a specified traffic policy.
type GetTrafficPolicyInstanceOutput struct {

	// A complex type that contains settings for the traffic policy instance.
	//
	// This member is required.
	TrafficPolicyInstance *types.TrafficPolicyInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTrafficPolicyInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetTrafficPolicyInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetTrafficPolicyInstance{}, middleware.After)
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
	if err = addOpGetTrafficPolicyInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTrafficPolicyInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTrafficPolicyInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetTrafficPolicyInstance",
	}
}
