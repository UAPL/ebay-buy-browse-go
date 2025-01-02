/*
Browse API

The Browse API has the following resources:<ul><li><b>item_summary:</b><br>Allows shoppers to search for specific items by keyword, GTIN, category, charity, product, image, or item aspects and refine the results by using filters, such as aspects, compatibility, and fields values, or UI parameters.</li><li><b>item:</b><br>Allows shoppers to retrieve the details of a specific item or all items in an item group, which is an item with variations such as color and size and check if a product is compatible with the specified item, such as if a specific car is compatible with a specific part.<br><br>This resource also provides a bridge between the eBay legacy APIs, such as the <a href=\"/api-docs/user-guides/static/finding-user-guide-landing.html\" target=\"_blank\">Finding</b>, and the RESTful APIs, which use different formats for the item IDs.</li></ul>The <b>item_summary</b>, <b>search_by_image</b>, and <b>item</b> resource calls require an <a href=\"/api-docs/static/oauth-client-credentials-grant.html\" target=\"_blank\">Application access token</a>.

API version: v1.19.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package buybrowse

import (
	"encoding/json"
)

// checks if the AuthenticityVerificationProgram type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticityVerificationProgram{}

// AuthenticityVerificationProgram A type that identifies whether the item is from a verified seller.
type AuthenticityVerificationProgram struct {
	// An indication that the item is from a verified seller.
	Description *string `json:"description,omitempty"`
	// The URL to the Authenticity Verification program terms of use.
	TermsWebUrl *string `json:"termsWebUrl,omitempty"`
}

// NewAuthenticityVerificationProgram instantiates a new AuthenticityVerificationProgram object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticityVerificationProgram() *AuthenticityVerificationProgram {
	this := AuthenticityVerificationProgram{}
	return &this
}

// NewAuthenticityVerificationProgramWithDefaults instantiates a new AuthenticityVerificationProgram object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticityVerificationProgramWithDefaults() *AuthenticityVerificationProgram {
	this := AuthenticityVerificationProgram{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthenticityVerificationProgram) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticityVerificationProgram) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthenticityVerificationProgram) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthenticityVerificationProgram) SetDescription(v string) {
	o.Description = &v
}

// GetTermsWebUrl returns the TermsWebUrl field value if set, zero value otherwise.
func (o *AuthenticityVerificationProgram) GetTermsWebUrl() string {
	if o == nil || IsNil(o.TermsWebUrl) {
		var ret string
		return ret
	}
	return *o.TermsWebUrl
}

// GetTermsWebUrlOk returns a tuple with the TermsWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticityVerificationProgram) GetTermsWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TermsWebUrl) {
		return nil, false
	}
	return o.TermsWebUrl, true
}

// HasTermsWebUrl returns a boolean if a field has been set.
func (o *AuthenticityVerificationProgram) HasTermsWebUrl() bool {
	if o != nil && !IsNil(o.TermsWebUrl) {
		return true
	}

	return false
}

// SetTermsWebUrl gets a reference to the given string and assigns it to the TermsWebUrl field.
func (o *AuthenticityVerificationProgram) SetTermsWebUrl(v string) {
	o.TermsWebUrl = &v
}

func (o AuthenticityVerificationProgram) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticityVerificationProgram) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TermsWebUrl) {
		toSerialize["termsWebUrl"] = o.TermsWebUrl
	}
	return toSerialize, nil
}

type NullableAuthenticityVerificationProgram struct {
	value *AuthenticityVerificationProgram
	isSet bool
}

func (v NullableAuthenticityVerificationProgram) Get() *AuthenticityVerificationProgram {
	return v.value
}

func (v *NullableAuthenticityVerificationProgram) Set(val *AuthenticityVerificationProgram) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticityVerificationProgram) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticityVerificationProgram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticityVerificationProgram(val *AuthenticityVerificationProgram) *NullableAuthenticityVerificationProgram {
	return &NullableAuthenticityVerificationProgram{value: val, isSet: true}
}

func (v NullableAuthenticityVerificationProgram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticityVerificationProgram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


