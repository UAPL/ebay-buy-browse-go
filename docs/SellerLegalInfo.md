# SellerLegalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The seller&#39;s business email address. | [optional] 
**Fax** | Pointer to **string** | The seller&#39; business fax number. | [optional] 
**Imprint** | Pointer to **string** | This is a free-form string created by the seller. This is information often found on business cards, such as address. This is information used by some countries. | [optional] 
**LegalContactFirstName** | Pointer to **string** | The seller&#39;s first name. | [optional] 
**LegalContactLastName** | Pointer to **string** | The seller&#39;s last name. | [optional] 
**Name** | Pointer to **string** | The name of the seller&#39;s business. | [optional] 
**Phone** | Pointer to **string** | The seller&#39;s business phone number. | [optional] 
**RegistrationNumber** | Pointer to **string** | The seller&#39;s registration number. This is information used by some countries. | [optional] 
**SellerProvidedLegalAddress** | Pointer to [**LegalAddress**](LegalAddress.md) |  | [optional] 
**TermsOfService** | Pointer to **string** | This is a free-form string created by the seller. This is the seller&#39;s terms or condition, which is in addition to the seller&#39;s return policies. | [optional] 
**VatDetails** | Pointer to [**[]VatDetail**](VatDetail.md) | An array of the seller&#39;s VAT (value added tax) IDs and the issuing country. VAT is a tax added by some European countries. | [optional] 
**EconomicOperator** | Pointer to [**EconomicOperator**](EconomicOperator.md) |  | [optional] 
**WeeeNumber** | Pointer to **string** | The Waste Electrical and Electronic Equipment (WEEE) registration number required for any seller to place electrical and electronic equipment on the market in Germany. This manufacturer number is assigned to the first distributors of electrical and electronic equipment and comprises a country code and an 8-digit sequence of digits (e.g. “WEEE Reg. No. DE 12345678”). | [optional] 

## Methods

### NewSellerLegalInfo

`func NewSellerLegalInfo() *SellerLegalInfo`

NewSellerLegalInfo instantiates a new SellerLegalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellerLegalInfoWithDefaults

`func NewSellerLegalInfoWithDefaults() *SellerLegalInfo`

NewSellerLegalInfoWithDefaults instantiates a new SellerLegalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SellerLegalInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SellerLegalInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SellerLegalInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SellerLegalInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFax

`func (o *SellerLegalInfo) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *SellerLegalInfo) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *SellerLegalInfo) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *SellerLegalInfo) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetImprint

`func (o *SellerLegalInfo) GetImprint() string`

GetImprint returns the Imprint field if non-nil, zero value otherwise.

### GetImprintOk

`func (o *SellerLegalInfo) GetImprintOk() (*string, bool)`

GetImprintOk returns a tuple with the Imprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImprint

`func (o *SellerLegalInfo) SetImprint(v string)`

SetImprint sets Imprint field to given value.

### HasImprint

`func (o *SellerLegalInfo) HasImprint() bool`

HasImprint returns a boolean if a field has been set.

### GetLegalContactFirstName

`func (o *SellerLegalInfo) GetLegalContactFirstName() string`

GetLegalContactFirstName returns the LegalContactFirstName field if non-nil, zero value otherwise.

### GetLegalContactFirstNameOk

`func (o *SellerLegalInfo) GetLegalContactFirstNameOk() (*string, bool)`

GetLegalContactFirstNameOk returns a tuple with the LegalContactFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalContactFirstName

`func (o *SellerLegalInfo) SetLegalContactFirstName(v string)`

SetLegalContactFirstName sets LegalContactFirstName field to given value.

### HasLegalContactFirstName

`func (o *SellerLegalInfo) HasLegalContactFirstName() bool`

HasLegalContactFirstName returns a boolean if a field has been set.

### GetLegalContactLastName

`func (o *SellerLegalInfo) GetLegalContactLastName() string`

GetLegalContactLastName returns the LegalContactLastName field if non-nil, zero value otherwise.

### GetLegalContactLastNameOk

`func (o *SellerLegalInfo) GetLegalContactLastNameOk() (*string, bool)`

GetLegalContactLastNameOk returns a tuple with the LegalContactLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalContactLastName

`func (o *SellerLegalInfo) SetLegalContactLastName(v string)`

SetLegalContactLastName sets LegalContactLastName field to given value.

### HasLegalContactLastName

`func (o *SellerLegalInfo) HasLegalContactLastName() bool`

HasLegalContactLastName returns a boolean if a field has been set.

### GetName

`func (o *SellerLegalInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SellerLegalInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SellerLegalInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SellerLegalInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *SellerLegalInfo) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SellerLegalInfo) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SellerLegalInfo) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SellerLegalInfo) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetRegistrationNumber

