/*
Browse API

<p>The Browse API has the following resources:</p>   <ul> <li><b> item_summary: </b> Lets shoppers search for specific items by keyword, GTIN, category, charity, product, or item aspects and refine the results by using filters, such as aspects, compatibility, and fields values.</li>  <li><b> search_by_image: </b><a href=\"https://developer.ebay.com/api-docs/static/versioning.html#experimental\" target=\"_blank\"><img src=\"/cms/img/docs/experimental-icon.svg\" class=\"legend-icon experimental-icon\" alt=\"Experimental Release\" title=\"Experimental Release\" />&nbsp;(Experimental)</a> Lets shoppers search for specific items by image. You can refine the results by using URI parameters and filters.</li>   <li><b> item: </b> <ul><li>Lets you retrieve the details of a specific item or all the items in an item group, which is an item with variations such as color and size and check if a product is compatible with the specified item, such as if a specific car is compatible with a specific part.</li> <li>Provides a bridge between the eBay legacy APIs, such as <b> Finding</b>, and the RESTful APIs, which use different formats for the item IDs.</li>  </ul> </li>  <li> <b> shopping_cart: </b> <a href=\"https://developer.ebay.com/api-docs/static/versioning.html#experimental\" target=\"_blank\"><img src=\"/cms/img/docs/experimental-icon.svg\" class=\"legend-icon experimental-icon\" alt=\"Experimental Release\" title=\"Experimental Release\" />&nbsp;(Experimental)</a> <a href=\"https://developer.ebay.com/api-docs/static/versioning.html#limited\" target=\"_blank\"> <img src=\"/cms/img/docs/partners-api.svg\" class=\"legend-icon partners-icon\" title=\"Limited Release\"  alt=\"Limited Release\" />(Limited Release)</a> Provides the ability for eBay members to see the contents of their eBay cart, and add, remove, and change the quantity of items in their eBay cart.&nbsp;&nbsp;<b> Note: </b> This resource is not available in the eBay API Explorer.</li></ul>       <p>The <b> item_summary</b>, <b> search_by_image</b>, and <b> item</b> resource calls require an <a href=\"/api-docs/static/oauth-client-credentials-grant.html\">Application access token</a>. The <b> shopping_cart</b> resource calls require a <a href=\"/api-docs/static/oauth-authorization-code-grant.html\">User access token</a>.</p>

API version: v1.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package buybrowse

import (
	"encoding/json"
)

// SellerLegalInfo The type that defines the fields for the contact information for a seller.
type SellerLegalInfo struct {
	// The seller's business email address.
	Email *string `json:"email,omitempty"`
	// The seller' business fax number.
	Fax *string `json:"fax,omitempty"`
	// This is a free-form string created by the seller. This is information often found on business cards, such as address. This is information used by some countries.
	Imprint *string `json:"imprint,omitempty"`
	// The seller's first name.
	LegalContactFirstName *string `json:"legalContactFirstName,omitempty"`
	// The seller's last name.
	LegalContactLastName *string `json:"legalContactLastName,omitempty"`
	// The name of the seller's business.
	Name *string `json:"name,omitempty"`
	// The seller's business phone number.
	Phone *string `json:"phone,omitempty"`
	// The seller's registration number. This is information used by some countries.
	RegistrationNumber *string `json:"registrationNumber,omitempty"`
	SellerProvidedLegalAddress *LegalAddress `json:"sellerProvidedLegalAddress,omitempty"`
	// This is a free-form string created by the seller. This is the seller's terms or condition, which is in addition to the seller's return policies.
	TermsOfService *string `json:"termsOfService,omitempty"`
	// An array of the seller's VAT (value added tax) IDs and the issuing country. VAT is a tax added by some European countries.
	VatDetails *[]VatDetail `json:"vatDetails,omitempty"`
}

// NewSellerLegalInfo instantiates a new SellerLegalInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSellerLegalInfo() *SellerLegalInfo {
	this := SellerLegalInfo{}
	return &this
}

// NewSellerLegalInfoWithDefaults instantiates a new SellerLegalInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSellerLegalInfoWithDefaults() *SellerLegalInfo {
	this := SellerLegalInfo{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SellerLegalInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellerLegalInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SellerLegalInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SellerLegalInfo) SetEmail(v string) {
	o.Email = &v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *SellerLegalInfo) GetFax() string {
	if o == nil || o.Fax == nil {
		var ret string
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellerLegalInfo) GetFaxOk() (*string, bool) {
	if o == nil || o.Fax == nil {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *SellerLegalInfo) HasFax() bool {
	if o != nil && o.Fax != nil {
		return true
	}

	return false
}

// SetFax gets a reference to the given string and assigns it to the Fax field.
func (o *SellerLegalInfo) SetFax(v string) {
	o.Fax = &v
}

// GetImprint returns the Imprint field value if set, zero value otherwise.
func (o *SellerLegalInfo) GetImprint() string {
	if o == nil || o.Imprint == nil {
		var ret string
		return ret
	}
	return *o.Imprint
}

// GetImprintOk returns a tuple with the Imprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellerLegalInfo) GetImprintOk() (*string, bool) {
	if o == nil || o.Imprint == nil {
		return nil, false
	}
	return o.Imprint, true
}

// HasImprint returns a boolean if a field has been set.
func (o *SellerLegalInfo) HasImprint() bool {
	if o != nil && o.Imprint != nil {
		return true
	}

	return false
}

// SetImprint gets a reference to the given string and assigns it to the Imprint field.
func (o *SellerLegalInfo) SetImprint(v string) {
	o.Imprint = &v
}

// GetLegalContactFirstName returns the LegalContactFirstName field value if set, zero value otherwise.
func (o *SellerLegalInfo) GetLegalContactFirstName() string {
	if o == nil || o.LegalContactFirstName == nil {
		var ret string
		return ret
	}
	return *o.LegalContactFirstName
}

// GetLegalContactFirstNameOk returns a tuple with the LegalContactFirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellerLegalInfo) GetLegalContactFirstNameOk() (*string, bool) {
	if o == nil || o.LegalContactFirstName == nil {
		return nil, false
	}
	return o.LegalContactFirstName, true
}

// HasLegalContactFirstName returns a boolean if a field has been set.
func (o *SellerLegalInfo) HasLegalContactFirstName() bool {
	if o != nil && o.LegalContactFirstName != nil {
		return true
	}

	return false
}

// SetLegalContactFirstName gets a reference to the given string and assigns it to the LegalContactFirstName field.
func (o *SellerLegalInfo) SetLegalContactFirstName(v string) {
	o.LegalContactFirstName = &v
}

// GetLegalContactLastName returns the LegalContactLastName field value if set, zero value otherwise.
func (o *SellerLegalInfo) GetLegalContactLastName() string {
	if o == nil || o.LegalContactLastName == nil {
		var ret string
		return ret
	}
	return *o.LegalContactLastName
}

// GetLegalContactLastNameOk returns a tuple with the LegalContactLastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellerLegalInfo) GetLegalContactLastNameOk() (*string, bool) {
	if o == nil || o.LegalContactLastName == nil {
		return nil, false
	}
	return o.LegalContactLastName, true
}

// HasLegalContactLastName returns a boolean if a field has been set.
func (o *SellerLegalInfo) HasLegalContactLastName() bool {
	if o != nil && o.LegalContactLastName != nil {
		return true
	}

	return false
}

// SetLegalContactLastName gets a reference to the given string and assigns it to the LegalContactLastName field.
func (o *SellerLegalInfo) SetLegalContactLastName(v string) {
	o.LegalContactLastName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SellerLegalInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellerLegalInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SellerLegalInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SellerLegalInfo) SetName(v string) {
	o.Name = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *SellerLegalInfo) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellerLegalInfo) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *SellerLegalInfo) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *SellerLegalInfo) SetPhone(v string) {
	o.Phone = &v
}

// GetRegistrationNumber returns the RegistrationNumber field value if set, zero value otherwise.
func (o *SellerLegalInfo) GetRegistrationNumber() string {
	if o == nil || o.RegistrationNumber == nil {
		var ret string
		return ret
	}
	return *o.RegistrationNumber
}

// GetRegistrationNumberOk returns a tuple with the RegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellerLegalInfo) GetRegistrationNumberOk() (*string, bool) {
	if o == nil || o.RegistrationNumber == nil {
		return nil, false
	}
	return o.RegistrationNumber, true
}

// HasRegistrationNumber returns a boolean if a field has been set.
func (o *SellerLegalInfo) HasRegistrationNumber() bool {
	if o != nil && o.RegistrationNumber != nil {
		return true
	}

	return false
}

// SetRegistrationNumber gets a reference to the given string and assigns it to the RegistrationNumber field.
func (o *SellerLegalInfo) SetRegistrationNumber(v string) {
	o.RegistrationNumber = &v
}

// GetSellerProvidedLegalAddress returns the SellerProvidedLegalAddress field value if set, zero value otherwise.
func (o *SellerLegalInfo) GetSellerProvidedLegalAddress() LegalAddress {
	if o == nil || o.SellerProvidedLegalAddress == nil {
		var ret LegalAddress
		return ret
	}
	return *o.SellerProvidedLegalAddress
}

// GetSellerProvidedLegalAddressOk returns a tuple with the SellerProvidedLegalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellerLegalInfo) GetSellerProvidedLegalAddressOk() (*LegalAddress, bool) {
	if o == nil || o.SellerProvidedLegalAddress == nil {
		return nil, false
	}
	return o.SellerProvidedLegalAddress, true
}

// HasSellerProvidedLegalAddress returns a boolean if a field has been set.
func (o *SellerLegalInfo) HasSellerProvidedLegalAddress() bool {
	if o != nil && o.SellerProvidedLegalAddress != nil {
		return true
	}

	return false
}

// SetSellerProvidedLegalAddress gets a reference to the given LegalAddress and assigns it to the SellerProvidedLegalAddress field.
func (o *SellerLegalInfo) SetSellerProvidedLegalAddress(v LegalAddress) {
	o.SellerProvidedLegalAddress = &v
}

// GetTermsOfService returns the TermsOfService field value if set, zero value otherwise.
func (o *SellerLegalInfo) GetTermsOfService() string {
	if o == nil || o.TermsOfService == nil {
		var ret string
		return ret
	}
	return *o.TermsOfService
}

// GetTermsOfServiceOk returns a tuple with the TermsOfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellerLegalInfo) GetTermsOfServiceOk() (*string, bool) {
	if o == nil || o.TermsOfService == nil {
		return nil, false
	}
	return o.TermsOfService, true
}

// HasTermsOfService returns a boolean if a field has been set.
func (o *SellerLegalInfo) HasTermsOfService() bool {
	if o != nil && o.TermsOfService != nil {
		return true
	}

	return false
}

// SetTermsOfService gets a reference to the given string and assigns it to the TermsOfService field.
func (o *SellerLegalInfo) SetTermsOfService(v string) {
	o.TermsOfService = &v
}

// GetVatDetails returns the VatDetails field value if set, zero value otherwise.
func (o *SellerLegalInfo) GetVatDetails() []VatDetail {
	if o == nil || o.VatDetails == nil {
		var ret []VatDetail
		return ret
	}
	return *o.VatDetails
}

// GetVatDetailsOk returns a tuple with the VatDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellerLegalInfo) GetVatDetailsOk() (*[]VatDetail, bool) {
	if o == nil || o.VatDetails == nil {
		return nil, false
	}
	return o.VatDetails, true
}

// HasVatDetails returns a boolean if a field has been set.
func (o *SellerLegalInfo) HasVatDetails() bool {
	if o != nil && o.VatDetails != nil {
		return true
	}

	return false
}

// SetVatDetails gets a reference to the given []VatDetail and assigns it to the VatDetails field.
func (o *SellerLegalInfo) SetVatDetails(v []VatDetail) {
	o.VatDetails = &v
}

func (o SellerLegalInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Fax != nil {
		toSerialize["fax"] = o.Fax
	}
	if o.Imprint != nil {
		toSerialize["imprint"] = o.Imprint
	}
	if o.LegalContactFirstName != nil {
		toSerialize["legalContactFirstName"] = o.LegalContactFirstName
	}
	if o.LegalContactLastName != nil {
		toSerialize["legalContactLastName"] = o.LegalContactLastName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.RegistrationNumber != nil {
		toSerialize["registrationNumber"] = o.RegistrationNumber
	}
	if o.SellerProvidedLegalAddress != nil {
		toSerialize["sellerProvidedLegalAddress"] = o.SellerProvidedLegalAddress
	}
	if o.TermsOfService != nil {
		toSerialize["termsOfService"] = o.TermsOfService
	}
	if o.VatDetails != nil {
		toSerialize["vatDetails"] = o.VatDetails
	}
	return json.Marshal(toSerialize)
}

type NullableSellerLegalInfo struct {
	value *SellerLegalInfo
	isSet bool
}

func (v NullableSellerLegalInfo) Get() *SellerLegalInfo {
	return v.value
}

func (v *NullableSellerLegalInfo) Set(val *SellerLegalInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSellerLegalInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSellerLegalInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSellerLegalInfo(val *SellerLegalInfo) *NullableSellerLegalInfo {
	return &NullableSellerLegalInfo{value: val, isSet: true}
}

func (v NullableSellerLegalInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSellerLegalInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


