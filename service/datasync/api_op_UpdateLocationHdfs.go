// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates some parameters of a previously created location for a Hadoop
// Distributed File System cluster.
func (c *Client) UpdateLocationHdfs(ctx context.Context, params *UpdateLocationHdfsInput, optFns ...func(*Options)) (*UpdateLocationHdfsOutput, error) {
	if params == nil {
		params = &UpdateLocationHdfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLocationHdfs", params, optFns, c.addOperationUpdateLocationHdfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLocationHdfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLocationHdfsInput struct {

	// The Amazon Resource Name (ARN) of the source HDFS cluster location.
	//
	// This member is required.
	LocationArn *string

	// The ARNs of the agents that are used to connect to the HDFS cluster.
	AgentArns []string

	// The type of authentication used to determine the identity of the user.
	AuthenticationType types.HdfsAuthenticationType

	// The size of the data blocks to write into the HDFS cluster.
	BlockSize *int32

	// The Kerberos key table (keytab) that contains mappings between the defined
	// Kerberos principal and the encrypted keys. You can load the keytab from a file
	// by providing the file's address. If you use the CLI, it performs base64 encoding
	// for you. Otherwise, provide the base64-encoded text.
	KerberosKeytab []byte

	// The krb5.conf file that contains the Kerberos configuration information. You
	// can load the krb5.conf file by providing the file's address. If you're using
	// the CLI, it performs the base64 encoding for you. Otherwise, provide the
	// base64-encoded text.
	KerberosKrb5Conf []byte

	// The Kerberos principal with access to the files and folders on the HDFS cluster.
	KerberosPrincipal *string

	// The URI of the HDFS cluster's Key Management Server (KMS).
	KmsKeyProviderUri *string

	// The NameNode that manages the HDFS namespace. The NameNode performs operations
	// such as opening, closing, and renaming files and directories. The NameNode
	// contains the information to map blocks of data to the DataNodes. You can use
	// only one NameNode.
	NameNodes []types.HdfsNameNode

	// The Quality of Protection (QOP) configuration specifies the Remote Procedure
	// Call (RPC) and data transfer privacy settings configured on the Hadoop
	// Distributed File System (HDFS) cluster.
	QopConfiguration *types.QopConfiguration

	// The number of DataNodes to replicate the data to when writing to the HDFS
	// cluster.
	ReplicationFactor *int32

	// The user name used to identify the client on the host operating system.
	SimpleUser *string

	// A subdirectory in the HDFS cluster. This subdirectory is used to read data from
	// or write data to the HDFS cluster.
	Subdirectory *string

	noSmithyDocumentSerde
}

type UpdateLocationHdfsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLocationHdfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLocationHdfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLocationHdfs{}, middleware.After)
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
	if err = addOpUpdateLocationHdfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLocationHdfs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLocationHdfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "UpdateLocationHdfs",
	}
}
