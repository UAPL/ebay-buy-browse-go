# CompanyAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** | The first line of the product manufacturer&#39;s street address. | [optional] 
**AddressLine2** | Pointer to **string** | The second line of the product manufacturer&#39;s street address. This field is not always used, but can be used for secondary address information such as &#39;Suite Number&#39; or &#39;Apt Number&#39;. | [optional] 
**City** | Pointer to **string** | The city of the product manufacturer&#39;s street address. | [optional] 
**CompanyName** | Pointer to **string** | The company name of the the product manufacturer. | [optional] 
**Country** | Pointer to **string** | The two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html \&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; standard code for the country of the address. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/ba:CountryCodeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**CountryName** | Pointer to **string** | The country name of the product manufacturer&#39;s street address. | [optional] 
**County** | Pointer to **string** | The county of the product manufacturer&#39;s street address. | [optional] 
**Email** | Pointer to **string** | The product manufacturer&#39;s business email address. | [optional] 
**Phone** | Pointer to **string** | The product manufacturer&#39;s business phone number. | [optional] 
**PostalCode** | Pointer to **string** | The postal code of the product manufacturer&#39;s street address. | [optional] 
**StateOrProvince** | Pointer to **string** | The state or province of the product manufacturer&#39;s street address. | [optional] 

## Methods

### NewCompanyAddress

`func NewCompanyAddress() *CompanyAddress`

NewCompanyAddress instantiates a new CompanyAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyAddressWithDefaults

`func NewCompanyAddressWithDefaults() *CompanyAddress`

NewCompanyAddressWithDefaults instantiates a new CompanyAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *CompanyAddress) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *CompanyAddress) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *CompanyAddress) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *CompanyAddress) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *CompanyAddress) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *CompanyAddress) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *CompanyAddress) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *CompanyAddress) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *CompanyAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CompanyAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CompanyAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CompanyAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *CompanyAddress) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyAddress) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyAddress) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyAddress) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *CompanyAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CompanyAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CompanyAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CompanyAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryName

`func (o *CompanyAddress) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *CompanyAddress) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *CompanyAddress) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *CompanyAddress) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetCounty

`func (o *CompanyAddress) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *CompanyAddress) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *CompanyAddress) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *CompanyAddress) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetEmail

`func (o *CompanyAddress) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CompanyAddress) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CompanyAddress) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CompanyAddress) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *CompanyAddress) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CompanyAddress) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CompanyAddress) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CompanyAddress) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *CompanyAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CompanyAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CompanyAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CompanyAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStateOrProvince

`func (o *CompanyAddress) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *CompanyAddress) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *CompanyAddress) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.

### HasStateOrProvince

`func (o *CompanyAddress) HasStateOrProvince() bool`

HasStateOrProvince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


