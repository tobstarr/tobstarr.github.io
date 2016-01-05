# ElasticSearch

## Not analyzed mapping

{{ require "src/es/template_not_analyzed.json" | code }}

curl --data-binary -X PUT http://127.0.0.1:9200/_template/template_not_analyzed

## sort by RandomK

see http://stackoverflow.com/questions/9796470/random-order-pagination-elasticsearch

	curl -XGET 'localhost:9200/_search' -d '{
		"query": {
			"function_score" : {
				"query" : { "match_all": {} },
				"random_score" : {}
			}
		}
	}'
