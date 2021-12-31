# SearchPagedCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoCorrections** | Pointer to [**AutoCorrections**](AutoCorrections.md) |  | [optional] 
**Href** | Pointer to **string** | The URI of the current page of results. &lt;br /&gt;&lt;br /&gt;The following example of the &lt;b&gt; search&lt;/b&gt; method returns items 1 thru 5 from the list of items found. &lt;br /&gt;&lt;br /&gt;&lt;code&gt;https://api.ebay.com/buy/v1/item_summary/search?q&#x3D;shirt&amp;limit&#x3D;5&amp;offset&#x3D;0&lt;/code&gt;. | [optional] 
**ItemSummaries** | Pointer to [**[]ItemSummary**](ItemSummary.md) | An array of the items on this page. The items are sorted according to the sorting method specified in the request. | [optional] 
**Limit** | Pointer to **int32** | The value of the &lt;b&gt;limit&lt;/b&gt; parameter submitted in the request, which is the maximum number of items to return on a page, from the result set. A result set is the complete set of items returned by the method. | [optional] 
**Next** | Pointer to **string** | The URI for the next page of results. This value is returned if there is an additional page of results to return from the result set. &lt;br /&gt;&lt;br /&gt;The following example of the &lt;b&gt; search&lt;/b&gt; method returns items 5 thru 10 from the list of items found.&lt;br /&gt; &lt;br /&gt;&lt;code&gt;https://api.ebay.com/buy/v1/item_summary/search?query&#x3D;t-shirts&amp;limit&#x3D;5&amp;offset&#x3D;10 &lt;/code&gt; | [optional] 
**Offset** | Pointer to **int32** | This value indicates the &lt;b&gt;offset&lt;/b&gt; used for current page of items being returned. Assume the initial request used an &lt;b&gt;offset&lt;/b&gt; of &lt;code&gt;0&lt;/code&gt; and a &lt;b&gt;limit&lt;/b&gt; of &lt;code&gt;3&lt;/code&gt;. Then in the first page of results, this value would be &lt;code&gt;0&lt;/code&gt;, and items 1-3 are returned. For the second page, this value is &lt;code&gt;3&lt;/code&gt; and so on. | [optional] 
**Prev** | Pointer to **string** | The URI for the previous page of results. This is returned if there is a previous page of results from the result set. &lt;br /&gt;&lt;br /&gt;The following example of the &lt;b&gt; search&lt;/b&gt; method returns items 1 thru 5 from the list of items found, which would be the first set of items returned.&lt;br /&gt; &lt;br /&gt;&lt;code&gt;https://api.ebay.com/buy/v1/item_summary/search?query&#x3D;t-shirts&amp;limit&#x3D;5&amp;offset&#x3D;0&lt;/code&gt; | [optional] 
**Refinement** | Pointer to [**Refinement**](Refinement.md) |  | [optional] 
**Total** | Pointer to **int32** | The total number of items that match the input criteria. | [optional] 
**Warnings** | Pointer to [**[]Error**](Error.md) | The container with all the warnings for the request. | [optional] 

## Methods

### NewSearchPagedCollection

`func NewSearchPagedCollection() *SearchPagedCollection`

NewSearchPagedCollection instantiates a new SearchPagedCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPagedCollectionWithDefaults

`func NewSearchPagedCollectionWithDefaults() *SearchPagedCollection`

NewSearchPagedCollectionWithDefaults instantiates a new SearchPagedCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoCorrections

`func (o *SearchPagedCollection) GetAutoCorrections() AutoCorrections`

GetAutoCorrections returns the AutoCorrections field if non-nil, zero value otherwise.

### GetAutoCorrectionsOk

`func (o *SearchPagedCollection) GetAutoCorrectionsOk() (*AutoCorrections, bool)`

GetAutoCorrectionsOk returns a tuple with the AutoCorrections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCorrections

`func (o *SearchPagedCollection) SetAutoCorrections(v AutoCorrections)`

SetAutoCorrections sets AutoCorrections field to given value.

### HasAutoCorrections

`func (o *SearchPagedCollection) HasAutoCorrections() bool`

HasAutoCorrections returns a boolean if a field has been set.

### GetHref

`func (o *SearchPagedCollection) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SearchPagedCollection) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SearchPagedCollection) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SearchPagedCollection) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItemSummaries

`func (o *SearchPagedCollection) GetItemSummaries() []ItemSummary`

GetItemSummaries returns the ItemSummaries field if non-nil, zero value otherwise.

### GetItemSummariesOk

`func (o *SearchPagedCollection) GetItemSummariesOk() (*[]ItemSummary, bool)`

GetItemSummariesOk returns a tuple with the ItemSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSummaries

`func (o *SearchPagedCollection) SetItemSummaries(v []ItemSummary)`

SetItemSummaries sets ItemSummaries field to given value.

### HasItemSummaries

`func (o *SearchPagedCollection) HasItemSummaries() bool`

HasItemSummaries returns a boolean if a field has been set.

### GetLimit

`func (o *SearchPagedCollection) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchPagedCollection) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchPagedCollection) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchPagedCollection) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNext

`func (o *SearchPagedCollection) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SearchPagedCollection) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SearchPagedCollection) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *SearchPagedCollection) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetOffset

`func (o *SearchPagedCollection) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SearchPagedCollection) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SearchPagedCollection) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SearchPagedCollection) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPrev

`func (o *SearchPagedCollection) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *SearchPagedCollection) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *SearchPagedCollection) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *SearchPagedCollection) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetRefinement

`func (o *SearchPagedCollection) GetRefinement() Refinement`

GetRefinement returns the Refinement field if non-nil, zero value otherwise.

### GetRefinementOk

`func (o *SearchPagedCollection) GetRefinementOk() (*Refinement, bool)`

GetRefinementOk returns a tuple with the Refinement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefinement

`func (o *SearchPagedCollection) SetRefinement(v Refinement)`

SetRefinement sets Refinement field to given value.

### HasRefinement

`func (o *SearchPagedCollection) HasRefinement() bool`

HasRefinement returns a boolean if a field has been set.

### GetTotal

`func (o *SearchPagedCollection) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchPagedCollection) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchPagedCollection) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SearchPagedCollection) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWarnings

`func (o *SearchPagedCollection) GetWarnings() []Error`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SearchPagedCollection) GetWarningsOk() (*[]Error, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SearchPagedCollection) SetWarnings(v []Error)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SearchPagedCollection) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


