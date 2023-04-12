// Code generated by smithy-go-codegen DO NOT EDIT.

package keyspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/keyspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The CreateTable operation adds a new table to the specified keyspace. Within a
// keyspace, table names must be unique. CreateTable is an asynchronous operation.
// When the request is received, the status of the table is set to CREATING . You
// can monitor the creation status of the new table by using the GetTable
// operation, which returns the current status of the table. You can start using a
// table when the status is ACTIVE . For more information, see Creating tables (https://docs.aws.amazon.com/keyspaces/latest/devguide/working-with-tables.html#tables-create)
// in the Amazon Keyspaces Developer Guide.
func (c *Client) CreateTable(ctx context.Context, params *CreateTableInput, optFns ...func(*Options)) (*CreateTableOutput, error) {
	if params == nil {
		params = &CreateTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTable", params, optFns, c.addOperationCreateTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTableInput struct {

	// The name of the keyspace that the table is going to be created in.
	//
	// This member is required.
	KeyspaceName *string

	// The schemaDefinition consists of the following parameters. For each column to
	// be created:
	//   - name - The name of the column.
	//   - type - An Amazon Keyspaces data type. For more information, see Data types (https://docs.aws.amazon.com/keyspaces/latest/devguide/cql.elements.html#cql.data-types)
	//   in the Amazon Keyspaces Developer Guide.
	// The primary key of the table consists of the following columns:
	//   - partitionKeys - The partition key can be a single column, or it can be a
	//   compound value composed of two or more columns. The partition key portion of the
	//   primary key is required and determines how Amazon Keyspaces stores your data.
	//   - name - The name of each partition key column.
	//   - clusteringKeys - The optional clustering column portion of your primary key
	//   determines how the data is clustered and sorted within each partition.
	//   - name - The name of the clustering column.
	//   - orderBy - Sets the ascendant ( ASC ) or descendant ( DESC ) order modifier.
	//   To define a column as static use staticColumns - Static columns store values
	//   that are shared by all rows in the same partition:
	//   - name - The name of the column.
	//   - type - An Amazon Keyspaces data type.
	//
	// This member is required.
	SchemaDefinition *types.SchemaDefinition

	// The name of the table.
	//
	// This member is required.
	TableName *string

	// Specifies the read/write throughput capacity mode for the table. The options
	// are:
	//   - throughputMode:PAY_PER_REQUEST and
	//   - throughputMode:PROVISIONED - Provisioned capacity mode requires
	//   readCapacityUnits and writeCapacityUnits as input.
	// The default is throughput_mode:PAY_PER_REQUEST . For more information, see
	// Read/write capacity modes (https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html)
	// in the Amazon Keyspaces Developer Guide.
	CapacitySpecification *types.CapacitySpecification

	// Enables client-side timestamps for the table. By default, the setting is
	// disabled. You can enable client-side timestamps with the following option:
	//   - status: "enabled"
	// Once client-side timestamps are enabled for a table, this setting cannot be
	// disabled.
	ClientSideTimestamps *types.ClientSideTimestamps

	// This parameter allows to enter a description of the table.
	Comment *types.Comment

	// The default Time to Live setting in seconds for the table. For more
	// information, see Setting the default TTL value for a table (https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_default_ttl)
	// in the Amazon Keyspaces Developer Guide.
	DefaultTimeToLive *int32

	// Specifies how the encryption key for encryption at rest is managed for the
	// table. You can choose one of the following KMS key (KMS key):
	//   - type:AWS_OWNED_KMS_KEY - This key is owned by Amazon Keyspaces.
	//   - type:CUSTOMER_MANAGED_KMS_KEY - This key is stored in your account and is
	//   created, owned, and managed by you. This option requires the
	//   kms_key_identifier of the KMS key in Amazon Resource Name (ARN) format as
	//   input.
	// The default is type:AWS_OWNED_KMS_KEY . For more information, see Encryption at
	// rest (https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html)
	// in the Amazon Keyspaces Developer Guide.
	EncryptionSpecification *types.EncryptionSpecification

	// Specifies if pointInTimeRecovery is enabled or disabled for the table. The
	// options are:
	//   - status=ENABLED
	//   - status=DISABLED
	// If it's not specified, the default is status=DISABLED . For more information,
	// see Point-in-time recovery (https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery.html)
	// in the Amazon Keyspaces Developer Guide.
	PointInTimeRecovery *types.PointInTimeRecovery

	// A list of key-value pair tags to be attached to the resource. For more
	// information, see Adding tags and labels to Amazon Keyspaces resources (https://docs.aws.amazon.com/keyspaces/latest/devguide/tagging-keyspaces.html)
	// in the Amazon Keyspaces Developer Guide.
	Tags []types.Tag

	// Enables Time to Live custom settings for the table. The options are:
	//   - status:enabled
	//   - status:disabled
	// The default is status:disabled . After ttl is enabled, you can't disable it for
	// the table. For more information, see Expiring data by using Amazon Keyspaces
	// Time to Live (TTL) (https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL.html)
	// in the Amazon Keyspaces Developer Guide.
	Ttl *types.TimeToLive

	noSmithyDocumentSerde
}

type CreateTableOutput struct {

	// The unique identifier of the table in the format of an Amazon Resource Name
	// (ARN).
	//
	// This member is required.
	ResourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateTable{}, middleware.After)
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
	if err = addOpCreateTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTable(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cassandra",
		OperationName: "CreateTable",
	}
}
