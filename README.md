# GitHub Labels cloner

Cloning GitHub labels from config files.

## Usage

```
  -configs value
    	Config path can contains multiple files
  -owner string
    	Repository owner (either username or organization name)
  -replace
    	Replace existed labels with config
  -repo string
    	Repository name; if not exist, will updates all repository from owner
```

## Example

```bash
GITHUB_TOKEN=$(gh auth token) ./ghlabel -replace \
  -configs .configs/default.yaml \
  -configs .configs/hacktoberfest.yaml \
  -configs .configs/release-please.yaml \
  -owner kc-workspace
```
