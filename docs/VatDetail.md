# VatDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuingCountry** | Pointer to **string** | The two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html \&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; standard of the country issuing the seller&#39;s VAT (value added tax) ID. VAT is a tax added by some European countries. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/ba:CountryCodeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**VatId** | Pointer to **string** | The seller&#39;s VAT (value added tax) ID. VAT is a tax added by some European countries. | [optional] 

## Methods

### NewVatDetail

`func NewVatDetail() *VatDetail`

NewVatDetail instantiates a new VatDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVatDetailWithDefaults

`func NewVatDetailWithDefaults() *VatDetail`

NewVatDetailWithDefaults instantiates a new VatDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuingCountry

`func (o *VatDetail) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *VatDetail) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *VatDetail) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *VatDetail) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### GetVatId

`func (o *VatDetail) GetVatId() string`

GetVatId returns the VatId field if non-nil, zero value otherwise.

### GetVatIdOk

`func (o *VatDetail) GetVatIdOk() (*string, bool)`

GetVatIdOk returns a tuple with the VatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatId

`func (o *VatDetail) SetVatId(v string)`

SetVatId sets VatId field to given value.

### HasVatId

`func (o *VatDetail) HasVatId() bool`

HasVatId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


