package gh

import (
	"fmt"
	"io"
	"net/http"

	"github.com/SadS4ndWiCh/gogh/pkg/htmlx"
)

type GithubUser struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	AvatarUrl   string `json:"avatar_url"`
	HtmlUrl     string `json:"html_url"`
	PublicRepos string `json:"public_repos"`
	Followers   string `json:"followers"`
	Following   string `json:"following"`
	Bio         string `json:"bio"`
	Blog        string `json:"-"`
}

func GetUser(user string) (*GithubUser, error) {
	url := fmt.Sprintf("https://github.com/%s", user)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	doc, err := htmlx.Load(content)
	if err != nil {
		return nil, err
	}

	sidebarEl := htmlx.GetElementByClassname(doc, "js-profile-editable-replace")
	if sidebarEl == nil {
		return nil, GHError{message: "Invalid HTML", status: INVALID_HTML}
	}

	nameEl := htmlx.GetElementByClassname(sidebarEl, "vcard-fullname")
	name, _ := htmlx.GetTextContent(nameEl)

	usernameEl := htmlx.GetElementByClassname(sidebarEl, "vcard-username")
	username, _ := htmlx.GetTextContent(usernameEl)

	bioEl := htmlx.GetElementByClassname(doc, "user-profile-bio")
	bio, _ := htmlx.GetTextContent(bioEl)

	avatarEl := htmlx.GetElementByClassname(sidebarEl, "avatar avatar-user width-full border color-bg-default")
	avatarUrl, _ := htmlx.GetAttribute(avatarEl, "src")

	followersHref := fmt.Sprintf("%s?tab=followers", url)
	followersContainerEl := htmlx.GetElementByAttribute(doc, "href", followersHref)
	followersEl := htmlx.GetElementByTagName(followersContainerEl, "span")
	followers, _ := htmlx.GetTextContent(followersEl)

	followingHref := fmt.Sprintf("%s?tab=following", url)
	followingContainerEl := htmlx.GetElementByAttribute(doc, "href", followingHref)
	followingEl := htmlx.GetElementByTagName(followingContainerEl, "span")
	following, _ := htmlx.GetTextContent(followingEl)

	publicReposContainerEl := htmlx.GetElementByAttribute(doc, "data-tab-item", "repositories")
	publicReposEl := htmlx.GetElementByClassname(publicReposContainerEl, "Counter")
	publicRepos, _ := htmlx.GetTextContent(publicReposEl)

	ghUser := &GithubUser{
		Login:       username,
		Name:        name,
		Bio:         bio,
		AvatarUrl:   avatarUrl,
		HtmlUrl:     url,
		PublicRepos: publicRepos,
		Followers:   followers,
		Following:   following,
		Blog:        "",
	}

	return ghUser, nil
}
