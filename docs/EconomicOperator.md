# EconomicOperator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | Pointer to **string** | The company name of the registered Economic Operator. | [optional] 
**AddressLine1** | Pointer to **string** | The first line of the registered Economic Operator&#39;s street address. | [optional] 
**AddressLine2** | Pointer to **string** | The second line, if any, of the registered Economic Operator&#39;s street address. This field is not always used, but can be used for &#39;Suite Number&#39; or &#39;Apt Number&#39;. | [optional] 
**City** | Pointer to **string** | The city of the registered Economic Operator&#39;s street address. | [optional] 
**StateOrProvince** | Pointer to **string** | The state or province of the registered Economic Operator&#39;s street address. | [optional] 
**PostalCode** | Pointer to **string** | The postal code of the registered Economic Operator&#39;s street address. | [optional] 
**Country** | Pointer to **string** | The two-letter &lt;a href&#x3D;\&quot;https://www.iso.org/iso-3166-country-codes.html \&quot; target&#x3D;\&quot;_blank\&quot;&gt;ISO 3166&lt;/a&gt; standard abbreviation of the country of the registered Economic Operator&#39;s address. | [optional] 
**Phone** | Pointer to **string** | The registered Economic Operator&#39;s business phone number. | [optional] 
**Email** | Pointer to **string** | The registered Economic Operator&#39;s business email address. | [optional] 

## Methods

### NewEconomicOperator

`func NewEconomicOperator() *EconomicOperator`

NewEconomicOperator instantiates a new EconomicOperator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEconomicOperatorWithDefaults

`func NewEconomicOperatorWithDefaults() *EconomicOperator`

NewEconomicOperatorWithDefaults instantiates a new EconomicOperator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyName

`func (o *EconomicOperator) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *EconomicOperator) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *EconomicOperator) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *EconomicOperator) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetAddressLine1

`func (o *EconomicOperator) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *EconomicOperator) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *EconomicOperator) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *EconomicOperator) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *EconomicOperator) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *EconomicOperator) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *EconomicOperator) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *EconomicOperator) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *EconomicOperator) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *EconomicOperator) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *EconomicOperator) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *EconomicOperator) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetStateOrProvince

`func (o *EconomicOperator) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *EconomicOperator) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *EconomicOperator) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.

### HasStateOrProvince

`func (o *EconomicOperator) HasStateOrProvince() bool`

HasStateOrProvince returns a boolean if a field has been set.

### GetPostalCode

`func (o *EconomicOperator) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *EconomicOperator) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *EconomicOperator) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *EconomicOperator) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *EconomicOperator) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *EconomicOperator) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *EconomicOperator) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *EconomicOperator) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPhone

`func (o *EconomicOperator) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *EconomicOperator) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *EconomicOperator) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *EconomicOperator) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *EconomicOperator) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EconomicOperator) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EconomicOperator) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EconomicOperator) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


