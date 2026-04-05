// Copyright 2021 The Casdoor Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package object

import (
	"encoding/gob"
	"fmt"
	"os"
	"slices"

	"github.com/casdoor/casdoor/conf"
	"github.com/casdoor/casdoor/util"
	"github.com/go-webauthn/webauthn/webauthn"
)

const (
	cheersAIDefaultOrganization = "CheersAI"
	cheersAIDefaultApplication  = "CheersAI-Desktop"
	cheersAILogoPath            = "/logo.png"
	cheersAIDefaultAffiliation  = "CheersAI Inc."
	cheersAIDefaultCurrency     = "CNY"
)

var cheersAIUserNavItems = []string{"/home-top", "/", "/shortcuts", "/apps"}

func InitDb() {
	existed := initBuiltInOrganization()
	if !existed {
		initBuiltInPermission()
		initBuiltInUser()
		initBuiltInCert()
		initBuiltInLdap()
	}

	initBuiltInProvider()
	initBuiltInApplication()
	ensureCheersAIDefaults()

	existed = initBuiltInApiModel()
	if !existed {
		initBuiltInApiAdapter()
		initBuiltInApiEnforcer()
		initBuiltInUserModel()
		initBuiltInUserAdapter()
		initBuiltInUserEnforcer()
	}

	initWebAuthn()
}

