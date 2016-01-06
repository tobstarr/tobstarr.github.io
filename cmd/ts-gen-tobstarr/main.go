package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tobstarr/tobstarr.github.io/cmd/ts-gen-tobstarr/Godeps/_workspace/src/github.com/dynport/dgtk/cli"
	"github.com/tobstarr/tobstarr.github.io/cmd/ts-gen-tobstarr/Godeps/_workspace/src/github.com/shurcooL/github_flavored_markdown"
)

var sources = map[string][]Source{
	"CNAME":                                FileSources("src/CNAME"),
	"aws-sdk-go.html":                      Layout(MarkdownSource(Render(FileSource("src/aws-sdk-go/index.md")))),
	"aws.html":                             Layout(MarkdownSource(Render(FileSource("src/aws/index.md")))),
	"cheats.html":                          Layout(FileSource("src/cheats.tpl")),
	"css/bootstrap.min.css":                FileSources("src/css/bootstrap.min.css"),
	"css/default.css":                      FileSources("src/css/default.css"),
	"docker.html":                          Layout(MarkdownSource(FileSource("src/articles/docker/index.md"))),
	"dotfiles.html":                        Layout(FileSource("src/dotfiles/index.tpl")),
	"es.html":                              Layout(MarkdownSource(Render(FileSource("src/es/index.md")))),
	"gnupg.html":                           Layout(MarkdownSource(FileSource("src/articles/gnupg/index.md"))),
	"go-build-flags.html":                  markdown("src/go-build-flags/index.md"),
	"goerrors.html":                        Layout(MarkdownSource(Render(FileSource(("src/go-errors/index.md"))))),
	"hello.txt":                            FileSources("src/hello.txt"),
	"id_rsa.pub":                           FileSources("src/id_rsa.pub"),
	"index.html":                           Layout(FileSource("src/index.tpl")),
	"js/jquery.dataTables.min.js":          FileSources("src/js/jquery.dataTables.min.js"),
	"js/jquery.min.js":                     FileSources("src/js/jquery.min.js"),
	"keybase.txt":                          FileSources("src/keybase.txt"),
	"loop.sh":                              Layout(MarkdownSource(Render(FileSource("src/loop/loop.sh")))),
	"nfl.html":                             Layout(FileSource("src/nfl.html")),
	"python-web-server.html":               Layout(MarkdownSource(FileSource("src/python-web-server/index.md"))),
	"qrcat/index.html":                     Layout(MarkdownSource(Render(FileSource("src/qrcat/index.md")))),
	"qrcat/qrcat.go":                       FileSources("src/qrcat/qrcat.go"),
	"qrcat/qrcat.png":                      FileSources("src/qrcat/qrcat.png"),
	"rkt.html":                             markdown("src/rkt/index.md"),
	"scratch/index.html":                   Layout(MarkdownSource(Render(FileSource("src/scratch/index.md")))),
	"setup_geminabox.sh":                   FileSources("src/speed-up-bundler-with-geminabox/setup_geminabox.sh"),
	"setup_go.sh":                          FileSources("src/articles/setup-go/install.sh"),
	"setup_vim.sh":                         FileSources("src/articles/setup-vim/setup_vim.sh"),
	"speed-up-bundler-with-geminabox.html": Layout(MarkdownSource(Render(FileSource("src/speed-up-bundler-with-geminabox/index.md")))),
	"systemd.html":                         Layout(MarkdownSource(Render(FileSource("src/systemd/index.md")))),
	"systemd_timer.html":                   Layout(MarkdownSource(Render(FileSource("src/systemd/timers.md")))),
	"tmux.conf":                            FileSources("src/tmux.conf"),
	"tobstarr.gpg":                         FileSources("src/tobstarr.gpg"),
	"versions.html":                        Layout(TemplateSource(versionsTpl, allVersions())),
	"vimrc":                                FileSources("src/vimrc"),
	"vimrc.conf":                           FileSources("src/vimrc"),
	"zsh.html":                             Layout(MarkdownSource(Render(FileSource("src/zsh/index.md")))),
	"zshrc.sh":                             FileSources("src/zsh/zshrc.sh"),
	"umlauts.html":                         FileSources("src/umlauts.html"),
}

func markdown(s string) []Source {
	return Layout(MarkdownSource(Render(FileSource(s))))
}

func chain(s string, funcs ...func(Source) Source) Source {
	return func(w io.Writer) error {
		s := FileSource(s)
		for _, f := range funcs {
			s = f(s)
		}
		return s(w)
	}
}

func main() {
	switch err := cli.RunActionWithArgs(&run{}); err {
	case nil, cli.ErrorHelpRequested, cli.ErrorNoRoute:
		// ignore
		return
	default:
		log.New(os.Stderr, "", 0).Fatal(err)
	}
}

type run struct {
	DoNotPush bool `cli:"opt --do-not-push"`
}

