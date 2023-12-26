/*
VirtyAPI

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the TokenData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenData{}

// TokenData struct for TokenData
type TokenData struct {
	Id *string `json:"id,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
	Role []string `json:"role,omitempty"`
}

// NewTokenData instantiates a new TokenData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenData() *TokenData {
	this := TokenData{}
	return &this
}

// NewTokenDataWithDefaults instantiates a new TokenData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenDataWithDefaults() *TokenData {
	this := TokenData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TokenData) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenData) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TokenData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TokenData) SetId(v string) {
	o.Id = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *TokenData) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenData) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TokenData) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *TokenData) SetScopes(v []string) {
	o.Scopes = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *TokenData) GetRole() []string {
	if o == nil || IsNil(o.Role) {
		var ret []string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenData) GetRoleOk() ([]string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *TokenData) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given []string and assigns it to the Role field.
func (o *TokenData) SetRole(v []string) {
	o.Role = v
}

func (o TokenData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableTokenData struct {
	value *TokenData
	isSet bool
}

func (v NullableTokenData) Get() *TokenData {
	return v.value
}

func (v *NullableTokenData) Set(val *TokenData) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenData) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenData(val *TokenData) *NullableTokenData {
	return &NullableTokenData{value: val, isSet: true}
}

func (v NullableTokenData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


