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

type ExternalFixedIPObservation struct {

	// The IP address to set on the router.
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Subnet in which the fixed IP belongs to.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type ExternalFixedIPParameters struct {

	// The IP address to set on the router.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Subnet in which the fixed IP belongs to.
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type RouterV2Observation struct {

	// Administrative up/down status for the router
	// (must be "true" or "false" if provided). Changing this updates the
	// admin_state_up of an existing router.
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// The collection of tags assigned on the router, which have been
	// explicitly and implicitly added.
	AllTags []*string `json:"allTags,omitempty" tf:"all_tags,omitempty"`

	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new router.
	AvailabilityZoneHints []*string `json:"availabilityZoneHints,omitempty" tf:"availability_zone_hints,omitempty"`

	// Human-readable description for the router.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates whether or not to create a
	// distributed router. The default policy setting in Neutron restricts
	// usage of this property to administrative users only.
	Distributed *bool `json:"distributed,omitempty" tf:"distributed,omitempty"`

	// Enable Source NAT for the router. Valid values are
	// "true" or "false". An external_network_id has to be set in order to
	// set this property. Changing this updates the enable_snat of the router.
	// Setting this value requires an ext-gw-mode extension to be enabled
	// in OpenStack Neutron.
	EnableSnat *bool `json:"enableSnat,omitempty" tf:"enable_snat,omitempty"`

	// An external fixed IP for the router. This
	// can be repeated. The structure is described below. An external_network_id
	// has to be set in order to set this property. Changing this updates the
	// external fixed IPs of the router.
	ExternalFixedIP []ExternalFixedIPObservation `json:"externalFixedIp,omitempty" tf:"external_fixed_ip,omitempty"`

	// (Deprecated - use external_network_id instead) The
	// network UUID of an external gateway for the router. A router with an
	// external gateway is required if any compute instances or load balancers
	// will be using floating IPs. Changing this updates the external gateway
	// of an existing router.
	ExternalGateway *string `json:"externalGateway,omitempty" tf:"external_gateway,omitempty"`

	// The network UUID of an external gateway
	// for the router. A router with an external gateway is required if any
	// compute instances or load balancers will be using floating IPs. Changing
	// this updates the external gateway of the router.
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	// A list of external subnet IDs to try over
	// each to obtain a fixed IP for the router. If a subnet ID in a list has
	// exhausted floating IP pool, the next subnet ID will be tried. This argument is
	// used only during the router creation and allows to set only one external fixed
	// IP. Conflicts with an external_fixed_ip argument.
	ExternalSubnetIds []*string `json:"externalSubnetIds,omitempty" tf:"external_subnet_ids,omitempty"`

	// ID of the router.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A unique name for the router. Changing this
	// updates the name of an existing router.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// router.
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A set of string tags for the router.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The owner of the floating IP. Required if admin wants
	// to create a router for another tenant. Changing this creates a new router.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional driver-specific options.
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`

	// Map of additional vendor-specific options.
	// Supported options are described below.
	VendorOptions []VendorOptionsObservation `json:"vendorOptions,omitempty" tf:"vendor_options,omitempty"`
}

type RouterV2Parameters struct {

	// Administrative up/down status for the router
	// (must be "true" or "false" if provided). Changing this updates the
	// admin_state_up of an existing router.
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// An availability zone is used to make
	// network resources highly available. Used for resources with high availability
	// so that they are scheduled on different availability zones. Changing this
	// creates a new router.
	// +kubebuilder:validation:Optional
	AvailabilityZoneHints []*string `json:"availabilityZoneHints,omitempty" tf:"availability_zone_hints,omitempty"`

	// Human-readable description for the router.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Indicates whether or not to create a
	// distributed router. The default policy setting in Neutron restricts
	// usage of this property to administrative users only.
	// +kubebuilder:validation:Optional
	Distributed *bool `json:"distributed,omitempty" tf:"distributed,omitempty"`

	// Enable Source NAT for the router. Valid values are
	// "true" or "false". An external_network_id has to be set in order to
	// set this property. Changing this updates the enable_snat of the router.
	// Setting this value requires an ext-gw-mode extension to be enabled
	// in OpenStack Neutron.
	// +kubebuilder:validation:Optional
	EnableSnat *bool `json:"enableSnat,omitempty" tf:"enable_snat,omitempty"`

	// An external fixed IP for the router. This
	// can be repeated. The structure is described below. An external_network_id
	// has to be set in order to set this property. Changing this updates the
	// external fixed IPs of the router.
	// +kubebuilder:validation:Optional
	ExternalFixedIP []ExternalFixedIPParameters `json:"externalFixedIp,omitempty" tf:"external_fixed_ip,omitempty"`

	// (Deprecated - use external_network_id instead) The
	// network UUID of an external gateway for the router. A router with an
	// external gateway is required if any compute instances or load balancers
	// will be using floating IPs. Changing this updates the external gateway
	// of an existing router.
	// +kubebuilder:validation:Optional
	ExternalGateway *string `json:"externalGateway,omitempty" tf:"external_gateway,omitempty"`

	// The network UUID of an external gateway
	// for the router. A router with an external gateway is required if any
	// compute instances or load balancers will be using floating IPs. Changing
	// this updates the external gateway of the router.
	// +kubebuilder:validation:Optional
	ExternalNetworkID *string `json:"externalNetworkId,omitempty" tf:"external_network_id,omitempty"`

	// A list of external subnet IDs to try over
	// each to obtain a fixed IP for the router. If a subnet ID in a list has
	// exhausted floating IP pool, the next subnet ID will be tried. This argument is
	// used only during the router creation and allows to set only one external fixed
	// IP. Conflicts with an external_fixed_ip argument.
	// +kubebuilder:validation:Optional
	ExternalSubnetIds []*string `json:"externalSubnetIds,omitempty" tf:"external_subnet_ids,omitempty"`

	// A unique name for the router. Changing this
	// updates the name of an existing router.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The region in which to obtain the V2 networking client.
	// A networking client is needed to create a router. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// router.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// A set of string tags for the router.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The owner of the floating IP. Required if admin wants
	// to create a router for another tenant. Changing this creates a new router.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Map of additional driver-specific options.
	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`

	// Map of additional vendor-specific options.
	// Supported options are described below.
	// +kubebuilder:validation:Optional
	VendorOptions []VendorOptionsParameters `json:"vendorOptions,omitempty" tf:"vendor_options,omitempty"`
}

type VendorOptionsObservation struct {

	// Boolean to control whether
	// the Router gateway is assigned during creation or updated after creation.
	SetRouterGatewayAfterCreate *bool `json:"setRouterGatewayAfterCreate,omitempty" tf:"set_router_gateway_after_create,omitempty"`
}

type VendorOptionsParameters struct {

	// Boolean to control whether
	// the Router gateway is assigned during creation or updated after creation.
	// +kubebuilder:validation:Optional
	SetRouterGatewayAfterCreate *bool `json:"setRouterGatewayAfterCreate,omitempty" tf:"set_router_gateway_after_create,omitempty"`
}

// RouterV2Spec defines the desired state of RouterV2
type RouterV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouterV2Parameters `json:"forProvider"`
}

// RouterV2Status defines the observed state of RouterV2.
type RouterV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouterV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouterV2 is the Schema for the RouterV2s API. Manages a V2 router resource within OpenStack.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,openstack}
type RouterV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouterV2Spec   `json:"spec"`
	Status            RouterV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterV2List contains a list of RouterV2s
type RouterV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouterV2 `json:"items"`
}

// Repository type metadata.
var (
	RouterV2_Kind             = "RouterV2"
	RouterV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouterV2_Kind}.String()
	RouterV2_KindAPIVersion   = RouterV2_Kind + "." + CRDGroupVersion.String()
	RouterV2_GroupVersionKind = CRDGroupVersion.WithKind(RouterV2_Kind)
)

func init() {
	SchemeBuilder.Register(&RouterV2{}, &RouterV2List{})
}
