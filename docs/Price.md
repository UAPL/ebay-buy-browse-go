# Price

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConvertedFromCurrency** | Pointer to **string** | The three-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 4217&lt;/a&gt; code representing the currency of the amount in the &lt;b&gt; convertedFromValue&lt;/b&gt; field. This value is the pre-conversion currency. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/ba:CurrencyCodeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**ConvertedFromValue** | Pointer to **string** | The monetary amount before any conversion is performed, in the currency specified by the &lt;b&gt; convertedFromCurrency&lt;/b&gt; field. This value is the pre-conversion amount. The &lt;b&gt; value&lt;/b&gt; field contains the converted amount of this value, in the currency specified by the &lt;b&gt; currency&lt;/b&gt; field. | [optional] 
**Currency** | Pointer to **string** | The three-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html\&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 4217&lt;/a&gt; code representing the currency of the amount in the &lt;b&gt; value&lt;/b&gt; field. If currency conversion/localization was performed, this is the post-conversion currency of the amount in the &lt;b&gt; value&lt;/b&gt; field.&lt;br /&gt;&lt;br /&gt;&lt;b&gt; Default:&lt;/b&gt; The currency of the user&#39;s country. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/ba:CurrencyCodeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Value** | Pointer to **string** | The amount of the currency specified in the &lt;b&gt; currency&lt;/b&gt; field. The value of &lt;b&gt; currency&lt;/b&gt; defaults to the standard currency used by the country of the eBay site offering the item. If currency conversion/localization was performed, this is the post-conversion amount.&lt;br /&gt;&lt;br /&gt;&lt;b&gt; Default:&lt;/b&gt; The currency of the user&#39;s country. | [optional] 

## Methods

### NewPrice

`func NewPrice() *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConvertedFromCurrency

`func (o *Price) GetConvertedFromCurrency() string`

GetConvertedFromCurrency returns the ConvertedFromCurrency field if non-nil, zero value otherwise.

### GetConvertedFromCurrencyOk

`func (o *Price) GetConvertedFromCurrencyOk() (*string, bool)`

GetConvertedFromCurrencyOk returns a tuple with the ConvertedFromCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedFromCurrency

`func (o *Price) SetConvertedFromCurrency(v string)`

SetConvertedFromCurrency sets ConvertedFromCurrency field to given value.

### HasConvertedFromCurrency

`func (o *Price) HasConvertedFromCurrency() bool`

HasConvertedFromCurrency returns a boolean if a field has been set.

### GetConvertedFromValue

`func (o *Price) GetConvertedFromValue() string`

GetConvertedFromValue returns the ConvertedFromValue field if non-nil, zero value otherwise.

### GetConvertedFromValueOk

`func (o *Price) GetConvertedFromValueOk() (*string, bool)`

GetConvertedFromValueOk returns a tuple with the ConvertedFromValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedFromValue

`func (o *Price) SetConvertedFromValue(v string)`

SetConvertedFromValue sets ConvertedFromValue field to given value.

### HasConvertedFromValue

`func (o *Price) HasConvertedFromValue() bool`

HasConvertedFromValue returns a boolean if a field has been set.

### GetCurrency

`func (o *Price) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Price) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Price) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Price) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetValue

`func (o *Price) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Price) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Price) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Price) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


