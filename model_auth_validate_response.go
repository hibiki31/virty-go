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

// checks if the AuthValidateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthValidateResponse{}

// AuthValidateResponse struct for AuthValidateResponse
type AuthValidateResponse struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	Username string `json:"username"`
}

type _AuthValidateResponse AuthValidateResponse

// NewAuthValidateResponse instantiates a new AuthValidateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthValidateResponse(accessToken string, tokenType string, username string) *AuthValidateResponse {
	this := AuthValidateResponse{}
	this.AccessToken = accessToken
	this.TokenType = tokenType
	this.Username = username
	return &this
}

// NewAuthValidateResponseWithDefaults instantiates a new AuthValidateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthValidateResponseWithDefaults() *AuthValidateResponse {
	this := AuthValidateResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *AuthValidateResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *AuthValidateResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *AuthValidateResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetTokenType returns the TokenType field value
func (o *AuthValidateResponse) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *AuthValidateResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *AuthValidateResponse) SetTokenType(v string) {
	o.TokenType = v
}

// GetUsername returns the Username field value
func (o *AuthValidateResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AuthValidateResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AuthValidateResponse) SetUsername(v string) {
	o.Username = v
}

func (o AuthValidateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthValidateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["token_type"] = o.TokenType
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *AuthValidateResponse) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_token",
		"token_type",
		"username",
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

	varAuthValidateResponse := _AuthValidateResponse{}

	err = json.Unmarshal(bytes, &varAuthValidateResponse)

	if err != nil {
		return err
	}

	*o = AuthValidateResponse(varAuthValidateResponse)

	return err
}

type NullableAuthValidateResponse struct {
	value *AuthValidateResponse
	isSet bool
}

func (v NullableAuthValidateResponse) Get() *AuthValidateResponse {
	return v.value
}

func (v *NullableAuthValidateResponse) Set(val *AuthValidateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthValidateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthValidateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthValidateResponse(val *AuthValidateResponse) *NullableAuthValidateResponse {
	return &NullableAuthValidateResponse{value: val, isSet: true}
}

func (v NullableAuthValidateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthValidateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


