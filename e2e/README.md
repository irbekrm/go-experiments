# e2e

## ginkgo

* bootstrap

`ginkgo bootstrap --nodot` // can be run to bootstrap without dot import, but will just explicitly import identifiers at the top level

Looks like not using `bootstrap` command at all is the only way how to avoid dot imports _and_ explicit import of all the `ginkgo/gomega` simbols at top level. Same with `ginkgo generate` command.