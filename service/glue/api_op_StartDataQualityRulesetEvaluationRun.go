// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Once you have a ruleset definition (either recommended or your own), you call
// this operation to evaluate the ruleset against a data source (Glue table). The
// evaluation computes results which you can retrieve with the GetDataQualityResult
// API.
func (c *Client) StartDataQualityRulesetEvaluationRun(ctx context.Context, params *StartDataQualityRulesetEvaluationRunInput, optFns ...func(*Options)) (*StartDataQualityRulesetEvaluationRunOutput, error) {
	if params == nil {
		params = &StartDataQualityRulesetEvaluationRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDataQualityRulesetEvaluationRun", params, optFns, c.addOperationStartDataQualityRulesetEvaluationRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDataQualityRulesetEvaluationRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDataQualityRulesetEvaluationRunInput struct {

	// The data source (Glue table) associated with this run.
	//
	// This member is required.
	DataSource *types.DataSource

	// An IAM role supplied to encrypt the results of the run.
	//
	// This member is required.
	Role *string

	// A list of ruleset names.
	//
	// This member is required.
	RulesetNames []string

	// Additional run options you can specify for an evaluation run.
	AdditionalRunOptions *types.DataQualityEvaluationRunAdditionalRunOptions

	// Used for idempotency and is recommended to be set to a random ID (such as a
	// UUID) to avoid creating or starting multiple instances of the same resource.
	ClientToken *string

	// The number of G.1X workers to be used in the run. The default is 5.
	NumberOfWorkers *int32

	// The timeout for a run in minutes. This is the maximum time that a run can
	// consume resources before it is terminated and enters TIMEOUT status. The
	// default is 2,880 minutes (48 hours).
	Timeout *int32

	noSmithyDocumentSerde
}

type StartDataQualityRulesetEvaluationRunOutput struct {

	// The unique run identifier associated with this run.
	RunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDataQualityRulesetEvaluationRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartDataQualityRulesetEvaluationRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartDataQualityRulesetEvaluationRun{}, middleware.After)
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
	if err = addOpStartDataQualityRulesetEvaluationRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDataQualityRulesetEvaluationRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartDataQualityRulesetEvaluationRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartDataQualityRulesetEvaluationRun",
	}
}
