// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deregisters an account in Audit Manager. Before you deregister, you can use the
// UpdateSettings (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_UpdateSettings.html)
// API operation to set your preferred data retention policy. By default, Audit
// Manager retains your data. If you want to delete your data, you can use the
// DeregistrationPolicy attribute to request the deletion of your data. For more
// information about data retention, see Data Protection (https://docs.aws.amazon.com/audit-manager/latest/userguide/data-protection.html)
// in the Audit Manager User Guide.
func (c *Client) DeregisterAccount(ctx context.Context, params *DeregisterAccountInput, optFns ...func(*Options)) (*DeregisterAccountOutput, error) {
	if params == nil {
		params = &DeregisterAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterAccount", params, optFns, c.addOperationDeregisterAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterAccountInput struct {
	noSmithyDocumentSerde
}

type DeregisterAccountOutput struct {

	// The registration status of the account.
	Status types.AccountStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeregisterAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeregisterAccount{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "DeregisterAccount",
	}
}
