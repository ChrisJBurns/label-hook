# Configuration

Below is a description of what the properties are within in the `config.json` file and how the label-hook application uses them.


| Property                      | Description                                                                          |
| ----------------------------- | ------------------------------------------------------------------------------------ | 
| `organisation`                | User or Organisation in which you want the label-hook application to operate within e.g. `ChrisJBurns`  | 
| `host`                        | The host you want the application to listen and serve on e.g. `localhost`            |
| `port`                        | The port you want the application to listen and serve on e.g. `8080`                 |
| `access_token`                | The GitHub Personal Access Token that is used by `go-github` in order to authenticate the requests sent to the GitHub API                     |
