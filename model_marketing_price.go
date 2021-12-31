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

// MarketingPrice The type that defines the fields that describe a seller discount.
type MarketingPrice struct {
	DiscountAmount *ConvertedAmount `json:"discountAmount,omitempty"`
	// This field expresses the percentage of the seller discount based on the value in the <b>  originalPrice</b> container.
	DiscountPercentage *string `json:"discountPercentage,omitempty"`
	OriginalPrice *ConvertedAmount `json:"originalPrice,omitempty"`
	// Indicates the pricing treatment (discount) that was applied to the price of the item. <br /><br /><span class=\"tablenote\"><b>Note: </b> The pricing treatment affects the way and where the discounted price can be displayed.</span> For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:PriceTreatmentEnum'>eBay API documentation</a>
	PriceTreatment *string `json:"priceTreatment,omitempty"`
}

// NewMarketingPrice instantiates a new MarketingPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingPrice() *MarketingPrice {
	this := MarketingPrice{}
	return &this
}

// NewMarketingPriceWithDefaults instantiates a new MarketingPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingPriceWithDefaults() *MarketingPrice {
	this := MarketingPrice{}
	return &this
}

// GetDiscountAmount returns the DiscountAmount field value if set, zero value otherwise.
func (o *MarketingPrice) GetDiscountAmount() ConvertedAmount {
	if o == nil || o.DiscountAmount == nil {
		var ret ConvertedAmount
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingPrice) GetDiscountAmountOk() (*ConvertedAmount, bool) {
	if o == nil || o.DiscountAmount == nil {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *MarketingPrice) HasDiscountAmount() bool {
	if o != nil && o.DiscountAmount != nil {
		return true
	}

	return false
}

// SetDiscountAmount gets a reference to the given ConvertedAmount and assigns it to the DiscountAmount field.
func (o *MarketingPrice) SetDiscountAmount(v ConvertedAmount) {
	o.DiscountAmount = &v
}

// GetDiscountPercentage returns the DiscountPercentage field value if set, zero value otherwise.
func (o *MarketingPrice) GetDiscountPercentage() string {
	if o == nil || o.DiscountPercentage == nil {
		var ret string
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingPrice) GetDiscountPercentageOk() (*string, bool) {
	if o == nil || o.DiscountPercentage == nil {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *MarketingPrice) HasDiscountPercentage() bool {
	if o != nil && o.DiscountPercentage != nil {
		return true
	}

	return false
}

// SetDiscountPercentage gets a reference to the given string and assigns it to the DiscountPercentage field.
func (o *MarketingPrice) SetDiscountPercentage(v string) {
	o.DiscountPercentage = &v
}

// GetOriginalPrice returns the OriginalPrice field value if set, zero value otherwise.
func (o *MarketingPrice) GetOriginalPrice() ConvertedAmount {
	if o == nil || o.OriginalPrice == nil {
		var ret ConvertedAmount
		return ret
	}
	return *o.OriginalPrice
}

// GetOriginalPriceOk returns a tuple with the OriginalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingPrice) GetOriginalPriceOk() (*ConvertedAmount, bool) {
	if o == nil || o.OriginalPrice == nil {
		return nil, false
	}
	return o.OriginalPrice, true
}

// HasOriginalPrice returns a boolean if a field has been set.
func (o *MarketingPrice) HasOriginalPrice() bool {
	if o != nil && o.OriginalPrice != nil {
		return true
	}

	return false
}

// SetOriginalPrice gets a reference to the given ConvertedAmount and assigns it to the OriginalPrice field.
func (o *MarketingPrice) SetOriginalPrice(v ConvertedAmount) {
	o.OriginalPrice = &v
}

// GetPriceTreatment returns the PriceTreatment field value if set, zero value otherwise.
func (o *MarketingPrice) GetPriceTreatment() string {
	if o == nil || o.PriceTreatment == nil {
		var ret string
		return ret
	}
	return *o.PriceTreatment
}

// GetPriceTreatmentOk returns a tuple with the PriceTreatment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingPrice) GetPriceTreatmentOk() (*string, bool) {
	if o == nil || o.PriceTreatment == nil {
		return nil, false
	}
	return o.PriceTreatment, true
}

// HasPriceTreatment returns a boolean if a field has been set.
func (o *MarketingPrice) HasPriceTreatment() bool {
	if o != nil && o.PriceTreatment != nil {
		return true
	}

	return false
}

// SetPriceTreatment gets a reference to the given string and assigns it to the PriceTreatment field.
func (o *MarketingPrice) SetPriceTreatment(v string) {
	o.PriceTreatment = &v
}

func (o MarketingPrice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DiscountAmount != nil {
		toSerialize["discountAmount"] = o.DiscountAmount
	}
	if o.DiscountPercentage != nil {
		toSerialize["discountPercentage"] = o.DiscountPercentage
	}
	if o.OriginalPrice != nil {
		toSerialize["originalPrice"] = o.OriginalPrice
	}
	if o.PriceTreatment != nil {
		toSerialize["priceTreatment"] = o.PriceTreatment
	}
	return json.Marshal(toSerialize)
}

type NullableMarketingPrice struct {
	value *MarketingPrice
	isSet bool
}

func (v NullableMarketingPrice) Get() *MarketingPrice {
	return v.value
}

func (v *NullableMarketingPrice) Set(val *MarketingPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingPrice(val *MarketingPrice) *NullableMarketingPrice {
	return &NullableMarketingPrice{value: val, isSet: true}
}

func (v NullableMarketingPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

