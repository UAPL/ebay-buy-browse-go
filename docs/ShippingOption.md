# ShippingOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalShippingCostPerUnit** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**CutOffDateUsedForEstimate** | Pointer to **string** | The deadline date that the item must be purchased by in order to be received by the buyer within the delivery window (&lt;b&gt; maxEstimatedDeliveryDate&lt;/b&gt; and  &lt;b&gt; minEstimatedDeliveryDate&lt;/b&gt; fields). This field is returned only for items that are eligible for &#39;Same Day Handling&#39;. For these items, the value of this field is what is displayed in the &lt;b&gt; Delivery&lt;/b&gt; line on the View Item page.  &lt;br&gt;&lt;br&gt;This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer. | [optional] 
**FulfilledThrough** | Pointer to **string** | If the item is being shipped by the eBay &lt;a href&#x3D;\&quot;https://pages.ebay.com/seller-center/shipping/global-shipping-program.html \&quot;&gt;Global Shipping program&lt;/a&gt;, this field returns &lt;code&gt;GLOBAL_SHIPPING&lt;/code&gt;.&lt;br&gt;&lt;br&gt;If the item is being shipped using the eBay International Shipping program, this field returns &lt;code&gt;INTERNATIONAL_SHIPPING&lt;/code&gt;. &lt;br&gt;&lt;br&gt;Otherwise, this field is null. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:FulfilledThroughEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**GuaranteedDelivery** | Pointer to **bool** | Although this field is still returned, it can be ignored since eBay Guaranteed Delivery is no longer a supported feature on any marketplace. This field may get removed from the schema in the future. | [optional] 
**ImportCharges** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**MaxEstimatedDeliveryDate** | Pointer to **string** | The end date of the delivery window (latest projected delivery date).  This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer. &lt;br&gt; &lt;br&gt; &lt;span class&#x3D;\&quot;tablenote\&quot;&gt; &lt;b&gt; Note: &lt;/b&gt; For the best accuracy, always include the location of where the item is be shipped in the &lt;code&gt; contextualLocation&lt;/code&gt; values of the &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Headers\&quot;&gt; &lt;code&gt;X-EBAY-C-ENDUSERCTX&lt;/code&gt;&lt;/a&gt; request header.&lt;/span&gt;  | [optional] 
**MinEstimatedDeliveryDate** | Pointer to **string** | The start date of the delivery window (earliest projected delivery date). This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer. &lt;br&gt; &lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt; &lt;b&gt; Note: &lt;/b&gt; For the best accuracy, always include the location of where the item is be shipped in the &lt;code&gt; contextualLocation&lt;/code&gt; values of the &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Headers\&quot;&gt; &lt;code&gt;X-EBAY-C-ENDUSERCTX&lt;/code&gt;&lt;/a&gt; request header.&lt;/span&gt; | [optional] 
**QuantityUsedForEstimate** | Pointer to **int32** | The number of items used when calculating the estimation information. | [optional] 
**ShippingCarrierCode** | Pointer to **string** | The name of the shipping provider, such as FedEx, or USPS. | [optional] 
**ShippingCost** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**ShippingCostType** | Pointer to **string** | Indicates the class of the shipping cost. &lt;br&gt;&lt;br&gt;&lt;b&gt; Valid Values: &lt;/b&gt; FIXED or CALCULATED &lt;br&gt;&lt;br&gt;Code so that your app gracefully handles any future changes to this list.  | [optional] 
**ShippingServiceCode** | Pointer to **string** | The type of shipping service. For example, USPS First Class. | [optional] 
**ShipToLocationUsedForEstimate** | Pointer to [**ShipToLocation**](ShipToLocation.md) |  | [optional] 
**TrademarkSymbol** | Pointer to **string** | Any trademark symbol, such as &amp;#8482; or &amp;reg;, that needs to be shown in superscript next to the shipping service name. | [optional] 
**Type** | Pointer to **string** | The type of a shipping option, such as EXPEDITED, ONE_DAY, STANDARD, ECONOMY, PICKUP, etc. | [optional] 

## Methods

### NewShippingOption

`func NewShippingOption() *ShippingOption`

NewShippingOption instantiates a new ShippingOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingOptionWithDefaults

`func NewShippingOptionWithDefaults() *ShippingOption`

NewShippingOptionWithDefaults instantiates a new ShippingOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalShippingCostPerUnit

`func (o *ShippingOption) GetAdditionalShippingCostPerUnit() ConvertedAmount`

GetAdditionalShippingCostPerUnit returns the AdditionalShippingCostPerUnit field if non-nil, zero value otherwise.

### GetAdditionalShippingCostPerUnitOk

`func (o *ShippingOption) GetAdditionalShippingCostPerUnitOk() (*ConvertedAmount, bool)`

GetAdditionalShippingCostPerUnitOk returns a tuple with the AdditionalShippingCostPerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalShippingCostPerUnit

`func (o *ShippingOption) SetAdditionalShippingCostPerUnit(v ConvertedAmount)`

