package acmeclient

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
)

type Settings struct {
	Enabled            string          `json:"enabled"`
	AutoRenewal        string          `json:"autoRenewal"`
	UpdateCron         api.SelectedMap `json:"UpdateCron"`
	Environment        api.SelectedMap `json:"environment"`
	ChallengePort      string          `json:"challengePort"`
	TLSChallengePort   string          `json:"TLSchallengePort"`
	RestartTimeout     string          `json:"restartTimeout"`
	HaproxyIntegration string          `json:"haproxyIntegration"`
	HaproxyAclRef      api.SelectedMap `json:"haproxyAclRef"`
	HaproxyActionRef   api.SelectedMap `json:"haproxyActionRef"`
	HaproxyServerRef   api.SelectedMap `json:"haproxyServerRef"`
	HaproxyBackendRef  api.SelectedMap `json:"haproxyBackendRef"`
	LogLevel           api.SelectedMap `json:"logLevel"`
	ShowIntro          string          `json:"showIntro"`
}

type SettingsWrapper struct {
	Settings Settings `json:"settings"`
}

type SettingsMonad struct {
	Acmeclient SettingsWrapper `json:"acmeclient"`
}

type SettingsSetResult struct {
	Result string `json:"result"`
}

func (c *Controller) SettingsGet(ctx context.Context) (*SettingsMonad, error) {
	callOpts := api.RPCOpts{
		BaseEndpoint:    "/acmeclient/settings/get",
		Method:          "GET",
		PathParameters:  []string{},
		QueryParameters: map[string]string{},
		BodyParameters:  map[string]interface{}{},
	}

	resultData := &SettingsMonad{}
	result, err := api.Call(c.Client(), ctx, callOpts, resultData)
	if err != nil {
		return nil, fmt.Errorf("Get call failed: %w", err)
	}
	return result, nil
}

func (c *Controller) SettingsSet(ctx context.Context, settings *Settings) (*SettingsSetResult, error) {
	bodyParams := map[string]interface{}{
		"acmeclient": SettingsWrapper{Settings: *settings},
	}

	callOpts := api.RPCOpts{
		BaseEndpoint:    "/acmeclient/settings/set",
		Method:          "POST",
		PathParameters:  []string{},
		QueryParameters: map[string]string{},
		BodyParameters:  bodyParams,
	}

	resultData := &SettingsSetResult{}
	result, err := api.Call(c.Client(), ctx, callOpts, resultData)
	if err != nil {
		return nil, fmt.Errorf("Set call failed: %w", err)
	}
	return result, nil
}
