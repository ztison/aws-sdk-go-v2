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

// Modifies the permissions for your VPC endpoint service. You can add or remove
// permissions for service consumers (Amazon Web Services accounts, users, and IAM
// roles) to connect to your endpoint service. If you grant permissions to all
// principals, the service is public. Any users who know the name of a public
// service can send a request to attach an endpoint. If the service does not
// require manual approval, attachments are automatically approved.
func (c *Client) ModifyVpcEndpointServicePermissions(ctx context.Context, params *ModifyVpcEndpointServicePermissionsInput, optFns ...func(*Options)) (*ModifyVpcEndpointServicePermissionsOutput, error) {
	if params == nil {
		params = &ModifyVpcEndpointServicePermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpcEndpointServicePermissions", params, optFns, c.addOperationModifyVpcEndpointServicePermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpcEndpointServicePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpcEndpointServicePermissionsInput struct {

	// The ID of the service.
	//
	// This member is required.
	ServiceId *string

	// The Amazon Resource Names (ARN) of the principals. Permissions are granted to
	// the principals in this list. To grant permissions to all principals, specify an
	// asterisk (*).
	AddAllowedPrincipals []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The Amazon Resource Names (ARN) of the principals. Permissions are revoked for
	// principals in this list.
	RemoveAllowedPrincipals []string

	noSmithyDocumentSerde
}

type ModifyVpcEndpointServicePermissionsOutput struct {

	// Information about the added principals.
	AddedPrincipals []types.AddedPrincipal

	// Returns true if the request succeeds; otherwise, it returns an error.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyVpcEndpointServicePermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpcEndpointServicePermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpcEndpointServicePermissions{}, middleware.After)
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
	if err = addOpModifyVpcEndpointServicePermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpcEndpointServicePermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyVpcEndpointServicePermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVpcEndpointServicePermissions",
	}
}
