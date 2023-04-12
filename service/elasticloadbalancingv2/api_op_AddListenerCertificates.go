// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds the specified SSL server certificate to the certificate list for the
// specified HTTPS or TLS listener. If the certificate in already in the
// certificate list, the call is successful but the certificate is not added again.
// For more information, see HTTPS listeners (https://docs.aws.amazon.com/elasticloadbalancing/latest/application/create-https-listener.html)
// in the Application Load Balancers Guide or TLS listeners (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html)
// in the Network Load Balancers Guide.
func (c *Client) AddListenerCertificates(ctx context.Context, params *AddListenerCertificatesInput, optFns ...func(*Options)) (*AddListenerCertificatesOutput, error) {
	if params == nil {
		params = &AddListenerCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddListenerCertificates", params, optFns, c.addOperationAddListenerCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddListenerCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddListenerCertificatesInput struct {

	// The certificate to add. You can specify one certificate per call. Set
	// CertificateArn to the certificate ARN but do not set IsDefault .
	//
	// This member is required.
	Certificates []types.Certificate

	// The Amazon Resource Name (ARN) of the listener.
	//
	// This member is required.
	ListenerArn *string

	noSmithyDocumentSerde
}

type AddListenerCertificatesOutput struct {

	// Information about the certificates in the certificate list.
	Certificates []types.Certificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddListenerCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAddListenerCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAddListenerCertificates{}, middleware.After)
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
	if err = addOpAddListenerCertificatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddListenerCertificates(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddListenerCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "AddListenerCertificates",
	}
}
