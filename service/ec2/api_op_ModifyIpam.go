// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modify the configurations of an IPAM.
func (c *Client) ModifyIpam(ctx context.Context, params *ModifyIpamInput, optFns ...func(*Options)) (*ModifyIpamOutput, error) {
	if params == nil {
		params = &ModifyIpamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyIpam", params, optFns, c.addOperationModifyIpamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyIpamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyIpamInput struct {

	// The ID of the IPAM you want to modify.
	//
	// This member is required.
	IpamId *string

	// Choose the operating Regions for the IPAM. Operating Regions are Amazon Web
	// Services Regions where the IPAM is allowed to manage IP address CIDRs. IPAM only
	// discovers and monitors resources in the Amazon Web Services Regions you select
	// as operating Regions. For more information about operating Regions, see Create
	// an IPAM (https://docs.aws.amazon.com/vpc/latest/ipam/create-ipam.html) in the
	// Amazon VPC IPAM User Guide.
	AddOperatingRegions []types.AddIpamOperatingRegion

	// The description of the IPAM you want to modify.
	Description *string

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The operating Regions to remove.
	RemoveOperatingRegions []types.RemoveIpamOperatingRegion

	noSmithyDocumentSerde
}

type ModifyIpamOutput struct {

	// The results of the modification.
	Ipam *types.Ipam

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyIpamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyIpam{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyIpam{}, middleware.After)
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
	if err = addOpModifyIpamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyIpam(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyIpam(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyIpam",
	}
}
