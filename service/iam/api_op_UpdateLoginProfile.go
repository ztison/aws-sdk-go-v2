// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the password for the specified IAM user. You can use the CLI, the
// Amazon Web Services API, or the Users page in the IAM console to change the
// password for any IAM user. Use ChangePassword to change your own password in
// the My Security Credentials page in the Amazon Web Services Management Console.
// For more information about modifying passwords, see Managing passwords (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html)
// in the IAM User Guide.
func (c *Client) UpdateLoginProfile(ctx context.Context, params *UpdateLoginProfileInput, optFns ...func(*Options)) (*UpdateLoginProfileOutput, error) {
	if params == nil {
		params = &UpdateLoginProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLoginProfile", params, optFns, c.addOperationUpdateLoginProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLoginProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLoginProfileInput struct {

	// The name of the user whose password you want to update. This parameter allows
	// (through its regex pattern (http://wikipedia.org/wiki/regex) ) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string

	// The new password for the specified IAM user. The regex pattern (http://wikipedia.org/wiki/regex)
	// used to validate this parameter is a string of characters consisting of the
	// following:
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	// However, the format can be further restricted by the account administrator by
	// setting a password policy on the Amazon Web Services account. For more
	// information, see UpdateAccountPasswordPolicy .
	Password *string

	// Allows this new password to be used only once by requiring the specified IAM
	// user to set a new password on next sign-in.
	PasswordResetRequired *bool

	noSmithyDocumentSerde
}

type UpdateLoginProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLoginProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateLoginProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateLoginProfile{}, middleware.After)
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
	if err = addOpUpdateLoginProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLoginProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLoginProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UpdateLoginProfile",
	}
}