func (r *run) Run() error {
	l := log.New(os.Stderr, "", 0)
	wd, err := createRelease(l)
	if err != nil || r.DoNotPush {
		return err
	}
	c := exec.Command("git", "add", ".")
	c.Dir = wd
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	if err := c.Run(); err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	c = exec.Command("git", "diff", "--cached")
	c.Dir = wd
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return err
	}
	fmt.Printf("commit message: ")
	line, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		return err
	}

	buf = &bytes.Buffer{}
	c = exec.Command("git", "commit", "-m", string(line))
	c.Dir = wd
	c.Stdout = buf
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	if err := c.Run(); err != nil {
		if strings.Contains(buf.String(), "nothing to commit, working directory clean") {
			l.Printf("nothing to commit")
			return nil
		}
		return fmt.Errorf("%s: %q", err, buf.String())
	}

	c = exec.Command("git", "push")
	c.Dir = wd
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return err
	}
	l.Printf("released!")
	return nil
}

func FileSources(paths ...string) (list []Source) {
	for _, p := range paths {
		list = append(list, FileSource(p))
	}
	return list

}

func loadHTMLFiles() (list map[string]struct{}, err error) {
	list = map[string]struct{}{}
	return list, filepath.Walk(".", func(p string, info os.FileInfo, err error) error {
		switch {
		case info.IsDir():
			switch p {
			case ".git", "cmd":
				return filepath.SkipDir
			}
			return nil
		case strings.HasSuffix(p, ".html"):
			list[p] = struct{}{}
		}
		return nil
	})
}

func Layout(s Source) []Source {
	return []Source{FileSource("src/header.tpl"), s, FileSource("src/footer.tpl")}
}

func TemplateSource(tpl string, payload interface{}) Source {
	return func(w io.Writer) error {
		t, err := template.New("").Parse(tpl)
		if err != nil {
			return err
		}
		return t.Execute(w, payload)
	}
}

func MarkdownSource(in Source) Source {
	return func(w io.Writer) error {
		buf := &bytes.Buffer{}
		if err := in(buf); err != nil {
			return err
		}
		b := github_flavored_markdown.Markdown(buf.Bytes())
		_, err := w.Write(b)
		return err
	}
}

func Render(in Source) Source {
	return func(w io.Writer) error {
		buf := &bytes.Buffer{}
		if err := in(buf); err != nil {
			return err
		}
		s, err := renderTemplate(buf.String())
		if err != nil {
			return fmt.Errorf("error rendering string: %q", err.Error())
		}
		_, err = io.WriteString(w, s)
		return err
	}
}

func getOrigin() (string, error) {
	b, err := exec.Command("git", "remote", "-v").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s:\b%s", err, b)
	}
	for _, line := range strings.Split(string(b), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 3 && fields[0] == "origin" && fields[2] == "(push)" {
			return fields[1], nil
		}
	}
	return "", fmt.Errorf("unable to extract origin from\n%s", b)
}

func createRelease(l Logger) (dir string, err error) {
	start := time.Now()
	res, err := exec.Command("git", "status", "--porcelain").CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s\n%s", res, err)
	}
	if len(res) > 0 {
		return "", fmt.Errorf("status not clean:\n%s", res)
	}
	wd := ".release"
	if _, err := os.Stat(filepath.Join(wd, ".git", "HEAD")); err != nil {
		origin, err := getOrigin()
		if err != nil {
			return "", err
		}
		c := exec.Command("git", "clone", "-b", "master", "--depth", "1", origin, wd)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			return "", err
		}
	}
	c := exec.Command("git", "reset", "origin/master")
	c.Dir = wd
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return "", err
	}
	err = filepath.Walk(wd, func(p string, info os.FileInfo, err error) error {
		if p == wd {
			return nil
		}
		if filepath.Base(p) == ".git" {
			return filepath.SkipDir
		}
		if err := os.RemoveAll(p); err != nil {
			return err
		}
		if info.IsDir() {
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	for path, sources := range sources {
		if err := writeFile(l, filepath.Join(wd, path), sources...); err != nil {
			return "", err
		}
	}
	l.Printf("finished in %.06f", time.Since(start).Seconds())
	return wd, nil
}

func FileSource(path string) Source {
	return func(w io.Writer) error {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		return err
	}
}

type Source func(w io.Writer) error

type Logger interface {
	Printf(string, ...interface{})
}

func writeFile(l Logger, out string, sources ...Source) error {
	//start := time.Now()
	buf := &bytes.Buffer{}
	for _, src := range sources {
		if err := src(buf); err != nil {
			return err
		}
	}
	if err := os.MkdirAll(filepath.Dir(out), 0755); err != nil {
		return err
	}
	tmp := out + ".tmp." + strconv.FormatInt(time.Now().UTC().UnixNano(), 36)
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)
	_, err = io.Copy(f, buf)
	if err != nil {
		return err
	}
	err = os.Rename(tmp, out)
	if err != nil {
		return err
	}
	//l.Printf("write file %s in %.6f", out, time.Since(start).Seconds())
	return nil
}
