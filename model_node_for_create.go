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

// checks if the NodeForCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeForCreate{}

// NodeForCreate struct for NodeForCreate
type NodeForCreate struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Domain string `json:"domain"`
	UserName string `json:"userName"`
	Port int32 `json:"port"`
	LibvirtRole bool `json:"libvirtRole"`
}

type _NodeForCreate NodeForCreate

// NewNodeForCreate instantiates a new NodeForCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeForCreate(name string, description string, domain string, userName string, port int32, libvirtRole bool) *NodeForCreate {
	this := NodeForCreate{}
	this.Name = name
	this.Description = description
	this.Domain = domain
	this.UserName = userName
	this.Port = port
	this.LibvirtRole = libvirtRole
	return &this
}

// NewNodeForCreateWithDefaults instantiates a new NodeForCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeForCreateWithDefaults() *NodeForCreate {
	this := NodeForCreate{}
	return &this
}

// GetName returns the Name field value
func (o *NodeForCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NodeForCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NodeForCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *NodeForCreate) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *NodeForCreate) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *NodeForCreate) SetDescription(v string) {
	o.Description = v
}

// GetDomain returns the Domain field value
func (o *NodeForCreate) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *NodeForCreate) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *NodeForCreate) SetDomain(v string) {
	o.Domain = v
}

// GetUserName returns the UserName field value
func (o *NodeForCreate) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *NodeForCreate) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *NodeForCreate) SetUserName(v string) {
	o.UserName = v
}

// GetPort returns the Port field value
func (o *NodeForCreate) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *NodeForCreate) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *NodeForCreate) SetPort(v int32) {
	o.Port = v
}

// GetLibvirtRole returns the LibvirtRole field value
func (o *NodeForCreate) GetLibvirtRole() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LibvirtRole
}

// GetLibvirtRoleOk returns a tuple with the LibvirtRole field value
// and a boolean to check if the value has been set.
func (o *NodeForCreate) GetLibvirtRoleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LibvirtRole, true
}

// SetLibvirtRole sets field value
func (o *NodeForCreate) SetLibvirtRole(v bool) {
	o.LibvirtRole = v
}

func (o NodeForCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeForCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["domain"] = o.Domain
	toSerialize["userName"] = o.UserName
	toSerialize["port"] = o.Port
	toSerialize["libvirtRole"] = o.LibvirtRole
	return toSerialize, nil
}

func (o *NodeForCreate) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"domain",
		"userName",
		"port",
		"libvirtRole",
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

	varNodeForCreate := _NodeForCreate{}

	err = json.Unmarshal(bytes, &varNodeForCreate)

	if err != nil {
		return err
	}

	*o = NodeForCreate(varNodeForCreate)

	return err
}

type NullableNodeForCreate struct {
	value *NodeForCreate
	isSet bool
}

func (v NullableNodeForCreate) Get() *NodeForCreate {
	return v.value
}

func (v *NullableNodeForCreate) Set(val *NodeForCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeForCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeForCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeForCreate(val *NodeForCreate) *NullableNodeForCreate {
	return &NullableNodeForCreate{value: val, isSet: true}
}

func (v NullableNodeForCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeForCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


