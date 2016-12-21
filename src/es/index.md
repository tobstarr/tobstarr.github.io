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

## ES 5 in linux

	sudo sysctl -w vm.max_map_count=262144

## Sort via query string

	sort=anio:desc

## Update API

	curl -XPOST 'http://localhost:9200/designs/shirt/1/_update' -d'
	{
		 "script" : "ctx._source.votes += 1"
	}'

	# add `?retry_on_conflict=5` to avoid version conflicts

## Drain nodes in a cluster

	curl -XPUT ${master}:9200/_cluster/settings -d '{
		"transient" :{
				"cluster.routing.allocation.exclude._ip" : "172.31.14.194,172.31.7.158,172.31.5.167,172.31.7.157"
		 }
	}';echo
