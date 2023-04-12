// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the user attributes and metadata for a user.
func (c *Client) GetUser(ctx context.Context, params *GetUserInput, optFns ...func(*Options)) (*GetUserOutput, error) {
	if params == nil {
		params = &GetUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUser", params, optFns, c.addOperationGetUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to get information about the user.
type GetUserInput struct {

	// A non-expired access token for the user whose information you want to query.
	//
	// This member is required.
	AccessToken *string

	noSmithyDocumentSerde
}

// Represents the response from the server from the request to get information
// about the user.
type GetUserOutput struct {

	// An array of name-value pairs representing user attributes. For custom
	// attributes, you must prepend the custom: prefix to the attribute name.
	//
	// This member is required.
	UserAttributes []types.AttributeType

	// The user name of the user you want to retrieve from the get user request.
	//
	// This member is required.
	Username *string

	// This response parameter is no longer supported. It provides information only
	// about SMS MFA configurations. It doesn't provide information about time-based
	// one-time password (TOTP) software token MFA configurations. To look up
	// information about either type of MFA configuration, use UserMFASettingList
	// instead.
	MFAOptions []types.MFAOptionType

	// The user's preferred MFA setting.
	PreferredMfaSetting *string

	// The MFA options that are activated for the user. The possible values in this
	// list are SMS_MFA and SOFTWARE_TOKEN_MFA .
	UserMFASettingList []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUser{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = addOpGetUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUser",
	}
}