func getBuiltInAccountItems() []*AccountItem {
	return []*AccountItem{
		{Name: "Organization", Visible: true, ViewRule: "Public", ModifyRule: "Admin"},
		{Name: "ID", Visible: true, ViewRule: "Public", ModifyRule: "Immutable"},
		{Name: "Name", Visible: true, ViewRule: "Public", ModifyRule: "Admin"},
		{Name: "Display name", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "First name", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Last name", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Avatar", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "User type", Visible: true, ViewRule: "Public", ModifyRule: "Admin"},
		{Name: "Password", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "Email", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Phone", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Country code", Visible: true, ViewRule: "Public", ModifyRule: "Admin"},
		{Name: "Country/Region", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Location", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Address", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Addresses", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Affiliation", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Title", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "ID card type", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "ID card", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "ID card info", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "Real name", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "ID verification", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "Homepage", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Bio", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Tag", Visible: true, ViewRule: "Public", ModifyRule: "Admin"},
		{Name: "Language", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Gender", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Birthday", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Education", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Balance", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "Balance credit", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "Balance currency", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "Cart", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "Transactions", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "Score", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Karma", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Ranking", Visible: true, ViewRule: "Public", ModifyRule: "Self"},
		{Name: "Signup application", Visible: true, ViewRule: "Public", ModifyRule: "Admin"},
		{Name: "Register type", Visible: true, ViewRule: "Public", ModifyRule: "Admin"},
		{Name: "Register source", Visible: true, ViewRule: "Public", ModifyRule: "Admin"},
		{Name: "Roles", Visible: true, ViewRule: "Public", ModifyRule: "Immutable"},
		{Name: "Permissions", Visible: true, ViewRule: "Public", ModifyRule: "Immutable"},
		{Name: "Groups", Visible: true, ViewRule: "Public", ModifyRule: "Admin"},
		{Name: "Consents", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "3rd-party logins", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "Properties", Visible: true, ViewRule: "Admin", ModifyRule: "Admin"},
		{Name: "Is admin", Visible: true, ViewRule: "Admin", ModifyRule: "Admin"},
		{Name: "Is forbidden", Visible: true, ViewRule: "Admin", ModifyRule: "Admin"},
		{Name: "Is deleted", Visible: true, ViewRule: "Admin", ModifyRule: "Admin"},
		{Name: "Multi-factor authentication", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "MFA items", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "WebAuthn credentials", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "Last change password time", Visible: true, ViewRule: "Admin", ModifyRule: "Admin"},
		{Name: "Managed accounts", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "Face ID", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "MFA accounts", Visible: true, ViewRule: "Self", ModifyRule: "Self"},
		{Name: "Need update password", Visible: true, ViewRule: "Admin", ModifyRule: "Admin"},
		{Name: "IP whitelist", Visible: true, ViewRule: "Admin", ModifyRule: "Admin"},
	}
}

func initBuiltInOrganization() bool {
	organization, err := getOrganization("admin", "built-in")
	if err != nil {
		panic(err)
	}

	if organization != nil {
		return true
	}

	organization = &Organization{
		Owner:              "admin",
		Name:               "built-in",
		CreatedTime:        util.GetCurrentTime(),
		DisplayName:        "CheersAI-SSO",
		WebsiteUrl:         "https://example.com",
		Logo:               fmt.Sprintf("%s/logo.png", conf.GetConfigString("staticBaseUrl")),
		LogoDark:           fmt.Sprintf("%s/logo-dark.svg", conf.GetConfigString("staticBaseUrl")),
		Favicon:            fmt.Sprintf("%s/favicon.png", conf.GetConfigString("staticBaseUrl")),
		PasswordType:       "bcrypt",
		PasswordOptions:    []string{"AtLeast6"},
		CountryCodes:       []string{"US", "ES", "FR", "DE", "GB", "CN", "JP", "KR", "VN", "ID", "SG", "IN"},
		DefaultAvatar:      cheersAILogoPath,
		UserTypes:          []string{},
		Tags:               []string{},
		Languages:          []string{"en", "es", "fr", "de", "ja", "zh", "vi", "pt", "tr", "pl", "uk"},
		InitScore:          2000,
		AccountItems:       getBuiltInAccountItems(),
		BalanceCurrency:    cheersAIDefaultCurrency,
		EnableSoftDeletion: false,
		IsProfilePublic:    false,
		UseEmailAsUsername: false,
		EnableTour:         true,
		DcrPolicy:          "open",
	}
	_, err = AddOrganization(organization)
	if err != nil {
		panic(err)
	}

	return false
}

func initBuiltInUser() {
	user, err := getUser("built-in", "admin")
	if err != nil {
		panic(err)
	}
	if user != nil {
		return
	}

	user = &User{
		Owner:             "built-in",
		Name:              "admin",
		CreatedTime:       util.GetCurrentTime(),
		Id:                util.GenerateId(),
		Type:              "normal-user",
		Password:          "123",
		DisplayName:       "Admin",
		Avatar:            cheersAILogoPath,
		Email:             "admin@example.com",
		Phone:             "12345678910",
		CountryCode:       "US",
		Address:           []string{},
		Affiliation:       cheersAIDefaultAffiliation,
		Tag:               "staff",
		Score:             2000,
		Ranking:           1,
		IsAdmin:           true,
		IsForbidden:       false,
		IsDeleted:         false,
		SignupApplication: "app-built-in",
		RegisterType:      "Add User",
		RegisterSource:    "built-in/admin",
		CreatedIp:         "127.0.0.1",
		Properties:        make(map[string]string),
	}
	_, err = AddUser(user, "en")
	if err != nil {
		panic(err)
	}
}

func initBuiltInApplication() {
	application, err := getApplication("admin", "app-built-in")
	if err != nil {
		panic(err)
	}

	if application != nil {
		if ensureRequiredApplicationProviders(application) {
			_, err = UpdateApplication(util.GetId("admin", "app-built-in"), application, true, "en")
			if err != nil {
				panic(err)
			}
		}
		return
	}

	application = &Application{
		Owner:          "admin",
		Name:           "app-built-in",
		CreatedTime:    util.GetCurrentTime(),
		DisplayName:    "CheersAI-SSO",
		Category:       "Default",
		Type:           "All",
		Scopes:         []*ScopeItem{},
		Logo:           fmt.Sprintf("%s/logo.png", conf.GetConfigString("staticBaseUrl")),
		Title:          "CheersAI-SSO",
		Favicon:        fmt.Sprintf("%s/favicon.png", conf.GetConfigString("staticBaseUrl")),
		HomepageUrl:    "",
		Organization:   "built-in",
		Cert:           "cert-built-in",
		EnablePassword: true,
		EnableSignUp:   true,
		Providers: []*ProviderItem{
			{Name: "provider_captcha_default", CanSignUp: false, CanSignIn: false, CanUnlink: false, Prompted: false, SignupGroup: "", Rule: "None", Provider: nil},
			{Name: "provider_storage_local_file_system", CanSignUp: false, CanSignIn: false, CanUnlink: false, Prompted: false, SignupGroup: "", Rule: "None", Provider: nil},
		},
		SigninMethods: []*SigninMethod{
			{Name: "Password", DisplayName: "Password", Rule: "All"},
			{Name: "Verification code", DisplayName: "Verification code", Rule: "All"},
			{Name: "WebAuthn", DisplayName: "WebAuthn", Rule: "None"},
			{Name: "Face ID", DisplayName: "Face ID", Rule: "None"},
		},
		SignupItems: []*SignupItem{
			{Name: "ID", Visible: false, Required: true, Prompted: false, Rule: "Random"},
			{Name: "Username", Visible: true, Required: true, Prompted: false, Rule: "None"},
			{Name: "Display name", Visible: true, Required: true, Prompted: false, Rule: "None"},
			{Name: "Password", Visible: true, Required: true, Prompted: false, Rule: "None"},
			{Name: "Confirm password", Visible: true, Required: true, Prompted: false, Rule: "None"},
			{Name: "Email", Visible: true, Required: false, Prompted: false, Rule: "Normal"},
			{Name: "Phone", Visible: true, Required: false, Prompted: false, Rule: "None"},
			{Name: "Agreement", Visible: true, Required: true, Prompted: false, Rule: "None"},
		},
		Tags:          []string{},
		RedirectUris:  []string{},
		TokenFormat:   "JWT",
		TokenFields:   []string{},
		ExpireInHours: 168,
		FormOffset:    2,

		CookieExpireInHours: 720,
	}
	_, err = AddApplication(application)
	if err != nil {
		panic(err)
	}
}

func readTokenFromFile() (string, string) {
	pemPath := "./object/token_jwt_key.pem"
	keyPath := "./object/token_jwt_key.key"
	pem, err := os.ReadFile(pemPath)
	if err != nil {
		return "", ""
	}
	key, err := os.ReadFile(keyPath)
	if err != nil {
		return "", ""
	}
	return string(pem), string(key)
}

func initBuiltInCert() {
	tokenJwtCertificate, tokenJwtPrivateKey := readTokenFromFile()
	cert, err := getCert("admin", "cert-built-in")
	if err != nil {
		panic(err)
	}

	if cert != nil {
		return
	}

	cert = &Cert{
		Owner:           "admin",
		Name:            "cert-built-in",
		CreatedTime:     util.GetCurrentTime(),
		DisplayName:     "Built-in Cert",
		Scope:           "JWT",
		Type:            "x509",
		CryptoAlgorithm: "RS256",
		BitSize:         4096,
		ExpireInYears:   20,
		Certificate:     tokenJwtCertificate,
		PrivateKey:      tokenJwtPrivateKey,
	}
	_, err = AddCert(cert)
	if err != nil {
		panic(err)
	}
}

func initBuiltInLdap() {
	ldap, err := GetLdap("ldap-built-in")
	if err != nil {
		panic(err)
	}

	if ldap != nil {
		return
	}

	ldap = &Ldap{
		Id:         "ldap-built-in",
		Owner:      "built-in",
		ServerName: "BuildIn LDAP Server",
		Host:       "example.com",
		Port:       389,
		Username:   "cn=buildin,dc=example,dc=com",
		Password:   "123",
		BaseDn:     "ou=BuildIn,dc=example,dc=com",
		AutoSync:   0,
		LastSync:   "",
	}
	_, err = AddLdap(ldap)
	if err != nil {
		panic(err)
	}
}

func ensureRequiredApplicationProviders(application *Application) bool {
	if application == nil {
		return false
	}

	requiredProviders := []string{
		"provider_captcha_default",
		"provider_storage_local_file_system",
	}
	modified := false
	existedProviders := map[string]bool{}

	for _, providerItem := range application.Providers {
		existedProviders[providerItem.Name] = true
	}

	for _, providerName := range requiredProviders {
		if existedProviders[providerName] {
			continue
		}

		application.Providers = append(application.Providers, &ProviderItem{
			Name:        providerName,
			CanSignUp:   false,
			CanSignIn:   false,
			CanUnlink:   false,
			Prompted:    false,
			SignupGroup: "",
			Rule:        "None",
			Provider:    nil,
		})
		modified = true
	}

	return modified
}

func ensureCheersAIDefaults() {
	oldDefaultAvatar := fmt.Sprintf("%s/img/casbin.svg", conf.GetConfigString("staticBaseUrl"))

	_, _ = ormer.Engine.Where("default_avatar = ?", oldDefaultAvatar).Cols("default_avatar").Update(&Organization{DefaultAvatar: cheersAILogoPath})
	_, _ = ormer.Engine.Where("avatar = ?", oldDefaultAvatar).Cols("avatar").Update(&User{Avatar: cheersAILogoPath})
	_, _ = ormer.Engine.Exec("update user set avatar = replace(avatar, 'https:/files/', '/files/') where avatar like 'https:/files/%'")
	_, _ = ormer.Engine.Exec("update user set avatar = replace(avatar, 'http:/files/', '/files/') where avatar like 'http:/files/%'")
	_, _ = ormer.Engine.Exec("update user set permanent_avatar = replace(permanent_avatar, 'https:/files/', '/files/') where permanent_avatar like 'https:/files/%'")
	_, _ = ormer.Engine.Exec("update user set permanent_avatar = replace(permanent_avatar, 'http:/files/', '/files/') where permanent_avatar like 'http:/files/%'")

	organization, err := getOrganization("admin", cheersAIDefaultOrganization)
	if err == nil && organization != nil {
		modified := false
		if application, appErr := getApplication("admin", cheersAIDefaultApplication); appErr == nil && application != nil {
			if ensureRequiredApplicationProviders(application) {
				_, _ = UpdateApplication(util.GetId("admin", cheersAIDefaultApplication), application, true, "en")
			}
		}
		if organization.DefaultAvatar == "" || organization.DefaultAvatar == oldDefaultAvatar {
			organization.DefaultAvatar = cheersAILogoPath
			modified = true
		}
		if organization.DefaultApplication == "" {
			if application, appErr := getApplication("admin", cheersAIDefaultApplication); appErr == nil && application != nil {
				organization.DefaultApplication = cheersAIDefaultApplication
				modified = true
			}
		}
		if organization.BalanceCurrency == "" || organization.BalanceCurrency == "USD" {
			organization.BalanceCurrency = cheersAIDefaultCurrency
			modified = true
		}
		if !slices.Equal(organization.UserNavItems, cheersAIUserNavItems) {
			organization.UserNavItems = append([]string{}, cheersAIUserNavItems...)
			modified = true
		}
		if modified {
			_, _ = UpdateOrganization(util.GetId("admin", cheersAIDefaultOrganization), organization, true)
		}
	}

	builtInOrganization, err := getOrganization("admin", "built-in")
	if err == nil && builtInOrganization != nil {
		modified := false
		if builtInOrganization.DefaultAvatar == "" || builtInOrganization.DefaultAvatar == oldDefaultAvatar {
			builtInOrganization.DefaultAvatar = cheersAILogoPath
			modified = true
		}
		if builtInOrganization.BalanceCurrency == "" || builtInOrganization.BalanceCurrency == "USD" {
			builtInOrganization.BalanceCurrency = cheersAIDefaultCurrency
			modified = true
		}
		if modified {
			_, _ = UpdateOrganization(util.GetId("admin", "built-in"), builtInOrganization, true)
		}
	}

	ensureUserDefaults("built-in", "admin", "app-built-in")
	ensureUserDefaults(cheersAIDefaultOrganization, "CheersAl_Admin", cheersAIDefaultApplication)
	ensureUserDefaults(cheersAIDefaultOrganization, "user_01", cheersAIDefaultApplication)
	ensureUserDefaults(cheersAIDefaultOrganization, "user_02", cheersAIDefaultApplication)
	ensureUserAdmin(cheersAIDefaultOrganization, "user_01")
	ensureApplicationPermission(cheersAIDefaultOrganization, "permission_CheersAl_admin", util.GetId(cheersAIDefaultOrganization, "CheersAl_Admin"), cheersAIDefaultApplication)
	ensureApplicationPermission(cheersAIDefaultOrganization, "permission_CheersAl_edit", util.GetId(cheersAIDefaultOrganization, "user_01"), cheersAIDefaultApplication)
	ensureApplicationPermission(cheersAIDefaultOrganization, "permission_CheersAl_member", util.GetId(cheersAIDefaultOrganization, "user_02"), cheersAIDefaultApplication)
}

func ensureUserDefaults(owner string, name string, signupApplication string) {
	user, err := getUser(owner, name)
	if err != nil || user == nil {
		return
	}

	oldDefaultAvatar := fmt.Sprintf("%s/img/casbin.svg", conf.GetConfigString("staticBaseUrl"))
	columns := []string{}

	if user.Avatar == "" || user.Avatar == oldDefaultAvatar {
		user.Avatar = cheersAILogoPath
		columns = append(columns, "avatar")
	}
	if user.Affiliation == "" || user.Affiliation == "Example Inc." {
		user.Affiliation = cheersAIDefaultAffiliation
		columns = append(columns, "affiliation")
	}
	if signupApplication != "" && (user.SignupApplication == "" || user.SignupApplication == "app-built-in") {
		user.SignupApplication = signupApplication
		columns = append(columns, "signup_application")
	}
	if user.BalanceCurrency == "" || user.BalanceCurrency == "USD" {
		user.BalanceCurrency = cheersAIDefaultCurrency
		columns = append(columns, "balance_currency")
	}

	if len(columns) == 0 {
		return
	}

	user.PermanentAvatar = "*"
	_, _ = UpdateUser(user.GetId(), user, columns, true)
}

func ensureUserAdmin(owner string, name string) {
	user, err := getUser(owner, name)
	if err != nil || user == nil || user.IsAdmin {
		return
	}

	user.IsAdmin = true
	_, _ = UpdateUser(user.GetId(), user, []string{"is_admin"}, true)
}

func ensureApplicationPermission(owner string, name string, userId string, application string) {
	permission, err := getPermission(owner, name)
	if err != nil {
		return
	}

	if permission == nil {
		permission = &Permission{
			Owner:        owner,
			Name:         name,
			DisplayName:  name,
			Model:        "user-model-built-in",
			ResourceType: "Application",
			Resources:    []string{application},
			Actions:      []string{"Read"},
			Users:        []string{userId},
			Roles:        []string{},
			Groups:       []string{},
			Effect:       "Allow",
			IsEnabled:    true,
			State:        "Approved",
		}
		_, _ = AddPermission(permission)
		return
	}

	modified := false
	if permission.Model == "" {
		permission.Model = "user-model-built-in"
		modified = true
	}
	if permission.ResourceType != "Application" {
		permission.ResourceType = "Application"
		modified = true
	}
	if !slices.Equal(permission.Resources, []string{application}) {
		permission.Resources = []string{application}
		modified = true
	}
	if !slices.Equal(permission.Actions, []string{"Read"}) {
		permission.Actions = []string{"Read"}
		modified = true
	}
	if !slices.Contains(permission.Users, userId) {
		permission.Users = []string{userId}
		modified = true
	}
	if permission.Effect != "Allow" {
		permission.Effect = "Allow"
		modified = true
	}
	if !permission.IsEnabled {
		permission.IsEnabled = true
		modified = true
	}
	if permission.State != "Approved" {
		permission.State = "Approved"
		modified = true
	}

	if modified {
		_, _ = UpdatePermission(permission.GetId(), permission)
	}
}

func initBuiltInProvider() {
	providers := []*Provider{
		{
			Owner:       "admin",
			Name:        "provider_captcha_default",
			CreatedTime: util.GetCurrentTime(),
			DisplayName: "Captcha Default",
			Category:    "Captcha",
			Type:        "Default",
		},
		{
			Owner:       "admin",
			Name:        "provider_balance",
			CreatedTime: util.GetCurrentTime(),
			DisplayName: "Balance",
			Category:    "Payment",
			Type:        "Balance",
		},
		{
			Owner:       "admin",
			Name:        "provider_payment_dummy",
			CreatedTime: util.GetCurrentTime(),
			DisplayName: "Dummy Payment",
			Category:    "Payment",
			Type:        "Dummy",
		},
		{
			Owner:       "admin",
			Name:        "provider_storage_local_file_system",
			CreatedTime: util.GetCurrentTime(),
			DisplayName: "Local File System",
			Category:    "Storage",
			Type:        "Local File System",
		},
	}

	for _, provider := range providers {
		existingProvider, err := GetProvider(util.GetId("admin", provider.Name))
		if err != nil {
			panic(err)
		}

		if existingProvider != nil {
			continue
		}

		_, err = AddProvider(provider)
		if err != nil {
			panic(err)
		}
	}
}

func initWebAuthn() {
	gob.Register(webauthn.SessionData{})
}

func initBuiltInUserModel() {
	model, err := GetModel("built-in/user-model-built-in")
	if err != nil {
		panic(err)
	}

	if model != nil {
		return
	}

	model = &Model{
		Owner:       "built-in",
		Name:        "user-model-built-in",
		CreatedTime: util.GetCurrentTime(),
		DisplayName: "Built-in Model",
		ModelText: `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act`,
	}
	_, err = AddModel(model)
	if err != nil {
		panic(err)
	}
}

func initBuiltInApiModel() bool {
	model, err := GetModel("built-in/api-model-built-in")
	if err != nil {
		panic(err)
	}

	if model != nil {
		return true
	}

	modelText := `[request_definition]
r = subOwner, subName, method, urlPath, objOwner, objName

[policy_definition]
p = subOwner, subName, method, urlPath, objOwner, objName

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = (r.subOwner == p.subOwner || p.subOwner == "*") && \
    (r.subName == p.subName || p.subName == "*" || r.subName != "anonymous" && p.subName == "!anonymous") && \
    (r.method == p.method || p.method == "*") && \
    (keyMatch2(r.urlPath, p.urlPath) || p.urlPath == "*") && \
    (r.objOwner == p.objOwner || p.objOwner == "*") && \
    (r.objName == p.objName || p.objName == "*") || \
    (r.subOwner == r.objOwner && r.subName == r.objName)`

	model = &Model{
		Owner:       "built-in",
		Name:        "api-model-built-in",
		CreatedTime: util.GetCurrentTime(),
		DisplayName: "API Model",
		ModelText:   modelText,
	}
	_, err = AddModel(model)
	if err != nil {
		panic(err)
	}
	return false
}

func initBuiltInPermission() {
	permission, err := GetPermission("built-in/permission-built-in")
	if err != nil {
		panic(err)
	}
	if permission != nil {
		return
	}

	permission = &Permission{
		Owner:        "built-in",
		Name:         "permission-built-in",
		CreatedTime:  util.GetCurrentTime(),
		DisplayName:  "Built-in Permission",
		Description:  "Built-in Permission",
		Users:        []string{"built-in/*"},
		Groups:       []string{},
		Roles:        []string{},
		Domains:      []string{},
		Model:        "built-in/user-model-built-in",
		Adapter:      "",
		ResourceType: "Application",
		Resources:    []string{"app-built-in"},
		Actions:      []string{"Read", "Write", "Admin"},
		Effect:       "Allow",
		IsEnabled:    true,
		Submitter:    "admin",
		Approver:     "admin",
		ApproveTime:  util.GetCurrentTime(),
		State:        "Approved",
	}
	_, err = AddPermission(permission)
	if err != nil {
		panic(err)
	}
}

func initBuiltInUserAdapter() {
	adapter, err := GetAdapter("built-in/user-adapter-built-in")
	if err != nil {
		panic(err)
	}

	if adapter != nil {
		return
	}

	adapter = &Adapter{
		Owner:       "built-in",
		Name:        "user-adapter-built-in",
		CreatedTime: util.GetCurrentTime(),
		Table:       "casbin_user_rule",
		UseSameDb:   true,
	}
	_, err = AddAdapter(adapter)
	if err != nil {
		panic(err)
	}
}

func initBuiltInApiAdapter() {
	adapter, err := GetAdapter("built-in/api-adapter-built-in")
	if err != nil {
		panic(err)
	}

	if adapter != nil {
		return
	}

	adapter = &Adapter{
		Owner:       "built-in",
		Name:        "api-adapter-built-in",
		CreatedTime: util.GetCurrentTime(),
		Table:       "casbin_api_rule",
		UseSameDb:   true,
	}
	_, err = AddAdapter(adapter)
	if err != nil {
		panic(err)
	}
}

func initBuiltInUserEnforcer() {
	enforcer, err := GetEnforcer("built-in/user-enforcer-built-in")
	if err != nil {
		panic(err)
	}

	if enforcer != nil {
		return
	}

	enforcer = &Enforcer{
		Owner:       "built-in",
		Name:        "user-enforcer-built-in",
		CreatedTime: util.GetCurrentTime(),
		DisplayName: "User Enforcer",
		Model:       "built-in/user-model-built-in",
		Adapter:     "built-in/user-adapter-built-in",
	}

	_, err = AddEnforcer(enforcer)
	if err != nil {
		panic(err)
	}
}

func initBuiltInApiEnforcer() {
	enforcer, err := GetEnforcer("built-in/api-enforcer-built-in")
	if err != nil {
		panic(err)
	}

	if enforcer != nil {
		return
	}

	enforcer = &Enforcer{
		Owner:       "built-in",
		Name:        "api-enforcer-built-in",
		CreatedTime: util.GetCurrentTime(),
		DisplayName: "API Enforcer",
		Model:       "built-in/api-model-built-in",
		Adapter:     "built-in/api-adapter-built-in",
	}

	_, err = AddEnforcer(enforcer)
	if err != nil {
		panic(err)
	}
}
