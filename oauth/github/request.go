package github

import (
	"resty.dev/v3"
)

func (a *oauth) getAccessToken(code string) (*accessToken, error) {
	client := resty.New()
	res, err := client.R().SetHeader("Accept", "application/json").SetQueryParams(map[string]string{
		"client_id":     a.config.AccessKeyID,
		"client_secret": a.config.AccessKeySecret,
		"code":          code,
		"redirect_uri":  a.basePath + callbackPath,
	}).SetResult(&accessToken{}).Post(accessTokenURL)
	if err != nil {
		return nil, err
	}

	return res.Result().(*accessToken), nil
}

func (a *oauth) getUserInfo(code, state string) error {
	tokenResp, err := a.getAccessToken(code)
	if err != nil {
		return err
	}
	client := resty.New()
	res, err := client.R().
		SetHeaders(map[string]string{
			"Accept":               "application/vnd.github+json",
			"X-GitHub-Api-Version": "2022-11-28",
		}).
		SetAuthToken(tokenResp.AccessToken).SetResult(&userInfo{}).
		Get(userAPIURL)
	if err != nil {
		return err
	}
	a.cache.setData(state, res.Result().(*userInfo))
	return nil
}
