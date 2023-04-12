// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a pending configuration change to a broker.
func (c *Client) UpdateBroker(ctx context.Context, params *UpdateBrokerInput, optFns ...func(*Options)) (*UpdateBrokerOutput, error) {
	if params == nil {
		params = &UpdateBrokerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBroker", params, optFns, c.addOperationUpdateBrokerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBrokerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates the broker using the specified properties.
type UpdateBrokerInput struct {

	// The unique ID that Amazon MQ generates for the broker.
	//
	// This member is required.
	BrokerId *string

	// Optional. The authentication strategy used to secure the broker. The default is
	// SIMPLE.
	AuthenticationStrategy types.AuthenticationStrategy

	// Enables automatic upgrades to new minor versions for brokers, as new versions
	// are released and supported by Amazon MQ. Automatic upgrades occur during the
	// scheduled maintenance window of the broker or after a manual broker reboot.
	AutoMinorVersionUpgrade bool

	// A list of information about the configuration.
	Configuration *types.ConfigurationId

	// The broker engine version. For a list of supported engine versions, see
	// Supported engines (https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/broker-engine.html)
	// .
	EngineVersion *string

	// The broker's host instance type to upgrade to. For a list of supported instance
	// types, see Broker instance types (https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/broker.html#broker-instance-types)
	// .
	HostInstanceType *string

	// Optional. The metadata of the LDAP server used to authenticate and authorize
	// connections to the broker. Does not apply to RabbitMQ brokers.
	LdapServerMetadata *types.LdapServerMetadataInput

	// Enables Amazon CloudWatch logging for brokers.
	Logs *types.Logs

	// The parameters that determine the WeeklyStartTime.
	MaintenanceWindowStartTime *types.WeeklyStartTime

	// The list of security groups (1 minimum, 5 maximum) that authorizes connections
	// to brokers.
	SecurityGroups []string

	noSmithyDocumentSerde
}

type UpdateBrokerOutput struct {

	// Optional. The authentication strategy used to secure the broker. The default is
	// SIMPLE.
	AuthenticationStrategy types.AuthenticationStrategy

	// The new boolean value that specifies whether broker engines automatically
	// upgrade to new minor versions as new versions are released and supported by
	// Amazon MQ.
	AutoMinorVersionUpgrade bool

	// Required. The unique ID that Amazon MQ generates for the broker.
	BrokerId *string

	// The ID of the updated configuration.
	Configuration *types.ConfigurationId

	// The broker engine version to upgrade to. For a list of supported engine
	// versions, see Supported engines (https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/broker-engine.html)
	// .
	EngineVersion *string

	// The broker's host instance type to upgrade to. For a list of supported instance
	// types, see Broker instance types (https://docs.aws.amazon.com//amazon-mq/latest/developer-guide/broker.html#broker-instance-types)
	// .
	HostInstanceType *string

	// Optional. The metadata of the LDAP server used to authenticate and authorize
	// connections to the broker. Does not apply to RabbitMQ brokers.
	LdapServerMetadata *types.LdapServerMetadataOutput

	// The list of information about logs to be enabled for the specified broker.
	Logs *types.Logs

	// The parameters that determine the WeeklyStartTime.
	MaintenanceWindowStartTime *types.WeeklyStartTime

	// The list of security groups (1 minimum, 5 maximum) that authorizes connections
	// to brokers.
	SecurityGroups []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBrokerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBroker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBroker{}, middleware.After)
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
	if err = addOpUpdateBrokerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBroker(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBroker(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "UpdateBroker",
	}
}
