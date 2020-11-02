package social

import (
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana/pkg/models"
	"github.com/patcharp/golib/one/identity"
	"golang.org/x/oauth2"
	"net/http"
)

type SocialOneIdOAuth struct {
	*SocialBase
	apiUrl  string
	teamIds []int
}

func (s *SocialOneIdOAuth) Type() int {
	return int(models.ONEID)
}

func (s *SocialOneIdOAuth) IsTeamMember(client *http.Client) bool {
	if len(s.teamIds) == 0 {
		return true
	}
	return false
}

func (s *SocialOneIdOAuth) UserInfo(client *http.Client, token *oauth2.Token) (*BasicUserInfo, error) {
	s.log.Info("Getting user info")
	response, err := HttpGet(client, s.apiUrl)
	if err != nil {
		return nil, fmt.Errorf("Error getting user info: %s", err)
	}
	var data identity.AccountProfile
	err = json.Unmarshal(response.Body, &data)
	if err != nil {
		return nil, fmt.Errorf("Error getting user info: %s", err)
	}
	return &BasicUserInfo{
		Id:    data.ID,
		Name:  fmt.Sprintf("%s%s %s", data.TitleTH, data.FirstNameTH, data.LastNameTH),
		Email: data.ThaiEmail1,
		Login: data.ThaiEmail1,
	}, nil
}
