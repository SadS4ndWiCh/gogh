# 🐙 gogh

An experiment to do web scraping with GO, creating my own functions to get elements in HTML on top of the `net/html` package.

## 🎏 Routes

#### Get user data
```
GET /users/{username}
```
```json
// "/users/SadS4ndWiCh"

{
  "id": "2507959",
  "login": "SadS4ndWiCh",
  "name": "Caio Vinícius",
  "avatar_url": "https://avatars.githubusercontent.com/u/71348567?v=4",
  "html_url": "https://github.com/SadS4ndWiCh",
  "public_repos": "35",
  "followers": "13",
  "following": "40",
  "bio": "world.execute(me);"
}
```
#### Get user repos 
```
GET /users/{username}/repos
```
| Query Param | Type     | Description                                    |
| :---------- | :------- | :--------------------------------------------- |
| `page`      | `number` | **Optional**: The current page. Default to `1` |
```json
// "/users/SadS4ndWiCh/repos"

[
  {
    "name": "gogh",
    "description": "🐙 An experiment to do web scraping with GO, creating my own functions to get elements in HTML on top of the `net/html` package.",
    "updated_at": "2024-02-18T01:10:35Z",
    "git_url": "git://github.com/SadS4ndWiCh/gogh.git",
    "ssh_url": "git@github.com:SadS4ndWiCh/gogh.git",
    "clone_url": "https://github.com/SadS4ndWiCh/gogh.git",
    "stargazers_count": 0,
    "language": "Go"
  },
  { ... }
]
```