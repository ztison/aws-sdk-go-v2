// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Prevents the user from signing in with the specified external (SAML or social)
// identity provider (IdP). If the user that you want to deactivate is a Amazon
// Cognito user pools native username + password user, they can't use their
// password to sign in. If the user to deactivate is a linked external IdP user,
// any link between that user and an existing user is removed. When the external
// user signs in again, and the user is no longer attached to the previously linked
// DestinationUser , the user must create a new user account. See
// AdminLinkProviderForUser (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AdminLinkProviderForUser.html)
// . This action is enabled only for admin access and requires developer
// credentials. The ProviderName must match the value specified when creating an
// IdP for the pool. To deactivate a native username + password user, the
// ProviderName value must be Cognito and the ProviderAttributeName must be
// Cognito_Subject . The ProviderAttributeValue must be the name that is used in
// the user pool for the user. The ProviderAttributeName must always be
// Cognito_Subject for social IdPs. The ProviderAttributeValue must always be the
// exact subject that was used when the user was originally linked as a source
// user. For de-linking a SAML identity, there are two scenarios. If the linked
// identity has not yet been used to sign in, the ProviderAttributeName and
// ProviderAttributeValue must be the same values that were used for the SourceUser
// when the identities were originally linked using AdminLinkProviderForUser call.
// (If the linking was done with ProviderAttributeName set to Cognito_Subject , the
// same applies here). However, if the user has already signed in, the
// ProviderAttributeName must be Cognito_Subject and ProviderAttributeValue must
// be the subject of the SAML assertion.
func (c *Client) AdminDisableProviderForUser(ctx context.Context, params *AdminDisableProviderForUserInput, optFns ...func(*Options)) (*AdminDisableProviderForUserOutput, error) {
	if params == nil {
		params = &AdminDisableProviderForUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminDisableProviderForUser", params, optFns, c.addOperationAdminDisableProviderForUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminDisableProviderForUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminDisableProviderForUserInput struct {

	// The user to be disabled.
	//
	// This member is required.
	User *types.ProviderUserIdentifierType

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	noSmithyDocumentSerde
}

type AdminDisableProviderForUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminDisableProviderForUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminDisableProviderForUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminDisableProviderForUser{}, middleware.After)
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
	if err = addOpAdminDisableProviderForUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminDisableProviderForUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminDisableProviderForUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminDisableProviderForUser",
	}
}
