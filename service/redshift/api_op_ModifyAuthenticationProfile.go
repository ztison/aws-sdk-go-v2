// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies an authentication profile.
func (c *Client) ModifyAuthenticationProfile(ctx context.Context, params *ModifyAuthenticationProfileInput, optFns ...func(*Options)) (*ModifyAuthenticationProfileOutput, error) {
	if params == nil {
		params = &ModifyAuthenticationProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyAuthenticationProfile", params, optFns, c.addOperationModifyAuthenticationProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyAuthenticationProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyAuthenticationProfileInput struct {

	// The new content of the authentication profile in JSON format. The maximum
	// length of the JSON string is determined by a quota for your account.
	//
	// This member is required.
	AuthenticationProfileContent *string

	// The name of the authentication profile to replace.
	//
	// This member is required.
	AuthenticationProfileName *string

	noSmithyDocumentSerde
}

type ModifyAuthenticationProfileOutput struct {

	// The updated content of the authentication profile in JSON format.
	AuthenticationProfileContent *string

	// The name of the authentication profile that was replaced.
	AuthenticationProfileName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyAuthenticationProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyAuthenticationProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyAuthenticationProfile{}, middleware.After)
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
	if err = addOpModifyAuthenticationProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyAuthenticationProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyAuthenticationProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifyAuthenticationProfile",
	}
}
