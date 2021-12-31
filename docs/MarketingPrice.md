# MarketingPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountAmount** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**DiscountPercentage** | Pointer to **string** | This field expresses the percentage of the seller discount based on the value in the &lt;b&gt;  originalPrice&lt;/b&gt; container. | [optional] 
**OriginalPrice** | Pointer to [**ConvertedAmount**](ConvertedAmount.md) |  | [optional] 
**PriceTreatment** | Pointer to **string** | Indicates the pricing treatment (discount) that was applied to the price of the item. &lt;br /&gt;&lt;br /&gt;&lt;span class&#x3D;\&quot;tablenote\&quot;&gt;&lt;b&gt;Note: &lt;/b&gt; The pricing treatment affects the way and where the discounted price can be displayed.&lt;/span&gt; For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:PriceTreatmentEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewMarketingPrice

`func NewMarketingPrice() *MarketingPrice`

NewMarketingPrice instantiates a new MarketingPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingPriceWithDefaults

`func NewMarketingPriceWithDefaults() *MarketingPrice`

NewMarketingPriceWithDefaults instantiates a new MarketingPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountAmount

`func (o *MarketingPrice) GetDiscountAmount() ConvertedAmount`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *MarketingPrice) GetDiscountAmountOk() (*ConvertedAmount, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *MarketingPrice) SetDiscountAmount(v ConvertedAmount)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *MarketingPrice) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *MarketingPrice) GetDiscountPercentage() string`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *MarketingPrice) GetDiscountPercentageOk() (*string, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *MarketingPrice) SetDiscountPercentage(v string)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *MarketingPrice) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetOriginalPrice

`func (o *MarketingPrice) GetOriginalPrice() ConvertedAmount`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *MarketingPrice) GetOriginalPriceOk() (*ConvertedAmount, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *MarketingPrice) SetOriginalPrice(v ConvertedAmount)`

SetOriginalPrice sets OriginalPrice field to given value.

### HasOriginalPrice

`func (o *MarketingPrice) HasOriginalPrice() bool`

HasOriginalPrice returns a boolean if a field has been set.

### GetPriceTreatment

`func (o *MarketingPrice) GetPriceTreatment() string`

GetPriceTreatment returns the PriceTreatment field if non-nil, zero value otherwise.

### GetPriceTreatmentOk

`func (o *MarketingPrice) GetPriceTreatmentOk() (*string, bool)`

GetPriceTreatmentOk returns a tuple with the PriceTreatment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTreatment

`func (o *MarketingPrice) SetPriceTreatment(v string)`

SetPriceTreatment sets PriceTreatment field to given value.

### HasPriceTreatment

`func (o *MarketingPrice) HasPriceTreatment() bool`

HasPriceTreatment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


