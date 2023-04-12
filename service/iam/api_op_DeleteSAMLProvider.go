// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a SAML provider resource in IAM. Deleting the provider resource from
// IAM does not update any roles that reference the SAML provider resource's ARN as
// a principal in their trust policies. Any attempt to assume a role that
// references a non-existent provider resource ARN fails. This operation requires
// Signature Version 4 (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html)
// .
func (c *Client) DeleteSAMLProvider(ctx context.Context, params *DeleteSAMLProviderInput, optFns ...func(*Options)) (*DeleteSAMLProviderOutput, error) {
	if params == nil {
		params = &DeleteSAMLProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSAMLProvider", params, optFns, c.addOperationDeleteSAMLProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSAMLProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSAMLProviderInput struct {

	// The Amazon Resource Name (ARN) of the SAML provider to delete.
	//
	// This member is required.
	SAMLProviderArn *string

	noSmithyDocumentSerde
}

type DeleteSAMLProviderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSAMLProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteSAMLProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteSAMLProvider{}, middleware.After)
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
	if err = addOpDeleteSAMLProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSAMLProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSAMLProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteSAMLProvider",
	}
}
