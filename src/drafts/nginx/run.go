package nginx

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	l := log.New(os.Stderr, "", 0)
	if err := run(l); err != nil {
		log.Fatal(err)
	}
}

func run(l *log.Logger) error {
	l.Printf("running")

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	c := exec.Command("docker", "run", "--rm", "-p", "8080:8080", "-p", "80:80", "-v", filepath.Join(wd, "nginx.conf")+":/etc/nginx/nginx.conf:ro", "nginx:1.9.3")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
