// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An extracted segment of the text that is an attribute of an entity, or
// otherwise related to an entity, such as the dosage of a medication taken. It
// contains information about the attribute such as id, begin and end offset within
// the input text, and the segment of the input text.
type Attribute struct {

	// The 0-based character offset in the input text that shows where the attribute
	// begins. The offset returns the UTF-8 code point in the string.
	BeginOffset *int32

	// The category of attribute.
	Category EntityType

	// The 0-based character offset in the input text that shows where the attribute
	// ends. The offset returns the UTF-8 code point in the string.
	EndOffset *int32

	// The numeric identifier for this attribute. This is a monotonically increasing
	// id unique within this response rather than a global unique identifier.
	Id *int32

	// The level of confidence that Comprehend Medical; has that this attribute is
	// correctly related to this entity.
	RelationshipScore *float32

	// The type of relationship between the entity and attribute. Type for the
	// relationship is OVERLAP , indicating that the entity occurred at the same time
	// as the Date_Expression .
	RelationshipType RelationshipType

	// The level of confidence that Comprehend Medical; has that the segment of text
	// is correctly recognized as an attribute.
	Score *float32

	// The segment of input text extracted as this attribute.
	Text *string

	// Contextual information for this attribute.
	Traits []Trait

	// The type of attribute.
	Type EntitySubType

	noSmithyDocumentSerde
}

// The number of characters in the input text to be analyzed.
type Characters struct {

	// The number of characters present in the input text document as processed by
	// Comprehend Medical.
	OriginalTextCharacters *int32

	noSmithyDocumentSerde
}

// Provides information for filtering a list of detection jobs.
type ComprehendMedicalAsyncJobFilter struct {

	// Filters on the name of the job.
	JobName *string

	// Filters the list of jobs based on job status. Returns only jobs with the
	// specified status.
	JobStatus JobStatus

	// Filters the list of jobs based on the time that the job was submitted for
	// processing. Returns only jobs submitted after the specified time. Jobs are
	// returned in descending order, newest to oldest.
	SubmitTimeAfter *time.Time

	// Filters the list of jobs based on the time that the job was submitted for
	// processing. Returns only jobs submitted before the specified time. Jobs are
	// returned in ascending order, oldest to newest.
	SubmitTimeBefore *time.Time

	noSmithyDocumentSerde
}

// Provides information about a detection job.
type ComprehendMedicalAsyncJobProperties struct {

	// The Amazon Resource Name (ARN) that gives Comprehend Medical; read access to
	// your input data.
	DataAccessRoleArn *string

	// The time that the detection job completed.
	EndTime *time.Time

	// The date and time that job metadata is deleted from the server. Output files in
	// your S3 bucket will not be deleted. After the metadata is deleted, the job will
	// no longer appear in the results of the ListEntitiesDetectionV2Job or the
	// ListPHIDetectionJobs operation.
	ExpirationTime *time.Time

	// The input data configuration that you supplied when you created the detection
	// job.
	InputDataConfig *InputDataConfig

	// The identifier assigned to the detection job.
	JobId *string

	// The name that you assigned to the detection job.
	JobName *string

	// The current status of the detection job. If the status is FAILED , the Message
	// field shows the reason for the failure.
	JobStatus JobStatus

	// The AWS Key Management Service key, if any, used to encrypt the output files.
	KMSKey *string

	// The language code of the input documents.
	LanguageCode LanguageCode

	// The path to the file that describes the results of a batch job.
	ManifestFilePath *string

	// A description of the status of a job.
	Message *string

	// The version of the model used to analyze the documents. The version number
	// looks like X.X.X. You can use this information to track the model used for a
	// particular batch of documents.
	ModelVersion *string

	// The output data configuration that you supplied when you created the detection
	// job.
	OutputDataConfig *OutputDataConfig

	// The time that the detection job was submitted for processing.
	SubmitTime *time.Time

	noSmithyDocumentSerde
}

