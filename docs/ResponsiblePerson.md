# ResponsiblePerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** | The first line of the Responsible Person&#39;s street address. | [optional] 
**AddressLine2** | Pointer to **string** | The second line of the Responsible Person&#39;s address. This field is not always used, but can be used for secondary address information such as &#39;Suite Number&#39; or &#39;Apt Number&#39;. | [optional] 
**City** | Pointer to **string** | The city of the Responsible Person&#39;s street address. | [optional] 
**CompanyName** | Pointer to **string** | The name of the the Responsible Person or entity. | [optional] 
**Country** | Pointer to **string** | The two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html \&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; standard of the country of the address. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/ba:CountryCodeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**CountryName** | Pointer to **string** | The country name of the Responsible Person&#39;s street address. | [optional] 
**County** | Pointer to **string** | The county of the Responsible Person&#39;s street address. | [optional] 
**Email** | Pointer to **string** | The email of the Responsible Person&#39;s street address. | [optional] 
**Phone** | Pointer to **string** | The phone number of the Responsible Person&#39;s street address. | [optional] 
**PostalCode** | Pointer to **string** | The postal code of the Responsible Person&#39;s street address. | [optional] 
**StateOrProvince** | Pointer to **string** | The state or province of the Responsible Person&#39;s street address. | [optional] 
**Types** | Pointer to **[]string** | The type(s) associated with the Responsible Person or entity. | [optional] 

## Methods

### NewResponsiblePerson

`func NewResponsiblePerson() *ResponsiblePerson`

NewResponsiblePerson instantiates a new ResponsiblePerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsiblePersonWithDefaults

`func NewResponsiblePersonWithDefaults() *ResponsiblePerson`

NewResponsiblePersonWithDefaults instantiates a new ResponsiblePerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *ResponsiblePerson) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *ResponsiblePerson) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *ResponsiblePerson) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *ResponsiblePerson) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *ResponsiblePerson) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ResponsiblePerson) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ResponsiblePerson) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ResponsiblePerson) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *ResponsiblePerson) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ResponsiblePerson) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ResponsiblePerson) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ResponsiblePerson) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *ResponsiblePerson) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *ResponsiblePerson) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *ResponsiblePerson) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *ResponsiblePerson) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *ResponsiblePerson) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ResponsiblePerson) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ResponsiblePerson) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ResponsiblePerson) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryName

`func (o *ResponsiblePerson) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *ResponsiblePerson) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *ResponsiblePerson) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *ResponsiblePerson) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetCounty

`func (o *ResponsiblePerson) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *ResponsiblePerson) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *ResponsiblePerson) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *ResponsiblePerson) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetEmail

`func (o *ResponsiblePerson) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ResponsiblePerson) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ResponsiblePerson) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ResponsiblePerson) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *ResponsiblePerson) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ResponsiblePerson) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ResponsiblePerson) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ResponsiblePerson) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *ResponsiblePerson) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ResponsiblePerson) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ResponsiblePerson) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ResponsiblePerson) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetStateOrProvince

`func (o *ResponsiblePerson) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *ResponsiblePerson) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *ResponsiblePerson) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.

### HasStateOrProvince

`func (o *ResponsiblePerson) HasStateOrProvince() bool`

HasStateOrProvince returns a boolean if a field has been set.

### GetTypes

`func (o *ResponsiblePerson) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ResponsiblePerson) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ResponsiblePerson) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *ResponsiblePerson) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


