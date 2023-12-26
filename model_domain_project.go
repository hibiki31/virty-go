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

// checks if the DomainProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainProject{}

// DomainProject struct for DomainProject
type DomainProject struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type _DomainProject DomainProject

// NewDomainProject instantiates a new DomainProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainProject(id string, name string) *DomainProject {
	this := DomainProject{}
	this.Id = id
	this.Name = name
	return &this
}

// NewDomainProjectWithDefaults instantiates a new DomainProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainProjectWithDefaults() *DomainProject {
	this := DomainProject{}
	return &this
}

// GetId returns the Id field value
func (o *DomainProject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DomainProject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DomainProject) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DomainProject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DomainProject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DomainProject) SetName(v string) {
	o.Name = v
}

func (o DomainProject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *DomainProject) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varDomainProject := _DomainProject{}

	err = json.Unmarshal(bytes, &varDomainProject)

	if err != nil {
		return err
	}

	*o = DomainProject(varDomainProject)

	return err
}

type NullableDomainProject struct {
	value *DomainProject
	isSet bool
}

func (v NullableDomainProject) Get() *DomainProject {
	return v.value
}

func (v *NullableDomainProject) Set(val *DomainProject) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainProject) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainProject(val *DomainProject) *NullableDomainProject {
	return &NullableDomainProject{value: val, isSet: true}
}

func (v NullableDomainProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


