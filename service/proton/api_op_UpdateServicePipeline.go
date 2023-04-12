// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update the service pipeline. There are four modes for updating a service
// pipeline. The deploymentType field defines the mode. NONE In this mode, a
// deployment doesn't occur. Only the requested metadata parameters are updated.
// CURRENT_VERSION In this mode, the service pipeline is deployed and updated with
// the new spec that you provide. Only requested parameters are updated. Don’t
// include major or minor version parameters when you use this deployment-type .
// MINOR_VERSION In this mode, the service pipeline is deployed and updated with
// the published, recommended (latest) minor version of the current major version
// in use, by default. You can specify a different minor version of the current
// major version in use. MAJOR_VERSION In this mode, the service pipeline is
// deployed and updated with the published, recommended (latest) major and minor
// version of the current template by default. You can specify a different major
// version that's higher than the major version in use and a minor version.
func (c *Client) UpdateServicePipeline(ctx context.Context, params *UpdateServicePipelineInput, optFns ...func(*Options)) (*UpdateServicePipelineOutput, error) {
	if params == nil {
		params = &UpdateServicePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateServicePipeline", params, optFns, c.addOperationUpdateServicePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServicePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServicePipelineInput struct {

	// The deployment type. There are four modes for updating a service pipeline. The
	// deploymentType field defines the mode. NONE In this mode, a deployment doesn't
	// occur. Only the requested metadata parameters are updated. CURRENT_VERSION In
	// this mode, the service pipeline is deployed and updated with the new spec that
	// you provide. Only requested parameters are updated. Don’t include major or minor
	// version parameters when you use this deployment-type . MINOR_VERSION In this
	// mode, the service pipeline is deployed and updated with the published,
	// recommended (latest) minor version of the current major version in use, by
	// default. You can specify a different minor version of the current major version
	// in use. MAJOR_VERSION In this mode, the service pipeline is deployed and
	// updated with the published, recommended (latest) major and minor version of the
	// current template, by default. You can specify a different major version that's
	// higher than the major version in use and a minor version.
	//
	// This member is required.
	DeploymentType types.DeploymentUpdateType

	// The name of the service to that the pipeline is associated with.
	//
	// This member is required.
	ServiceName *string

	// The spec for the service pipeline to update.
	//
	// This value conforms to the media type: application/yaml
	//
	// This member is required.
	Spec *string

	// The major version of the service template that was used to create the service
	// that the pipeline is associated with.
	TemplateMajorVersion *string

	// The minor version of the service template that was used to create the service
	// that the pipeline is associated with.
	TemplateMinorVersion *string

	noSmithyDocumentSerde
}

type UpdateServicePipelineOutput struct {

	// The pipeline details that are returned by Proton.
	//
	// This member is required.
	Pipeline *types.ServicePipeline

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServicePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateServicePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateServicePipeline{}, middleware.After)
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
	if err = addOpUpdateServicePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServicePipeline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateServicePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "UpdateServicePipeline",
	}
}
