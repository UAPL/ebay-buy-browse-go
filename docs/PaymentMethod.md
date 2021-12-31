# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodType** | Pointer to **string** | The payment method type, such as credit card or cash. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:PaymentMethodTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 
**PaymentMethodBrands** | Pointer to [**[]PaymentMethodBrand**](PaymentMethodBrand.md) | The payment method brands, including the payment method brand type and logo image. | [optional] 
**PaymentInstructions** | Pointer to **[]string** | The payment instructions for the buyer, such as &lt;i&gt;cash in person&lt;/i&gt; or &lt;i&gt;contact seller&lt;/i&gt;. | [optional] 
**SellerInstructions** | Pointer to **[]string** | The seller instructions to the buyer, such as &lt;i&gt;accepts credit cards&lt;/i&gt; or &lt;i&gt;see description&lt;/i&gt;. | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod() *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodType

`func (o *PaymentMethod) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *PaymentMethod) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *PaymentMethod) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *PaymentMethod) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### GetPaymentMethodBrands

`func (o *PaymentMethod) GetPaymentMethodBrands() []PaymentMethodBrand`

GetPaymentMethodBrands returns the PaymentMethodBrands field if non-nil, zero value otherwise.

### GetPaymentMethodBrandsOk

`func (o *PaymentMethod) GetPaymentMethodBrandsOk() (*[]PaymentMethodBrand, bool)`

GetPaymentMethodBrandsOk returns a tuple with the PaymentMethodBrands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodBrands

`func (o *PaymentMethod) SetPaymentMethodBrands(v []PaymentMethodBrand)`

SetPaymentMethodBrands sets PaymentMethodBrands field to given value.

### HasPaymentMethodBrands

`func (o *PaymentMethod) HasPaymentMethodBrands() bool`

HasPaymentMethodBrands returns a boolean if a field has been set.

### GetPaymentInstructions

`func (o *PaymentMethod) GetPaymentInstructions() []string`

GetPaymentInstructions returns the PaymentInstructions field if non-nil, zero value otherwise.

### GetPaymentInstructionsOk

`func (o *PaymentMethod) GetPaymentInstructionsOk() (*[]string, bool)`

GetPaymentInstructionsOk returns a tuple with the PaymentInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstructions

`func (o *PaymentMethod) SetPaymentInstructions(v []string)`

SetPaymentInstructions sets PaymentInstructions field to given value.

### HasPaymentInstructions

`func (o *PaymentMethod) HasPaymentInstructions() bool`

HasPaymentInstructions returns a boolean if a field has been set.

### GetSellerInstructions

`func (o *PaymentMethod) GetSellerInstructions() []string`

GetSellerInstructions returns the SellerInstructions field if non-nil, zero value otherwise.

### GetSellerInstructionsOk

`func (o *PaymentMethod) GetSellerInstructionsOk() (*[]string, bool)`

GetSellerInstructionsOk returns a tuple with the SellerInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerInstructions

`func (o *PaymentMethod) SetSellerInstructions(v []string)`

SetSellerInstructions sets SellerInstructions field to given value.

### HasSellerInstructions

`func (o *PaymentMethod) HasSellerInstructions() bool`

HasSellerInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


