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

// checks if the DomainProjectForUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainProjectForUpdate{}

// DomainProjectForUpdate struct for DomainProjectForUpdate
type DomainProjectForUpdate struct {
	Uuid string `json:"uuid"`
	ProjectId string `json:"projectId"`
}

type _DomainProjectForUpdate DomainProjectForUpdate

// NewDomainProjectForUpdate instantiates a new DomainProjectForUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainProjectForUpdate(uuid string, projectId string) *DomainProjectForUpdate {
	this := DomainProjectForUpdate{}
	this.Uuid = uuid
	this.ProjectId = projectId
	return &this
}

// NewDomainProjectForUpdateWithDefaults instantiates a new DomainProjectForUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainProjectForUpdateWithDefaults() *DomainProjectForUpdate {
	this := DomainProjectForUpdate{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *DomainProjectForUpdate) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *DomainProjectForUpdate) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *DomainProjectForUpdate) SetUuid(v string) {
	o.Uuid = v
}

// GetProjectId returns the ProjectId field value
func (o *DomainProjectForUpdate) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DomainProjectForUpdate) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DomainProjectForUpdate) SetProjectId(v string) {
	o.ProjectId = v
}

func (o DomainProjectForUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainProjectForUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

func (o *DomainProjectForUpdate) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
		"projectId",
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

	varDomainProjectForUpdate := _DomainProjectForUpdate{}

	err = json.Unmarshal(bytes, &varDomainProjectForUpdate)

	if err != nil {
		return err
	}

	*o = DomainProjectForUpdate(varDomainProjectForUpdate)

	return err
}

type NullableDomainProjectForUpdate struct {
	value *DomainProjectForUpdate
	isSet bool
}

func (v NullableDomainProjectForUpdate) Get() *DomainProjectForUpdate {
	return v.value
}

func (v *NullableDomainProjectForUpdate) Set(val *DomainProjectForUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainProjectForUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainProjectForUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainProjectForUpdate(val *DomainProjectForUpdate) *NullableDomainProjectForUpdate {
	return &NullableDomainProjectForUpdate{value: val, isSet: true}
}

func (v NullableDomainProjectForUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainProjectForUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


