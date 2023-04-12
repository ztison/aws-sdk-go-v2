// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a PackageDescription (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageDescription.html)
// object that contains information about the requested package.
func (c *Client) DescribePackage(ctx context.Context, params *DescribePackageInput, optFns ...func(*Options)) (*DescribePackageOutput, error) {
	if params == nil {
		params = &DescribePackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePackage", params, optFns, c.addOperationDescribePackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePackageInput struct {

	// The name of the domain that contains the repository that contains the package.
	//
	// This member is required.
	Domain *string

	// A format that specifies the type of the requested package.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the requested package.
	//
	// This member is required.
	Package *string

	// The name of the repository that contains the requested package.
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The namespace of the requested package. The package component that specifies
	// its namespace depends on its type. For example:
	//   - The namespace of a Maven package is its groupId . The namespace is required
	//   when requesting Maven packages.
	//   - The namespace of an npm package is its scope .
	//   - Python and NuGet packages do not contain a corresponding component,
	//   packages of those formats do not have a namespace.
	//   - The namespace of a generic package is its namespace .
	Namespace *string

	noSmithyDocumentSerde
}

type DescribePackageOutput struct {

	// A PackageDescription (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageDescription.html)
	// object that contains information about the requested package.
	//
	// This member is required.
	Package *types.PackageDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribePackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribePackage{}, middleware.After)
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
	if err = addOpDescribePackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "DescribePackage",
	}
}