// Provides information about an extracted medical entity.
type Entity struct {

	// The extracted attributes that relate to this entity.
	Attributes []Attribute

	// The 0-based character offset in the input text that shows where the entity
	// begins. The offset returns the UTF-8 code point in the string.
	BeginOffset *int32

	// The category of the entity.
	Category EntityType

	// The 0-based character offset in the input text that shows where the entity
	// ends. The offset returns the UTF-8 code point in the string.
	EndOffset *int32

	// The numeric identifier for the entity. This is a monotonically increasing id
	// unique within this response rather than a global unique identifier.
	Id *int32

	// The level of confidence that Comprehend Medical; has in the accuracy of the
	// detection.
	Score *float32

	// The segment of input text extracted as this entity.
	Text *string

	// Contextual information for the entity.
	Traits []Trait

	// Describes the specific type of entity with category of entities.
	Type EntitySubType

	noSmithyDocumentSerde
}

// The detected attributes that relate to an entity. This includes an extracted
// segment of the text that is an attribute of an entity, or otherwise related to
// an entity. InferICD10CM detects the following attributes: Direction , System,
// Organ or Site , and Acuity .
type ICD10CMAttribute struct {

	// The 0-based character offset in the input text that shows where the attribute
	// begins. The offset returns the UTF-8 code point in the string.
	BeginOffset *int32

	// The category of attribute. Can be either of DX_NAME or TIME_EXPRESSION .
	Category ICD10CMEntityType

	// The 0-based character offset in the input text that shows where the attribute
	// ends. The offset returns the UTF-8 code point in the string.
	EndOffset *int32

	// The numeric identifier for this attribute. This is a monotonically increasing
	// id unique within this response rather than a global unique identifier.
	Id *int32

	// The level of confidence that Amazon Comprehend Medical has that this attribute
	// is correctly related to this entity.
	RelationshipScore *float32

	// The type of relationship between the entity and attribute. Type for the
	// relationship can be either of OVERLAP or SYSTEM_ORGAN_SITE .
	RelationshipType ICD10CMRelationshipType

	// The level of confidence that Amazon Comprehend Medical has that the segment of
	// text is correctly recognized as an attribute.
	Score *float32

	// The segment of input text which contains the detected attribute.
	Text *string

	// The contextual information for the attribute. The traits recognized by
	// InferICD10CM are DIAGNOSIS , SIGN , SYMPTOM , and NEGATION .
	Traits []ICD10CMTrait

	// The type of attribute. InferICD10CM detects entities of the type DX_NAME .
	Type ICD10CMAttributeType

	noSmithyDocumentSerde
}

// The ICD-10-CM concepts that the entity could refer to, along with a score
// indicating the likelihood of the match.
type ICD10CMConcept struct {

	// The ICD-10-CM code that identifies the concept found in the knowledge base from
	// the Centers for Disease Control.
	Code *string

	// The long description of the ICD-10-CM code in the ontology.
	Description *string

	// The level of confidence that Amazon Comprehend Medical has that the entity is
	// accurately linked to an ICD-10-CM concept.
	Score *float32

	noSmithyDocumentSerde
}

// The collection of medical entities extracted from the input text and their
// associated information. For each entity, the response provides the entity text,
// the entity category, where the entity text begins and ends, and the level of
// confidence that Amazon Comprehend Medical has in the detection and analysis.
// Attributes and traits of the entity are also returned.
type ICD10CMEntity struct {

	// The detected attributes that relate to the entity. An extracted segment of the
	// text that is an attribute of an entity, or otherwise related to an entity, such
	// as the nature of a medical condition.
	Attributes []ICD10CMAttribute

	// The 0-based character offset in the input text that shows where the entity
	// begins. The offset returns the UTF-8 code point in the string.
	BeginOffset *int32

	// The category of the entity. InferICD10CM detects entities in the
	// MEDICAL_CONDITION category.
	Category ICD10CMEntityCategory

	// The 0-based character offset in the input text that shows where the entity
	// ends. The offset returns the UTF-8 code point in the string.
	EndOffset *int32

	// The ICD-10-CM concepts that the entity could refer to, along with a score
	// indicating the likelihood of the match.
	ICD10CMConcepts []ICD10CMConcept

	// The numeric identifier for the entity. This is a monotonically increasing id
	// unique within this response rather than a global unique identifier.
	Id *int32

	// The level of confidence that Amazon Comprehend Medical has in the accuracy of
	// the detection.
	Score *float32

	// The segment of input text that is matched to the detected entity.
	Text *string

	// Provides Contextual information for the entity. The traits recognized by
	// InferICD10CM are DIAGNOSIS , SIGN , SYMPTOM , and NEGATION.
	Traits []ICD10CMTrait

	// Describes the specific type of entity with category of entities. InferICD10CM
	// detects entities of the type DX_NAME and TIME_EXPRESSION .
	Type ICD10CMEntityType

	noSmithyDocumentSerde
}

