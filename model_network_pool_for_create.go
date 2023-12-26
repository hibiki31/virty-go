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

// checks if the NetworkPoolForCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkPoolForCreate{}

// NetworkPoolForCreate struct for NetworkPoolForCreate
type NetworkPoolForCreate struct {
	Name string `json:"name"`
}

type _NetworkPoolForCreate NetworkPoolForCreate

// NewNetworkPoolForCreate instantiates a new NetworkPoolForCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkPoolForCreate(name string) *NetworkPoolForCreate {
	this := NetworkPoolForCreate{}
	this.Name = name
	return &this
}

// NewNetworkPoolForCreateWithDefaults instantiates a new NetworkPoolForCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkPoolForCreateWithDefaults() *NetworkPoolForCreate {
	this := NetworkPoolForCreate{}
	return &this
}

// GetName returns the Name field value
func (o *NetworkPoolForCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkPoolForCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkPoolForCreate) SetName(v string) {
	o.Name = v
}

func (o NetworkPoolForCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkPoolForCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *NetworkPoolForCreate) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varNetworkPoolForCreate := _NetworkPoolForCreate{}

	err = json.Unmarshal(bytes, &varNetworkPoolForCreate)

	if err != nil {
		return err
	}

	*o = NetworkPoolForCreate(varNetworkPoolForCreate)

	return err
}

type NullableNetworkPoolForCreate struct {
	value *NetworkPoolForCreate
	isSet bool
}

func (v NullableNetworkPoolForCreate) Get() *NetworkPoolForCreate {
	return v.value
}

func (v *NullableNetworkPoolForCreate) Set(val *NetworkPoolForCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkPoolForCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkPoolForCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkPoolForCreate(val *NetworkPoolForCreate) *NullableNetworkPoolForCreate {
	return &NullableNetworkPoolForCreate{value: val, isSet: true}
}

func (v NullableNetworkPoolForCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkPoolForCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

