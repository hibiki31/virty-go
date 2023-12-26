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

// checks if the NetworkPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkPage{}

// NetworkPage struct for NetworkPage
type NetworkPage struct {
	Count int32 `json:"count"`
	Data []Network `json:"data"`
}

type _NetworkPage NetworkPage

// NewNetworkPage instantiates a new NetworkPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPage(count int32, data []Network) *NetworkPage {
	this := NetworkPage{}
	this.Count = count
	this.Data = data
	return &this
}

// NewNetworkPageWithDefaults instantiates a new NetworkPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPageWithDefaults() *NetworkPage {
	this := NetworkPage{}
	return &this
}

// GetCount returns the Count field value
func (o *NetworkPage) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *NetworkPage) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *NetworkPage) SetCount(v int32) {
	o.Count = v
}

// GetData returns the Data field value
func (o *NetworkPage) GetData() []Network {
	if o == nil {
		var ret []Network
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NetworkPage) GetDataOk() ([]Network, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *NetworkPage) SetData(v []Network) {
	o.Data = v
}

func (o NetworkPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *NetworkPage) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"data",
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

	varNetworkPage := _NetworkPage{}

	err = json.Unmarshal(bytes, &varNetworkPage)

	if err != nil {
		return err
	}

	*o = NetworkPage(varNetworkPage)

	return err
}

type NullableNetworkPage struct {
	value *NetworkPage
	isSet bool
}

func (v NullableNetworkPage) Get() *NetworkPage {
	return v.value
}

func (v *NullableNetworkPage) Set(val *NetworkPage) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPage) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPage(val *NetworkPage) *NullableNetworkPage {
	return &NullableNetworkPage{value: val, isSet: true}
}

func (v NullableNetworkPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


