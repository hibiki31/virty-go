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

// checks if the NodeRoleForUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeRoleForUpdate{}

// NodeRoleForUpdate struct for NodeRoleForUpdate
type NodeRoleForUpdate struct {
	NodeName string `json:"nodeName"`
	RoleName string `json:"roleName"`
	ExtraJson map[string]interface{} `json:"extraJson,omitempty"`
}

type _NodeRoleForUpdate NodeRoleForUpdate

// NewNodeRoleForUpdate instantiates a new NodeRoleForUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeRoleForUpdate(nodeName string, roleName string) *NodeRoleForUpdate {
	this := NodeRoleForUpdate{}
	this.NodeName = nodeName
	this.RoleName = roleName
	return &this
}

// NewNodeRoleForUpdateWithDefaults instantiates a new NodeRoleForUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeRoleForUpdateWithDefaults() *NodeRoleForUpdate {
	this := NodeRoleForUpdate{}
	return &this
}

// GetNodeName returns the NodeName field value
func (o *NodeRoleForUpdate) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *NodeRoleForUpdate) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *NodeRoleForUpdate) SetNodeName(v string) {
	o.NodeName = v
}

// GetRoleName returns the RoleName field value
func (o *NodeRoleForUpdate) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *NodeRoleForUpdate) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *NodeRoleForUpdate) SetRoleName(v string) {
	o.RoleName = v
}

// GetExtraJson returns the ExtraJson field value if set, zero value otherwise.
func (o *NodeRoleForUpdate) GetExtraJson() map[string]interface{} {
	if o == nil || IsNil(o.ExtraJson) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExtraJson
}

// GetExtraJsonOk returns a tuple with the ExtraJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeRoleForUpdate) GetExtraJsonOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExtraJson) {
		return map[string]interface{}{}, false
	}
	return o.ExtraJson, true
}

// HasExtraJson returns a boolean if a field has been set.
func (o *NodeRoleForUpdate) HasExtraJson() bool {
	if o != nil && !IsNil(o.ExtraJson) {
		return true
	}

	return false
}

// SetExtraJson gets a reference to the given map[string]interface{} and assigns it to the ExtraJson field.
func (o *NodeRoleForUpdate) SetExtraJson(v map[string]interface{}) {
	o.ExtraJson = v
}

func (o NodeRoleForUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeRoleForUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodeName"] = o.NodeName
	toSerialize["roleName"] = o.RoleName
	if !IsNil(o.ExtraJson) {
		toSerialize["extraJson"] = o.ExtraJson
	}
	return toSerialize, nil
}

func (o *NodeRoleForUpdate) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nodeName",
		"roleName",
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

	varNodeRoleForUpdate := _NodeRoleForUpdate{}

	err = json.Unmarshal(bytes, &varNodeRoleForUpdate)

	if err != nil {
		return err
	}

	*o = NodeRoleForUpdate(varNodeRoleForUpdate)

	return err
}

type NullableNodeRoleForUpdate struct {
	value *NodeRoleForUpdate
	isSet bool
}

func (v NullableNodeRoleForUpdate) Get() *NodeRoleForUpdate {
	return v.value
}

func (v *NullableNodeRoleForUpdate) Set(val *NodeRoleForUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeRoleForUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeRoleForUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeRoleForUpdate(val *NodeRoleForUpdate) *NullableNodeRoleForUpdate {
	return &NullableNodeRoleForUpdate{value: val, isSet: true}
}

func (v NullableNodeRoleForUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeRoleForUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


