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

// checks if the ShippingOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingOption{}

// ShippingOption The type that defines the fields for the details of a shipping provider.
type ShippingOption struct {
	AdditionalShippingCostPerUnit *ConvertedAmount `json:"additionalShippingCostPerUnit,omitempty"`
	// The deadline date that the item must be purchased by in order to be received by the buyer within the delivery window (<b> maxEstimatedDeliveryDate</b> and  <b> minEstimatedDeliveryDate</b> fields). This field is returned only for items that are eligible for 'Same Day Handling'. For these items, the value of this field is what is displayed in the <b> Delivery</b> line on the View Item page.  <br><br>This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer.
	CutOffDateUsedForEstimate *string `json:"cutOffDateUsedForEstimate,omitempty"`
	// If the item is being shipped by the eBay <a href=\"https://pages.ebay.com/seller-center/shipping/global-shipping-program.html \">Global Shipping program</a>, this field returns <code>GLOBAL_SHIPPING</code>.<br><br>If the item is being shipped using the eBay International Shipping program, this field returns <code>INTERNATIONAL_SHIPPING</code>. <br><br>Otherwise, this field is null. For implementation help, refer to <a href='https://developer.ebay.com/api-docs/buy/browse/types/gct:FulfilledThroughEnum'>eBay API documentation</a>
	FulfilledThrough *string `json:"fulfilledThrough,omitempty"`
	// Although this field is still returned, it can be ignored since eBay Guaranteed Delivery is no longer a supported feature on any marketplace. This field may get removed from the schema in the future.
	GuaranteedDelivery *bool `json:"guaranteedDelivery,omitempty"`
	ImportCharges *ConvertedAmount `json:"importCharges,omitempty"`
	// The end date of the delivery window (latest projected delivery date).  This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer. <br> <br> <span class=\"tablenote\"> <b> Note: </b> For the best accuracy, always include the location of where the item is be shipped in the <code> contextualLocation</code> values of the <a href=\"/api-docs/buy/static/api-browse.html#Headers\"> <code>X-EBAY-C-ENDUSERCTX</code></a> request header.</span> 
	MaxEstimatedDeliveryDate *string `json:"maxEstimatedDeliveryDate,omitempty"`
	// The start date of the delivery window (earliest projected delivery date). This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer. <br> <br><span class=\"tablenote\"> <b> Note: </b> For the best accuracy, always include the location of where the item is be shipped in the <code> contextualLocation</code> values of the <a href=\"/api-docs/buy/static/api-browse.html#Headers\"> <code>X-EBAY-C-ENDUSERCTX</code></a> request header.</span>
	MinEstimatedDeliveryDate *string `json:"minEstimatedDeliveryDate,omitempty"`
	// The number of items used when calculating the estimation information.
	QuantityUsedForEstimate *int32 `json:"quantityUsedForEstimate,omitempty"`
	// The name of the shipping provider, such as FedEx, or USPS.
	ShippingCarrierCode *string `json:"shippingCarrierCode,omitempty"`
	ShippingCost *ConvertedAmount `json:"shippingCost,omitempty"`
	// Indicates the class of the shipping cost. <br><br><b> Valid Values: </b> FIXED or CALCULATED <br><br>Code so that your app gracefully handles any future changes to this list. 
	ShippingCostType *string `json:"shippingCostType,omitempty"`
	// The type of shipping service. For example, USPS First Class.
	ShippingServiceCode *string `json:"shippingServiceCode,omitempty"`
	ShipToLocationUsedForEstimate *ShipToLocation `json:"shipToLocationUsedForEstimate,omitempty"`
	// Any trademark symbol, such as &#8482; or &reg;, that needs to be shown in superscript next to the shipping service name.
	TrademarkSymbol *string `json:"trademarkSymbol,omitempty"`
	// The type of a shipping option, such as EXPEDITED, ONE_DAY, STANDARD, ECONOMY, PICKUP, etc.
	Type *string `json:"type,omitempty"`
}

// NewShippingOption instantiates a new ShippingOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingOption() *ShippingOption {
	this := ShippingOption{}
	return &this
}

// NewShippingOptionWithDefaults instantiates a new ShippingOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingOptionWithDefaults() *ShippingOption {
	this := ShippingOption{}
	return &this
}

