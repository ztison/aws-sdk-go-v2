// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes all versions of the bot, including the $LATEST version. To delete a
// specific version of the bot, use the DeleteBotVersion operation. The DeleteBot
// operation doesn't immediately remove the bot schema. Instead, it is marked for
// deletion and removed later. Amazon Lex stores utterances indefinitely for
// improving the ability of your bot to respond to user inputs. These utterances
// are not removed when the bot is deleted. To remove the utterances, use the
// DeleteUtterances operation. If a bot has an alias, you can't delete it. Instead,
// the DeleteBot operation returns a ResourceInUseException exception that
// includes a reference to the alias that refers to the bot. To remove the
// reference to the bot, delete the alias. If you get the same exception again,
// delete the referring alias until the DeleteBot operation is successful. This
// operation requires permissions for the lex:DeleteBot action.
func (c *Client) DeleteBot(ctx context.Context, params *DeleteBotInput, optFns ...func(*Options)) (*DeleteBotOutput, error) {
	if params == nil {
		params = &DeleteBotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBot", params, optFns, c.addOperationDeleteBotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBotInput struct {

	// The name of the bot. The name is case sensitive.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DeleteBotOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteBot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteBot{}, middleware.After)
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
	if err = addOpDeleteBotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteBot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DeleteBot",
	}
}
