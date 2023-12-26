/*
VirtyAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package virty-go

import (
	"encoding/json"
	"fmt"
)

// checks if the NetworkPortgroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkPortgroup{}

// NetworkPortgroup struct for NetworkPortgroup
type NetworkPortgroup struct {
	Name string `json:"name"`
	VlanId *string `json:"vlanId,omitempty"`
	IsDefault bool `json:"isDefault"`
}

type _NetworkPortgroup NetworkPortgroup

// NewNetworkPortgroup instantiates a new NetworkPortgroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPortgroup(name string, isDefault bool) *NetworkPortgroup {
	this := NetworkPortgroup{}
	this.Name = name
	this.IsDefault = isDefault
	return &this
}

// NewNetworkPortgroupWithDefaults instantiates a new NetworkPortgroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPortgroupWithDefaults() *NetworkPortgroup {
	this := NetworkPortgroup{}
	return &this
}

// GetName returns the Name field value
func (o *NetworkPortgroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkPortgroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkPortgroup) SetName(v string) {
	o.Name = v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *NetworkPortgroup) GetVlanId() string {
	if o == nil || IsNil(o.VlanId) {
		var ret string
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkPortgroup) GetVlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *NetworkPortgroup) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given string and assigns it to the VlanId field.
func (o *NetworkPortgroup) SetVlanId(v string) {
	o.VlanId = &v
}

// GetIsDefault returns the IsDefault field value
func (o *NetworkPortgroup) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *NetworkPortgroup) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *NetworkPortgroup) SetIsDefault(v bool) {
	o.IsDefault = v
}

func (o NetworkPortgroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkPortgroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	toSerialize["isDefault"] = o.IsDefault
	return toSerialize, nil
}

func (o *NetworkPortgroup) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"isDefault",
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

	varNetworkPortgroup := _NetworkPortgroup{}

	err = json.Unmarshal(bytes, &varNetworkPortgroup)

	if err != nil {
		return err
	}

	*o = NetworkPortgroup(varNetworkPortgroup)

	return err
}

type NullableNetworkPortgroup struct {
	value *NetworkPortgroup
	isSet bool
}

func (v NullableNetworkPortgroup) Get() *NetworkPortgroup {
	return v.value
}

func (v *NullableNetworkPortgroup) Set(val *NetworkPortgroup) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPortgroup) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPortgroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPortgroup(val *NetworkPortgroup) *NullableNetworkPortgroup {
	return &NullableNetworkPortgroup{value: val, isSet: true}
}

func (v NullableNetworkPortgroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPortgroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

