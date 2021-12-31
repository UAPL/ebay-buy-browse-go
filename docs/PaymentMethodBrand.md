# PaymentMethodBrand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodBrandType** | Pointer to **string** | The payment method brand, such as Visa or PayPal. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:PaymentMethodBrandEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**LogoImage** | Pointer to [**Image**](Image.md) |  | [optional] 

## Methods

### NewPaymentMethodBrand

`func NewPaymentMethodBrand() *PaymentMethodBrand`

NewPaymentMethodBrand instantiates a new PaymentMethodBrand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodBrandWithDefaults

`func NewPaymentMethodBrandWithDefaults() *PaymentMethodBrand`

NewPaymentMethodBrandWithDefaults instantiates a new PaymentMethodBrand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodBrandType

`func (o *PaymentMethodBrand) GetPaymentMethodBrandType() string`

GetPaymentMethodBrandType returns the PaymentMethodBrandType field if non-nil, zero value otherwise.

### GetPaymentMethodBrandTypeOk

`func (o *PaymentMethodBrand) GetPaymentMethodBrandTypeOk() (*string, bool)`

GetPaymentMethodBrandTypeOk returns a tuple with the PaymentMethodBrandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodBrandType

`func (o *PaymentMethodBrand) SetPaymentMethodBrandType(v string)`

SetPaymentMethodBrandType sets PaymentMethodBrandType field to given value.

### HasPaymentMethodBrandType

`func (o *PaymentMethodBrand) HasPaymentMethodBrandType() bool`

HasPaymentMethodBrandType returns a boolean if a field has been set.

### GetLogoImage

`func (o *PaymentMethodBrand) GetLogoImage() Image`

GetLogoImage returns the LogoImage field if non-nil, zero value otherwise.

### GetLogoImageOk

`func (o *PaymentMethodBrand) GetLogoImageOk() (*Image, bool)`

GetLogoImageOk returns a tuple with the LogoImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoImage

`func (o *PaymentMethodBrand) SetLogoImage(v Image)`

SetLogoImage sets LogoImage field to given value.

### HasLogoImage

`func (o *PaymentMethodBrand) HasLogoImage() bool`

HasLogoImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


