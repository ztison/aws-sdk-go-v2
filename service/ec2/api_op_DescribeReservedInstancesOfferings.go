// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes Reserved Instance offerings that are available for purchase. With
// Reserved Instances, you purchase the right to launch instances for a period of
// time. During that time period, you do not receive insufficient capacity errors,
// and you pay a lower usage rate than the rate charged for On-Demand instances for
// the actual time used. If you have listed your own Reserved Instances for sale in
// the Reserved Instance Marketplace, they will be excluded from these results.
// This is to ensure that you do not purchase your own Reserved Instances. For more
// information, see Reserved Instance Marketplace (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html)
// in the Amazon EC2 User Guide.
func (c *Client) DescribeReservedInstancesOfferings(ctx context.Context, params *DescribeReservedInstancesOfferingsInput, optFns ...func(*Options)) (*DescribeReservedInstancesOfferingsOutput, error) {
	if params == nil {
		params = &DescribeReservedInstancesOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedInstancesOfferings", params, optFns, c.addOperationDescribeReservedInstancesOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedInstancesOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeReservedInstancesOfferings.
type DescribeReservedInstancesOfferingsInput struct {

	// The Availability Zone in which the Reserved Instance can be used.
	AvailabilityZone *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters.
	//   - availability-zone - The Availability Zone where the Reserved Instance can be
	//   used.
	//   - duration - The duration of the Reserved Instance (for example, one year or
	//   three years), in seconds ( 31536000 | 94608000 ).
	//   - fixed-price - The purchase price of the Reserved Instance (for example,
	//   9800.0).
	//   - instance-type - The instance type that is covered by the reservation.
	//   - marketplace - Set to true to show only Reserved Instance Marketplace
	//   offerings. When this filter is not used, which is the default behavior, all
	//   offerings from both Amazon Web Services and the Reserved Instance Marketplace
	//   are listed.
	//   - product-description - The Reserved Instance product platform description.
	//   Instances that include (Amazon VPC) in the product platform description will
	//   only be displayed to EC2-Classic account holders and are for use with Amazon
	//   VPC. ( Linux/UNIX | Linux/UNIX (Amazon VPC) | SUSE Linux | SUSE Linux (Amazon
	//   VPC) | Red Hat Enterprise Linux | Red Hat Enterprise Linux (Amazon VPC) | Red
	//   Hat Enterprise Linux with HA (Amazon VPC) | Windows | Windows (Amazon VPC) |
	//   Windows with SQL Server Standard | Windows with SQL Server Standard (Amazon
	//   VPC) | Windows with SQL Server Web | Windows with SQL Server Web (Amazon VPC)
	//   | Windows with SQL Server Enterprise | Windows with SQL Server Enterprise
	//   (Amazon VPC) )
	//   - reserved-instances-offering-id - The Reserved Instances offering ID.
	//   - scope - The scope of the Reserved Instance ( Availability Zone or Region ).
	//   - usage-price - The usage price of the Reserved Instance, per hour (for
	//   example, 0.84).
	Filters []types.Filter

	// Include Reserved Instance Marketplace offerings in the response.
	IncludeMarketplace *bool

	// The tenancy of the instances covered by the reservation. A Reserved Instance
	// with a tenancy of dedicated is applied to instances that run in a VPC on
	// single-tenant hardware (i.e., Dedicated Instances). Important: The host value
	// cannot be used with this parameter. Use the default or dedicated values only.
	// Default: default
	InstanceTenancy types.Tenancy

	// The instance type that the reservation will cover (for example, m1.small ). For
	// more information, see Instance types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html)
	// in the Amazon EC2 User Guide.
	InstanceType types.InstanceType

	// The maximum duration (in seconds) to filter when searching for offerings.
	// Default: 94608000 (3 years)
	MaxDuration *int64

	// The maximum number of instances to filter when searching for offerings.
	// Default: 20
	MaxInstanceCount *int32

	// The maximum number of results to return for the request in a single page. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. The maximum is 100. Default: 100
	MaxResults *int32

	// The minimum duration (in seconds) to filter when searching for offerings.
	// Default: 2592000 (1 month)
	MinDuration *int64

	// The token to retrieve the next page of results.
	NextToken *string

	// The offering class of the Reserved Instance. Can be standard or convertible .
	OfferingClass types.OfferingClassType

	// The Reserved Instance offering type. If you are using tools that predate the
	// 2011-11-01 API version, you only have access to the Medium Utilization Reserved
	// Instance offering type.
	OfferingType types.OfferingTypeValues

	// The Reserved Instance product platform description. Instances that include
	// (Amazon VPC) in the description are for use with Amazon VPC.
	ProductDescription types.RIProductDescription

	// One or more Reserved Instances offering IDs.
	ReservedInstancesOfferingIds []string

	noSmithyDocumentSerde
}

// Contains the output of DescribeReservedInstancesOfferings.
type DescribeReservedInstancesOfferingsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// A list of Reserved Instances offerings.
	ReservedInstancesOfferings []types.ReservedInstancesOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedInstancesOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeReservedInstancesOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeReservedInstancesOfferings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedInstancesOfferings(options.Region), middleware.Before); err != nil {
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

// DescribeReservedInstancesOfferingsAPIClient is a client that implements the
// DescribeReservedInstancesOfferings operation.
type DescribeReservedInstancesOfferingsAPIClient interface {
	DescribeReservedInstancesOfferings(context.Context, *DescribeReservedInstancesOfferingsInput, ...func(*Options)) (*DescribeReservedInstancesOfferingsOutput, error)
}

var _ DescribeReservedInstancesOfferingsAPIClient = (*Client)(nil)

// DescribeReservedInstancesOfferingsPaginatorOptions is the paginator options for
// DescribeReservedInstancesOfferings
type DescribeReservedInstancesOfferingsPaginatorOptions struct {
	// The maximum number of results to return for the request in a single page. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. The maximum is 100. Default: 100
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReservedInstancesOfferingsPaginator is a paginator for
// DescribeReservedInstancesOfferings
type DescribeReservedInstancesOfferingsPaginator struct {
	options   DescribeReservedInstancesOfferingsPaginatorOptions
	client    DescribeReservedInstancesOfferingsAPIClient
	params    *DescribeReservedInstancesOfferingsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReservedInstancesOfferingsPaginator returns a new
// DescribeReservedInstancesOfferingsPaginator
func NewDescribeReservedInstancesOfferingsPaginator(client DescribeReservedInstancesOfferingsAPIClient, params *DescribeReservedInstancesOfferingsInput, optFns ...func(*DescribeReservedInstancesOfferingsPaginatorOptions)) *DescribeReservedInstancesOfferingsPaginator {
	if params == nil {
		params = &DescribeReservedInstancesOfferingsInput{}
	}

	options := DescribeReservedInstancesOfferingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReservedInstancesOfferingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReservedInstancesOfferingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReservedInstancesOfferings page.
func (p *DescribeReservedInstancesOfferingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReservedInstancesOfferingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeReservedInstancesOfferings(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeReservedInstancesOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeReservedInstancesOfferings",
	}
}
