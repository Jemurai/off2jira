# OFF To JIRA

Push off formatted findings to JIRA.

## Releasing

off2jira works to follow golang best practices.  Therefore, when updating, we need to do the following:

- `go get -u` 
- `go mod tidy`
- `git commit -m "change with version"`
- `git tag v1.0.6`
- `git push origin v1.0.6`

Run the build.sh and get the different types of artifacts and include them in the release.