`func (o *SellerLegalInfo) GetRegistrationNumber() string`

GetRegistrationNumber returns the RegistrationNumber field if non-nil, zero value otherwise.

### GetRegistrationNumberOk

`func (o *SellerLegalInfo) GetRegistrationNumberOk() (*string, bool)`

GetRegistrationNumberOk returns a tuple with the RegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNumber

`func (o *SellerLegalInfo) SetRegistrationNumber(v string)`

SetRegistrationNumber sets RegistrationNumber field to given value.

### HasRegistrationNumber

`func (o *SellerLegalInfo) HasRegistrationNumber() bool`

HasRegistrationNumber returns a boolean if a field has been set.

### GetSellerProvidedLegalAddress

`func (o *SellerLegalInfo) GetSellerProvidedLegalAddress() LegalAddress`

GetSellerProvidedLegalAddress returns the SellerProvidedLegalAddress field if non-nil, zero value otherwise.

### GetSellerProvidedLegalAddressOk

`func (o *SellerLegalInfo) GetSellerProvidedLegalAddressOk() (*LegalAddress, bool)`

GetSellerProvidedLegalAddressOk returns a tuple with the SellerProvidedLegalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerProvidedLegalAddress

`func (o *SellerLegalInfo) SetSellerProvidedLegalAddress(v LegalAddress)`

SetSellerProvidedLegalAddress sets SellerProvidedLegalAddress field to given value.

### HasSellerProvidedLegalAddress

`func (o *SellerLegalInfo) HasSellerProvidedLegalAddress() bool`

HasSellerProvidedLegalAddress returns a boolean if a field has been set.

### GetTermsOfService

`func (o *SellerLegalInfo) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *SellerLegalInfo) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *SellerLegalInfo) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *SellerLegalInfo) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetVatDetails

`func (o *SellerLegalInfo) GetVatDetails() []VatDetail`

GetVatDetails returns the VatDetails field if non-nil, zero value otherwise.

### GetVatDetailsOk

`func (o *SellerLegalInfo) GetVatDetailsOk() (*[]VatDetail, bool)`

GetVatDetailsOk returns a tuple with the VatDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatDetails

`func (o *SellerLegalInfo) SetVatDetails(v []VatDetail)`

SetVatDetails sets VatDetails field to given value.

### HasVatDetails

`func (o *SellerLegalInfo) HasVatDetails() bool`

HasVatDetails returns a boolean if a field has been set.

### GetEconomicOperator

`func (o *SellerLegalInfo) GetEconomicOperator() EconomicOperator`

GetEconomicOperator returns the EconomicOperator field if non-nil, zero value otherwise.

### GetEconomicOperatorOk

`func (o *SellerLegalInfo) GetEconomicOperatorOk() (*EconomicOperator, bool)`

GetEconomicOperatorOk returns a tuple with the EconomicOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEconomicOperator

`func (o *SellerLegalInfo) SetEconomicOperator(v EconomicOperator)`

SetEconomicOperator sets EconomicOperator field to given value.

### HasEconomicOperator

`func (o *SellerLegalInfo) HasEconomicOperator() bool`

HasEconomicOperator returns a boolean if a field has been set.

### GetWeeeNumber

`func (o *SellerLegalInfo) GetWeeeNumber() string`

GetWeeeNumber returns the WeeeNumber field if non-nil, zero value otherwise.

### GetWeeeNumberOk

`func (o *SellerLegalInfo) GetWeeeNumberOk() (*string, bool)`

GetWeeeNumberOk returns a tuple with the WeeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeeNumber

`func (o *SellerLegalInfo) SetWeeeNumber(v string)`

SetWeeeNumber sets WeeeNumber field to given value.

### HasWeeeNumber

`func (o *SellerLegalInfo) HasWeeeNumber() bool`

HasWeeeNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


