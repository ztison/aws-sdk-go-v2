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

// Updates the identity information for the specified user. We strongly recommend
// limiting who has the ability to invoke UpdateUserIdentityInfo . Someone with
// that ability can change the login credentials of other users by changing their
// email address. This poses a security risk to your organization. They can change
// the email address of a user to the attacker's email address, and then reset the
// password through email. For more information, see Best Practices for Security
// Profiles (https://docs.aws.amazon.com/connect/latest/adminguide/security-profile-best-practices.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) UpdateUserIdentityInfo(ctx context.Context, params *UpdateUserIdentityInfoInput, optFns ...func(*Options)) (*UpdateUserIdentityInfoOutput, error) {
	if params == nil {
		params = &UpdateUserIdentityInfoInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUserIdentityInfo", params, optFns, c.addOperationUpdateUserIdentityInfoMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserIdentityInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserIdentityInfoInput struct {

	// The identity information for the user.
	//
	// This member is required.
	IdentityInfo *types.UserIdentityInfo

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The identifier of the user account.
	//
	// This member is required.
	UserId *string

	noSmithyDocumentSerde
}

type UpdateUserIdentityInfoOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateUserIdentityInfoMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateUserIdentityInfo{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateUserIdentityInfo{}, middleware.After)
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
	if err = addOpUpdateUserIdentityInfoValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserIdentityInfo(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateUserIdentityInfo(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateUserIdentityInfo",
	}
}
