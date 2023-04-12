// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Subscribes an endpoint to an Amazon SNS topic. If the endpoint type is HTTP/S
// or email, or if the endpoint and the topic are not in the same Amazon Web
// Services account, the endpoint owner must run the ConfirmSubscription action to
// confirm the subscription. You call the ConfirmSubscription action with the
// token from the subscription response. Confirmation tokens are valid for three
// days. This action is throttled at 100 transactions per second (TPS).
func (c *Client) Subscribe(ctx context.Context, params *SubscribeInput, optFns ...func(*Options)) (*SubscribeOutput, error) {
	if params == nil {
		params = &SubscribeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Subscribe", params, optFns, c.addOperationSubscribeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SubscribeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for Subscribe action.
type SubscribeInput struct {

	// The protocol that you want to use. Supported protocols include:
	//   - http – delivery of JSON-encoded message via HTTP POST
	//   - https – delivery of JSON-encoded message via HTTPS POST
	//   - email – delivery of message via SMTP
	//   - email-json – delivery of JSON-encoded message via SMTP
	//   - sms – delivery of message via SMS
	//   - sqs – delivery of JSON-encoded message to an Amazon SQS queue
	//   - application – delivery of JSON-encoded message to an EndpointArn for a
	//   mobile app and device
	//   - lambda – delivery of JSON-encoded message to an Lambda function
	//   - firehose – delivery of JSON-encoded message to an Amazon Kinesis Data
	//   Firehose delivery stream.
	//
	// This member is required.
	Protocol *string

	// The ARN of the topic you want to subscribe to.
	//
	// This member is required.
	TopicArn *string

	// A map of attributes with their corresponding values. The following lists the
	// names, descriptions, and values of the special request parameters that the
	// Subscribe action uses:
	//   - DeliveryPolicy – The policy that defines how Amazon SNS retries failed
	//   deliveries to HTTP/S endpoints.
	//   - FilterPolicy – The simple JSON object that lets your subscriber receive only
	//   a subset of messages, rather than receiving every message published to the
	//   topic.
	//   - FilterPolicyScope – This attribute lets you choose the filtering scope by
	//   using one of the following string value types:
	//   - MessageAttributes (default) – The filter is applied on the message
	//   attributes.
	//   - MessageBody – The filter is applied on the message body.
	//   - RawMessageDelivery – When set to true , enables raw message delivery to
	//   Amazon SQS or HTTP/S endpoints. This eliminates the need for the endpoints to
	//   process JSON formatting, which is otherwise created for Amazon SNS metadata.
	//   - RedrivePolicy – When specified, sends undeliverable messages to the
	//   specified Amazon SQS dead-letter queue. Messages that can't be delivered due to
	//   client errors (for example, when the subscribed endpoint is unreachable) or
	//   server errors (for example, when the service that powers the subscribed endpoint
	//   becomes unavailable) are held in the dead-letter queue for further analysis or
	//   reprocessing.
	// The following attribute applies only to Amazon Kinesis Data Firehose delivery
	// stream subscriptions:
	//   - SubscriptionRoleArn – The ARN of the IAM role that has the following:
	//   - Permission to write to the Kinesis Data Firehose delivery stream
	//   - Amazon SNS listed as a trusted entity Specifying a valid ARN for this
	//   attribute is required for Kinesis Data Firehose delivery stream subscriptions.
	//   For more information, see Fanout to Kinesis Data Firehose delivery streams (https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html)
	//   in the Amazon SNS Developer Guide.
	Attributes map[string]string

	// The endpoint that you want to receive notifications. Endpoints vary by
	// protocol:
	//   - For the http protocol, the (public) endpoint is a URL beginning with http://
	//   .
	//   - For the https protocol, the (public) endpoint is a URL beginning with
	//   https:// .
	//   - For the email protocol, the endpoint is an email address.
	//   - For the email-json protocol, the endpoint is an email address.
	//   - For the sms protocol, the endpoint is a phone number of an SMS-enabled
	//   device.
	//   - For the sqs protocol, the endpoint is the ARN of an Amazon SQS queue.
	//   - For the application protocol, the endpoint is the EndpointArn of a mobile
	//   app and device.
	//   - For the lambda protocol, the endpoint is the ARN of an Lambda function.
	//   - For the firehose protocol, the endpoint is the ARN of an Amazon Kinesis Data
	//   Firehose delivery stream.
	Endpoint *string

	// Sets whether the response from the Subscribe request includes the subscription
	// ARN, even if the subscription is not yet confirmed. If you set this parameter to
	// true , the response includes the ARN in all cases, even if the subscription is
	// not yet confirmed. In addition to the ARN for confirmed subscriptions, the
	// response also includes the pending subscription ARN value for subscriptions
	// that aren't yet confirmed. A subscription becomes confirmed when the subscriber
	// calls the ConfirmSubscription action with a confirmation token. The default
	// value is false .
	ReturnSubscriptionArn bool

	noSmithyDocumentSerde
}

// Response for Subscribe action.
type SubscribeOutput struct {

	// The ARN of the subscription if it is confirmed, or the string "pending
	// confirmation" if the subscription requires confirmation. However, if the API
	// request parameter ReturnSubscriptionArn is true, then the value is always the
	// subscription ARN, even if the subscription requires confirmation.
	SubscriptionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSubscribeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSubscribe{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSubscribe{}, middleware.After)
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
	if err = addOpSubscribeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSubscribe(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSubscribe(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "Subscribe",
	}
}
