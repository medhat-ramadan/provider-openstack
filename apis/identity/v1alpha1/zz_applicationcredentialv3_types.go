/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccessRulesObservation struct {

	// (Computed) The ID of the existing access rule. The access rule ID of
	// another application credential can be provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AccessRulesParameters struct {

	// The request method that the application credential is
	// permitted to use for a given API endpoint. Allowed values: POST, GET,
	// HEAD, PATCH, PUT and DELETE.
	// +kubebuilder:validation:Required
	Method *string `json:"method" tf:"method,omitempty"`

	// The API path that the application credential is permitted
	// to access. May use named wildcards such as {tag} or the unnamed wildcard
	// * to match against any string in the path up to a /, or the recursive
	// wildcard ** to include / in the matched path.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// The service type identifier for the service that the
	// application credential is granted to access. Must be a service type that is
	// listed in the service catalog and not a code name for a service. E.g.
	// identity, compute, volumev3, image, network,
	// object-store, sharev2, dns, key-manager, monitoring, etc.
	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`
}

type ApplicationCredentialV3Observation struct {

	// A collection of one or more access rules, which
	// this application credential allows to follow. The structure is described
	// below. Changing this creates a new application credential.
	// +kubebuilder:validation:Optional
	AccessRules []AccessRulesObservation `json:"accessRules,omitempty" tf:"access_rules,omitempty"`

	// (Computed) The ID of the existing access rule. The access rule ID of
	// another application credential can be provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the project the application credential was created
	// for and that authentication requests using this application credential will
	// be scoped to.
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`
}

type ApplicationCredentialV3Parameters struct {

	// A collection of one or more access rules, which
	// this application credential allows to follow. The structure is described
	// below. Changing this creates a new application credential.
	// +kubebuilder:validation:Optional
	AccessRules []AccessRulesParameters `json:"accessRules,omitempty" tf:"access_rules,omitempty"`

	// A description of the application credential.
	// Changing this creates a new application credential.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The expiration time of the application credential
	// in the RFC3339 timestamp format (e.g. 2019-03-09T12:58:49Z). If omitted,
	// an application credential will never expire. Changing this creates a new
	// application credential.
	// +kubebuilder:validation:Optional
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	// A name of the application credential. Changing this
	// creates a new application credential.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The region in which to obtain the V3 Keystone client.
	// If omitted, the region argument of the provider is used. Changing this
	// creates a new application credential.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A collection of one or more role names, which this
	// application credential has to be associated with its project. If omitted,
	// all the current user's roles within the scoped project will be inherited by
	// a new application credential. Changing this creates a new application
	// credential.
	// +kubebuilder:validation:Optional
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`

	// The secret for the application credential. If omitted,
	// it will be generated by the server. Changing this creates a new application
	// credential.
	// +kubebuilder:validation:Optional
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// A flag indicating whether the application
	// credential may be used for creation or destruction of other application
	// credentials or trusts. Changing this creates a new application credential.
	// +kubebuilder:validation:Optional
	Unrestricted *bool `json:"unrestricted,omitempty" tf:"unrestricted,omitempty"`
}

// ApplicationCredentialV3Spec defines the desired state of ApplicationCredentialV3
type ApplicationCredentialV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationCredentialV3Parameters `json:"forProvider"`
}

// ApplicationCredentialV3Status defines the observed state of ApplicationCredentialV3.
type ApplicationCredentialV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationCredentialV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationCredentialV3 is the Schema for the ApplicationCredentialV3s API. Manages a V3 Application Credential resource within OpenStack Keystone.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type ApplicationCredentialV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationCredentialV3Spec   `json:"spec"`
	Status            ApplicationCredentialV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationCredentialV3List contains a list of ApplicationCredentialV3s
type ApplicationCredentialV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationCredentialV3 `json:"items"`
}

// Repository type metadata.
var (
	ApplicationCredentialV3_Kind             = "ApplicationCredentialV3"
	ApplicationCredentialV3_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationCredentialV3_Kind}.String()
	ApplicationCredentialV3_KindAPIVersion   = ApplicationCredentialV3_Kind + "." + CRDGroupVersion.String()
	ApplicationCredentialV3_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationCredentialV3_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationCredentialV3{}, &ApplicationCredentialV3List{})
}
