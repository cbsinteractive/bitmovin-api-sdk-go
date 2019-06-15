package account
import (
    "github.com/cbsinteractive/bitmovin-api-sdk-go/common"
    
    "github.com/cbsinteractive/bitmovin-api-sdk-go/model"
    "github.com/cbsinteractive/bitmovin-api-sdk-go/pagination"
)

type AccountOrganizationsApi struct {
    apiClient *common.ApiClient
    Groups *AccountOrganizationsGroupsApi
}

func NewAccountOrganizationsApi(configs ...func(*common.ApiClient)) (*AccountOrganizationsApi, error) {
	apiClient, err := common.NewApiClient(configs...)
	if err != nil {
		return nil, err
	}

    api := &AccountOrganizationsApi{apiClient: apiClient}

    groupsApi, err := NewAccountOrganizationsGroupsApi(configs...)
    api.Groups = groupsApi

	if err != nil {
		return nil, err
	}

	return api, nil
}

func (api *AccountOrganizationsApi) Create(organization model.Organization) (*model.Organization, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *model.Organization
    err := api.apiClient.Post("/account/organizations", &organization, &responseModel, reqParams)
    return responseModel, err
}

func (api *AccountOrganizationsApi) Delete(organizationId string) (*model.BitmovinResponse, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
    }

    var responseModel *model.BitmovinResponse
    err := api.apiClient.Delete("/account/organizations/{organization_id}", &responseModel, reqParams)
    return responseModel, err
}

func (api *AccountOrganizationsApi) Get(organizationId string) (*model.Organization, error) {
    reqParams := func(params *common.RequestParams) {
        params.PathParams["organization_id"] = organizationId
    }

    var responseModel *model.Organization
    err := api.apiClient.Get("/account/organizations/{organization_id}", &responseModel, reqParams)
    return responseModel, err
}

func (api *AccountOrganizationsApi) List() (*pagination.OrganizationsListPagination, error) {
    reqParams := func(params *common.RequestParams) {
    }

    var responseModel *pagination.OrganizationsListPagination
    err := api.apiClient.Get("/account/organizations", &responseModel, reqParams)
    return responseModel, err
}

