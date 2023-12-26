/*
VirtyAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package virty-go

import (
	"encoding/json"
)

// checks if the NetworkProviderForCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkProviderForCreate{}

// NetworkProviderForCreate struct for NetworkProviderForCreate
type NetworkProviderForCreate struct {
	Name *string `json:"name,omitempty"`
	DnsDomain *string `json:"dnsDomain,omitempty"`
	NetworkAddress *string `json:"networkAddress,omitempty"`
	NetworkPrefix *string `json:"networkPrefix,omitempty"`
	GatewayAddress *string `json:"gatewayAddress,omitempty"`
	DhcpStart *string `json:"dhcpStart,omitempty"`
	DhcpEnd *string `json:"dhcpEnd,omitempty"`
	NetworkNode *string `json:"networkNode,omitempty"`
}

// NewNetworkProviderForCreate instantiates a new NetworkProviderForCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProviderForCreate() *NetworkProviderForCreate {
	this := NetworkProviderForCreate{}
	return &this
}

// NewNetworkProviderForCreateWithDefaults instantiates a new NetworkProviderForCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProviderForCreateWithDefaults() *NetworkProviderForCreate {
	this := NetworkProviderForCreate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkProviderForCreate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProviderForCreate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkProviderForCreate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkProviderForCreate) SetName(v string) {
	o.Name = &v
}

// GetDnsDomain returns the DnsDomain field value if set, zero value otherwise.
func (o *NetworkProviderForCreate) GetDnsDomain() string {
	if o == nil || IsNil(o.DnsDomain) {
		var ret string
		return ret
	}
	return *o.DnsDomain
}

// GetDnsDomainOk returns a tuple with the DnsDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProviderForCreate) GetDnsDomainOk() (*string, bool) {
	if o == nil || IsNil(o.DnsDomain) {
		return nil, false
	}
	return o.DnsDomain, true
}

// HasDnsDomain returns a boolean if a field has been set.
func (o *NetworkProviderForCreate) HasDnsDomain() bool {
	if o != nil && !IsNil(o.DnsDomain) {
		return true
	}

	return false
}

// SetDnsDomain gets a reference to the given string and assigns it to the DnsDomain field.
func (o *NetworkProviderForCreate) SetDnsDomain(v string) {
	o.DnsDomain = &v
}

// GetNetworkAddress returns the NetworkAddress field value if set, zero value otherwise.
func (o *NetworkProviderForCreate) GetNetworkAddress() string {
	if o == nil || IsNil(o.NetworkAddress) {
		var ret string
		return ret
	}
	return *o.NetworkAddress
}

// GetNetworkAddressOk returns a tuple with the NetworkAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProviderForCreate) GetNetworkAddressOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkAddress) {
		return nil, false
	}
	return o.NetworkAddress, true
}

// HasNetworkAddress returns a boolean if a field has been set.
func (o *NetworkProviderForCreate) HasNetworkAddress() bool {
	if o != nil && !IsNil(o.NetworkAddress) {
		return true
	}

	return false
}

// SetNetworkAddress gets a reference to the given string and assigns it to the NetworkAddress field.
func (o *NetworkProviderForCreate) SetNetworkAddress(v string) {
	o.NetworkAddress = &v
}

// GetNetworkPrefix returns the NetworkPrefix field value if set, zero value otherwise.
func (o *NetworkProviderForCreate) GetNetworkPrefix() string {
	if o == nil || IsNil(o.NetworkPrefix) {
		var ret string
		return ret
	}
	return *o.NetworkPrefix
}

// GetNetworkPrefixOk returns a tuple with the NetworkPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProviderForCreate) GetNetworkPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkPrefix) {
		return nil, false
	}
	return o.NetworkPrefix, true
}

// HasNetworkPrefix returns a boolean if a field has been set.
func (o *NetworkProviderForCreate) HasNetworkPrefix() bool {
	if o != nil && !IsNil(o.NetworkPrefix) {
		return true
	}

	return false
}

// SetNetworkPrefix gets a reference to the given string and assigns it to the NetworkPrefix field.
func (o *NetworkProviderForCreate) SetNetworkPrefix(v string) {
	o.NetworkPrefix = &v
}

// GetGatewayAddress returns the GatewayAddress field value if set, zero value otherwise.
func (o *NetworkProviderForCreate) GetGatewayAddress() string {
	if o == nil || IsNil(o.GatewayAddress) {
		var ret string
		return ret
	}
	return *o.GatewayAddress
}

// GetGatewayAddressOk returns a tuple with the GatewayAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProviderForCreate) GetGatewayAddressOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayAddress) {
		return nil, false
	}
	return o.GatewayAddress, true
}

// HasGatewayAddress returns a boolean if a field has been set.
func (o *NetworkProviderForCreate) HasGatewayAddress() bool {
	if o != nil && !IsNil(o.GatewayAddress) {
		return true
	}

	return false
}

// SetGatewayAddress gets a reference to the given string and assigns it to the GatewayAddress field.
func (o *NetworkProviderForCreate) SetGatewayAddress(v string) {
	o.GatewayAddress = &v
}

// GetDhcpStart returns the DhcpStart field value if set, zero value otherwise.
func (o *NetworkProviderForCreate) GetDhcpStart() string {
	if o == nil || IsNil(o.DhcpStart) {
		var ret string
		return ret
	}
	return *o.DhcpStart
}

// GetDhcpStartOk returns a tuple with the DhcpStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProviderForCreate) GetDhcpStartOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpStart) {
		return nil, false
	}
	return o.DhcpStart, true
}

// HasDhcpStart returns a boolean if a field has been set.
func (o *NetworkProviderForCreate) HasDhcpStart() bool {
	if o != nil && !IsNil(o.DhcpStart) {
		return true
	}

	return false
}

// SetDhcpStart gets a reference to the given string and assigns it to the DhcpStart field.
func (o *NetworkProviderForCreate) SetDhcpStart(v string) {
	o.DhcpStart = &v
}

// GetDhcpEnd returns the DhcpEnd field value if set, zero value otherwise.
func (o *NetworkProviderForCreate) GetDhcpEnd() string {
	if o == nil || IsNil(o.DhcpEnd) {
		var ret string
		return ret
	}
	return *o.DhcpEnd
}

// GetDhcpEndOk returns a tuple with the DhcpEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProviderForCreate) GetDhcpEndOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpEnd) {
		return nil, false
	}
	return o.DhcpEnd, true
}

// HasDhcpEnd returns a boolean if a field has been set.
func (o *NetworkProviderForCreate) HasDhcpEnd() bool {
	if o != nil && !IsNil(o.DhcpEnd) {
		return true
	}

	return false
}

// SetDhcpEnd gets a reference to the given string and assigns it to the DhcpEnd field.
func (o *NetworkProviderForCreate) SetDhcpEnd(v string) {
	o.DhcpEnd = &v
}

// GetNetworkNode returns the NetworkNode field value if set, zero value otherwise.
func (o *NetworkProviderForCreate) GetNetworkNode() string {
	if o == nil || IsNil(o.NetworkNode) {
		var ret string
		return ret
	}
	return *o.NetworkNode
}

// GetNetworkNodeOk returns a tuple with the NetworkNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProviderForCreate) GetNetworkNodeOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkNode) {
		return nil, false
	}
	return o.NetworkNode, true
}

// HasNetworkNode returns a boolean if a field has been set.
func (o *NetworkProviderForCreate) HasNetworkNode() bool {
	if o != nil && !IsNil(o.NetworkNode) {
		return true
	}

	return false
}

// SetNetworkNode gets a reference to the given string and assigns it to the NetworkNode field.
func (o *NetworkProviderForCreate) SetNetworkNode(v string) {
	o.NetworkNode = &v
}

func (o NetworkProviderForCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkProviderForCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DnsDomain) {
		toSerialize["dnsDomain"] = o.DnsDomain
	}
	if !IsNil(o.NetworkAddress) {
		toSerialize["networkAddress"] = o.NetworkAddress
	}
	if !IsNil(o.NetworkPrefix) {
		toSerialize["networkPrefix"] = o.NetworkPrefix
	}
	if !IsNil(o.GatewayAddress) {
		toSerialize["gatewayAddress"] = o.GatewayAddress
	}
	if !IsNil(o.DhcpStart) {
		toSerialize["dhcpStart"] = o.DhcpStart
	}
	if !IsNil(o.DhcpEnd) {
		toSerialize["dhcpEnd"] = o.DhcpEnd
	}
	if !IsNil(o.NetworkNode) {
		toSerialize["networkNode"] = o.NetworkNode
	}
	return toSerialize, nil
}

type NullableNetworkProviderForCreate struct {
	value *NetworkProviderForCreate
	isSet bool
}

func (v NullableNetworkProviderForCreate) Get() *NetworkProviderForCreate {
	return v.value
}

func (v *NullableNetworkProviderForCreate) Set(val *NetworkProviderForCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkProviderForCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkProviderForCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkProviderForCreate(val *NetworkProviderForCreate) *NullableNetworkProviderForCreate {
	return &NullableNetworkProviderForCreate{value: val, isSet: true}
}

func (v NullableNetworkProviderForCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkProviderForCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


