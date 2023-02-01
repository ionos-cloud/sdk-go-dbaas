/*
 * IONOS DBaaS PostgreSQL REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional PostgreSQL database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// DBUser Credentials for the database user to be created.
type DBUser struct {
	// The username for the initial PostgreSQL user. Some system usernames are restricted (e.g. \"postgres\", \"admin\", \"standby\").
	Username *string `json:"username"`
	Password *string `json:"password"`
}

// NewDBUser instantiates a new DBUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDBUser(username string, password string) *DBUser {
	this := DBUser{}

	this.Username = &username
	this.Password = &password

	return &this
}

// NewDBUserWithDefaults instantiates a new DBUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDBUserWithDefaults() *DBUser {
	this := DBUser{}
	return &this
}

// GetUsername returns the Username field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DBUser) GetUsername() *string {
	if o == nil {
		return nil
	}

	return o.Username

}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DBUser) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Username, true
}

// SetUsername sets field value
func (o *DBUser) SetUsername(v string) {

	o.Username = &v

}

// HasUsername returns a boolean if a field has been set.
func (o *DBUser) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// GetPassword returns the Password field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DBUser) GetPassword() *string {
	if o == nil {
		return nil
	}

	return o.Password

}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DBUser) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Password, true
}

// SetPassword sets field value
func (o *DBUser) SetPassword(v string) {

	o.Password = &v

}

// HasPassword returns a boolean if a field has been set.
func (o *DBUser) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

func (o DBUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}

	if o.Password != nil {
		toSerialize["password"] = o.Password
	}

	return json.Marshal(toSerialize)
}

type NullableDBUser struct {
	value *DBUser
	isSet bool
}

func (v NullableDBUser) Get() *DBUser {
	return v.value
}

func (v *NullableDBUser) Set(val *DBUser) {
	v.value = val
	v.isSet = true
}

func (v NullableDBUser) IsSet() bool {
	return v.isSet
}

func (v *NullableDBUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDBUser(val *DBUser) *NullableDBUser {
	return &NullableDBUser{value: val, isSet: true}
}

func (v NullableDBUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDBUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
