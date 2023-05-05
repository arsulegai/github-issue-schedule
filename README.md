# GitHub Issue Schedule

Create issues to project repositories as per the scheduled
configuration.

## Environment Variables

| Env Variable | Description | Default |
|--- | --- | ---|
| GITHUB_TOKEN | Token that can give access to configured repositories | |
| CONFIG_FILE | Input configuration | config.yaml |

## Contributions

Each commit must include `Signed-off-by:`
in the commit message (run `git commit -s` to auto-sign).
This sign off means you agree the commit satisfies the
[Developer Certificate of Origin(DCO)](https://developercertificate.org/).

### Build
 
`Makefile` is provided to ease up the development, and auto-perform lint
checks as well as run tests if any. This is to maintain a good quality on
the codebase.

Run the following command from your terminal

```shell
make
```

In case of build failure because of format check, run the following command

```shell
make format
```

### Version

Increment version by raising a PR against [VERSION](./VERSION) file.
Use the same tag while creating a release/tag.
