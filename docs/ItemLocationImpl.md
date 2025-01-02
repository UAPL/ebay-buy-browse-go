# ItemLocationImpl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** | The first line of the street address. | [optional] 
**AddressLine2** | Pointer to **string** | The second line of the street address. This field may contain such values as an apartment or suite number. | [optional] 
**City** | Pointer to **string** | The city in which the item is located.&lt;br&gt;&lt;br&gt;&lt;b&gt;Restriction:&lt;/b&gt; This field is populated in the &lt;code&gt;search&lt;/code&gt; method response &lt;i&gt;only&lt;/i&gt; when &lt;code&gt;fieldgroups&lt;/code&gt; &#x3D; &lt;code&gt;EXTENDED&lt;/code&gt;. | [optional] 
**Country** | Pointer to **string** | The two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html \&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; standard code that indicates the country in which the item is located. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/ba:CountryCodeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**County** | Pointer to **string** | The county in which the item is located. | [optional] 
**PostalCode** | Pointer to **string** | The postal code (or zip code in US) where the item is located. Sellers set a postal code for items when they are listed. The postal code is used for calculating proximity searches. It is anonymized when returned in &lt;code&gt;itemLocation.postalCode&lt;/code&gt; via the API. | [optional] 
**StateOrProvince** | Pointer to **string** | The state or province in which the item is located. | [optional] 

## Methods

### NewItemLocationImpl

`func NewItemLocationImpl() *ItemLocationImpl`

NewItemLocationImpl instantiates a new ItemLocationImpl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemLocationImplWithDefaults

`func NewItemLocationImplWithDefaults() *ItemLocationImpl`

NewItemLocationImplWithDefaults instantiates a new ItemLocationImpl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *ItemLocationImpl) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *ItemLocationImpl) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *ItemLocationImpl) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *ItemLocationImpl) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *ItemLocationImpl) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ItemLocationImpl) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ItemLocationImpl) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ItemLocationImpl) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *ItemLocationImpl) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ItemLocationImpl) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ItemLocationImpl) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ItemLocationImpl) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *ItemLocationImpl) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ItemLocationImpl) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ItemLocationImpl) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ItemLocationImpl) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCounty

`func (o *ItemLocationImpl) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *ItemLocationImpl) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *ItemLocationImpl) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *ItemLocationImpl) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetPostalCode

`func (o *ItemLocationImpl) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ItemLocationImpl) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ItemLocationImpl) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ItemLocationImpl) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStateOrProvince

`func (o *ItemLocationImpl) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *ItemLocationImpl) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *ItemLocationImpl) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.

### HasStateOrProvince

`func (o *ItemLocationImpl) HasStateOrProvince() bool`

HasStateOrProvince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


