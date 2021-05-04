# using go mod
1. `cd <existing-go-project>/<existing-folder>`
1. `go mod init github.com/<username>/<existing-go-project>/<existing-folder>`
1. time flies, run `go mod tidy`

# using go lang dep (deprecated)
1. `cd <existing-go-project>`
1. `go init` -- only first time, this will create Gopkg.toml & Gopkg.lock
1. time flies, edit Gopkg.toml
1. `go ensure`
