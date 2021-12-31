# AspectGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aspects** | Pointer to [**[]Aspect**](Aspect.md) | An array of the name/value pairs for the aspects of the product. For example: BRAND/Apple | [optional] 
**LocalizedGroupName** | Pointer to **string** | The name of a group of aspects. &lt;br /&gt;&lt;br /&gt;In the following example, &lt;b&gt; Product Identifiers&lt;/b&gt; and &lt;b&gt; Process&lt;/b&gt; are product aspect group names. Under the group name are the product aspect name/value pairs. &lt;p&gt;&lt;b&gt; Product Identifiers&lt;/b&gt; &lt;br /&gt;&amp;nbsp;&amp;nbsp;&amp;nbsp;Brand/Apple &lt;br /&gt;&amp;nbsp;&amp;nbsp;&amp;nbsp;Product Family/iMac&lt;/p&gt; &lt;p&gt;&lt;b&gt; Processor&lt;/b&gt;&lt;br /&gt;&amp;nbsp;&amp;nbsp;&amp;nbsp;Processor Type/Intel &lt;br /&gt;&amp;nbsp;&amp;nbsp;&amp;nbsp;Processor Speed/3.10&lt;/p&gt;  | [optional] 

## Methods

### NewAspectGroup

`func NewAspectGroup() *AspectGroup`

NewAspectGroup instantiates a new AspectGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAspectGroupWithDefaults

`func NewAspectGroupWithDefaults() *AspectGroup`

NewAspectGroupWithDefaults instantiates a new AspectGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspects

`func (o *AspectGroup) GetAspects() []Aspect`

GetAspects returns the Aspects field if non-nil, zero value otherwise.

### GetAspectsOk

`func (o *AspectGroup) GetAspectsOk() (*[]Aspect, bool)`

GetAspectsOk returns a tuple with the Aspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspects

`func (o *AspectGroup) SetAspects(v []Aspect)`

SetAspects sets Aspects field to given value.

### HasAspects

`func (o *AspectGroup) HasAspects() bool`

HasAspects returns a boolean if a field has been set.

### GetLocalizedGroupName

`func (o *AspectGroup) GetLocalizedGroupName() string`

GetLocalizedGroupName returns the LocalizedGroupName field if non-nil, zero value otherwise.

### GetLocalizedGroupNameOk

`func (o *AspectGroup) GetLocalizedGroupNameOk() (*string, bool)`

GetLocalizedGroupNameOk returns a tuple with the LocalizedGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedGroupName

`func (o *AspectGroup) SetLocalizedGroupName(v string)`

SetLocalizedGroupName sets LocalizedGroupName field to given value.

### HasLocalizedGroupName

`func (o *AspectGroup) HasLocalizedGroupName() bool`

HasLocalizedGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


