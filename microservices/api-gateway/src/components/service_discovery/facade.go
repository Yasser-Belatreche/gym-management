package service_discovery

import (
	"encoding/json"
	"gym-management-api-gateway/src/lib/primitives/application_specific"
	"net/http"
)

type facade struct {
	serviceDiscoveryUrl string
	apiSecret           string
}

func (f *facade) GetAuthService() (*Service, error) {
	url, err := f.GetServiceUrl("auth")
	if err != nil {
		return nil, err
	}

	return &Service{
		Url:       url,
		ApiSecret: f.apiSecret,
	}, nil
}

func (f *facade) GetMembershipsService() (*Service, error) {
	url, err := f.GetServiceUrl("memberships")
	if err != nil {
		return nil, err
	}

	return &Service{
		Url:       url,
		ApiSecret: f.apiSecret,
	}, nil
}

func (f *facade) GetGymsService() (*Service, error) {
	url, err := f.GetServiceUrl("gyms")
	if err != nil {
		return nil, err
	}

	return &Service{
		Url:       url,
		ApiSecret: f.apiSecret,
	}, nil
}

func (f *facade) GetHealth() (map[string]interface{}, error) {
	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, f.serviceDiscoveryUrl+"/api/v1/health", nil)
	if err != nil {
		return nil, application_specific.NewUnknownException("ERROR_CREATING_REQUEST", err.Error(), nil)
	}

	request.Header.Add("X-Api-Secret", f.apiSecret)

	resp, err := client.Do(request)
	if err != nil {
		return nil, application_specific.NewUnknownException("ERROR_SENDING_HTTP_REQUEST", err.Error(), nil)
	}

	var response map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, application_specific.NewUnknownException("ERROR_DECODING_RESPONSE", err.Error(), nil)
	}

	return response, nil
}

func (f *facade) GetServiceUrl(service string) (string, error) {
	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, f.serviceDiscoveryUrl+"/api/v1/services/"+service, nil)
	if err != nil {
		return "", application_specific.NewUnknownException("ERROR_CREATING_REQUEST", err.Error(), nil)
	}

	request.Header.Add("X-Api-Secret", f.apiSecret)

	resp, err := client.Do(request)
	if err != nil {
		return "", application_specific.NewUnknownException("ERROR_SENDING_HTTP_REQUEST", err.Error(), nil)
	}

	if resp.StatusCode != http.StatusOK {
		var response errorResponse

		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return "", application_specific.NewUnknownException("ERROR_DECODING_RESPONSE", err.Error(), nil)
		}

		return "", application_specific.NewUnknownException("ERROR_SENDING_HTTP_REQUEST", response.Error, map[string]interface{}{
			"response": response,
		})
	}

	var response successResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", application_specific.NewUnknownException("ERROR_DECODING_RESPONSE", err.Error(), nil)
	}

	return response.Url, nil
}

type successResponse struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type errorResponse struct {
	Error string `json:"error"`
}
