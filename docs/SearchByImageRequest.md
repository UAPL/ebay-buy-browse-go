# SearchByImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | Pointer to **string** | The Base64 string of the image.&lt;br&gt;&lt;br&gt;To get the Base64 image string, you can use sites such as &lt;a href&#x3D;\&quot;https://codebeautify.org/image-to-base64-converter \&quot; target&#x3D;\&quot;_blank\&quot;&gt;https://codebeautify.org/image-to-base64-converter&lt;/a&gt;. | [optional] 

## Methods

### NewSearchByImageRequest

`func NewSearchByImageRequest() *SearchByImageRequest`

NewSearchByImageRequest instantiates a new SearchByImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchByImageRequestWithDefaults

`func NewSearchByImageRequestWithDefaults() *SearchByImageRequest`

NewSearchByImageRequestWithDefaults instantiates a new SearchByImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *SearchByImageRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *SearchByImageRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *SearchByImageRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *SearchByImageRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


