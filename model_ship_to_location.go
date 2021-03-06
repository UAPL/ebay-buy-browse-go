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

// ShipToLocation The type that defines the fields for the country and postal code of where an item is to be shipped.
type ShipToLocation struct {
	// The two-letter <a href=\"https://www.iso.org/iso-3166-country-codes.html\" target=\"_blank\">ISO 3166</a> standard of the country for where the item is to be shipped. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:CountryCodeEnum'>eBay API documentation</a>
	Country *string `json:"country,omitempty"`
	// The zip code (postal code) for where the item is to be shipped.
	PostalCode *string `json:"postalCode,omitempty"`
}

// NewShipToLocation instantiates a new ShipToLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipToLocation() *ShipToLocation {
	this := ShipToLocation{}
	return &this
}

// NewShipToLocationWithDefaults instantiates a new ShipToLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipToLocationWithDefaults() *ShipToLocation {
	this := ShipToLocation{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *ShipToLocation) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipToLocation) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *ShipToLocation) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *ShipToLocation) SetCountry(v string) {
	o.Country = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *ShipToLocation) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipToLocation) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *ShipToLocation) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *ShipToLocation) SetPostalCode(v string) {
	o.PostalCode = &v
}

func (o ShipToLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.PostalCode != nil {
		toSerialize["postalCode"] = o.PostalCode
	}
	return json.Marshal(toSerialize)
}

type NullableShipToLocation struct {
	value *ShipToLocation
	isSet bool
}

func (v NullableShipToLocation) Get() *ShipToLocation {
	return v.value
}

func (v *NullableShipToLocation) Set(val *ShipToLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableShipToLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableShipToLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipToLocation(val *ShipToLocation) *NullableShipToLocation {
	return &NullableShipToLocation{value: val, isSet: true}
}

func (v NullableShipToLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipToLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


