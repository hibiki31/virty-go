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

// checks if the NetworkForCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkForCreate{}

// NetworkForCreate struct for NetworkForCreate
type NetworkForCreate struct {
	Name string `json:"name"`
	NodeName string `json:"nodeName"`
	// brdige or ovs
	Type string `json:"type"`
	BridgeDevice *string `json:"bridgeDevice,omitempty"`
}

type _NetworkForCreate NetworkForCreate

// NewNetworkForCreate instantiates a new NetworkForCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkForCreate(name string, nodeName string, type_ string) *NetworkForCreate {
	this := NetworkForCreate{}
	this.Name = name
	this.NodeName = nodeName
	this.Type = type_
	return &this
}

// NewNetworkForCreateWithDefaults instantiates a new NetworkForCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkForCreateWithDefaults() *NetworkForCreate {
	this := NetworkForCreate{}
	return &this
}

// GetName returns the Name field value
func (o *NetworkForCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkForCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkForCreate) SetName(v string) {
	o.Name = v
}

// GetNodeName returns the NodeName field value
func (o *NetworkForCreate) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *NetworkForCreate) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *NetworkForCreate) SetNodeName(v string) {
	o.NodeName = v
}

// GetType returns the Type field value
func (o *NetworkForCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NetworkForCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NetworkForCreate) SetType(v string) {
	o.Type = v
}

// GetBridgeDevice returns the BridgeDevice field value if set, zero value otherwise.
func (o *NetworkForCreate) GetBridgeDevice() string {
	if o == nil || IsNil(o.BridgeDevice) {
		var ret string
		return ret
	}
	return *o.BridgeDevice
}

// GetBridgeDeviceOk returns a tuple with the BridgeDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkForCreate) GetBridgeDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.BridgeDevice) {
		return nil, false
	}
	return o.BridgeDevice, true
}

// HasBridgeDevice returns a boolean if a field has been set.
func (o *NetworkForCreate) HasBridgeDevice() bool {
	if o != nil && !IsNil(o.BridgeDevice) {
		return true
	}

	return false
}

// SetBridgeDevice gets a reference to the given string and assigns it to the BridgeDevice field.
func (o *NetworkForCreate) SetBridgeDevice(v string) {
	o.BridgeDevice = &v
}

func (o NetworkForCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkForCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["nodeName"] = o.NodeName
	toSerialize["type"] = o.Type
	if !IsNil(o.BridgeDevice) {
		toSerialize["bridgeDevice"] = o.BridgeDevice
	}
	return toSerialize, nil
}

func (o *NetworkForCreate) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"nodeName",
		"type",
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

	varNetworkForCreate := _NetworkForCreate{}

	err = json.Unmarshal(bytes, &varNetworkForCreate)

	if err != nil {
		return err
	}

	*o = NetworkForCreate(varNetworkForCreate)

	return err
}

type NullableNetworkForCreate struct {
	value *NetworkForCreate
	isSet bool
}

func (v NullableNetworkForCreate) Get() *NetworkForCreate {
	return v.value
}

func (v *NullableNetworkForCreate) Set(val *NetworkForCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkForCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkForCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkForCreate(val *NetworkForCreate) *NullableNetworkForCreate {
	return &NullableNetworkForCreate{value: val, isSet: true}
}

func (v NullableNetworkForCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkForCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

