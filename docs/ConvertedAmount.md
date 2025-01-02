# ConvertedAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConvertedFromCurrency** | Pointer to **string** | The three-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html \&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 4217&lt;/a&gt; code representing the currency of the amount in the &lt;code&gt;convertedFromValue&lt;/code&gt; field. This value is required or returned only if currency conversion/localization is required, and represents the pre-conversion currency. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/ba:CurrencyCodeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**ConvertedFromValue** | Pointer to **string** | The monetary amount before any conversion is performed, in the currency specified by the &lt;code&gt;convertedFromCurrency&lt;/code&gt; field. This value is required or returned only if currency conversion/localization is required. The &lt;code&gt;value&lt;/code&gt; field contains the converted amount of this value, in the currency specified by the &lt;code&gt;currency&lt;/code&gt; field. | [optional] 
**Currency** | Pointer to **string** | The three-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html \&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 4217&lt;/a&gt; code representing the currency of the amount in the &lt;code&gt;value&lt;/code&gt; field. If currency conversion/localization is required, this is the post-conversion currency of the amount in the &lt;code&gt;value&lt;/code&gt; field.&lt;br&gt;&lt;br&gt;&lt;b&gt;Default:&lt;/b&gt; The currency of the authenticated user&#39;s country. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/ba:CurrencyCodeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Value** | Pointer to **string** | The monetary amount in the currency specified by the &lt;code&gt;currency&lt;/code&gt; field. If currency conversion/localization is required, this value is the converted amount, and the &lt;code&gt;convertedFromValue&lt;/code&gt; field contains the amount in the original currency. | [optional] 

## Methods

### NewConvertedAmount

`func NewConvertedAmount() *ConvertedAmount`

NewConvertedAmount instantiates a new ConvertedAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertedAmountWithDefaults

`func NewConvertedAmountWithDefaults() *ConvertedAmount`

NewConvertedAmountWithDefaults instantiates a new ConvertedAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConvertedFromCurrency

`func (o *ConvertedAmount) GetConvertedFromCurrency() string`

GetConvertedFromCurrency returns the ConvertedFromCurrency field if non-nil, zero value otherwise.

### GetConvertedFromCurrencyOk

`func (o *ConvertedAmount) GetConvertedFromCurrencyOk() (*string, bool)`

GetConvertedFromCurrencyOk returns a tuple with the ConvertedFromCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedFromCurrency

`func (o *ConvertedAmount) SetConvertedFromCurrency(v string)`

SetConvertedFromCurrency sets ConvertedFromCurrency field to given value.

### HasConvertedFromCurrency

`func (o *ConvertedAmount) HasConvertedFromCurrency() bool`

HasConvertedFromCurrency returns a boolean if a field has been set.

### GetConvertedFromValue

`func (o *ConvertedAmount) GetConvertedFromValue() string`

GetConvertedFromValue returns the ConvertedFromValue field if non-nil, zero value otherwise.

### GetConvertedFromValueOk

`func (o *ConvertedAmount) GetConvertedFromValueOk() (*string, bool)`

GetConvertedFromValueOk returns a tuple with the ConvertedFromValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedFromValue

`func (o *ConvertedAmount) SetConvertedFromValue(v string)`

SetConvertedFromValue sets ConvertedFromValue field to given value.

### HasConvertedFromValue

`func (o *ConvertedAmount) HasConvertedFromValue() bool`

HasConvertedFromValue returns a boolean if a field has been set.

### GetCurrency

`func (o *ConvertedAmount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ConvertedAmount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ConvertedAmount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ConvertedAmount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetValue

`func (o *ConvertedAmount) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConvertedAmount) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConvertedAmount) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConvertedAmount) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


