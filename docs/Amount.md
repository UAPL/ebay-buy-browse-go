# Amount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The list of valid currencies. Each &lt;a href&#x3D;\&quot;https://www.iso.org/iso-4217-currency-codes.html \&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 4217&lt;/a&gt; currency code includes the currency name followed by the numeric value.&lt;br&gt;&lt;br&gt;For example, the Canadian Dollar code (CAD) would take the following form: &lt;code&gt;Canadian Dollar, 124&lt;/code&gt;. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/ba:CurrencyCodeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**Value** | Pointer to **string** | The value of the discounted amount. | [optional] 

## Methods

### NewAmount

`func NewAmount() *Amount`

NewAmount instantiates a new Amount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmountWithDefaults

`func NewAmountWithDefaults() *Amount`

NewAmountWithDefaults instantiates a new Amount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Amount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Amount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Amount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Amount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetValue

`func (o *Amount) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Amount) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Amount) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Amount) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


