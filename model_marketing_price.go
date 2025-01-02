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

// checks if the MarketingPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketingPrice{}

// MarketingPrice The type that defines the fields that describe a seller discount.
type MarketingPrice struct {
	DiscountAmount *ConvertedAmount `json:"discountAmount,omitempty"`
	// This field expresses the percentage of the seller discount based on the value in the <code>originalPrice</code> container.
	DiscountPercentage *string `json:"discountPercentage,omitempty"`
	OriginalPrice *ConvertedAmount `json:"originalPrice,omitempty"`
	// Indicates the pricing treatment (discount) that was applied to the price of the item.<br><br><span class=\"tablenote\"><b>Note:</b> The pricing treatment affects the way and where the discounted price can be displayed.</span> For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:PriceTreatmentEnum'>eBay API documentation</a>
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
	if o == nil || IsNil(o.DiscountAmount) {
		var ret ConvertedAmount
		return ret
	}
	return *o.DiscountAmount
}

// GetDiscountAmountOk returns a tuple with the DiscountAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingPrice) GetDiscountAmountOk() (*ConvertedAmount, bool) {
	if o == nil || IsNil(o.DiscountAmount) {
		return nil, false
	}
	return o.DiscountAmount, true
}

// HasDiscountAmount returns a boolean if a field has been set.
func (o *MarketingPrice) HasDiscountAmount() bool {
	if o != nil && !IsNil(o.DiscountAmount) {
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
	if o == nil || IsNil(o.DiscountPercentage) {
		var ret string
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingPrice) GetDiscountPercentageOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountPercentage) {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *MarketingPrice) HasDiscountPercentage() bool {
	if o != nil && !IsNil(o.DiscountPercentage) {
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
	if o == nil || IsNil(o.OriginalPrice) {
		var ret ConvertedAmount
		return ret
	}
	return *o.OriginalPrice
}

// GetOriginalPriceOk returns a tuple with the OriginalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingPrice) GetOriginalPriceOk() (*ConvertedAmount, bool) {
	if o == nil || IsNil(o.OriginalPrice) {
		return nil, false
	}
	return o.OriginalPrice, true
}

// HasOriginalPrice returns a boolean if a field has been set.
func (o *MarketingPrice) HasOriginalPrice() bool {
	if o != nil && !IsNil(o.OriginalPrice) {
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
	if o == nil || IsNil(o.PriceTreatment) {
		var ret string
		return ret
	}
	return *o.PriceTreatment
}

// GetPriceTreatmentOk returns a tuple with the PriceTreatment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingPrice) GetPriceTreatmentOk() (*string, bool) {
	if o == nil || IsNil(o.PriceTreatment) {
		return nil, false
	}
	return o.PriceTreatment, true
}

// HasPriceTreatment returns a boolean if a field has been set.
func (o *MarketingPrice) HasPriceTreatment() bool {
	if o != nil && !IsNil(o.PriceTreatment) {
		return true
	}

	return false
}

// SetPriceTreatment gets a reference to the given string and assigns it to the PriceTreatment field.
func (o *MarketingPrice) SetPriceTreatment(v string) {
	o.PriceTreatment = &v
}

func (o MarketingPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketingPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiscountAmount) {
		toSerialize["discountAmount"] = o.DiscountAmount
	}
	if !IsNil(o.DiscountPercentage) {
		toSerialize["discountPercentage"] = o.DiscountPercentage
	}
	if !IsNil(o.OriginalPrice) {
		toSerialize["originalPrice"] = o.OriginalPrice
	}
	if !IsNil(o.PriceTreatment) {
		toSerialize["priceTreatment"] = o.PriceTreatment
	}
	return toSerialize, nil
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


