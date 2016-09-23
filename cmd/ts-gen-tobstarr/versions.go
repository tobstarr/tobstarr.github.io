package main

import (
	"bytes"
	"html/template"
	"sort"
)

type version struct {
	Name    string `json:"name,omitempty"`
	Website string `json:"website,omitempty"`
	Docker  string `json:"docker,omitempty"`
}

var versionsMap = map[string]*version{
	"Consul":          {Website: "https://www.consul.io/downloads.html"},
	"CoreOS":          {Website: "https://coreos.com/releases/"},
	"Docker Compose":  {Website: "https://github.com/docker/compose/releases"},
	"Docker Registry": {Docker: "https://hub.docker.com/_/registry/"},
	"Docker":          {Website: "https://github.com/docker/docker/releases"},
	"ElasticSearch":   {Website: "https://www.elastic.co/downloads/elasticsearch", Docker: "https://registry.hub.docker.com/_/elasticsearch/"},
	"Geminabox":       {Website: "https://github.com/geminabox/geminabox/releases"},
	"Golang":          {Website: "https://golang.org/dl/", Docker: "https://registry.hub.docker.com/_/golang/"},
	"JQuery":          {Website: "http://jquery.com/"},
	"Java":            {Docker: "https://registry.hub.docker.com/_/java/"},
	"Jenkins":         {Website: "https://jenkins-ci.org/changelog"},
	"Kibana":          {Website: "https://www.elastic.co/downloads/kibana", Docker: "https://registry.hub.docker.com/_/kibana/"},
	"Kubernetes":      {Website: "https://github.com/kubernetes/kubernetes/releases"},
	"Linux Kernel":    {Website: "https://www.kernel.org"},
	"Minikube":        {Website: "https://github.com/kubernetes/minikube/releases"},
	"Mongo":           {Docker: "https://registry.hub.docker.com/_/mongo/"},
	"MySQL":           {Website: "http://dev.mysql.com/downloads/mysql/", Docker: "https://registry.hub.docker.com/_/mysql/"},
	"Nginx":           {Website: "http://nginx.org/en/download.html", Docker: "https://registry.hub.docker.com/_/nginx/"},
	"PHP":             {Docker: "https://registry.hub.docker.com/_/php/"},
	"Percona":         {Docker: "https://hub.docker.com/_/percona/"},
	"PostgreSQL":      {Website: "http://www.postgresql.org/ftp/source/", Docker: "https://registry.hub.docker.com/_/postgres/"},
	"Redis":           {Website: "http://redis.io/download", Docker: "https://registry.hub.docker.com/_/redis/"},
	"Ruby":            {Website: "https://www.ruby-lang.org/en/downloads/", Docker: "https://registry.hub.docker.com/_/ruby/"},
	"Selenium":        {Website: "http://www.seleniumhq.org/download/"},
	"SyslogNG":        {Website: "https://github.com/balabit/syslog-ng/releases"},
	"rkt":             {Website: "https://github.com/coreos/rkt/releases"},
}

var versionsTpl = `<h2>Versions</h2>

<table class="table table-striped">
<thead>
	<tr>
		<th>Name</th>
		<th>Original</th>
		<th>Docker</th>
	</tr>
</thead>
<tbody>
{{ range . }}
<tr>
  <td>{{ .Name }}
  <td>{{ if .Website }}<a href="{{ .Website }}">{{ .Name }}</a>{{ end }}
  <td>{{ if .Docker }}<a href="{{ .Docker }}">{{ .Name }}</a>{{ end }}
{{ end }}
</tbody>
</table>

<script type="text/javascript" charset="utf-8">
$(document).ready(function() {
	$('table').DataTable({pageLength: 100});
})
</script>
`

func versionsSrc() (string, error) {
	t, err := template.New("").Parse(versionsTpl)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, allVersions()); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func allVersions() (list versions) {
	for name, c := range versionsMap {
		c.Name = name
		list = append(list, c)
	}
	sort.Sort(list)
	return list
}

type versions []*version

func (list versions) Len() int {
	return len(list)
}

func (list versions) Swap(a, b int) {
	list[a], list[b] = list[b], list[a]
}

func (list versions) Less(a, b int) bool {
	return list[a].Name < list[b].Name
}
