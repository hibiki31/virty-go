/*
VirtyAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the NodeInterfaceIpv6Info type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeInterfaceIpv6Info{}

// NodeInterfaceIpv6Info struct for NodeInterfaceIpv6Info
type NodeInterfaceIpv6Info struct {
	Address string `json:"address"`
	Prefixlen int32 `json:"prefixlen"`
}

type _NodeInterfaceIpv6Info NodeInterfaceIpv6Info

// NewNodeInterfaceIpv6Info instantiates a new NodeInterfaceIpv6Info object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeInterfaceIpv6Info(address string, prefixlen int32) *NodeInterfaceIpv6Info {
	this := NodeInterfaceIpv6Info{}
	this.Address = address
	this.Prefixlen = prefixlen
	return &this
}

// NewNodeInterfaceIpv6InfoWithDefaults instantiates a new NodeInterfaceIpv6Info object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeInterfaceIpv6InfoWithDefaults() *NodeInterfaceIpv6Info {
	this := NodeInterfaceIpv6Info{}
	return &this
}

// GetAddress returns the Address field value
func (o *NodeInterfaceIpv6Info) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *NodeInterfaceIpv6Info) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *NodeInterfaceIpv6Info) SetAddress(v string) {
	o.Address = v
}

// GetPrefixlen returns the Prefixlen field value
func (o *NodeInterfaceIpv6Info) GetPrefixlen() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Prefixlen
}

// GetPrefixlenOk returns a tuple with the Prefixlen field value
// and a boolean to check if the value has been set.
func (o *NodeInterfaceIpv6Info) GetPrefixlenOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefixlen, true
}

// SetPrefixlen sets field value
func (o *NodeInterfaceIpv6Info) SetPrefixlen(v int32) {
	o.Prefixlen = v
}

func (o NodeInterfaceIpv6Info) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeInterfaceIpv6Info) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["prefixlen"] = o.Prefixlen
	return toSerialize, nil
}

func (o *NodeInterfaceIpv6Info) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"prefixlen",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNodeInterfaceIpv6Info := _NodeInterfaceIpv6Info{}

	err = json.Unmarshal(bytes, &varNodeInterfaceIpv6Info)

	if err != nil {
		return err
	}

	*o = NodeInterfaceIpv6Info(varNodeInterfaceIpv6Info)

	return err
}

type NullableNodeInterfaceIpv6Info struct {
	value *NodeInterfaceIpv6Info
	isSet bool
}

func (v NullableNodeInterfaceIpv6Info) Get() *NodeInterfaceIpv6Info {
	return v.value
}

func (v *NullableNodeInterfaceIpv6Info) Set(val *NodeInterfaceIpv6Info) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeInterfaceIpv6Info) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeInterfaceIpv6Info) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeInterfaceIpv6Info(val *NodeInterfaceIpv6Info) *NullableNodeInterfaceIpv6Info {
	return &NullableNodeInterfaceIpv6Info{value: val, isSet: true}
}

func (v NullableNodeInterfaceIpv6Info) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeInterfaceIpv6Info) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


