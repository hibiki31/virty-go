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

// checks if the UserScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserScope{}

// UserScope struct for UserScope
type UserScope struct {
	Name string `json:"name"`
}

type _UserScope UserScope

// NewUserScope instantiates a new UserScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserScope(name string) *UserScope {
	this := UserScope{}
	this.Name = name
	return &this
}

// NewUserScopeWithDefaults instantiates a new UserScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserScopeWithDefaults() *UserScope {
	this := UserScope{}
	return &this
}

// GetName returns the Name field value
func (o *UserScope) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserScope) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserScope) SetName(v string) {
	o.Name = v
}

func (o UserScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *UserScope) UnmarshalJSON(bytes []byte) (err error) {
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

	varUserScope := _UserScope{}

	err = json.Unmarshal(bytes, &varUserScope)

	if err != nil {
		return err
	}

	*o = UserScope(varUserScope)

	return err
}

type NullableUserScope struct {
	value *UserScope
	isSet bool
}

func (v NullableUserScope) Get() *UserScope {
	return v.value
}

func (v *NullableUserScope) Set(val *UserScope) {
	v.value = val
	v.isSet = true
}

func (v NullableUserScope) IsSet() bool {
	return v.isSet
}

func (v *NullableUserScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserScope(val *UserScope) *NullableUserScope {
	return &NullableUserScope{value: val, isSet: true}
}

func (v NullableUserScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

