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

// Enables additional CloudWatch metrics for the specified CloudFront
// distribution. The additional metrics incur an additional cost. For more
// information, see Viewing additional CloudFront distribution metrics (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/viewing-cloudfront-metrics.html#monitoring-console.distributions-additional)
// in the Amazon CloudFront Developer Guide.
func (c *Client) CreateMonitoringSubscription(ctx context.Context, params *CreateMonitoringSubscriptionInput, optFns ...func(*Options)) (*CreateMonitoringSubscriptionOutput, error) {
	if params == nil {
		params = &CreateMonitoringSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMonitoringSubscription", params, optFns, c.addOperationCreateMonitoringSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMonitoringSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMonitoringSubscriptionInput struct {

	// The ID of the distribution that you are enabling metrics for.
	//
	// This member is required.
	DistributionId *string

	// A monitoring subscription. This structure contains information about whether
	// additional CloudWatch metrics are enabled for a given CloudFront distribution.
	//
	// This member is required.
	MonitoringSubscription *types.MonitoringSubscription

	noSmithyDocumentSerde
}

type CreateMonitoringSubscriptionOutput struct {

	// A monitoring subscription. This structure contains information about whether
	// additional CloudWatch metrics are enabled for a given CloudFront distribution.
	MonitoringSubscription *types.MonitoringSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMonitoringSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateMonitoringSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateMonitoringSubscription{}, middleware.After)
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
	if err = addOpCreateMonitoringSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMonitoringSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMonitoringSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateMonitoringSubscription",
	}
}
