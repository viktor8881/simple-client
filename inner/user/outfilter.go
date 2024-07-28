package user

import (
	"encoding/json"
	genclient "simple-client/generated/http/client"
	generated "simple-client/generated/http/server"
)

func convertListUserToResponse(dtoOut *genclient.ProxyResponse) *generated.ListUserResponse {
	var dest generated.ListUserResponse
	//dest.Users = make([]*generated.User, len(dtoOut.Users))
	//
	//for i, user := range dtoOut.Users {
	//	dest.Users[i] = &generated.User{
	//		Id:   user.Id,
	//		Name: user.Name,
	//	}
	//}

	return &dest
}

func сonvertProxyResponseToGetUserResponse(proxyResp generated.ProxyResponse) (*generated.GetUserResponse, error) {
	var getUserResp *generated.GetUserResponse

	jsonData, err := json.Marshal(proxyResp)
	if err != nil {
		return getUserResp, err
	}

	err = json.Unmarshal(jsonData, &getUserResp)
	if err != nil {
		return getUserResp, err
	}

	return getUserResp, nil
}
