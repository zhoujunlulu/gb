package gb

import (
	"go/build"
	"path/filepath"
	"strings"
)

// Project represents a gb project. A gb project has a simlar layout to
// a $GOPATH workspace. Each gb project has a standard directory layout
// starting at the project root, which we'll refer too as $PROJECT.
//
//     $PROJECT/                       - the project root
//     $PROJECT/.gogo/                 - used internally by gogo and identifies
//                                       the root of the project.
//     $PROJECT/src/                   - base directory for the source of packages
//     $PROJECT/bin/                   - base directory for the compiled binaries
type Project struct {
	rootdir string
}

// NewContext returns a new build context from this project.
func (p *Project) NewContext(tc Toolchain) *Context {
	context := build.Default
	context.GOPATH = togopath(p.Srcdirs())
	return &Context{
		Project: p,
		Context: &context,
		tc:      tc,
		workdir: mktmpdir(),
	}
}

func togopath(srcdirs []string) string {
	var s []string
	for _, srcdir := range srcdirs {
		s = append(s, filepath.Dir(srcdir))
	}
	return strings.Join(s, ":")
}

func NewProject(root string) *Project {
	return &Project{
		rootdir: root,
	}
}

// Pkgdir returns the path to precompiled packages.
func (p *Project) Pkgdir() string {
	return filepath.Join(p.rootdir, "pkg")
}

// Projectdir returns the path root of this project.
func (p *Project) Projectdir() string {
	return p.rootdir
}

// Srcdirs returns the path to the source directories.
// The first source directory will always be
// filepath.Join(Projectdir(), "src)
// but there may be additional directories.
func (p *Project) Srcdirs() []string {
	srcdirs := []string{filepath.Join(p.Projectdir(), "src")}
	srcdirs = append(srcdirs, filepath.Join(p.Projectdir(), "vendor", "src"))
	return srcdirs
}