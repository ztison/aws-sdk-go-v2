// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a description of the specified model package, which is used to create
// SageMaker models or list them on Amazon Web Services Marketplace. To create
// models in SageMaker, buyers can subscribe to model packages listed on Amazon Web
// Services Marketplace.
func (c *Client) DescribeModelPackage(ctx context.Context, params *DescribeModelPackageInput, optFns ...func(*Options)) (*DescribeModelPackageOutput, error) {
	if params == nil {
		params = &DescribeModelPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeModelPackage", params, optFns, c.addOperationDescribeModelPackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeModelPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeModelPackageInput struct {

	// The name or Amazon Resource Name (ARN) of the model package to describe. When
	// you specify a name, the name must have 1 to 63 characters. Valid characters are
	// a-z, A-Z, 0-9, and - (hyphen).
	//
	// This member is required.
	ModelPackageName *string

	noSmithyDocumentSerde
}

type DescribeModelPackageOutput struct {

	// A timestamp specifying when the model package was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the model package.
	//
	// This member is required.
	ModelPackageArn *string

	// The name of the model package being described.
	//
	// This member is required.
	ModelPackageName *string

	// The current status of the model package.
	//
	// This member is required.
	ModelPackageStatus types.ModelPackageStatus

	// Details about the current status of the model package.
	//
	// This member is required.
	ModelPackageStatusDetails *types.ModelPackageStatusDetails

	// An array of additional Inference Specification objects. Each additional
	// Inference Specification specifies artifacts based on this model package that can
	// be used on inference endpoints. Generally used with SageMaker Neo to store the
	// compiled artifacts.
	AdditionalInferenceSpecifications []types.AdditionalInferenceSpecificationDefinition

	// A description provided for the model approval.
	ApprovalDescription *string

	// Whether the model package is certified for listing on Amazon Web Services
	// Marketplace.
	CertifyForMarketplace bool

	// Information about the user who created or modified an experiment, trial, trial
	// component, lineage group, project, or model card.
	CreatedBy *types.UserContext

	// The metadata properties associated with the model package versions.
	CustomerMetadataProperties map[string]string

	// The machine learning domain of the model package you specified. Common machine
	// learning domains include computer vision and natural language processing.
	Domain *string

	// Represents the drift check baselines that can be used when the model monitor is
	// set using the model package. For more information, see the topic on Drift
	// Detection against Previous Baselines in SageMaker Pipelines (https://docs.aws.amazon.com/sagemaker/latest/dg/pipelines-quality-clarify-baseline-lifecycle.html#pipelines-quality-clarify-baseline-drift-detection)
	// in the Amazon SageMaker Developer Guide.
	DriftCheckBaselines *types.DriftCheckBaselines

	// Details about inference jobs that can be run with models based on this model
	// package.
	InferenceSpecification *types.InferenceSpecification

	// Information about the user who created or modified an experiment, trial, trial
	// component, lineage group, project, or model card.
	LastModifiedBy *types.UserContext

	// The last time that the model package was modified.
	LastModifiedTime *time.Time

	// Metadata properties of the tracking entity, trial, or trial component.
	MetadataProperties *types.MetadataProperties

	// The approval status of the model package.
	ModelApprovalStatus types.ModelApprovalStatus

	// Metrics for the model.
	ModelMetrics *types.ModelMetrics

	// A brief summary of the model package.
	ModelPackageDescription *string

	// If the model is a versioned model, the name of the model group that the
	// versioned model belongs to.
	ModelPackageGroupName *string

	// The version of the model package.
	ModelPackageVersion *int32

	// The Amazon Simple Storage Service (Amazon S3) path where the sample payload are
	// stored. This path points to a single gzip compressed tar archive (.tar.gz
	// suffix).
	SamplePayloadUrl *string

	// Details about the algorithm that was used to create the model package.
	SourceAlgorithmSpecification *types.SourceAlgorithmSpecification

	// The machine learning task you specified that your model package accomplishes.
	// Common machine learning tasks include object detection and image classification.
	Task *string

	// Configurations for one or more transform jobs that SageMaker runs to test the
	// model package.
	ValidationSpecification *types.ModelPackageValidationSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeModelPackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeModelPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeModelPackage{}, middleware.After)
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
	if err = addOpDescribeModelPackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeModelPackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeModelPackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeModelPackage",
	}
}
