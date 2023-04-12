// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/outposts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a site for an Outpost.
func (c *Client) CreateSite(ctx context.Context, params *CreateSiteInput, optFns ...func(*Options)) (*CreateSiteOutput, error) {
	if params == nil {
		params = &CreateSiteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSite", params, optFns, c.addOperationCreateSiteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSiteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSiteInput struct {

	// The name of the site.
	//
	// This member is required.
	Name *string

	// The description of the site.
	Description *string

	// Additional information that you provide about site access requirements,
	// electrician scheduling, personal protective equipment, or regulation of
	// equipment materials that could affect your installation process.
	Notes *string

	// The location to install and power on the hardware. This address might be
	// different from the shipping address.
	OperatingAddress *types.Address

	// Information about the physical and logistical details for the rack at this
	// site. For more information about hardware requirements for racks, see Network
	// readiness checklist (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-requirements.html#checklist)
	// in the Amazon Web Services Outposts User Guide.
	RackPhysicalProperties *types.RackPhysicalProperties

	// The location to ship the hardware. This address might be different from the
	// operating address.
	ShippingAddress *types.Address

	// The tags to apply to a site.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateSiteOutput struct {

	// Information about a site.
	Site *types.Site

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSiteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSite{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSite{}, middleware.After)
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
	if err = addOpCreateSiteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSite(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSite(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "outposts",
		OperationName: "CreateSite",
	}
}
