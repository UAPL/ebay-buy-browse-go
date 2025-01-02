# LegalAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** | The first line of the street address. | [optional] 
**AddressLine2** | Pointer to **string** | The second line of the street address. This field is not always used, but can be used for &#39;Suite Number&#39; or &#39;Apt Number&#39;. | [optional] 
**City** | Pointer to **string** | The city of the address. | [optional] 
**Country** | Pointer to **string** | The two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html \&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; standard code for the country of the address. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/ba:CountryCodeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**CountryName** | Pointer to **string** | The name of the country of the address. | [optional] 
**County** | Pointer to **string** | The name of the county of the address. | [optional] 
**PostalCode** | Pointer to **string** | The postal code of the address. | [optional] 
**StateOrProvince** | Pointer to **string** | The state or province of the address. | [optional] 

## Methods

### NewLegalAddress

`func NewLegalAddress() *LegalAddress`

NewLegalAddress instantiates a new LegalAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalAddressWithDefaults

`func NewLegalAddressWithDefaults() *LegalAddress`

NewLegalAddressWithDefaults instantiates a new LegalAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *LegalAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *LegalAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *LegalAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *LegalAddress) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *LegalAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *LegalAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *LegalAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *LegalAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *LegalAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *LegalAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *LegalAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *LegalAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *LegalAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *LegalAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *LegalAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *LegalAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryName

`func (o *LegalAddress) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *LegalAddress) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *LegalAddress) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *LegalAddress) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetCounty

`func (o *LegalAddress) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *LegalAddress) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *LegalAddress) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *LegalAddress) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetPostalCode

`func (o *LegalAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *LegalAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *LegalAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *LegalAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStateOrProvince

`func (o *LegalAddress) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *LegalAddress) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *LegalAddress) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.

### HasStateOrProvince

`func (o *LegalAddress) HasStateOrProvince() bool`

HasStateOrProvince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


