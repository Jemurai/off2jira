# OFF To JIRA

Push off formatted findings to JIRA.

## Running

1. Get your findings into OFF format (see [github.com/owasp/off](https://github.com/owasp/off))
1. `go get github.com/jemurai/off2jira`
1. Set environment variables
1. `off2jira off-file.json`

## Setup

Set your environment with:

- AUTOM8D_JIRA_SERVER - The Jira url.  Eg. `jemurai.atlassian.net`.
- AUTOM8D_JIRA_EMAIL - Your JIRA user email.
- AUTOM8D_JIRA_TOKEN - An API Token.  See [JIRA Documentation](https://developer.atlassian.com/cloud/jira/platform/basic-auth-for-rest-apis/)
- AUTOM8D_JIRA_PROJECT - The name of the JIRA Project to put the issue in.
- AUTOM8D_JIRA_ISSUE_TYPE (Optional, defaults to `Bug`)

See also the `example_env.sh` file.

## Releasing

off2jira works to follow golang best practices.  Therefore, when updating, we need to do the following:

- `go get -u` 
- `go mod tidy`
- `git commit -m "change with version"`
- `git tag v1.0.6`
- `git push origin v1.0.6`

Run the build.sh and get the different types of artifacts and include them in the release.
