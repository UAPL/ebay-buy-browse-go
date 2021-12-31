# Aspect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalizedName** | Pointer to **string** | The text representing the name of the aspect for the name/value pair, such as Brand. | [optional] 
**LocalizedValues** | Pointer to **[]string** | The text representing the value of the aspect for the name/value pair, such as Apple. | [optional] 

## Methods

### NewAspect

`func NewAspect() *Aspect`

NewAspect instantiates a new Aspect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAspectWithDefaults

`func NewAspectWithDefaults() *Aspect`

NewAspectWithDefaults instantiates a new Aspect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalizedName

`func (o *Aspect) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *Aspect) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *Aspect) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *Aspect) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### GetLocalizedValues

`func (o *Aspect) GetLocalizedValues() []string`

GetLocalizedValues returns the LocalizedValues field if non-nil, zero value otherwise.

### GetLocalizedValuesOk

`func (o *Aspect) GetLocalizedValuesOk() (*[]string, bool)`

GetLocalizedValuesOk returns a tuple with the LocalizedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedValues

`func (o *Aspect) SetLocalizedValues(v []string)`

SetLocalizedValues sets LocalizedValues field to given value.

### HasLocalizedValues

`func (o *Aspect) HasLocalizedValues() bool`

HasLocalizedValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