// Contextual information for the entity. The traits recognized by InferICD10CM
// are DIAGNOSIS , SIGN , SYMPTOM , and NEGATION .
type ICD10CMTrait struct {

	// Provides a name or contextual description about the trait.
	Name ICD10CMTraitName

	// The level of confidence that Comprehend Medical; has that the segment of text
	// is correctly recognized as a trait.
	Score *float32

	noSmithyDocumentSerde
}

// The input properties for an entities detection job. This includes the name of
// the S3 bucket and the path to the files to be analyzed.
type InputDataConfig struct {

	// The URI of the S3 bucket that contains the input data. The bucket must be in
	// the same region as the API endpoint that you are calling. Each file in the
	// document collection must be less than 40 KB. You can store a maximum of 30 GB in
	// the bucket.
	//
	// This member is required.
	S3Bucket *string

	// The path to the input data files in the S3 bucket.
	S3Key *string

	noSmithyDocumentSerde
}

// The output properties for a detection job.
type OutputDataConfig struct {

	// When you use the OutputDataConfig object with asynchronous operations, you
	// specify the Amazon S3 location where you want to write the output data. The URI
	// must be in the same region as the API endpoint that you are calling. The
	// location is used as the prefix for the actual location of the output.
	//
	// This member is required.
	S3Bucket *string

	// The path to the output data files in the S3 bucket. Comprehend Medical; creates
	// an output directory using the job ID so that the output from one job does not
	// overwrite the output of another.
	S3Key *string

	noSmithyDocumentSerde
}

// The extracted attributes that relate to this entity. The attributes recognized
// by InferRxNorm are DOSAGE , DURATION , FORM , FREQUENCY , RATE , ROUTE_OR_MODE .
type RxNormAttribute struct {

	// The 0-based character offset in the input text that shows where the attribute
	// begins. The offset returns the UTF-8 code point in the string.
	BeginOffset *int32

	// The 0-based character offset in the input text that shows where the attribute
	// ends. The offset returns the UTF-8 code point in the string.
	EndOffset *int32

	// The numeric identifier for this attribute. This is a monotonically increasing
	// id unique within this response rather than a global unique identifier.
	Id *int32

	// The level of confidence that Amazon Comprehend Medical has that the attribute
	// is accurately linked to an entity.
	RelationshipScore *float32

	// The level of confidence that Comprehend Medical has that the segment of text is
	// correctly recognized as an attribute.
	Score *float32

	// The segment of input text which corresponds to the detected attribute.
	Text *string

	// Contextual information for the attribute. InferRxNorm recognizes the trait
	// NEGATION for attributes, i.e. that the patient is not taking a specific dose or
	// form of a medication.
	Traits []RxNormTrait

	// The type of attribute. The types of attributes recognized by InferRxNorm are
	// BRAND_NAME and GENERIC_NAME .
	Type RxNormAttributeType

	noSmithyDocumentSerde
}

// The RxNorm concept that the entity could refer to, along with a score
// indicating the likelihood of the match.
type RxNormConcept struct {

	// RxNorm concept ID, also known as the RxCUI.
	Code *string

	// The description of the RxNorm concept.
	Description *string

	// The level of confidence that Amazon Comprehend Medical has that the entity is
	// accurately linked to the reported RxNorm concept.
	Score *float32

	noSmithyDocumentSerde
}

