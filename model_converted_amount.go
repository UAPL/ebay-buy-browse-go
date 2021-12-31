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

// ConvertedAmount This type defines the monetary value of an amount. It can provide the amount in both the currency used on the eBay site where an item is being offered and the conversion of that value into another currency, if applicable.
type ConvertedAmount struct {
	// The three-letter <a href=\"https://www.iso.org/iso-4217-currency-codes.html\" target=\"_blank\">ISO 4217</a> code representing the currency of the amount in the <b> convertedFromValue</b> field. This value is required or returned only if currency conversion/localization is required, and represents the pre-conversion currency. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:CurrencyCodeEnum'>eBay API documentation</a>
	ConvertedFromCurrency *string `json:"convertedFromCurrency,omitempty"`
	// The monetary amount before any conversion is performed, in the currency specified by the <b> convertedFromCurrency</b> field. This value is required or returned only if currency conversion/localization is required. The <b> value</b> field contains the converted amount of this value, in the currency specified by the <b> currency</b> field.
	ConvertedFromValue *string `json:"convertedFromValue,omitempty"`
	// The three-letter <a href=\"https://www.iso.org/iso-4217-currency-codes.html\" target=\"_blank\">ISO 4217</a> code representing the currency of the amount in the <b> value</b> field. If currency conversion/localization is required, this is the post-conversion currency of the amount in the <b> value</b> field.   <br /><br /><b> Default:</b> The currency of the authenticated user's country. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/ba:CurrencyCodeEnum'>eBay API documentation</a>
	Currency *string `json:"currency,omitempty"`
	// The monetary amount, in the currency specified by the <b> currency</b> field. If currency conversion/localization is required, this value is the converted amount, and the <b> convertedFromValue</b> field contains the amount in the original currency. 
	Value *string `json:"value,omitempty"`
}

// NewConvertedAmount instantiates a new ConvertedAmount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvertedAmount() *ConvertedAmount {
	this := ConvertedAmount{}
	return &this
}

// NewConvertedAmountWithDefaults instantiates a new ConvertedAmount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvertedAmountWithDefaults() *ConvertedAmount {
	this := ConvertedAmount{}
	return &this
}

// GetConvertedFromCurrency returns the ConvertedFromCurrency field value if set, zero value otherwise.
func (o *ConvertedAmount) GetConvertedFromCurrency() string {
	if o == nil || o.ConvertedFromCurrency == nil {
		var ret string
		return ret
	}
	return *o.ConvertedFromCurrency
}

// GetConvertedFromCurrencyOk returns a tuple with the ConvertedFromCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertedAmount) GetConvertedFromCurrencyOk() (*string, bool) {
	if o == nil || o.ConvertedFromCurrency == nil {
		return nil, false
	}
	return o.ConvertedFromCurrency, true
}

// HasConvertedFromCurrency returns a boolean if a field has been set.
func (o *ConvertedAmount) HasConvertedFromCurrency() bool {
	if o != nil && o.ConvertedFromCurrency != nil {
		return true
	}

	return false
}

// SetConvertedFromCurrency gets a reference to the given string and assigns it to the ConvertedFromCurrency field.
func (o *ConvertedAmount) SetConvertedFromCurrency(v string) {
	o.ConvertedFromCurrency = &v
}

// GetConvertedFromValue returns the ConvertedFromValue field value if set, zero value otherwise.
func (o *ConvertedAmount) GetConvertedFromValue() string {
	if o == nil || o.ConvertedFromValue == nil {
		var ret string
		return ret
	}
	return *o.ConvertedFromValue
}

// GetConvertedFromValueOk returns a tuple with the ConvertedFromValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertedAmount) GetConvertedFromValueOk() (*string, bool) {
	if o == nil || o.ConvertedFromValue == nil {
		return nil, false
	}
	return o.ConvertedFromValue, true
}

// HasConvertedFromValue returns a boolean if a field has been set.
func (o *ConvertedAmount) HasConvertedFromValue() bool {
	if o != nil && o.ConvertedFromValue != nil {
		return true
	}

	return false
}

// SetConvertedFromValue gets a reference to the given string and assigns it to the ConvertedFromValue field.
func (o *ConvertedAmount) SetConvertedFromValue(v string) {
	o.ConvertedFromValue = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ConvertedAmount) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertedAmount) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ConvertedAmount) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ConvertedAmount) SetCurrency(v string) {
	o.Currency = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConvertedAmount) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertedAmount) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConvertedAmount) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ConvertedAmount) SetValue(v string) {
	o.Value = &v
}

func (o ConvertedAmount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConvertedFromCurrency != nil {
		toSerialize["convertedFromCurrency"] = o.ConvertedFromCurrency
	}
	if o.ConvertedFromValue != nil {
		toSerialize["convertedFromValue"] = o.ConvertedFromValue
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableConvertedAmount struct {
	value *ConvertedAmount
	isSet bool
}

func (v NullableConvertedAmount) Get() *ConvertedAmount {
	return v.value
}

func (v *NullableConvertedAmount) Set(val *ConvertedAmount) {
	v.value = val
	v.isSet = true
}

func (v NullableConvertedAmount) IsSet() bool {
	return v.isSet
}

func (v *NullableConvertedAmount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvertedAmount(val *ConvertedAmount) *NullableConvertedAmount {
	return &NullableConvertedAmount{value: val, isSet: true}
}

func (v NullableConvertedAmount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvertedAmount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

