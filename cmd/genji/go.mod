module github.com/genjidb/genji/cmd/genji

go 1.15

require (
	github.com/agnivade/levenshtein v1.1.0
	github.com/c-bata/go-prompt v0.2.5
	github.com/genjidb/genji v0.12.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	go.uber.org/multierr v1.6.0
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys v0.0.0-20200930185726-fdedc70b468f // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/genjidb/genji v0.12.0 => ../../
