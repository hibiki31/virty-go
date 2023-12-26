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

// checks if the SSHKeyPair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SSHKeyPair{}

// SSHKeyPair struct for SSHKeyPair
type SSHKeyPair struct {
	PrivateKey string `json:"privateKey"`
	PublicKey string `json:"publicKey"`
}

type _SSHKeyPair SSHKeyPair

// NewSSHKeyPair instantiates a new SSHKeyPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSHKeyPair(privateKey string, publicKey string) *SSHKeyPair {
	this := SSHKeyPair{}
	this.PrivateKey = privateKey
	this.PublicKey = publicKey
	return &this
}

// NewSSHKeyPairWithDefaults instantiates a new SSHKeyPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHKeyPairWithDefaults() *SSHKeyPair {
	this := SSHKeyPair{}
	return &this
}

// GetPrivateKey returns the PrivateKey field value
func (o *SSHKeyPair) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *SSHKeyPair) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *SSHKeyPair) SetPrivateKey(v string) {
	o.PrivateKey = v
}

// GetPublicKey returns the PublicKey field value
func (o *SSHKeyPair) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *SSHKeyPair) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *SSHKeyPair) SetPublicKey(v string) {
	o.PublicKey = v
}

func (o SSHKeyPair) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SSHKeyPair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["privateKey"] = o.PrivateKey
	toSerialize["publicKey"] = o.PublicKey
	return toSerialize, nil
}

func (o *SSHKeyPair) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"privateKey",
		"publicKey",
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

	varSSHKeyPair := _SSHKeyPair{}

	err = json.Unmarshal(bytes, &varSSHKeyPair)

	if err != nil {
		return err
	}

	*o = SSHKeyPair(varSSHKeyPair)

	return err
}

type NullableSSHKeyPair struct {
	value *SSHKeyPair
	isSet bool
}

func (v NullableSSHKeyPair) Get() *SSHKeyPair {
	return v.value
}

func (v *NullableSSHKeyPair) Set(val *SSHKeyPair) {
	v.value = val
	v.isSet = true
}

func (v NullableSSHKeyPair) IsSet() bool {
	return v.isSet
}

func (v *NullableSSHKeyPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSHKeyPair(val *SSHKeyPair) *NullableSSHKeyPair {
	return &NullableSSHKeyPair{value: val, isSet: true}
}

func (v NullableSSHKeyPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSHKeyPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


