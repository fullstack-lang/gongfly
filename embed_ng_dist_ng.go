package gongfly

import "embed"

// NgDistNg is the export of angular distribution. This allows
// embedding of the pages in the web server
//
//go:embed ng-github.com-fullstack-lang-gongfly/dist/ng-github.com-fullstack-lang-gongfly
var NgDistNg embed.FS
