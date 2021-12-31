# SellerCustomPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The seller-defined description of the policy. | [optional] 
**Label** | Pointer to **string** | The seller-defined label for an individual custom policy. | [optional] 
**Type** | Pointer to **string** | The type of custom policy, such as PRODUCT_COMPLIANCE or TAKE_BACK. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/gct:SellerCustomPolicyTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewSellerCustomPolicy

`func NewSellerCustomPolicy() *SellerCustomPolicy`

NewSellerCustomPolicy instantiates a new SellerCustomPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellerCustomPolicyWithDefaults

`func NewSellerCustomPolicyWithDefaults() *SellerCustomPolicy`

NewSellerCustomPolicyWithDefaults instantiates a new SellerCustomPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SellerCustomPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SellerCustomPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SellerCustomPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SellerCustomPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *SellerCustomPolicy) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SellerCustomPolicy) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SellerCustomPolicy) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SellerCustomPolicy) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *SellerCustomPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SellerCustomPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SellerCustomPolicy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SellerCustomPolicy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