// GetAdditionalShippingCostPerUnit returns the AdditionalShippingCostPerUnit field value if set, zero value otherwise.
func (o *ShippingOption) GetAdditionalShippingCostPerUnit() ConvertedAmount {
	if o == nil || IsNil(o.AdditionalShippingCostPerUnit) {
		var ret ConvertedAmount
		return ret
	}
	return *o.AdditionalShippingCostPerUnit
}

// GetAdditionalShippingCostPerUnitOk returns a tuple with the AdditionalShippingCostPerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetAdditionalShippingCostPerUnitOk() (*ConvertedAmount, bool) {
	if o == nil || IsNil(o.AdditionalShippingCostPerUnit) {
		return nil, false
	}
	return o.AdditionalShippingCostPerUnit, true
}

// HasAdditionalShippingCostPerUnit returns a boolean if a field has been set.
func (o *ShippingOption) HasAdditionalShippingCostPerUnit() bool {
	if o != nil && !IsNil(o.AdditionalShippingCostPerUnit) {
		return true
	}

	return false
}

// SetAdditionalShippingCostPerUnit gets a reference to the given ConvertedAmount and assigns it to the AdditionalShippingCostPerUnit field.
func (o *ShippingOption) SetAdditionalShippingCostPerUnit(v ConvertedAmount) {
	o.AdditionalShippingCostPerUnit = &v
}

// GetCutOffDateUsedForEstimate returns the CutOffDateUsedForEstimate field value if set, zero value otherwise.
func (o *ShippingOption) GetCutOffDateUsedForEstimate() string {
	if o == nil || IsNil(o.CutOffDateUsedForEstimate) {
		var ret string
		return ret
	}
	return *o.CutOffDateUsedForEstimate
}

// GetCutOffDateUsedForEstimateOk returns a tuple with the CutOffDateUsedForEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetCutOffDateUsedForEstimateOk() (*string, bool) {
	if o == nil || IsNil(o.CutOffDateUsedForEstimate) {
		return nil, false
	}
	return o.CutOffDateUsedForEstimate, true
}

// HasCutOffDateUsedForEstimate returns a boolean if a field has been set.
func (o *ShippingOption) HasCutOffDateUsedForEstimate() bool {
	if o != nil && !IsNil(o.CutOffDateUsedForEstimate) {
		return true
	}

	return false
}

// SetCutOffDateUsedForEstimate gets a reference to the given string and assigns it to the CutOffDateUsedForEstimate field.
func (o *ShippingOption) SetCutOffDateUsedForEstimate(v string) {
	o.CutOffDateUsedForEstimate = &v
}

// GetFulfilledThrough returns the FulfilledThrough field value if set, zero value otherwise.
func (o *ShippingOption) GetFulfilledThrough() string {
	if o == nil || IsNil(o.FulfilledThrough) {
		var ret string
		return ret
	}
	return *o.FulfilledThrough
}

// GetFulfilledThroughOk returns a tuple with the FulfilledThrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetFulfilledThroughOk() (*string, bool) {
	if o == nil || IsNil(o.FulfilledThrough) {
		return nil, false
	}
	return o.FulfilledThrough, true
}

// HasFulfilledThrough returns a boolean if a field has been set.
func (o *ShippingOption) HasFulfilledThrough() bool {
	if o != nil && !IsNil(o.FulfilledThrough) {
		return true
	}

	return false
}

// SetFulfilledThrough gets a reference to the given string and assigns it to the FulfilledThrough field.
func (o *ShippingOption) SetFulfilledThrough(v string) {
	o.FulfilledThrough = &v
}

// GetGuaranteedDelivery returns the GuaranteedDelivery field value if set, zero value otherwise.
func (o *ShippingOption) GetGuaranteedDelivery() bool {
	if o == nil || IsNil(o.GuaranteedDelivery) {
		var ret bool
		return ret
	}
	return *o.GuaranteedDelivery
}

// GetGuaranteedDeliveryOk returns a tuple with the GuaranteedDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetGuaranteedDeliveryOk() (*bool, bool) {
	if o == nil || IsNil(o.GuaranteedDelivery) {
		return nil, false
	}
	return o.GuaranteedDelivery, true
}

// HasGuaranteedDelivery returns a boolean if a field has been set.
func (o *ShippingOption) HasGuaranteedDelivery() bool {
	if o != nil && !IsNil(o.GuaranteedDelivery) {
		return true
	}

	return false
}

// SetGuaranteedDelivery gets a reference to the given bool and assigns it to the GuaranteedDelivery field.
func (o *ShippingOption) SetGuaranteedDelivery(v bool) {
	o.GuaranteedDelivery = &v
}

