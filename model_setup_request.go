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

// checks if the SetupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetupRequest{}

// SetupRequest struct for SetupRequest
type SetupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type _SetupRequest SetupRequest

// NewSetupRequest instantiates a new SetupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetupRequest(username string, password string) *SetupRequest {
	this := SetupRequest{}
	this.Username = username
	this.Password = password
	return &this
}

// NewSetupRequestWithDefaults instantiates a new SetupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetupRequestWithDefaults() *SetupRequest {
	this := SetupRequest{}
	return &this
}

// GetUsername returns the Username field value
func (o *SetupRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SetupRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *SetupRequest) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *SetupRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SetupRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SetupRequest) SetPassword(v string) {
	o.Password = v
}

func (o SetupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *SetupRequest) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
		"password",
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

	varSetupRequest := _SetupRequest{}

	err = json.Unmarshal(bytes, &varSetupRequest)

	if err != nil {
		return err
	}

	*o = SetupRequest(varSetupRequest)

	return err
}

type NullableSetupRequest struct {
	value *SetupRequest
	isSet bool
}

func (v NullableSetupRequest) Get() *SetupRequest {
	return v.value
}

func (v *NullableSetupRequest) Set(val *SetupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetupRequest(val *SetupRequest) *NullableSetupRequest {
	return &NullableSetupRequest{value: val, isSet: true}
}

func (v NullableSetupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