// The collection of medical entities extracted from the input text and their
// associated information. For each entity, the response provides the entity text,
// the entity category, where the entity text begins and ends, and the level of
// confidence that Amazon Comprehend Medical has in the detection and analysis.
// Attributes and traits of the entity are also returned.
type RxNormEntity struct {

	// The extracted attributes that relate to the entity. The attributes recognized
	// by InferRxNorm are DOSAGE , DURATION , FORM , FREQUENCY , RATE , ROUTE_OR_MODE ,
	// and STRENGTH .
	Attributes []RxNormAttribute

	// The 0-based character offset in the input text that shows where the entity
	// begins. The offset returns the UTF-8 code point in the string.
	BeginOffset *int32

	// The category of the entity. The recognized categories are GENERIC or BRAND_NAME .
	Category RxNormEntityCategory

	// The 0-based character offset in the input text that shows where the entity
	// ends. The offset returns the UTF-8 code point in the string.
	EndOffset *int32

	// The numeric identifier for the entity. This is a monotonically increasing id
	// unique within this response rather than a global unique identifier.
	Id *int32

	// The RxNorm concepts that the entity could refer to, along with a score
	// indicating the likelihood of the match.
	RxNormConcepts []RxNormConcept

	// The level of confidence that Amazon Comprehend Medical has in the accuracy of
	// the detected entity.
	Score *float32

	// The segment of input text extracted from which the entity was detected.
	Text *string

	// Contextual information for the entity.
	Traits []RxNormTrait

	// Describes the specific type of entity. For InferRxNorm, the recognized entity
	// type is MEDICATION .
	Type RxNormEntityType

	noSmithyDocumentSerde
}

// The contextual information for the entity. InferRxNorm recognizes the trait
// NEGATION , which is any indication that the patient is not taking a medication.
type RxNormTrait struct {

	// Provides a name or contextual description about the trait.
	Name RxNormTraitName

	// The level of confidence that Amazon Comprehend Medical has in the accuracy of
	// the detected trait.
	Score *float32

	noSmithyDocumentSerde
}

// The extracted attributes that relate to an entity. An extracted segment of the
// text that is an attribute of an entity, or otherwise related to an entity, such
// as the dosage of a medication taken.
type SNOMEDCTAttribute struct {

	// The 0-based character offset in the input text that shows where the attribute
	// begins. The offset returns the UTF-8 code point in the string.
	BeginOffset *int32

	// The category of the detected attribute. Possible categories include
	// MEDICAL_CONDITION, ANATOMY, and TEST_TREATMENT_PROCEDURE.
	Category SNOMEDCTEntityCategory

	// The 0-based character offset in the input text that shows where the attribute
	// ends. The offset returns the UTF-8 code point in the string.
	EndOffset *int32

	// The numeric identifier for this attribute. This is a monotonically increasing
	// id unique within this response rather than a global unique identifier.
	Id *int32

	// The level of confidence that Comprehend Medical has that this attribute is
	// correctly related to this entity.
	RelationshipScore *float32

	// The type of relationship that exists between the entity and the related
	// attribute.
	RelationshipType SNOMEDCTRelationshipType

	// The SNOMED-CT concepts specific to an attribute, along with a score indicating
	// the likelihood of the match.
	SNOMEDCTConcepts []SNOMEDCTConcept

	// The level of confidence that Comprehend Medical has that the segment of text is
	// correctly recognized as an attribute.
	Score *float32

	// The segment of input text extracted as this attribute.
	Text *string

	// Contextual information for an attribute. Examples include signs, symptoms,
	// diagnosis, and negation.
	Traits []SNOMEDCTTrait

	// The type of attribute. Possible types include DX_NAME, ACUITY, DIRECTION,
	// SYSTEM_ORGAN_SITE,TEST_NAME, TEST_VALUE, TEST_UNIT, PROCEDURE_NAME, and
	// TREATMENT_NAME.
	Type SNOMEDCTAttributeType

	noSmithyDocumentSerde
}

