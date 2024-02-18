package gh

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/SadS4ndWiCh/gogh/pkg/htmlx"
	"golang.org/x/net/html"
)

type GithubRepository struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	UpdatedAt       string `json:"updated_at"`
	GitUrl          string `json:"git_url"`
	SSHUrl          string `json:"ssh_url"`
	CloneUrl        string `json:"clone_url"`
	StargazersCount int    `json:"stargazers_count"`
	Language        string `json:"language"`
}

func GetRepositories(user string, page int) ([]*GithubRepository, error) {
	url := fmt.Sprintf("https://github.com/%s?page=%d&tab=repositories", user, page)
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

	repositoriesContainerEl := htmlx.GetElementByAttribute(doc, "data-filterable-for", "your-repos-filter")
	if repositoriesContainerEl == nil {
		return nil, GHError{message: "Invalid HTML", status: INVALID_HTML}
	}

	var repositories []*GithubRepository
	repositoriesListEl := htmlx.GetElementsByTagName(repositoriesContainerEl, "li")
	for _, repoEl := range repositoriesListEl {
		repository := getRepository(repoEl, user)
		repositories = append(repositories, repository)
	}

	return repositories, nil
}

func getRepository(repoEl *html.Node, user string) *GithubRepository {
	nameEl := htmlx.GetElementByAttribute(repoEl, "itemprop", "name codeRepository")
	name, _ := htmlx.GetTextContent(nameEl)

	var description string
	descriptionEl := htmlx.GetElementByAttribute(repoEl, "itemprop", "description")
	if descriptionEl != nil {
		description, _ = htmlx.GetTextContent(descriptionEl)
	}

	var language string
	languageEl := htmlx.GetElementByAttribute(repoEl, "itemprop", "programmingLanguage")
	if languageEl != nil {
		language, _ = htmlx.GetTextContent(languageEl)
	}

	updatedAtEl := htmlx.GetElementByTagName(repoEl, "relative-time")
	updatedAt, _ := htmlx.GetAttribute(updatedAtEl, "datetime")

	stargazersCount := getStars(repoEl)

	return &GithubRepository{
		Name:            name,
		Description:     description,
		Language:        language,
		UpdatedAt:       updatedAt,
		StargazersCount: stargazersCount,
		GitUrl:          fmt.Sprintf("git://github.com/%s/%s.git", user, name),
		SSHUrl:          fmt.Sprintf("git@github.com:%s/%s.git", user, name),
		CloneUrl:        fmt.Sprintf("https://github.com/%s/%s.git", user, name),
	}
}

func getStars(repoEl *html.Node) int {
	linksEl := htmlx.GetElementsByClassname(repoEl, "Link--muted mr-3")
	for _, linkEl := range linksEl {
		linkHref, exists := htmlx.GetAttribute(linkEl, "href")
		if !exists || !strings.HasSuffix(linkHref, "stargazers") {
			continue
		}

		stargazersStr, exists := htmlx.GetTextContent(linkEl.LastChild)
		if !exists {
			return 0
		}

		stargazers, err := strconv.Atoi(strings.TrimSpace(stargazersStr))
		if err != nil {
			return 0
		}

		return stargazers
	}

	return 0
}
