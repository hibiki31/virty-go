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

// checks if the ProjectForUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectForUpdate{}

// ProjectForUpdate struct for ProjectForUpdate
type ProjectForUpdate struct {
	ProjectId string `json:"projectId"`
	UserId string `json:"userId"`
}

type _ProjectForUpdate ProjectForUpdate

// NewProjectForUpdate instantiates a new ProjectForUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectForUpdate(projectId string, userId string) *ProjectForUpdate {
	this := ProjectForUpdate{}
	this.ProjectId = projectId
	this.UserId = userId
	return &this
}

// NewProjectForUpdateWithDefaults instantiates a new ProjectForUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectForUpdateWithDefaults() *ProjectForUpdate {
	this := ProjectForUpdate{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *ProjectForUpdate) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *ProjectForUpdate) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *ProjectForUpdate) SetProjectId(v string) {
	o.ProjectId = v
}

// GetUserId returns the UserId field value
func (o *ProjectForUpdate) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ProjectForUpdate) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ProjectForUpdate) SetUserId(v string) {
	o.UserId = v
}

func (o ProjectForUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectForUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *ProjectForUpdate) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"projectId",
		"userId",
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

	varProjectForUpdate := _ProjectForUpdate{}

	err = json.Unmarshal(bytes, &varProjectForUpdate)

	if err != nil {
		return err
	}

	*o = ProjectForUpdate(varProjectForUpdate)

	return err
}

type NullableProjectForUpdate struct {
	value *ProjectForUpdate
	isSet bool
}

func (v NullableProjectForUpdate) Get() *ProjectForUpdate {
	return v.value
}

func (v *NullableProjectForUpdate) Set(val *ProjectForUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectForUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectForUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectForUpdate(val *ProjectForUpdate) *NullableProjectForUpdate {
	return &NullableProjectForUpdate{value: val, isSet: true}
}

func (v NullableProjectForUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectForUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