// The SNOMED-CT concepts that the entity could refer to, along with a score
// indicating the likelihood of the match.
type SNOMEDCTConcept struct {

	// The numeric ID for the SNOMED-CT concept.
	Code *string

	// The description of the SNOMED-CT concept.
	Description *string

	// The level of confidence Comprehend Medical has that the entity should be linked
	// to the identified SNOMED-CT concept.
	Score *float32

	noSmithyDocumentSerde
}

// The information about the revision of the SNOMED-CT ontology in the response.
// Specifically, the details include the SNOMED-CT edition, language, and version
// date.
type SNOMEDCTDetails struct {

	// The edition of SNOMED-CT used. The edition used for the InferSNOMEDCT editions
	// is the US edition.
	Edition *string

	// The language used in the SNOMED-CT ontology. All Amazon Comprehend Medical
	// operations are US English (en).
	Language *string

	// The version date of the SNOMED-CT ontology used.
	VersionDate *string

	noSmithyDocumentSerde
}

// The collection of medical entities extracted from the input text and their
// associated information. For each entity, the response provides the entity text,
// the entity category, where the entity text begins and ends, and the level of
// confidence that Comprehend Medical has in the detection and analysis. Attributes
// and traits of the entity are also returned.
type SNOMEDCTEntity struct {

	// An extracted segment of the text that is an attribute of an entity, or
	// otherwise related to an entity, such as the dosage of a medication taken.
	Attributes []SNOMEDCTAttribute

	// The 0-based character offset in the input text that shows where the entity
	// begins. The offset returns the UTF-8 code point in the string.
	BeginOffset *int32

	// The category of the detected entity. Possible categories are MEDICAL_CONDITION,
	// ANATOMY, or TEST_TREATMENT_PROCEDURE.
	Category SNOMEDCTEntityCategory

	// The 0-based character offset in the input text that shows where the entity
	// ends. The offset returns the UTF-8 code point in the string.
	EndOffset *int32

	// The numeric identifier for the entity. This is a monotonically increasing id
	// unique within this response rather than a global unique identifier.
	Id *int32

	// The SNOMED concepts that the entity could refer to, along with a score
	// indicating the likelihood of the match.
	SNOMEDCTConcepts []SNOMEDCTConcept

	// The level of confidence that Comprehend Medical has in the accuracy of the
	// detected entity.
	Score *float32

	// The segment of input text extracted as this entity.
	Text *string

	// Contextual information for the entity.
	Traits []SNOMEDCTTrait

	// Describes the specific type of entity with category of entities. Possible types
	// include DX_NAME, ACUITY, DIRECTION, SYSTEM_ORGAN_SITE, TEST_NAME, TEST_VALUE,
	// TEST_UNIT, PROCEDURE_NAME, or TREATMENT_NAME.
	Type SNOMEDCTEntityType

	noSmithyDocumentSerde
}

// Contextual information for an entity.
type SNOMEDCTTrait struct {

	// The name or contextual description of a detected trait.
	Name SNOMEDCTTraitName

	// The level of confidence that Comprehend Medical has in the accuracy of a
	// detected trait.
	Score *float32

	noSmithyDocumentSerde
}

// Provides contextual information about the extracted entity.
type Trait struct {

	// Provides a name or contextual description about the trait.
	Name AttributeName

	// The level of confidence that Comprehend Medical; has in the accuracy of this
	// trait.
	Score *float32

	noSmithyDocumentSerde
}

// An attribute that was extracted, but Comprehend Medical; was unable to relate
// to an entity.
type UnmappedAttribute struct {

	// The specific attribute that has been extracted but not mapped to an entity.
	Attribute *Attribute

	// The type of the unmapped attribute, could be one of the following values:
	// "MEDICATION", "MEDICAL_CONDITION", "ANATOMY", "TEST_AND_TREATMENT_PROCEDURE" or
	// "PROTECTED_HEALTH_INFORMATION".
	Type EntityType

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