SetAdditionalShippingCostPerUnit sets AdditionalShippingCostPerUnit field to given value.

### HasAdditionalShippingCostPerUnit

`func (o *ShippingOption) HasAdditionalShippingCostPerUnit() bool`

HasAdditionalShippingCostPerUnit returns a boolean if a field has been set.

### GetCutOffDateUsedForEstimate

`func (o *ShippingOption) GetCutOffDateUsedForEstimate() string`

GetCutOffDateUsedForEstimate returns the CutOffDateUsedForEstimate field if non-nil, zero value otherwise.

### GetCutOffDateUsedForEstimateOk

`func (o *ShippingOption) GetCutOffDateUsedForEstimateOk() (*string, bool)`

GetCutOffDateUsedForEstimateOk returns a tuple with the CutOffDateUsedForEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutOffDateUsedForEstimate

`func (o *ShippingOption) SetCutOffDateUsedForEstimate(v string)`

SetCutOffDateUsedForEstimate sets CutOffDateUsedForEstimate field to given value.

### HasCutOffDateUsedForEstimate

`func (o *ShippingOption) HasCutOffDateUsedForEstimate() bool`

HasCutOffDateUsedForEstimate returns a boolean if a field has been set.

### GetFulfilledThrough

`func (o *ShippingOption) GetFulfilledThrough() string`

GetFulfilledThrough returns the FulfilledThrough field if non-nil, zero value otherwise.

### GetFulfilledThroughOk

`func (o *ShippingOption) GetFulfilledThroughOk() (*string, bool)`

GetFulfilledThroughOk returns a tuple with the FulfilledThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilledThrough

`func (o *ShippingOption) SetFulfilledThrough(v string)`

SetFulfilledThrough sets FulfilledThrough field to given value.

### HasFulfilledThrough

`func (o *ShippingOption) HasFulfilledThrough() bool`

HasFulfilledThrough returns a boolean if a field has been set.

### GetGuaranteedDelivery

`func (o *ShippingOption) GetGuaranteedDelivery() bool`

GetGuaranteedDelivery returns the GuaranteedDelivery field if non-nil, zero value otherwise.

### GetGuaranteedDeliveryOk

`func (o *ShippingOption) GetGuaranteedDeliveryOk() (*bool, bool)`

GetGuaranteedDeliveryOk returns a tuple with the GuaranteedDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedDelivery

`func (o *ShippingOption) SetGuaranteedDelivery(v bool)`

SetGuaranteedDelivery sets GuaranteedDelivery field to given value.

### HasGuaranteedDelivery

`func (o *ShippingOption) HasGuaranteedDelivery() bool`

HasGuaranteedDelivery returns a boolean if a field has been set.

### GetImportCharges

`func (o *ShippingOption) GetImportCharges() ConvertedAmount`

GetImportCharges returns the ImportCharges field if non-nil, zero value otherwise.

### GetImportChargesOk

`func (o *ShippingOption) GetImportChargesOk() (*ConvertedAmount, bool)`

GetImportChargesOk returns a tuple with the ImportCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportCharges

`func (o *ShippingOption) SetImportCharges(v ConvertedAmount)`

SetImportCharges sets ImportCharges field to given value.

### HasImportCharges

`func (o *ShippingOption) HasImportCharges() bool`

HasImportCharges returns a boolean if a field has been set.

### GetMaxEstimatedDeliveryDate

`func (o *ShippingOption) GetMaxEstimatedDeliveryDate() string`

GetMaxEstimatedDeliveryDate returns the MaxEstimatedDeliveryDate field if non-nil, zero value otherwise.

### GetMaxEstimatedDeliveryDateOk

`func (o *ShippingOption) GetMaxEstimatedDeliveryDateOk() (*string, bool)`

GetMaxEstimatedDeliveryDateOk returns a tuple with the MaxEstimatedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEstimatedDeliveryDate

`func (o *ShippingOption) SetMaxEstimatedDeliveryDate(v string)`

SetMaxEstimatedDeliveryDate sets MaxEstimatedDeliveryDate field to given value.

### HasMaxEstimatedDeliveryDate

`func (o *ShippingOption) HasMaxEstimatedDeliveryDate() bool`

HasMaxEstimatedDeliveryDate returns a boolean if a field has been set.

### GetMinEstimatedDeliveryDate

`func (o *ShippingOption) GetMinEstimatedDeliveryDate() string`

GetMinEstimatedDeliveryDate returns the MinEstimatedDeliveryDate field if non-nil, zero value otherwise.

### GetMinEstimatedDeliveryDateOk

`func (o *ShippingOption) GetMinEstimatedDeliveryDateOk() (*string, bool)`

GetMinEstimatedDeliveryDateOk returns a tuple with the MinEstimatedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEstimatedDeliveryDate

`func (o *ShippingOption) SetMinEstimatedDeliveryDate(v string)`

SetMinEstimatedDeliveryDate sets MinEstimatedDeliveryDate field to given value.

### HasMinEstimatedDeliveryDate

`func (o *ShippingOption) HasMinEstimatedDeliveryDate() bool`

