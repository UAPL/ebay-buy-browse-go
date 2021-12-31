# AdditionalProductIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductIdentity** | Pointer to [**[]ProductIdentity**](ProductIdentity.md) | An array of the product identifier/value pairs for the product associated with the item. This is returned if the seller has associated the eBay Product Identifier (ePID) with the item and the request has &lt;b&gt; fieldgroups&lt;/b&gt; set to &lt;code&gt;PRODUCT&lt;/code&gt;. &lt;br /&gt;&lt;br /&gt;The following table shows what is returned, based on the item information provided by the seller, when the &lt;b&gt; fieldgroups&lt;/b&gt; set to &lt;code&gt;PRODUCT&lt;/code&gt;.        &lt;br /&gt;&lt;br /&gt;&lt;div style&#x3D;\&quot;overflow-x:auto;\&quot;&gt; &lt;table border&#x3D;1&gt; &lt;tr&gt; &lt;th&gt; ePID Provided &lt;/th&gt;  &lt;th&gt; Product&amp;nbsp;ID(s) Provided&lt;/th&gt; &lt;th&gt; Response&lt;/th&gt; &lt;/tr&gt; &lt;tr&gt; &lt;td&gt; No &lt;/td&gt;  &lt;td&gt; No &lt;/td&gt; &lt;td&gt; The &lt;b&gt; AdditionalProductIdentity&lt;/b&gt; container is &lt;i&gt; not&lt;/i&gt; returned.&lt;/td&gt;&lt;/tr&gt;   &lt;tr&gt; &lt;td&gt; No &lt;/td&gt;  &lt;td&gt; Yes &lt;/td&gt;  &lt;td&gt; The &lt;b&gt; AdditionalProductIdentity&lt;/b&gt; container is &lt;i&gt; not&lt;/i&gt; returned but the product identifiers specified by the seller are returned in the &lt;b&gt; localizedAspects&lt;/b&gt; container. &lt;/td&gt;  &lt;/tr&gt;   &lt;tr&gt; &lt;td&gt; Yes &lt;/td&gt;  &lt;td&gt; No &lt;/td&gt; &lt;td&gt;  The &lt;b&gt; AdditionalProductIdentity&lt;/b&gt; container is returned listing the product identifiers of the product.&lt;/td&gt;&lt;/tr&gt;   &lt;tr&gt; &lt;td&gt; Yes &lt;/td&gt;  &lt;td&gt; Yes &lt;/td&gt; &lt;td&gt; The &lt;b&gt; AdditionalProductIdentity&lt;/b&gt; container is returned listing all the product identifiers of the product and the product identifiers specified by the seller are returned in the &lt;b&gt; localizedAspects&lt;/b&gt; container.&lt;/td&gt; &lt;/tr&gt;   &lt;/table&gt; &lt;/div&gt; | [optional] 

## Methods

### NewAdditionalProductIdentity

`func NewAdditionalProductIdentity() *AdditionalProductIdentity`

NewAdditionalProductIdentity instantiates a new AdditionalProductIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalProductIdentityWithDefaults

`func NewAdditionalProductIdentityWithDefaults() *AdditionalProductIdentity`

NewAdditionalProductIdentityWithDefaults instantiates a new AdditionalProductIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductIdentity

`func (o *AdditionalProductIdentity) GetProductIdentity() []ProductIdentity`

GetProductIdentity returns the ProductIdentity field if non-nil, zero value otherwise.

### GetProductIdentityOk

`func (o *AdditionalProductIdentity) GetProductIdentityOk() (*[]ProductIdentity, bool)`

GetProductIdentityOk returns a tuple with the ProductIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductIdentity

`func (o *AdditionalProductIdentity) SetProductIdentity(v []ProductIdentity)`

SetProductIdentity sets ProductIdentity field to given value.

### HasProductIdentity

`func (o *AdditionalProductIdentity) HasProductIdentity() bool`

HasProductIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