// GetImportCharges returns the ImportCharges field value if set, zero value otherwise.
func (o *ShippingOption) GetImportCharges() ConvertedAmount {
	if o == nil || IsNil(o.ImportCharges) {
		var ret ConvertedAmount
		return ret
	}
	return *o.ImportCharges
}

// GetImportChargesOk returns a tuple with the ImportCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetImportChargesOk() (*ConvertedAmount, bool) {
	if o == nil || IsNil(o.ImportCharges) {
		return nil, false
	}
	return o.ImportCharges, true
}

// HasImportCharges returns a boolean if a field has been set.
func (o *ShippingOption) HasImportCharges() bool {
	if o != nil && !IsNil(o.ImportCharges) {
		return true
	}

	return false
}

// SetImportCharges gets a reference to the given ConvertedAmount and assigns it to the ImportCharges field.
func (o *ShippingOption) SetImportCharges(v ConvertedAmount) {
	o.ImportCharges = &v
}

// GetMaxEstimatedDeliveryDate returns the MaxEstimatedDeliveryDate field value if set, zero value otherwise.
func (o *ShippingOption) GetMaxEstimatedDeliveryDate() string {
	if o == nil || IsNil(o.MaxEstimatedDeliveryDate) {
		var ret string
		return ret
	}
	return *o.MaxEstimatedDeliveryDate
}

// GetMaxEstimatedDeliveryDateOk returns a tuple with the MaxEstimatedDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetMaxEstimatedDeliveryDateOk() (*string, bool) {
	if o == nil || IsNil(o.MaxEstimatedDeliveryDate) {
		return nil, false
	}
	return o.MaxEstimatedDeliveryDate, true
}

// HasMaxEstimatedDeliveryDate returns a boolean if a field has been set.
func (o *ShippingOption) HasMaxEstimatedDeliveryDate() bool {
	if o != nil && !IsNil(o.MaxEstimatedDeliveryDate) {
		return true
	}

	return false
}

// SetMaxEstimatedDeliveryDate gets a reference to the given string and assigns it to the MaxEstimatedDeliveryDate field.
func (o *ShippingOption) SetMaxEstimatedDeliveryDate(v string) {
	o.MaxEstimatedDeliveryDate = &v
}

// GetMinEstimatedDeliveryDate returns the MinEstimatedDeliveryDate field value if set, zero value otherwise.
func (o *ShippingOption) GetMinEstimatedDeliveryDate() string {
	if o == nil || IsNil(o.MinEstimatedDeliveryDate) {
		var ret string
		return ret
	}
	return *o.MinEstimatedDeliveryDate
}

// GetMinEstimatedDeliveryDateOk returns a tuple with the MinEstimatedDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetMinEstimatedDeliveryDateOk() (*string, bool) {
	if o == nil || IsNil(o.MinEstimatedDeliveryDate) {
		return nil, false
	}
	return o.MinEstimatedDeliveryDate, true
}

// HasMinEstimatedDeliveryDate returns a boolean if a field has been set.
func (o *ShippingOption) HasMinEstimatedDeliveryDate() bool {
	if o != nil && !IsNil(o.MinEstimatedDeliveryDate) {
		return true
	}

	return false
}

// SetMinEstimatedDeliveryDate gets a reference to the given string and assigns it to the MinEstimatedDeliveryDate field.
func (o *ShippingOption) SetMinEstimatedDeliveryDate(v string) {
	o.MinEstimatedDeliveryDate = &v
}

// GetQuantityUsedForEstimate returns the QuantityUsedForEstimate field value if set, zero value otherwise.
func (o *ShippingOption) GetQuantityUsedForEstimate() int32 {
	if o == nil || IsNil(o.QuantityUsedForEstimate) {
		var ret int32
		return ret
	}
	return *o.QuantityUsedForEstimate
}

// GetQuantityUsedForEstimateOk returns a tuple with the QuantityUsedForEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetQuantityUsedForEstimateOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityUsedForEstimate) {
		return nil, false
	}
	return o.QuantityUsedForEstimate, true
}

// HasQuantityUsedForEstimate returns a boolean if a field has been set.
func (o *ShippingOption) HasQuantityUsedForEstimate() bool {
	if o != nil && !IsNil(o.QuantityUsedForEstimate) {
		return true
	}

	return false
}

