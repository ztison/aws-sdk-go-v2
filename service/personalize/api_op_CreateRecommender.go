// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a recommender with the recipe (a Domain dataset group use case) you
// specify. You create recommenders for a Domain dataset group and specify the
// recommender's Amazon Resource Name (ARN) when you make a GetRecommendations (https://docs.aws.amazon.com/personalize/latest/dg/API_RS_GetRecommendations.html)
// request. Minimum recommendation requests per second When you create a
// recommender, you can configure the recommender's minimum recommendation requests
// per second. The minimum recommendation requests per second (
// minRecommendationRequestsPerSecond ) specifies the baseline recommendation
// request throughput provisioned by Amazon Personalize. The default
// minRecommendationRequestsPerSecond is 1 . A recommendation request is a single
// GetRecommendations operation. Request throughput is measured in requests per
// second and Amazon Personalize uses your requests per second to derive your
// requests per hour and the price of your recommender usage. If your requests per
// second increases beyond minRecommendationRequestsPerSecond , Amazon Personalize
// auto-scales the provisioned capacity up and down, but never below
// minRecommendationRequestsPerSecond . There's a short time delay while the
// capacity is increased that might cause loss of requests. Your bill is the
// greater of either the minimum requests per hour (based on
// minRecommendationRequestsPerSecond) or the actual number of requests. The actual
// request throughput used is calculated as the average requests/second within a
// one-hour window. We recommend starting with the default
// minRecommendationRequestsPerSecond , track your usage using Amazon CloudWatch
// metrics, and then increase the minRecommendationRequestsPerSecond as necessary.
// Status A recommender can be in one of the following states:
//   - CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//   - STOP PENDING > STOP IN_PROGRESS > INACTIVE > START PENDING > START
//     IN_PROGRESS > ACTIVE
//   - DELETE PENDING > DELETE IN_PROGRESS
//
// To get the recommender status, call DescribeRecommender (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeRecommender.html)
// . Wait until the status of the recommender is ACTIVE before asking the
// recommender for recommendations. Related APIs
//   - ListRecommenders (https://docs.aws.amazon.com/personalize/latest/dg/API_ListRecommenders.html)
//   - DescribeRecommender (https://docs.aws.amazon.com/personalize/latest/dg/API_DescribeRecommender.html)
//   - UpdateRecommender (https://docs.aws.amazon.com/personalize/latest/dg/API_UpdateRecommender.html)
//   - DeleteRecommender (https://docs.aws.amazon.com/personalize/latest/dg/API_DeleteRecommender.html)
func (c *Client) CreateRecommender(ctx context.Context, params *CreateRecommenderInput, optFns ...func(*Options)) (*CreateRecommenderOutput, error) {
	if params == nil {
		params = &CreateRecommenderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRecommender", params, optFns, c.addOperationCreateRecommenderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRecommenderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRecommenderInput struct {

	// The Amazon Resource Name (ARN) of the destination domain dataset group for the
	// recommender.
	//
	// This member is required.
	DatasetGroupArn *string

	// The name of the recommender.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the recipe that the recommender will use. For
	// a recommender, a recipe is a Domain dataset group use case. Only Domain dataset
	// group use cases can be used to create a recommender. For information about use
	// cases see Choosing recommender use cases (https://docs.aws.amazon.com/personalize/latest/dg/domain-use-cases.html)
	// .
	//
	// This member is required.
	RecipeArn *string

	// The configuration details of the recommender.
	RecommenderConfig *types.RecommenderConfig

	// A list of tags (https://docs.aws.amazon.com/personalize/latest/dev/tagging-resources.html)
	// to apply to the recommender.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateRecommenderOutput struct {

	// The Amazon Resource Name (ARN) of the recommender.
	RecommenderArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRecommenderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRecommender{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRecommender{}, middleware.After)
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
	if err = addOpCreateRecommenderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRecommender(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRecommender(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateRecommender",
	}
}
