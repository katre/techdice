workspace(name = "go_downloader")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Set up rules_go.
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "87f0fb9747854cb76a0a82430adccb6269f7d394237104a4523b51061c469171",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.23.1/rules_go-v0.23.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.23.1/rules_go-v0.23.1.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

# Set up gazelle.
http_archive(
    name = "bazel_gazelle",
    sha256 = "bfd86b3cbe855d6c16c6fce60d76bd51f5c8dbc9cfcaef7a2bb5c1aafd0710e8",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.0/bazel-gazelle-v0.21.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.0/bazel-gazelle-v0.21.0.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "discordgo",
    importpath = "github.com/bwmarrin/discordgo",
    sum = "h1:EVujVeX9y7hGISnw3SrhOTNiPtrAfihWdJetkumLZng=",
    version = "v0.20.0",
)

go_repository(
    name = "dgrouter",
    importpath = "github.com/Necroforger/dgrouter",
    sum = "h1:xuNW6yNogKnVFtgYWZIBNxlcnMMlIhegVxwNRDp8HhM=",
    version = "v0.0.0-20190528143456-040421b5a83e",
)

go_repository(
    name = "testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:nOGnQDM7FYENwehXlg/kFVnos3rEvtKTjRvOWSzb6H4=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:clyUAQHOM3G0M3f5vQj7LuJrETvjVot3Z5el9nffUtU=",
    version = "v2.3.0",
)

go_repository(
    name = "aoc",
    importpath = "github.com/katre/advent_of_code",
    sum = "h1:SO4n7ng3x05R14b/H0S57cbfavn1a8QaT4t533+oklg=",
    version = "v0.0.0-20200107154730-a463264398b4",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    sum = "h1:+/TMaTYc4QFitKJxsQ7Yye35DkWvkdLcvGKqM+x0Ufc=",
    version = "v1.4.2",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:cg5LA/zNPRzIXIWSCxQW10Rvpy94aQh3LT/ShoCpkHw=",
    version = "v0.0.0-20200510223506-06a226fb4e37",
)
