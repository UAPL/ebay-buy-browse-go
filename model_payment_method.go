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

// checks if the PaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethod{}

// PaymentMethod struct for PaymentMethod
type PaymentMethod struct {
	// The payment method type, such as credit card or cash. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:PaymentMethodTypeEnum'>eBay API documentation</a>
	PaymentMethodType *string `json:"paymentMethodType,omitempty"`
	// The payment method brands, including the payment method brand type and logo image.
	PaymentMethodBrands []PaymentMethodBrand `json:"paymentMethodBrands,omitempty"`
	// The payment instructions for the buyer, such as <i>cash in person</i> or <i>contact seller</i>.
	PaymentInstructions []string `json:"paymentInstructions,omitempty"`
	// The seller instructions to the buyer, such as <i>accepts credit cards</i> or <i>see description</i>.
	SellerInstructions []string `json:"sellerInstructions,omitempty"`
}

// NewPaymentMethod instantiates a new PaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethod() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodWithDefaults() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// GetPaymentMethodType returns the PaymentMethodType field value if set, zero value otherwise.
func (o *PaymentMethod) GetPaymentMethodType() string {
	if o == nil || IsNil(o.PaymentMethodType) {
		var ret string
		return ret
	}
	return *o.PaymentMethodType
}

// GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetPaymentMethodTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethodType) {
		return nil, false
	}
	return o.PaymentMethodType, true
}

// HasPaymentMethodType returns a boolean if a field has been set.
func (o *PaymentMethod) HasPaymentMethodType() bool {
	if o != nil && !IsNil(o.PaymentMethodType) {
		return true
	}

	return false
}

// SetPaymentMethodType gets a reference to the given string and assigns it to the PaymentMethodType field.
func (o *PaymentMethod) SetPaymentMethodType(v string) {
	o.PaymentMethodType = &v
}

// GetPaymentMethodBrands returns the PaymentMethodBrands field value if set, zero value otherwise.
func (o *PaymentMethod) GetPaymentMethodBrands() []PaymentMethodBrand {
	if o == nil || IsNil(o.PaymentMethodBrands) {
		var ret []PaymentMethodBrand
		return ret
	}
	return o.PaymentMethodBrands
}

// GetPaymentMethodBrandsOk returns a tuple with the PaymentMethodBrands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetPaymentMethodBrandsOk() ([]PaymentMethodBrand, bool) {
	if o == nil || IsNil(o.PaymentMethodBrands) {
		return nil, false
	}
	return o.PaymentMethodBrands, true
}

// HasPaymentMethodBrands returns a boolean if a field has been set.
func (o *PaymentMethod) HasPaymentMethodBrands() bool {
	if o != nil && !IsNil(o.PaymentMethodBrands) {
		return true
	}

	return false
}

// SetPaymentMethodBrands gets a reference to the given []PaymentMethodBrand and assigns it to the PaymentMethodBrands field.
func (o *PaymentMethod) SetPaymentMethodBrands(v []PaymentMethodBrand) {
	o.PaymentMethodBrands = v
}

// GetPaymentInstructions returns the PaymentInstructions field value if set, zero value otherwise.
func (o *PaymentMethod) GetPaymentInstructions() []string {
	if o == nil || IsNil(o.PaymentInstructions) {
		var ret []string
		return ret
	}
	return o.PaymentInstructions
}

// GetPaymentInstructionsOk returns a tuple with the PaymentInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetPaymentInstructionsOk() ([]string, bool) {
	if o == nil || IsNil(o.PaymentInstructions) {
		return nil, false
	}
	return o.PaymentInstructions, true
}

// HasPaymentInstructions returns a boolean if a field has been set.
func (o *PaymentMethod) HasPaymentInstructions() bool {
	if o != nil && !IsNil(o.PaymentInstructions) {
		return true
	}

	return false
}

// SetPaymentInstructions gets a reference to the given []string and assigns it to the PaymentInstructions field.
func (o *PaymentMethod) SetPaymentInstructions(v []string) {
	o.PaymentInstructions = v
}

// GetSellerInstructions returns the SellerInstructions field value if set, zero value otherwise.
func (o *PaymentMethod) GetSellerInstructions() []string {
	if o == nil || IsNil(o.SellerInstructions) {
		var ret []string
		return ret
	}
	return o.SellerInstructions
}

// GetSellerInstructionsOk returns a tuple with the SellerInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetSellerInstructionsOk() ([]string, bool) {
	if o == nil || IsNil(o.SellerInstructions) {
		return nil, false
	}
	return o.SellerInstructions, true
}

// HasSellerInstructions returns a boolean if a field has been set.
func (o *PaymentMethod) HasSellerInstructions() bool {
	if o != nil && !IsNil(o.SellerInstructions) {
		return true
	}

	return false
}

// SetSellerInstructions gets a reference to the given []string and assigns it to the SellerInstructions field.
func (o *PaymentMethod) SetSellerInstructions(v []string) {
	o.SellerInstructions = v
}

func (o PaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentMethodType) {
		toSerialize["paymentMethodType"] = o.PaymentMethodType
	}
	if !IsNil(o.PaymentMethodBrands) {
		toSerialize["paymentMethodBrands"] = o.PaymentMethodBrands
	}
	if !IsNil(o.PaymentInstructions) {
		toSerialize["paymentInstructions"] = o.PaymentInstructions
	}
	if !IsNil(o.SellerInstructions) {
		toSerialize["sellerInstructions"] = o.SellerInstructions
	}
	return toSerialize, nil
}

type NullablePaymentMethod struct {
	value *PaymentMethod
	isSet bool
}

func (v NullablePaymentMethod) Get() *PaymentMethod {
	return v.value
}

func (v *NullablePaymentMethod) Set(val *PaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethod(val *PaymentMethod) *NullablePaymentMethod {
	return &NullablePaymentMethod{value: val, isSet: true}
}

func (v NullablePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


