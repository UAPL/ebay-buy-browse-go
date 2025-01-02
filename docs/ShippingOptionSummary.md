# ShippingOptionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuaranteedDelivery** | Pointer to **bool** | Although this field is still returned, it can be ignored since eBay Guaranteed Delivery is no longer a supported feature on any marketplace. This field may get removed from the schema in the future. | [optional] 
**MaxEstimatedDeliveryDate** | Pointer to **string** | The end date of the delivery window (latest projected delivery date). This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer.&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt; &lt;b&gt;Note: &lt;/b&gt; For the best accuracy, always include the &lt;code&gt;contextualLocation&lt;/code&gt; values in the &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Headers\&quot; target&#x3D;\&quot;_blank\&quot;&gt;&lt;code&gt;X-EBAY-C-ENDUSERCTX&lt;/code&gt;&lt;/a&gt; request header.&lt;/span&gt; | [optional] 
**MinEstimatedDeliveryDate** | Pointer to **string** | The start date of the delivery window (earliest projected delivery date). This value is returned in UTC format (yyyy-MM-ddThh:mm:ss.sssZ), which you can convert into the local time of the buyer.&lt;br&gt;&lt;br&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note:&lt;/b&gt; For the best accuracy, always include the &lt;code&gt;contextualLocation&lt;/code&gt; values in the &lt;a href&#x3D;\&quot;/api-docs/buy/static/api-browse.html#Headers\&quot; target&#x3D;\&quot;_blank\&quot;&gt;&lt;code&gt;X-EBAY-C-ENDUSERCTX&lt;/code&gt;&lt;/a&gt; request header.&lt;/span&gt; | [optional] 
**ShippingCost** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**ShippingCostType** | Pointer to **string** | Indicates the type of shipping used to ship the item. Possible values are &lt;code&gt;FIXED&lt;/code&gt; (flat-rate shipping) and &lt;code&gt;CALCULATED&lt;/code&gt; (shipping cost calculated based on item and buyer location). | [optional] 

## Methods

### NewShippingOptionSummary

`func NewShippingOptionSummary() *ShippingOptionSummary`

NewShippingOptionSummary instantiates a new ShippingOptionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingOptionSummaryWithDefaults

`func NewShippingOptionSummaryWithDefaults() *ShippingOptionSummary`

NewShippingOptionSummaryWithDefaults instantiates a new ShippingOptionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuaranteedDelivery

`func (o *ShippingOptionSummary) GetGuaranteedDelivery() bool`

GetGuaranteedDelivery returns the GuaranteedDelivery field if non-nil, zero value otherwise.

### GetGuaranteedDeliveryOk

`func (o *ShippingOptionSummary) GetGuaranteedDeliveryOk() (*bool, bool)`

GetGuaranteedDeliveryOk returns a tuple with the GuaranteedDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedDelivery

`func (o *ShippingOptionSummary) SetGuaranteedDelivery(v bool)`

SetGuaranteedDelivery sets GuaranteedDelivery field to given value.

### HasGuaranteedDelivery

`func (o *ShippingOptionSummary) HasGuaranteedDelivery() bool`

HasGuaranteedDelivery returns a boolean if a field has been set.

### GetMaxEstimatedDeliveryDate

`func (o *ShippingOptionSummary) GetMaxEstimatedDeliveryDate() string`

GetMaxEstimatedDeliveryDate returns the MaxEstimatedDeliveryDate field if non-nil, zero value otherwise.

### GetMaxEstimatedDeliveryDateOk

`func (o *ShippingOptionSummary) GetMaxEstimatedDeliveryDateOk() (*string, bool)`

GetMaxEstimatedDeliveryDateOk returns a tuple with the MaxEstimatedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEstimatedDeliveryDate

`func (o *ShippingOptionSummary) SetMaxEstimatedDeliveryDate(v string)`

SetMaxEstimatedDeliveryDate sets MaxEstimatedDeliveryDate field to given value.

### HasMaxEstimatedDeliveryDate

`func (o *ShippingOptionSummary) HasMaxEstimatedDeliveryDate() bool`

HasMaxEstimatedDeliveryDate returns a boolean if a field has been set.

### GetMinEstimatedDeliveryDate

`func (o *ShippingOptionSummary) GetMinEstimatedDeliveryDate() string`

GetMinEstimatedDeliveryDate returns the MinEstimatedDeliveryDate field if non-nil, zero value otherwise.

### GetMinEstimatedDeliveryDateOk

`func (o *ShippingOptionSummary) GetMinEstimatedDeliveryDateOk() (*string, bool)`

GetMinEstimatedDeliveryDateOk returns a tuple with the MinEstimatedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEstimatedDeliveryDate

`func (o *ShippingOptionSummary) SetMinEstimatedDeliveryDate(v string)`

SetMinEstimatedDeliveryDate sets MinEstimatedDeliveryDate field to given value.

### HasMinEstimatedDeliveryDate

`func (o *ShippingOptionSummary) HasMinEstimatedDeliveryDate() bool`

HasMinEstimatedDeliveryDate returns a boolean if a field has been set.

### GetShippingCost

`func (o *ShippingOptionSummary) GetShippingCost() ConvertedAmount`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *ShippingOptionSummary) GetShippingCostOk() (*ConvertedAmount, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *ShippingOptionSummary) SetShippingCost(v ConvertedAmount)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *ShippingOptionSummary) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### GetShippingCostType

`func (o *ShippingOptionSummary) GetShippingCostType() string`

GetShippingCostType returns the ShippingCostType field if non-nil, zero value otherwise.

### GetShippingCostTypeOk

`func (o *ShippingOptionSummary) GetShippingCostTypeOk() (*string, bool)`

GetShippingCostTypeOk returns a tuple with the ShippingCostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCostType

`func (o *ShippingOptionSummary) SetShippingCostType(v string)`

SetShippingCostType sets ShippingCostType field to given value.

### HasShippingCostType

`func (o *ShippingOptionSummary) HasShippingCostType() bool`

HasShippingCostType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