// SetQuantityUsedForEstimate gets a reference to the given int32 and assigns it to the QuantityUsedForEstimate field.
func (o *ShippingOption) SetQuantityUsedForEstimate(v int32) {
	o.QuantityUsedForEstimate = &v
}

// GetShippingCarrierCode returns the ShippingCarrierCode field value if set, zero value otherwise.
func (o *ShippingOption) GetShippingCarrierCode() string {
	if o == nil || IsNil(o.ShippingCarrierCode) {
		var ret string
		return ret
	}
	return *o.ShippingCarrierCode
}

// GetShippingCarrierCodeOk returns a tuple with the ShippingCarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetShippingCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingCarrierCode) {
		return nil, false
	}
	return o.ShippingCarrierCode, true
}

// HasShippingCarrierCode returns a boolean if a field has been set.
func (o *ShippingOption) HasShippingCarrierCode() bool {
	if o != nil && !IsNil(o.ShippingCarrierCode) {
		return true
	}

	return false
}

// SetShippingCarrierCode gets a reference to the given string and assigns it to the ShippingCarrierCode field.
func (o *ShippingOption) SetShippingCarrierCode(v string) {
	o.ShippingCarrierCode = &v
}

// GetShippingCost returns the ShippingCost field value if set, zero value otherwise.
func (o *ShippingOption) GetShippingCost() ConvertedAmount {
	if o == nil || IsNil(o.ShippingCost) {
		var ret ConvertedAmount
		return ret
	}
	return *o.ShippingCost
}

// GetShippingCostOk returns a tuple with the ShippingCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetShippingCostOk() (*ConvertedAmount, bool) {
	if o == nil || IsNil(o.ShippingCost) {
		return nil, false
	}
	return o.ShippingCost, true
}

// HasShippingCost returns a boolean if a field has been set.
func (o *ShippingOption) HasShippingCost() bool {
	if o != nil && !IsNil(o.ShippingCost) {
		return true
	}

	return false
}

// SetShippingCost gets a reference to the given ConvertedAmount and assigns it to the ShippingCost field.
func (o *ShippingOption) SetShippingCost(v ConvertedAmount) {
	o.ShippingCost = &v
}

// GetShippingCostType returns the ShippingCostType field value if set, zero value otherwise.
func (o *ShippingOption) GetShippingCostType() string {
	if o == nil || IsNil(o.ShippingCostType) {
		var ret string
		return ret
	}
	return *o.ShippingCostType
}

// GetShippingCostTypeOk returns a tuple with the ShippingCostType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetShippingCostTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingCostType) {
		return nil, false
	}
	return o.ShippingCostType, true
}

// HasShippingCostType returns a boolean if a field has been set.
func (o *ShippingOption) HasShippingCostType() bool {
	if o != nil && !IsNil(o.ShippingCostType) {
		return true
	}

	return false
}

// SetShippingCostType gets a reference to the given string and assigns it to the ShippingCostType field.
func (o *ShippingOption) SetShippingCostType(v string) {
	o.ShippingCostType = &v
}

// GetShippingServiceCode returns the ShippingServiceCode field value if set, zero value otherwise.
func (o *ShippingOption) GetShippingServiceCode() string {
	if o == nil || IsNil(o.ShippingServiceCode) {
		var ret string
		return ret
	}
	return *o.ShippingServiceCode
}

// GetShippingServiceCodeOk returns a tuple with the ShippingServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetShippingServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingServiceCode) {
		return nil, false
	}
	return o.ShippingServiceCode, true
}

// HasShippingServiceCode returns a boolean if a field has been set.
func (o *ShippingOption) HasShippingServiceCode() bool {
	if o != nil && !IsNil(o.ShippingServiceCode) {
		return true
	}

	return false
}

// SetShippingServiceCode gets a reference to the given string and assigns it to the ShippingServiceCode field.
func (o *ShippingOption) SetShippingServiceCode(v string) {
	o.ShippingServiceCode = &v
}

// GetShipToLocationUsedForEstimate returns the ShipToLocationUsedForEstimate field value if set, zero value otherwise.
func (o *ShippingOption) GetShipToLocationUsedForEstimate() ShipToLocation {
	if o == nil || IsNil(o.ShipToLocationUsedForEstimate) {
		var ret ShipToLocation
		return ret
	}
	return *o.ShipToLocationUsedForEstimate
}

