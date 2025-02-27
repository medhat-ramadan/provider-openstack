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

type KeypairV2Observation struct {

	// The fingerprint of the public key.
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A unique name for the keypair. Changing this creates a new
	// keypair.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A pregenerated OpenSSH-formatted public key.
	// Changing this creates a new keypair. If a public key is not specified, then
	// a public/private key pair will be automatically generated. If a pair is
	// created, then destroying this resource means you will lose access to that
	// keypair forever.
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the region argument of the provider is used.
	// Changing this creates a new keypair.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// This allows administrative users to operate key-pairs
	// of specified user ID. For this feature your need to have openstack microversion
	// 2.10 (Liberty) or later.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Map of additional options.
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

type KeypairV2Parameters struct {

	// A unique name for the keypair. Changing this creates a new
	// keypair.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A pregenerated OpenSSH-formatted public key.
	// Changing this creates a new keypair. If a public key is not specified, then
	// a public/private key pair will be automatically generated. If a pair is
	// created, then destroying this resource means you will lose access to that
	// keypair forever.
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// The region in which to obtain the V2 Compute client.
	// Keypairs are associated with accounts, but a Compute client is needed to
	// create one. If omitted, the region argument of the provider is used.
	// Changing this creates a new keypair.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// This allows administrative users to operate key-pairs
	// of specified user ID. For this feature your need to have openstack microversion
	// 2.10 (Liberty) or later.
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Map of additional options.
	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// KeypairV2Spec defines the desired state of KeypairV2
type KeypairV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KeypairV2Parameters `json:"forProvider"`
}

// KeypairV2Status defines the observed state of KeypairV2.
type KeypairV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KeypairV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KeypairV2 is the Schema for the KeypairV2s API. Manages a V2 keypair resource within OpenStack.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type KeypairV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   KeypairV2Spec   `json:"spec"`
	Status KeypairV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeypairV2List contains a list of KeypairV2s
type KeypairV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeypairV2 `json:"items"`
}

// Repository type metadata.
var (
	KeypairV2_Kind             = "KeypairV2"
	KeypairV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KeypairV2_Kind}.String()
	KeypairV2_KindAPIVersion   = KeypairV2_Kind + "." + CRDGroupVersion.String()
	KeypairV2_GroupVersionKind = CRDGroupVersion.WithKind(KeypairV2_Kind)
)

func init() {
	SchemeBuilder.Register(&KeypairV2{}, &KeypairV2List{})
}
