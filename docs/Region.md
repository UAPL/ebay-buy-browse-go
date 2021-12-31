# Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionName** | Pointer to **string** | A localized text string that indicates the name of the region. Taxes are generally charged at the state/province level or at the country level in the case of VAT tax.  | [optional] 
**RegionType** | Pointer to **string** | An enumeration value that indicates the type of region for the tax jurisdiction. &lt;br&gt;&lt;br&gt;&lt;b&gt; Valid Values: &lt;/b&gt; &lt;ul&gt;&lt;li&gt;&lt;b&gt; STATE_OR_PROVINCE &lt;/b&gt; - Indicates the region is a state or province within a country, such as California or New York in the US, or Ontario or Alberta in Canada.&lt;/li&gt;&lt;li&gt;&lt;b&gt; COUNTRY &lt;/b&gt; - Indicates the region is a single country.&lt;/li&gt;&lt;/ul&gt;  Code so that your app gracefully handles any future changes to this list. For implementation help, refer to &lt;a href&#x3D;&#39;https://developer.ebay.com/api-docs/buy/browse/types/ba:RegionTypeEnum&#39;&gt;eBay API documentation&lt;/a&gt; | [optional] 

## Methods

### NewRegion

`func NewRegion() *Region`

NewRegion instantiates a new Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionWithDefaults

`func NewRegionWithDefaults() *Region`

NewRegionWithDefaults instantiates a new Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionName

`func (o *Region) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *Region) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *Region) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *Region) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### GetRegionType

`func (o *Region) GetRegionType() string`

GetRegionType returns the RegionType field if non-nil, zero value otherwise.

### GetRegionTypeOk

`func (o *Region) GetRegionTypeOk() (*string, bool)`

GetRegionTypeOk returns a tuple with the RegionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionType

`func (o *Region) SetRegionType(v string)`

SetRegionType sets RegionType field to given value.

### HasRegionType

`func (o *Region) HasRegionType() bool`

HasRegionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