// GetShipToLocationUsedForEstimateOk returns a tuple with the ShipToLocationUsedForEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetShipToLocationUsedForEstimateOk() (*ShipToLocation, bool) {
	if o == nil || IsNil(o.ShipToLocationUsedForEstimate) {
		return nil, false
	}
	return o.ShipToLocationUsedForEstimate, true
}

// HasShipToLocationUsedForEstimate returns a boolean if a field has been set.
func (o *ShippingOption) HasShipToLocationUsedForEstimate() bool {
	if o != nil && !IsNil(o.ShipToLocationUsedForEstimate) {
		return true
	}

	return false
}

// SetShipToLocationUsedForEstimate gets a reference to the given ShipToLocation and assigns it to the ShipToLocationUsedForEstimate field.
func (o *ShippingOption) SetShipToLocationUsedForEstimate(v ShipToLocation) {
	o.ShipToLocationUsedForEstimate = &v
}

// GetTrademarkSymbol returns the TrademarkSymbol field value if set, zero value otherwise.
func (o *ShippingOption) GetTrademarkSymbol() string {
	if o == nil || IsNil(o.TrademarkSymbol) {
		var ret string
		return ret
	}
	return *o.TrademarkSymbol
}

// GetTrademarkSymbolOk returns a tuple with the TrademarkSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetTrademarkSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.TrademarkSymbol) {
		return nil, false
	}
	return o.TrademarkSymbol, true
}

// HasTrademarkSymbol returns a boolean if a field has been set.
func (o *ShippingOption) HasTrademarkSymbol() bool {
	if o != nil && !IsNil(o.TrademarkSymbol) {
		return true
	}

	return false
}

// SetTrademarkSymbol gets a reference to the given string and assigns it to the TrademarkSymbol field.
func (o *ShippingOption) SetTrademarkSymbol(v string) {
	o.TrademarkSymbol = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ShippingOption) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingOption) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ShippingOption) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ShippingOption) SetType(v string) {
	o.Type = &v
}

func (o ShippingOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShippingOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalShippingCostPerUnit) {
		toSerialize["additionalShippingCostPerUnit"] = o.AdditionalShippingCostPerUnit
	}
	if !IsNil(o.CutOffDateUsedForEstimate) {
		toSerialize["cutOffDateUsedForEstimate"] = o.CutOffDateUsedForEstimate
	}
	if !IsNil(o.FulfilledThrough) {
		toSerialize["fulfilledThrough"] = o.FulfilledThrough
	}
	if !IsNil(o.GuaranteedDelivery) {
		toSerialize["guaranteedDelivery"] = o.GuaranteedDelivery
	}
	if !IsNil(o.ImportCharges) {
		toSerialize["importCharges"] = o.ImportCharges
	}
	if !IsNil(o.MaxEstimatedDeliveryDate) {
		toSerialize["maxEstimatedDeliveryDate"] = o.MaxEstimatedDeliveryDate
	}
	if !IsNil(o.MinEstimatedDeliveryDate) {
		toSerialize["minEstimatedDeliveryDate"] = o.MinEstimatedDeliveryDate
	}
	if !IsNil(o.QuantityUsedForEstimate) {
		toSerialize["quantityUsedForEstimate"] = o.QuantityUsedForEstimate
	}
	if !IsNil(o.ShippingCarrierCode) {
		toSerialize["shippingCarrierCode"] = o.ShippingCarrierCode
	}
	if !IsNil(o.ShippingCost) {
		toSerialize["shippingCost"] = o.ShippingCost
	}
	if !IsNil(o.ShippingCostType) {
		toSerialize["shippingCostType"] = o.ShippingCostType
	}
	if !IsNil(o.ShippingServiceCode) {
		toSerialize["shippingServiceCode"] = o.ShippingServiceCode
	}
	if !IsNil(o.ShipToLocationUsedForEstimate) {
		toSerialize["shipToLocationUsedForEstimate"] = o.ShipToLocationUsedForEstimate
	}
	if !IsNil(o.TrademarkSymbol) {
		toSerialize["trademarkSymbol"] = o.TrademarkSymbol
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableShippingOption struct {
	value *ShippingOption
	isSet bool
}

func (v NullableShippingOption) Get() *ShippingOption {
	return v.value
}

func (v *NullableShippingOption) Set(val *ShippingOption) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingOption) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingOption(val *ShippingOption) *NullableShippingOption {
	return &NullableShippingOption{value: val, isSet: true}
}

func (v NullableShippingOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


