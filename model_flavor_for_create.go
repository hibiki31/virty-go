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

// checks if the FlavorForCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlavorForCreate{}

// FlavorForCreate struct for FlavorForCreate
type FlavorForCreate struct {
	Name string `json:"name"`
	Os string `json:"os"`
	ManualUrl string `json:"manualUrl"`
	Icon string `json:"icon"`
	CloudInitReady bool `json:"cloudInitReady"`
	Description string `json:"description"`
}

type _FlavorForCreate FlavorForCreate

// NewFlavorForCreate instantiates a new FlavorForCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlavorForCreate(name string, os string, manualUrl string, icon string, cloudInitReady bool, description string) *FlavorForCreate {
	this := FlavorForCreate{}
	this.Name = name
	this.Os = os
	this.ManualUrl = manualUrl
	this.Icon = icon
	this.CloudInitReady = cloudInitReady
	this.Description = description
	return &this
}

// NewFlavorForCreateWithDefaults instantiates a new FlavorForCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlavorForCreateWithDefaults() *FlavorForCreate {
	this := FlavorForCreate{}
	return &this
}

// GetName returns the Name field value
func (o *FlavorForCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FlavorForCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FlavorForCreate) SetName(v string) {
	o.Name = v
}

// GetOs returns the Os field value
func (o *FlavorForCreate) GetOs() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Os
}

// GetOsOk returns a tuple with the Os field value
// and a boolean to check if the value has been set.
func (o *FlavorForCreate) GetOsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Os, true
}

// SetOs sets field value
func (o *FlavorForCreate) SetOs(v string) {
	o.Os = v
}

// GetManualUrl returns the ManualUrl field value
func (o *FlavorForCreate) GetManualUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManualUrl
}

// GetManualUrlOk returns a tuple with the ManualUrl field value
// and a boolean to check if the value has been set.
func (o *FlavorForCreate) GetManualUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManualUrl, true
}

// SetManualUrl sets field value
func (o *FlavorForCreate) SetManualUrl(v string) {
	o.ManualUrl = v
}

// GetIcon returns the Icon field value
func (o *FlavorForCreate) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *FlavorForCreate) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *FlavorForCreate) SetIcon(v string) {
	o.Icon = v
}

// GetCloudInitReady returns the CloudInitReady field value
func (o *FlavorForCreate) GetCloudInitReady() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CloudInitReady
}

// GetCloudInitReadyOk returns a tuple with the CloudInitReady field value
// and a boolean to check if the value has been set.
func (o *FlavorForCreate) GetCloudInitReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudInitReady, true
}

// SetCloudInitReady sets field value
func (o *FlavorForCreate) SetCloudInitReady(v bool) {
	o.CloudInitReady = v
}

// GetDescription returns the Description field value
func (o *FlavorForCreate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FlavorForCreate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *FlavorForCreate) SetDescription(v string) {
	o.Description = v
}

func (o FlavorForCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlavorForCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["os"] = o.Os
	toSerialize["manualUrl"] = o.ManualUrl
	toSerialize["icon"] = o.Icon
	toSerialize["cloudInitReady"] = o.CloudInitReady
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

func (o *FlavorForCreate) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"os",
		"manualUrl",
		"icon",
		"cloudInitReady",
		"description",
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

	varFlavorForCreate := _FlavorForCreate{}

	err = json.Unmarshal(bytes, &varFlavorForCreate)

	if err != nil {
		return err
	}

	*o = FlavorForCreate(varFlavorForCreate)

	return err
}

type NullableFlavorForCreate struct {
	value *FlavorForCreate
	isSet bool
}

func (v NullableFlavorForCreate) Get() *FlavorForCreate {
	return v.value
}

func (v *NullableFlavorForCreate) Set(val *FlavorForCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableFlavorForCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableFlavorForCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlavorForCreate(val *FlavorForCreate) *NullableFlavorForCreate {
	return &NullableFlavorForCreate{value: val, isSet: true}
}

func (v NullableFlavorForCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlavorForCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


