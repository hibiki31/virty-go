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

// checks if the ImageDownloadForCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageDownloadForCreate{}

// ImageDownloadForCreate struct for ImageDownloadForCreate
type ImageDownloadForCreate struct {
	StorageUuid string `json:"storageUuid"`
	ImageUrl string `json:"imageUrl"`
}

type _ImageDownloadForCreate ImageDownloadForCreate

// NewImageDownloadForCreate instantiates a new ImageDownloadForCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageDownloadForCreate(storageUuid string, imageUrl string) *ImageDownloadForCreate {
	this := ImageDownloadForCreate{}
	this.StorageUuid = storageUuid
	this.ImageUrl = imageUrl
	return &this
}

// NewImageDownloadForCreateWithDefaults instantiates a new ImageDownloadForCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageDownloadForCreateWithDefaults() *ImageDownloadForCreate {
	this := ImageDownloadForCreate{}
	return &this
}

// GetStorageUuid returns the StorageUuid field value
func (o *ImageDownloadForCreate) GetStorageUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageUuid
}

// GetStorageUuidOk returns a tuple with the StorageUuid field value
// and a boolean to check if the value has been set.
func (o *ImageDownloadForCreate) GetStorageUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageUuid, true
}

// SetStorageUuid sets field value
func (o *ImageDownloadForCreate) SetStorageUuid(v string) {
	o.StorageUuid = v
}

// GetImageUrl returns the ImageUrl field value
func (o *ImageDownloadForCreate) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *ImageDownloadForCreate) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *ImageDownloadForCreate) SetImageUrl(v string) {
	o.ImageUrl = v
}

func (o ImageDownloadForCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageDownloadForCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["storageUuid"] = o.StorageUuid
	toSerialize["imageUrl"] = o.ImageUrl
	return toSerialize, nil
}

func (o *ImageDownloadForCreate) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"storageUuid",
		"imageUrl",
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

	varImageDownloadForCreate := _ImageDownloadForCreate{}

	err = json.Unmarshal(bytes, &varImageDownloadForCreate)

	if err != nil {
		return err
	}

	*o = ImageDownloadForCreate(varImageDownloadForCreate)

	return err
}

type NullableImageDownloadForCreate struct {
	value *ImageDownloadForCreate
	isSet bool
}

func (v NullableImageDownloadForCreate) Get() *ImageDownloadForCreate {
	return v.value
}

func (v *NullableImageDownloadForCreate) Set(val *ImageDownloadForCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableImageDownloadForCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableImageDownloadForCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageDownloadForCreate(val *ImageDownloadForCreate) *NullableImageDownloadForCreate {
	return &NullableImageDownloadForCreate{value: val, isSet: true}
}

func (v NullableImageDownloadForCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageDownloadForCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

