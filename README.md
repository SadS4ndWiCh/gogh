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