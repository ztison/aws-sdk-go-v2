// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanagerusersubscriptions

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanagerusersubscriptions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates additional product configuration settings for the registered identity
// provider.
func (c *Client) UpdateIdentityProviderSettings(ctx context.Context, params *UpdateIdentityProviderSettingsInput, optFns ...func(*Options)) (*UpdateIdentityProviderSettingsOutput, error) {
	if params == nil {
		params = &UpdateIdentityProviderSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIdentityProviderSettings", params, optFns, c.addOperationUpdateIdentityProviderSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIdentityProviderSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIdentityProviderSettingsInput struct {

	// Details about an identity provider.
	//
	// This member is required.
	IdentityProvider types.IdentityProvider

	// The name of the user-based subscription product.
	//
	// This member is required.
	Product *string

	// Updates the registered identity provider’s product related configuration
	// settings. You can update any combination of settings in a single operation such
	// as the:
	//   - Subnets which you want to add to provision VPC endpoints.
	//   - Subnets which you want to remove the VPC endpoints from.
	//   - Security group ID which permits traffic to the VPC endpoints.
	//
	// This member is required.
	UpdateSettings *types.UpdateSettings

	noSmithyDocumentSerde
}

type UpdateIdentityProviderSettingsOutput struct {

	// Describes an identity provider.
	//
	// This member is required.
	IdentityProviderSummary *types.IdentityProviderSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIdentityProviderSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateIdentityProviderSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateIdentityProviderSettings{}, middleware.After)
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
	if err = addOpUpdateIdentityProviderSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIdentityProviderSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateIdentityProviderSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager-user-subscriptions",
		OperationName: "UpdateIdentityProviderSettings",
	}
}
