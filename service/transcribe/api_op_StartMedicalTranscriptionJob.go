// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Transcribes the audio from a medical dictation or conversation and applies any
// additional Request Parameters you choose to include in your request. In addition
// to many standard transcription features, Amazon Transcribe Medical provides you
// with a robust medical vocabulary and, optionally, content identification, which
// adds flags to personal health information (PHI). To learn more about these
// features, refer to How Amazon Transcribe Medical works (https://docs.aws.amazon.com/transcribe/latest/dg/how-it-works-med.html)
// . To make a StartMedicalTranscriptionJob request, you must first upload your
// media file into an Amazon S3 bucket; you can then specify the S3 location of the
// file using the Media parameter. You must include the following parameters in
// your StartMedicalTranscriptionJob request:
//   - region : The Amazon Web Services Region where you are making your request.
//     For a list of Amazon Web Services Regions supported with Amazon Transcribe,
//     refer to Amazon Transcribe endpoints and quotas (https://docs.aws.amazon.com/general/latest/gr/transcribe.html)
//     .
//   - MedicalTranscriptionJobName : A custom name you create for your
//     transcription job that is unique within your Amazon Web Services account.
//   - Media ( MediaFileUri ): The Amazon S3 location of your media file.
//   - LanguageCode : This must be en-US .
//   - OutputBucketName : The Amazon S3 bucket where you want your transcript
//     stored. If you want your output stored in a sub-folder of this bucket, you must
//     also include OutputKey .
//   - Specialty : This must be PRIMARYCARE .
//   - Type : Choose whether your audio is a conversation or a dictation.
func (c *Client) StartMedicalTranscriptionJob(ctx context.Context, params *StartMedicalTranscriptionJobInput, optFns ...func(*Options)) (*StartMedicalTranscriptionJobOutput, error) {
	if params == nil {
		params = &StartMedicalTranscriptionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMedicalTranscriptionJob", params, optFns, c.addOperationStartMedicalTranscriptionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMedicalTranscriptionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMedicalTranscriptionJobInput struct {

	// The language code that represents the language spoken in the input media file.
	// US English ( en-US ) is the only valid value for medical transcription jobs. Any
	// other value you enter for language code results in a BadRequestException error.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// Describes the Amazon S3 location of the media file you want to use in your
	// request. For information on supported media formats, refer to the MediaFormat (https://docs.aws.amazon.com/APIReference/API_StartTranscriptionJob.html#transcribe-StartTranscriptionJob-request-MediaFormat)
	// parameter or the Media formats (https://docs.aws.amazon.com/transcribe/latest/dg/how-input.html#how-input-audio)
	// section in the Amazon S3 Developer Guide.
	//
	// This member is required.
	Media *types.Media

	// A unique name, chosen by you, for your medical transcription job. The name that
	// you specify is also used as the default name of your transcription output file.
	// If you want to specify a different name for your transcription output, use the
	// OutputKey parameter. This name is case sensitive, cannot contain spaces, and
	// must be unique within an Amazon Web Services account. If you try to create a new
	// job with the same name as an existing job, you get a ConflictException error.
	//
	// This member is required.
	MedicalTranscriptionJobName *string

	// The name of the Amazon S3 bucket where you want your medical transcription
	// output stored. Do not include the S3:// prefix of the specified bucket. If you
	// want your output to go to a sub-folder of this bucket, specify it using the
	// OutputKey parameter; OutputBucketName only accepts the name of a bucket. For
	// example, if you want your output stored in S3://DOC-EXAMPLE-BUCKET , set
	// OutputBucketName to DOC-EXAMPLE-BUCKET . However, if you want your output stored
	// in S3://DOC-EXAMPLE-BUCKET/test-files/ , set OutputBucketName to
	// DOC-EXAMPLE-BUCKET and OutputKey to test-files/ . Note that Amazon Transcribe
	// must have permission to use the specified location. You can change Amazon S3
	// permissions using the Amazon Web Services Management Console (https://console.aws.amazon.com/s3)
	// . See also Permissions Required for IAM User Roles (https://docs.aws.amazon.com/transcribe/latest/dg/security_iam_id-based-policy-examples.html#auth-role-iam-user)
	// .
	//
	// This member is required.
	OutputBucketName *string

	// Specify the predominant medical specialty represented in your media. For batch
	// transcriptions, PRIMARYCARE is the only valid value. If you require additional
	// specialties, refer to .
	//
	// This member is required.
	Specialty types.Specialty

	// Specify whether your input media contains only one person ( DICTATION ) or
	// contains a conversation between two people ( CONVERSATION ). For example,
	// DICTATION could be used for a medical professional wanting to transcribe voice
	// memos; CONVERSATION could be used for transcribing the doctor-patient dialogue
	// during the patient's office visit.
	//
	// This member is required.
	Type types.Type

	// Labels all personal health information (PHI) identified in your transcript. For
	// more information, see Identifying personal health information (PHI) in a
	// transcription (https://docs.aws.amazon.com/transcribe/latest/dg/phi-id.html) .
	ContentIdentificationType types.MedicalContentIdentificationType

	// A map of plain text, non-secret key:value pairs, known as encryption context
	// pairs, that provide an added layer of security for your data. For more
	// information, see KMS encryption context (https://docs.aws.amazon.com/transcribe/latest/dg/key-management.html#kms-context)
	// and Asymmetric keys in KMS (https://docs.aws.amazon.com/transcribe/latest/dg/symmetric-asymmetric.html)
	// .
	KMSEncryptionContext map[string]string

	// Specify the format of your input media file.
	MediaFormat types.MediaFormat

	// The sample rate, in hertz, of the audio track in your input media file. If you
	// don't specify the media sample rate, Amazon Transcribe Medical determines it for
	// you. If you specify the sample rate, it must match the rate detected by Amazon
	// Transcribe Medical; if there's a mismatch between the value that you specify and
	// the value detected, your job fails. Therefore, in most cases, it's advised to
	// omit MediaSampleRateHertz and let Amazon Transcribe Medical determine the
	// sample rate.
	MediaSampleRateHertz *int32

	// The KMS key you want to use to encrypt your medical transcription output. If
	// using a key located in the current Amazon Web Services account, you can specify
	// your KMS key in one of four ways:
	//   - Use the KMS key ID itself. For example, 1234abcd-12ab-34cd-56ef-1234567890ab
	//   .
	//   - Use an alias for the KMS key ID. For example, alias/ExampleAlias .
	//   - Use the Amazon Resource Name (ARN) for the KMS key ID. For example,
	//   arn:aws:kms:region:account-ID:key/1234abcd-12ab-34cd-56ef-1234567890ab .
	//   - Use the ARN for the KMS key alias. For example,
	//   arn:aws:kms:region:account-ID:alias/ExampleAlias .
	// If using a key located in a different Amazon Web Services account than the
	// current Amazon Web Services account, you can specify your KMS key in one of two
	// ways:
	//   - Use the ARN for the KMS key ID. For example,
	//   arn:aws:kms:region:account-ID:key/1234abcd-12ab-34cd-56ef-1234567890ab .
	//   - Use the ARN for the KMS key alias. For example,
	//   arn:aws:kms:region:account-ID:alias/ExampleAlias .
	// If you don't specify an encryption key, your output is encrypted with the
	// default Amazon S3 key (SSE-S3). If you specify a KMS key to encrypt your output,
	// you must also specify an output location using the OutputLocation parameter.
	// Note that the role making the request must have permission to use the specified
	// KMS key.
	OutputEncryptionKMSKeyId *string

	// Use in combination with OutputBucketName to specify the output location of your
	// transcript and, optionally, a unique name for your output file. The default name
	// for your transcription output is the same as the name you specified for your
	// medical transcription job ( MedicalTranscriptionJobName ). Here are some
	// examples of how you can use OutputKey :
	//   - If you specify 'DOC-EXAMPLE-BUCKET' as the OutputBucketName and
	//   'my-transcript.json' as the OutputKey , your transcription output path is
	//   s3://DOC-EXAMPLE-BUCKET/my-transcript.json .
	//   - If you specify 'my-first-transcription' as the MedicalTranscriptionJobName ,
	//   'DOC-EXAMPLE-BUCKET' as the OutputBucketName , and 'my-transcript' as the
	//   OutputKey , your transcription output path is
	//   s3://DOC-EXAMPLE-BUCKET/my-transcript/my-first-transcription.json .
	//   - If you specify 'DOC-EXAMPLE-BUCKET' as the OutputBucketName and
	//   'test-files/my-transcript.json' as the OutputKey , your transcription output
	//   path is s3://DOC-EXAMPLE-BUCKET/test-files/my-transcript.json .
	//   - If you specify 'my-first-transcription' as the MedicalTranscriptionJobName ,
	//   'DOC-EXAMPLE-BUCKET' as the OutputBucketName , and 'test-files/my-transcript'
	//   as the OutputKey , your transcription output path is
	//   s3://DOC-EXAMPLE-BUCKET/test-files/my-transcript/my-first-transcription.json .
	// If you specify the name of an Amazon S3 bucket sub-folder that doesn't exist,
	// one is created for you.
	OutputKey *string

	// Specify additional optional settings in your request, including channel
	// identification, alternative transcriptions, and speaker partitioning. You can
	// use that to apply custom vocabularies to your transcription job.
	Settings *types.MedicalTranscriptionSetting

	// Adds one or more custom tags, each in the form of a key:value pair, to a new
	// medical transcription job at the time you start this new job. To learn more
	// about using tags with Amazon Transcribe, refer to Tagging resources (https://docs.aws.amazon.com/transcribe/latest/dg/tagging.html)
	// .
	Tags []types.Tag

	noSmithyDocumentSerde
}

type StartMedicalTranscriptionJobOutput struct {

	// Provides detailed information about the current medical transcription job,
	// including job status and, if applicable, failure reason.
	MedicalTranscriptionJob *types.MedicalTranscriptionJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMedicalTranscriptionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartMedicalTranscriptionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMedicalTranscriptionJob{}, middleware.After)
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
	if err = addOpStartMedicalTranscriptionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMedicalTranscriptionJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMedicalTranscriptionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "StartMedicalTranscriptionJob",
	}
}