HasMinEstimatedDeliveryDate returns a boolean if a field has been set.

### GetQuantityUsedForEstimate

`func (o *ShippingOption) GetQuantityUsedForEstimate() int32`

GetQuantityUsedForEstimate returns the QuantityUsedForEstimate field if non-nil, zero value otherwise.

### GetQuantityUsedForEstimateOk

`func (o *ShippingOption) GetQuantityUsedForEstimateOk() (*int32, bool)`

GetQuantityUsedForEstimateOk returns a tuple with the QuantityUsedForEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUsedForEstimate

`func (o *ShippingOption) SetQuantityUsedForEstimate(v int32)`

SetQuantityUsedForEstimate sets QuantityUsedForEstimate field to given value.

### HasQuantityUsedForEstimate

`func (o *ShippingOption) HasQuantityUsedForEstimate() bool`

HasQuantityUsedForEstimate returns a boolean if a field has been set.

### GetShippingCarrierCode

`func (o *ShippingOption) GetShippingCarrierCode() string`

GetShippingCarrierCode returns the ShippingCarrierCode field if non-nil, zero value otherwise.

### GetShippingCarrierCodeOk

`func (o *ShippingOption) GetShippingCarrierCodeOk() (*string, bool)`

GetShippingCarrierCodeOk returns a tuple with the ShippingCarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCarrierCode

`func (o *ShippingOption) SetShippingCarrierCode(v string)`

SetShippingCarrierCode sets ShippingCarrierCode field to given value.

### HasShippingCarrierCode

`func (o *ShippingOption) HasShippingCarrierCode() bool`

HasShippingCarrierCode returns a boolean if a field has been set.

### GetShippingCost

`func (o *ShippingOption) GetShippingCost() ConvertedAmount`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *ShippingOption) GetShippingCostOk() (*ConvertedAmount, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *ShippingOption) SetShippingCost(v ConvertedAmount)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *ShippingOption) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### GetShippingCostType

`func (o *ShippingOption) GetShippingCostType() string`

GetShippingCostType returns the ShippingCostType field if non-nil, zero value otherwise.

### GetShippingCostTypeOk

`func (o *ShippingOption) GetShippingCostTypeOk() (*string, bool)`

GetShippingCostTypeOk returns a tuple with the ShippingCostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCostType

`func (o *ShippingOption) SetShippingCostType(v string)`

SetShippingCostType sets ShippingCostType field to given value.

### HasShippingCostType

`func (o *ShippingOption) HasShippingCostType() bool`

HasShippingCostType returns a boolean if a field has been set.

### GetShippingServiceCode

`func (o *ShippingOption) GetShippingServiceCode() string`

GetShippingServiceCode returns the ShippingServiceCode field if non-nil, zero value otherwise.

### GetShippingServiceCodeOk

`func (o *ShippingOption) GetShippingServiceCodeOk() (*string, bool)`

GetShippingServiceCodeOk returns a tuple with the ShippingServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingServiceCode

`func (o *ShippingOption) SetShippingServiceCode(v string)`

SetShippingServiceCode sets ShippingServiceCode field to given value.

### HasShippingServiceCode

`func (o *ShippingOption) HasShippingServiceCode() bool`

HasShippingServiceCode returns a boolean if a field has been set.

### GetShipToLocationUsedForEstimate

`func (o *ShippingOption) GetShipToLocationUsedForEstimate() ShipToLocation`

GetShipToLocationUsedForEstimate returns the ShipToLocationUsedForEstimate field if non-nil, zero value otherwise.

### GetShipToLocationUsedForEstimateOk

`func (o *ShippingOption) GetShipToLocationUsedForEstimateOk() (*ShipToLocation, bool)`

GetShipToLocationUsedForEstimateOk returns a tuple with the ShipToLocationUsedForEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToLocationUsedForEstimate

`func (o *ShippingOption) SetShipToLocationUsedForEstimate(v ShipToLocation)`

SetShipToLocationUsedForEstimate sets ShipToLocationUsedForEstimate field to given value.

### HasShipToLocationUsedForEstimate

`func (o *ShippingOption) HasShipToLocationUsedForEstimate() bool`

HasShipToLocationUsedForEstimate returns a boolean if a field has been set.

### GetTrademarkSymbol

`func (o *ShippingOption) GetTrademarkSymbol() string`

GetTrademarkSymbol returns the TrademarkSymbol field if non-nil, zero value otherwise.

### GetTrademarkSymbolOk

`func (o *ShippingOption) GetTrademarkSymbolOk() (*string, bool)`

GetTrademarkSymbolOk returns a tuple with the TrademarkSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrademarkSymbol

`func (o *ShippingOption) SetTrademarkSymbol(v string)`

SetTrademarkSymbol sets TrademarkSymbol field to given value.

### HasTrademarkSymbol

`func (o *ShippingOption) HasTrademarkSymbol() bool`

HasTrademarkSymbol returns a boolean if a field has been set.

### GetType

`func (o *ShippingOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShippingOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShippingOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ShippingOption) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


