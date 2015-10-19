# ElasticSearch

## Not analyzed mapping

{{ require "src/es/template_not_analyzed.json" }}

curl -X PUT http://127.0.0.1:9200/_template/template_not_analyzed
