# Taxes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EbayCollectAndRemitTax** | Pointer to **bool** | This field is only returned if &lt;code&gt;true&lt;/code&gt;, and indicates that eBay will collect tax (sales tax, Goods and Services tax, or VAT) for at least one line item in the order, and remit the tax to the taxing authority of the buyer&#39;s residence.  | [optional] 
**IncludedInPrice** | Pointer to **bool** | This indicates if tax was applied for the cost of the item. | [optional] 
**ShippingAndHandlingTaxed** | Pointer to **bool** | This indicates if tax is applied for the shipping cost. | [optional] 
**TaxJurisdiction** | Pointer to [**TaxJurisdiction**](TaxJurisdiction.md) |  | [optional] 
**TaxPercentage** | Pointer to **string** | The percentage of tax. | [optional] 
**TaxType** | Pointer to **string** | This field indicates the type of tax that may be collected for the item. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:TaxType&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewTaxes

`func NewTaxes() *Taxes`

NewTaxes instantiates a new Taxes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxesWithDefaults

`func NewTaxesWithDefaults() *Taxes`

NewTaxesWithDefaults instantiates a new Taxes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEbayCollectAndRemitTax

`func (o *Taxes) GetEbayCollectAndRemitTax() bool`

GetEbayCollectAndRemitTax returns the EbayCollectAndRemitTax field if non-nil, zero value otherwise.

### GetEbayCollectAndRemitTaxOk

`func (o *Taxes) GetEbayCollectAndRemitTaxOk() (*bool, bool)`

GetEbayCollectAndRemitTaxOk returns a tuple with the EbayCollectAndRemitTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayCollectAndRemitTax

`func (o *Taxes) SetEbayCollectAndRemitTax(v bool)`

SetEbayCollectAndRemitTax sets EbayCollectAndRemitTax field to given value.

### HasEbayCollectAndRemitTax

`func (o *Taxes) HasEbayCollectAndRemitTax() bool`

HasEbayCollectAndRemitTax returns a boolean if a field has been set.

### GetIncludedInPrice

`func (o *Taxes) GetIncludedInPrice() bool`

GetIncludedInPrice returns the IncludedInPrice field if non-nil, zero value otherwise.

### GetIncludedInPriceOk

`func (o *Taxes) GetIncludedInPriceOk() (*bool, bool)`

GetIncludedInPriceOk returns a tuple with the IncludedInPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedInPrice

`func (o *Taxes) SetIncludedInPrice(v bool)`

SetIncludedInPrice sets IncludedInPrice field to given value.

### HasIncludedInPrice

`func (o *Taxes) HasIncludedInPrice() bool`

HasIncludedInPrice returns a boolean if a field has been set.

### GetShippingAndHandlingTaxed

`func (o *Taxes) GetShippingAndHandlingTaxed() bool`

GetShippingAndHandlingTaxed returns the ShippingAndHandlingTaxed field if non-nil, zero value otherwise.

### GetShippingAndHandlingTaxedOk

`func (o *Taxes) GetShippingAndHandlingTaxedOk() (*bool, bool)`

GetShippingAndHandlingTaxedOk returns a tuple with the ShippingAndHandlingTaxed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAndHandlingTaxed

`func (o *Taxes) SetShippingAndHandlingTaxed(v bool)`

SetShippingAndHandlingTaxed sets ShippingAndHandlingTaxed field to given value.

### HasShippingAndHandlingTaxed

`func (o *Taxes) HasShippingAndHandlingTaxed() bool`

HasShippingAndHandlingTaxed returns a boolean if a field has been set.

### GetTaxJurisdiction

`func (o *Taxes) GetTaxJurisdiction() TaxJurisdiction`

GetTaxJurisdiction returns the TaxJurisdiction field if non-nil, zero value otherwise.

### GetTaxJurisdictionOk

`func (o *Taxes) GetTaxJurisdictionOk() (*TaxJurisdiction, bool)`

GetTaxJurisdictionOk returns a tuple with the TaxJurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxJurisdiction

`func (o *Taxes) SetTaxJurisdiction(v TaxJurisdiction)`

SetTaxJurisdiction sets TaxJurisdiction field to given value.

### HasTaxJurisdiction

`func (o *Taxes) HasTaxJurisdiction() bool`

HasTaxJurisdiction returns a boolean if a field has been set.

### GetTaxPercentage

`func (o *Taxes) GetTaxPercentage() string`

GetTaxPercentage returns the TaxPercentage field if non-nil, zero value otherwise.

### GetTaxPercentageOk

`func (o *Taxes) GetTaxPercentageOk() (*string, bool)`

GetTaxPercentageOk returns a tuple with the TaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPercentage

`func (o *Taxes) SetTaxPercentage(v string)`

SetTaxPercentage sets TaxPercentage field to given value.

### HasTaxPercentage

`func (o *Taxes) HasTaxPercentage() bool`

HasTaxPercentage returns a boolean if a field has been set.

### GetTaxType

`func (o *Taxes) GetTaxType() string`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *Taxes) GetTaxTypeOk() (*string, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *Taxes) SetTaxType(v string)`

SetTaxType sets TaxType field to given value.

### HasTaxType

`func (o *Taxes) HasTaxType() bool`

HasTaxType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


