# CreateProjectDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Build** | Pointer to [**ProjectBuild**](ProjectBuild.md) |  | [optional] 
**EnvVars** | Pointer to **map[string]string** |  | [optional] 
**ExistingProjectConfigName** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**CreateProjectConfigSourceDTO**](CreateProjectConfigSourceDTO.md) |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateProjectDTO

`func NewCreateProjectDTO() *CreateProjectDTO`

NewCreateProjectDTO instantiates a new CreateProjectDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectDTOWithDefaults

`func NewCreateProjectDTOWithDefaults() *CreateProjectDTO`

NewCreateProjectDTOWithDefaults instantiates a new CreateProjectDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuild

`func (o *CreateProjectDTO) GetBuild() ProjectBuild`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *CreateProjectDTO) GetBuildOk() (*ProjectBuild, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *CreateProjectDTO) SetBuild(v ProjectBuild)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *CreateProjectDTO) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetEnvVars

`func (o *CreateProjectDTO) GetEnvVars() map[string]string`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *CreateProjectDTO) GetEnvVarsOk() (*map[string]string, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *CreateProjectDTO) SetEnvVars(v map[string]string)`

SetEnvVars sets EnvVars field to given value.

### HasEnvVars

`func (o *CreateProjectDTO) HasEnvVars() bool`

HasEnvVars returns a boolean if a field has been set.

### GetExistingProjectConfigName

`func (o *CreateProjectDTO) GetExistingProjectConfigName() string`

GetExistingProjectConfigName returns the ExistingProjectConfigName field if non-nil, zero value otherwise.

### GetExistingProjectConfigNameOk

`func (o *CreateProjectDTO) GetExistingProjectConfigNameOk() (*string, bool)`

GetExistingProjectConfigNameOk returns a tuple with the ExistingProjectConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingProjectConfigName

`func (o *CreateProjectDTO) SetExistingProjectConfigName(v string)`

SetExistingProjectConfigName sets ExistingProjectConfigName field to given value.

### HasExistingProjectConfigName

`func (o *CreateProjectDTO) HasExistingProjectConfigName() bool`

HasExistingProjectConfigName returns a boolean if a field has been set.

### GetImage

`func (o *CreateProjectDTO) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CreateProjectDTO) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CreateProjectDTO) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CreateProjectDTO) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *CreateProjectDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProjectDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProjectDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateProjectDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *CreateProjectDTO) GetSource() CreateProjectConfigSourceDTO`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateProjectDTO) GetSourceOk() (*CreateProjectConfigSourceDTO, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateProjectDTO) SetSource(v CreateProjectConfigSourceDTO)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreateProjectDTO) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUser

`func (o *CreateProjectDTO) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateProjectDTO) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateProjectDTO) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CreateProjectDTO) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


