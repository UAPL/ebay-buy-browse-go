# HazardousMaterialsLabels

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInformation** | Pointer to **string** | Additional information about the hazardous materials labels. | [optional] 
**Pictograms** | Pointer to [**[]HazardPictogram**](HazardPictogram.md) | An array of hazard pictograms that apply to the item. | [optional] 
**SignalWord** | Pointer to **string** | The signal word for the hazardous materials label (such as Danger or Warning). | [optional] 
**SignalWordId** | Pointer to **string** | The ID of the signal word for the hazardous materials label. | [optional] 
**Statements** | Pointer to [**[]HazardStatement**](HazardStatement.md) | An array of hazard statements for the item. | [optional] 

## Methods

### NewHazardousMaterialsLabels

`func NewHazardousMaterialsLabels() *HazardousMaterialsLabels`

NewHazardousMaterialsLabels instantiates a new HazardousMaterialsLabels object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHazardousMaterialsLabelsWithDefaults

`func NewHazardousMaterialsLabelsWithDefaults() *HazardousMaterialsLabels`

NewHazardousMaterialsLabelsWithDefaults instantiates a new HazardousMaterialsLabels object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInformation

`func (o *HazardousMaterialsLabels) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *HazardousMaterialsLabels) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *HazardousMaterialsLabels) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *HazardousMaterialsLabels) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetPictograms

`func (o *HazardousMaterialsLabels) GetPictograms() []HazardPictogram`

GetPictograms returns the Pictograms field if non-nil, zero value otherwise.

### GetPictogramsOk

`func (o *HazardousMaterialsLabels) GetPictogramsOk() (*[]HazardPictogram, bool)`

GetPictogramsOk returns a tuple with the Pictograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictograms

`func (o *HazardousMaterialsLabels) SetPictograms(v []HazardPictogram)`

SetPictograms sets Pictograms field to given value.

### HasPictograms

`func (o *HazardousMaterialsLabels) HasPictograms() bool`

HasPictograms returns a boolean if a field has been set.

### GetSignalWord

`func (o *HazardousMaterialsLabels) GetSignalWord() string`

GetSignalWord returns the SignalWord field if non-nil, zero value otherwise.

### GetSignalWordOk

`func (o *HazardousMaterialsLabels) GetSignalWordOk() (*string, bool)`

GetSignalWordOk returns a tuple with the SignalWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalWord

`func (o *HazardousMaterialsLabels) SetSignalWord(v string)`

SetSignalWord sets SignalWord field to given value.

### HasSignalWord

`func (o *HazardousMaterialsLabels) HasSignalWord() bool`

HasSignalWord returns a boolean if a field has been set.

### GetSignalWordId

`func (o *HazardousMaterialsLabels) GetSignalWordId() string`

GetSignalWordId returns the SignalWordId field if non-nil, zero value otherwise.

### GetSignalWordIdOk

`func (o *HazardousMaterialsLabels) GetSignalWordIdOk() (*string, bool)`

GetSignalWordIdOk returns a tuple with the SignalWordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalWordId

`func (o *HazardousMaterialsLabels) SetSignalWordId(v string)`

SetSignalWordId sets SignalWordId field to given value.

### HasSignalWordId

`func (o *HazardousMaterialsLabels) HasSignalWordId() bool`

HasSignalWordId returns a boolean if a field has been set.

### GetStatements

`func (o *HazardousMaterialsLabels) GetStatements() []HazardStatement`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *HazardousMaterialsLabels) GetStatementsOk() (*[]HazardStatement, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *HazardousMaterialsLabels) SetStatements(v []HazardStatement)`

SetStatements sets Statements field to given value.

### HasStatements

`func (o *HazardousMaterialsLabels) HasStatements() bool`

HasStatements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


