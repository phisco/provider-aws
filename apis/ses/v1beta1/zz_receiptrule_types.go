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

type AddHeaderActionObservation struct {
}

type AddHeaderActionParameters struct {

	// The name of the header to add
	// +kubebuilder:validation:Required
	HeaderName *string `json:"headerName" tf:"header_name,omitempty"`

	// The value of the header to add
	// +kubebuilder:validation:Required
	HeaderValue *string `json:"headerValue" tf:"header_value,omitempty"`

	// The position of the action in the receipt rule
	// +kubebuilder:validation:Required
	Position *float64 `json:"position" tf:"position,omitempty"`
}

type BounceActionObservation struct {
}

type BounceActionParameters struct {

	// The message to send
	// +kubebuilder:validation:Required
	Message *string `json:"message" tf:"message,omitempty"`

	// The position of the action in the receipt rule
	// +kubebuilder:validation:Required
	Position *float64 `json:"position" tf:"position,omitempty"`

	// The RFC 5321 SMTP reply code
	// +kubebuilder:validation:Required
	SMTPReplyCode *string `json:"smtpReplyCode" tf:"smtp_reply_code,omitempty"`

	// The email address of the sender
	// +kubebuilder:validation:Required
	Sender *string `json:"sender" tf:"sender,omitempty"`

	// The RFC 3463 SMTP enhanced status code
	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// The ARN of an SNS topic to notify
	// +kubebuilder:validation:Optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type LambdaActionObservation struct {
}

type LambdaActionParameters struct {

	// The ARN of the Lambda function to invoke
	// +kubebuilder:validation:Required
	FunctionArn *string `json:"functionArn" tf:"function_arn,omitempty"`

	// Event or RequestResponse
	// +kubebuilder:validation:Optional
	InvocationType *string `json:"invocationType,omitempty" tf:"invocation_type,omitempty"`

	// The position of the action in the receipt rule
	// +kubebuilder:validation:Required
	Position *float64 `json:"position" tf:"position,omitempty"`

	// The ARN of an SNS topic to notify
	// +kubebuilder:validation:Optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type ReceiptRuleObservation struct {

	// The SES receipt rule ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The SES receipt rule name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ReceiptRuleParameters struct {

	// A list of Add Header Action blocks. Documented below.
	// +kubebuilder:validation:Optional
	AddHeaderAction []AddHeaderActionParameters `json:"addHeaderAction,omitempty" tf:"add_header_action,omitempty"`

	// The name of the rule to place this rule after
	// +kubebuilder:validation:Optional
	After *string `json:"after,omitempty" tf:"after,omitempty"`

	// A list of Bounce Action blocks. Documented below.
	// +kubebuilder:validation:Optional
	BounceAction []BounceActionParameters `json:"bounceAction,omitempty" tf:"bounce_action,omitempty"`

	// If true, the rule will be enabled
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// A list of Lambda Action blocks. Documented below.
	// +kubebuilder:validation:Optional
	LambdaAction []LambdaActionParameters `json:"lambdaAction,omitempty" tf:"lambda_action,omitempty"`

	// The name of the rule
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A list of email addresses
	// +kubebuilder:validation:Optional
	Recipients []*string `json:"recipients,omitempty" tf:"recipients,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the rule set
	// +kubebuilder:validation:Required
	RuleSetName *string `json:"ruleSetName" tf:"rule_set_name,omitempty"`

	// A list of S3 Action blocks. Documented below.
	// +kubebuilder:validation:Optional
	S3Action []S3ActionParameters `json:"s3Action,omitempty" tf:"s3_action,omitempty"`

	// If true, incoming emails will be scanned for spam and viruses
	// +kubebuilder:validation:Optional
	ScanEnabled *bool `json:"scanEnabled,omitempty" tf:"scan_enabled,omitempty"`

	// A list of SNS Action blocks. Documented below.
	// +kubebuilder:validation:Optional
	SnsAction []SnsActionParameters `json:"snsAction,omitempty" tf:"sns_action,omitempty"`

	// A list of Stop Action blocks. Documented below.
	// +kubebuilder:validation:Optional
	StopAction []StopActionParameters `json:"stopAction,omitempty" tf:"stop_action,omitempty"`

	// Require or Optional
	// +kubebuilder:validation:Optional
	TLSPolicy *string `json:"tlsPolicy,omitempty" tf:"tls_policy,omitempty"`

	// A list of WorkMail Action blocks. Documented below.
	// +kubebuilder:validation:Optional
	WorkmailAction []WorkmailActionParameters `json:"workmailAction,omitempty" tf:"workmail_action,omitempty"`
}

type S3ActionObservation struct {
}

type S3ActionParameters struct {

	// The name of the S3 bucket
	// +kubebuilder:validation:Required
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`

	// The ARN of the KMS key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// The key prefix of the S3 bucket
	// +kubebuilder:validation:Optional
	ObjectKeyPrefix *string `json:"objectKeyPrefix,omitempty" tf:"object_key_prefix,omitempty"`

	// The position of the action in the receipt rule
	// +kubebuilder:validation:Required
	Position *float64 `json:"position" tf:"position,omitempty"`

	// The ARN of an SNS topic to notify
	// +kubebuilder:validation:Optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type SnsActionObservation struct {
}

type SnsActionParameters struct {

	// The encoding to use for the email within the Amazon SNS notification. Default value is UTF-8.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The position of the action in the receipt rule
	// +kubebuilder:validation:Required
	Position *float64 `json:"position" tf:"position,omitempty"`

	// The ARN of an SNS topic to notify
	// +kubebuilder:validation:Required
	TopicArn *string `json:"topicArn" tf:"topic_arn,omitempty"`
}

type StopActionObservation struct {
}

type StopActionParameters struct {

	// The position of the action in the receipt rule
	// +kubebuilder:validation:Required
	Position *float64 `json:"position" tf:"position,omitempty"`

	// The scope to apply. The only acceptable value is RuleSet.
	// +kubebuilder:validation:Required
	Scope *string `json:"scope" tf:"scope,omitempty"`

	// The ARN of an SNS topic to notify
	// +kubebuilder:validation:Optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type WorkmailActionObservation struct {
}

type WorkmailActionParameters struct {

	// The ARN of the WorkMail organization
	// +kubebuilder:validation:Required
	OrganizationArn *string `json:"organizationArn" tf:"organization_arn,omitempty"`

	// The position of the action in the receipt rule
	// +kubebuilder:validation:Required
	Position *float64 `json:"position" tf:"position,omitempty"`

	// The ARN of an SNS topic to notify
	// +kubebuilder:validation:Optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

// ReceiptRuleSpec defines the desired state of ReceiptRule
type ReceiptRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReceiptRuleParameters `json:"forProvider"`
}

// ReceiptRuleStatus defines the observed state of ReceiptRule.
type ReceiptRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReceiptRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReceiptRule is the Schema for the ReceiptRules API. Provides an SES receipt rule resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ReceiptRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReceiptRuleSpec   `json:"spec"`
	Status            ReceiptRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReceiptRuleList contains a list of ReceiptRules
type ReceiptRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReceiptRule `json:"items"`
}

// Repository type metadata.
var (
	ReceiptRule_Kind             = "ReceiptRule"
	ReceiptRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReceiptRule_Kind}.String()
	ReceiptRule_KindAPIVersion   = ReceiptRule_Kind + "." + CRDGroupVersion.String()
	ReceiptRule_GroupVersionKind = CRDGroupVersion.WithKind(ReceiptRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ReceiptRule{}, &ReceiptRuleList{})
}
