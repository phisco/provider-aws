/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AuthorizationConfigObservation struct {
}

type AuthorizationConfigParameters struct {

	// The authorization type that the HTTP endpoint requires. Default values is AWS_IAM.
	// +kubebuilder:validation:Optional
	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type,omitempty"`

	// The Identity and Access Management (IAM) settings. See AWS IAM Config.
	// +kubebuilder:validation:Optional
	AwsIAMConfig []AwsIAMConfigParameters `json:"awsIamConfig,omitempty" tf:"aws_iam_config,omitempty"`
}

type AwsIAMConfigObservation struct {
}

type AwsIAMConfigParameters struct {

	// The signing Amazon Web Services Region for IAM authorization.
	// +kubebuilder:validation:Optional
	SigningRegion *string `json:"signingRegion,omitempty" tf:"signing_region,omitempty"`

	// The signing service name for IAM authorization.
	// +kubebuilder:validation:Optional
	SigningServiceName *string `json:"signingServiceName,omitempty" tf:"signing_service_name,omitempty"`
}

type DatasourceObservation struct {

	// The ARN
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DatasourceParameters struct {

	// The API ID for the GraphQL API for the DataSource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appsync/v1beta1.GraphQLAPI
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// Reference to a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDRef *v1.Reference `json:"apiIdRef,omitempty" tf:"-"`

	// Selector for a GraphQLAPI in appsync to populate apiId.
	// +kubebuilder:validation:Optional
	APIIDSelector *v1.Selector `json:"apiIdSelector,omitempty" tf:"-"`

	// A description of the DataSource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// DynamoDB settings. See below
	// +kubebuilder:validation:Optional
	DynamodbConfig []DynamodbConfigParameters `json:"dynamodbConfig,omitempty" tf:"dynamodb_config,omitempty"`

	// Amazon Elasticsearch settings. See below
	// +kubebuilder:validation:Optional
	ElasticsearchConfig []ElasticsearchConfigParameters `json:"elasticsearchConfig,omitempty" tf:"elasticsearch_config,omitempty"`

	// HTTP settings. See below
	// +kubebuilder:validation:Optional
	HTTPConfig []HTTPConfigParameters `json:"httpConfig,omitempty" tf:"http_config,omitempty"`

	// AWS Lambda settings. See below
	// +kubebuilder:validation:Optional
	LambdaConfig []LambdaConfigParameters `json:"lambdaConfig,omitempty" tf:"lambda_config,omitempty"`

	// AWS Region for RDS HTTP endpoint. Defaults to current region.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// AWS RDS settings. See Relational Database Config
	// +kubebuilder:validation:Optional
	RelationalDatabaseConfig []RelationalDatabaseConfigParameters `json:"relationalDatabaseConfig,omitempty" tf:"relational_database_config,omitempty"`

	// The IAM service role ARN for the data source.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn,omitempty"`

	// Reference to a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnRef *v1.Reference `json:"serviceRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate serviceRoleArn.
	// +kubebuilder:validation:Optional
	ServiceRoleArnSelector *v1.Selector `json:"serviceRoleArnSelector,omitempty" tf:"-"`

	// The type of the DataSource. Valid values: AWS_LAMBDA, AMAZON_DYNAMODB, AMAZON_ELASTICSEARCH, HTTP, NONE, RELATIONAL_DATABASE.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type DeltaSyncConfigObservation struct {
}

type DeltaSyncConfigParameters struct {

	// +kubebuilder:validation:Optional
	BaseTableTTL *float64 `json:"baseTableTtl,omitempty" tf:"base_table_ttl,omitempty"`

	// A user-supplied name for the DataSource.
	// +kubebuilder:validation:Required
	DeltaSyncTableName *string `json:"deltaSyncTableName" tf:"delta_sync_table_name,omitempty"`

	// +kubebuilder:validation:Optional
	DeltaSyncTableTTL *float64 `json:"deltaSyncTableTtl,omitempty" tf:"delta_sync_table_ttl,omitempty"`
}

type DynamodbConfigObservation struct {
}

type DynamodbConfigParameters struct {

	// +kubebuilder:validation:Optional
	DeltaSyncConfig []DeltaSyncConfigParameters `json:"deltaSyncConfig,omitempty" tf:"delta_sync_config,omitempty"`

	// AWS region of the DynamoDB table. Defaults to current region.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Name of the DynamoDB table.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/dynamodb/v1beta1.Table
	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`

	// Reference to a Table in dynamodb to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameRef *v1.Reference `json:"tableNameRef,omitempty" tf:"-"`

	// Selector for a Table in dynamodb to populate tableName.
	// +kubebuilder:validation:Optional
	TableNameSelector *v1.Selector `json:"tableNameSelector,omitempty" tf:"-"`

	// Set to true to use Amazon Cognito credentials with this data source.
	// +kubebuilder:validation:Optional
	UseCallerCredentials *bool `json:"useCallerCredentials,omitempty" tf:"use_caller_credentials,omitempty"`

	// +kubebuilder:validation:Optional
	Versioned *bool `json:"versioned,omitempty" tf:"versioned,omitempty"`
}

type ElasticsearchConfigObservation struct {
}

type ElasticsearchConfigParameters struct {

	// HTTP endpoint of the Elasticsearch domain.
	// +kubebuilder:validation:Required
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`

	// AWS region of Elasticsearch domain. Defaults to current region.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type HTTPConfigObservation struct {
}

type HTTPConfigParameters struct {

	// The authorization configuration in case the HTTP endpoint requires authorization. See Authorization Config.
	// +kubebuilder:validation:Optional
	AuthorizationConfig []AuthorizationConfigParameters `json:"authorizationConfig,omitempty" tf:"authorization_config,omitempty"`

	// HTTP URL.
	// +kubebuilder:validation:Required
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`
}

type HTTPEndpointConfigObservation struct {
}

type HTTPEndpointConfigParameters struct {

	// AWS secret store ARN for database credentials.
	// +kubebuilder:validation:Required
	AwsSecretStoreArn *string `json:"awsSecretStoreArn" tf:"aws_secret_store_arn,omitempty"`

	// Amazon RDS cluster identifier.
	// +kubebuilder:validation:Required
	DBClusterIdentifier *string `json:"dbClusterIdentifier" tf:"db_cluster_identifier,omitempty"`

	// Logical database name.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// AWS Region for RDS HTTP endpoint. Defaults to current region.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Logical schema name.
	// +kubebuilder:validation:Optional
	Schema *string `json:"schema,omitempty" tf:"schema,omitempty"`
}

type LambdaConfigObservation struct {
}

type LambdaConfigParameters struct {

	// The ARN for the Lambda function.
	// +kubebuilder:validation:Required
	FunctionArn *string `json:"functionArn" tf:"function_arn,omitempty"`
}

type RelationalDatabaseConfigObservation struct {
}

type RelationalDatabaseConfigParameters struct {

	// The Amazon RDS HTTP endpoint configuration. See HTTP Endpoint Config.
	// +kubebuilder:validation:Optional
	HTTPEndpointConfig []HTTPEndpointConfigParameters `json:"httpEndpointConfig,omitempty" tf:"http_endpoint_config,omitempty"`

	// Source type for the relational database. Valid values: RDS_HTTP_ENDPOINT.
	// +kubebuilder:validation:Optional
	SourceType *string `json:"sourceType,omitempty" tf:"source_type,omitempty"`
}

// DatasourceSpec defines the desired state of Datasource
type DatasourceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatasourceParameters `json:"forProvider"`
}

// DatasourceStatus defines the observed state of Datasource.
type DatasourceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatasourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Datasource is the Schema for the Datasources API. Provides an AppSync DataSource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Datasource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasourceSpec   `json:"spec"`
	Status            DatasourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasourceList contains a list of Datasources
type DatasourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Datasource `json:"items"`
}

// Repository type metadata.
var (
	Datasource_Kind             = "Datasource"
	Datasource_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Datasource_Kind}.String()
	Datasource_KindAPIVersion   = Datasource_Kind + "." + CRDGroupVersion.String()
	Datasource_GroupVersionKind = CRDGroupVersion.WithKind(Datasource_Kind)
)

func init() {
	SchemeBuilder.Register(&Datasource{}, &DatasourceList{})
}
