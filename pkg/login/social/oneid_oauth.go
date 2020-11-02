package social

import (
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana/pkg/models"
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
	var data = struct {
		ID          string `json:"id"`
		FirstNameTH string `json:"first_name_th"`
		LastNameTH  string `json:"last_name_th"`
		TitleTH     string `json:"account_title_th"`
		ThaiEmail   string `json:"thai_email"`
	}{}
	err = json.Unmarshal(response.Body, &data)
	if err != nil {
		return nil, fmt.Errorf("Error getting user info: %s", err)
	}
	return &BasicUserInfo{
		Id:    data.ID,
		Name:  fmt.Sprintf("%s%s %s", data.TitleTH, data.FirstNameTH, data.LastNameTH),
		Email: data.ThaiEmail,
		Login: data.ThaiEmail,
	}, nil
}
