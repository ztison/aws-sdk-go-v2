// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Evaluates whether a pull request has met all the conditions specified in its
// associated approval rules.
func (c *Client) EvaluatePullRequestApprovalRules(ctx context.Context, params *EvaluatePullRequestApprovalRulesInput, optFns ...func(*Options)) (*EvaluatePullRequestApprovalRulesOutput, error) {
	if params == nil {
		params = &EvaluatePullRequestApprovalRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EvaluatePullRequestApprovalRules", params, optFns, c.addOperationEvaluatePullRequestApprovalRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EvaluatePullRequestApprovalRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EvaluatePullRequestApprovalRulesInput struct {

	// The system-generated ID of the pull request you want to evaluate.
	//
	// This member is required.
	PullRequestId *string

	// The system-generated ID for the pull request revision. To retrieve the most
	// recent revision ID for a pull request, use GetPullRequest .
	//
	// This member is required.
	RevisionId *string

	noSmithyDocumentSerde
}

type EvaluatePullRequestApprovalRulesOutput struct {

	// The result of the evaluation, including the names of the rules whose conditions
	// have been met (if any), the names of the rules whose conditions have not been
	// met (if any), whether the pull request is in the approved state, and whether the
	// pull request approval rule has been set aside by an override.
	//
	// This member is required.
	Evaluation *types.Evaluation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEvaluatePullRequestApprovalRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEvaluatePullRequestApprovalRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEvaluatePullRequestApprovalRules{}, middleware.After)
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
	if err = addOpEvaluatePullRequestApprovalRulesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEvaluatePullRequestApprovalRules(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEvaluatePullRequestApprovalRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "EvaluatePullRequestApprovalRules",
	}
}
