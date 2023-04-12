// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns existing recommendation preferences, such as enhanced infrastructure
// metrics. Use the scope parameter to specify which preferences to return. You
// can specify to return preferences for an organization, a specific account ID, or
// a specific EC2 instance or Auto Scaling group Amazon Resource Name (ARN). For
// more information, see Activating enhanced infrastructure metrics (https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html)
// in the Compute Optimizer User Guide.
func (c *Client) GetRecommendationPreferences(ctx context.Context, params *GetRecommendationPreferencesInput, optFns ...func(*Options)) (*GetRecommendationPreferencesOutput, error) {
	if params == nil {
		params = &GetRecommendationPreferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRecommendationPreferences", params, optFns, c.addOperationGetRecommendationPreferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRecommendationPreferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRecommendationPreferencesInput struct {

	// The target resource type of the recommendation preference for which to return
	// preferences. The Ec2Instance option encompasses standalone instances and
	// instances that are part of Auto Scaling groups. The AutoScalingGroup option
	// encompasses only instances that are part of an Auto Scaling group. The valid
	// values for this parameter are Ec2Instance and AutoScalingGroup .
	//
	// This member is required.
	ResourceType types.ResourceType

	// The maximum number of recommendation preferences to return with a single
	// request. To retrieve the remaining results, make another request with the
	// returned nextToken value.
	MaxResults *int32

	// The token to advance to the next page of recommendation preferences.
	NextToken *string

	// An object that describes the scope of the recommendation preference to return.
	// You can return recommendation preferences that are created at the organization
	// level (for management accounts of an organization only), account level, and
	// resource level. For more information, see Activating enhanced infrastructure
	// metrics (https://docs.aws.amazon.com/compute-optimizer/latest/ug/enhanced-infrastructure-metrics.html)
	// in the Compute Optimizer User Guide.
	Scope *types.Scope

	noSmithyDocumentSerde
}

type GetRecommendationPreferencesOutput struct {

	// The token to use to advance to the next page of recommendation preferences.
	// This value is null when there are no more pages of recommendation preferences to
	// return.
	NextToken *string

	// An array of objects that describe recommendation preferences.
	RecommendationPreferencesDetails []types.RecommendationPreferencesDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRecommendationPreferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetRecommendationPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetRecommendationPreferences{}, middleware.After)
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
	if err = addOpGetRecommendationPreferencesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecommendationPreferences(options.Region), middleware.Before); err != nil {
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

// GetRecommendationPreferencesAPIClient is a client that implements the
// GetRecommendationPreferences operation.
type GetRecommendationPreferencesAPIClient interface {
	GetRecommendationPreferences(context.Context, *GetRecommendationPreferencesInput, ...func(*Options)) (*GetRecommendationPreferencesOutput, error)
}

var _ GetRecommendationPreferencesAPIClient = (*Client)(nil)

// GetRecommendationPreferencesPaginatorOptions is the paginator options for
// GetRecommendationPreferences
type GetRecommendationPreferencesPaginatorOptions struct {
	// The maximum number of recommendation preferences to return with a single
	// request. To retrieve the remaining results, make another request with the
	// returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetRecommendationPreferencesPaginator is a paginator for
// GetRecommendationPreferences
type GetRecommendationPreferencesPaginator struct {
	options   GetRecommendationPreferencesPaginatorOptions
	client    GetRecommendationPreferencesAPIClient
	params    *GetRecommendationPreferencesInput
	nextToken *string
	firstPage bool
}

// NewGetRecommendationPreferencesPaginator returns a new
// GetRecommendationPreferencesPaginator
func NewGetRecommendationPreferencesPaginator(client GetRecommendationPreferencesAPIClient, params *GetRecommendationPreferencesInput, optFns ...func(*GetRecommendationPreferencesPaginatorOptions)) *GetRecommendationPreferencesPaginator {
	if params == nil {
		params = &GetRecommendationPreferencesInput{}
	}

	options := GetRecommendationPreferencesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetRecommendationPreferencesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetRecommendationPreferencesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetRecommendationPreferences page.
func (p *GetRecommendationPreferencesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetRecommendationPreferencesOutput, error) {
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

	result, err := p.client.GetRecommendationPreferences(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetRecommendationPreferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "GetRecommendationPreferences",
	}
}
