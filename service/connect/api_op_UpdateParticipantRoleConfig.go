// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates timeouts for when human chat participants are to be considered idle,
// and when agents are automatically disconnected from a chat due to idleness. You
// can set four timers:
//   - Customer idle timeout
//   - Customer auto-disconnect timeout
//   - Agent idle timeout
//   - Agent auto-disconnect timeout
//
// For more information about how chat timeouts work, see Set up chat timeouts for
// human participants (https://docs.aws.amazon.com/connect/latest/adminguide/setup-chat-timeouts.html)
// .
func (c *Client) UpdateParticipantRoleConfig(ctx context.Context, params *UpdateParticipantRoleConfigInput, optFns ...func(*Options)) (*UpdateParticipantRoleConfigOutput, error) {
	if params == nil {
		params = &UpdateParticipantRoleConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateParticipantRoleConfig", params, optFns, c.addOperationUpdateParticipantRoleConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateParticipantRoleConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateParticipantRoleConfigInput struct {

	// The Amazon Connect channel you want to configure.
	//
	// This member is required.
	ChannelConfiguration types.UpdateParticipantRoleConfigChannelInfo

	// The identifier of the contact in this instance of Amazon Connect.
	//
	// This member is required.
	ContactId *string

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	noSmithyDocumentSerde
}

type UpdateParticipantRoleConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateParticipantRoleConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateParticipantRoleConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateParticipantRoleConfig{}, middleware.After)
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
	if err = addOpUpdateParticipantRoleConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateParticipantRoleConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateParticipantRoleConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateParticipantRoleConfig",
	}
}
