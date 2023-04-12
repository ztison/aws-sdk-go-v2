// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a delegation signer (DS) record in the registry zone for this domain
// name.
func (c *Client) DisassociateDelegationSignerFromDomain(ctx context.Context, params *DisassociateDelegationSignerFromDomainInput, optFns ...func(*Options)) (*DisassociateDelegationSignerFromDomainOutput, error) {
	if params == nil {
		params = &DisassociateDelegationSignerFromDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateDelegationSignerFromDomain", params, optFns, c.addOperationDisassociateDelegationSignerFromDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateDelegationSignerFromDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateDelegationSignerFromDomainInput struct {

	// Name of the domain.
	//
	// This member is required.
	DomainName *string

	// An internal identification number assigned to each DS record after it’s
	// created. You can retrieve it as part of DNSSEC information returned by
	// GetDomainDetail (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetDomainDetail.html)
	// .
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DisassociateDelegationSignerFromDomainOutput struct {

	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html)
	// .
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateDelegationSignerFromDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateDelegationSignerFromDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateDelegationSignerFromDomain{}, middleware.After)
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
	if err = addOpDisassociateDelegationSignerFromDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateDelegationSignerFromDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateDelegationSignerFromDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "DisassociateDelegationSignerFromDomain",
	}
}
