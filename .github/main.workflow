workflow "Release on tag" {
  on = "push"
  resolves = ["goreleaser"]
}

action "is-tag" {
  uses = "actions/bin/filter@master"
  args = "tag"
}

action "goreleaser" {
  uses = "docker://goreleaser/goreleaser"
  needs = ["is-tag"]
  secrets = [
    "GITHUB_TOKEN"
  ]
  env = {
    CGO_ENABLED = 0
    GO111MODULE = "ON"
  }
  args = "release"
}
