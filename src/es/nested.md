# Nested

## Mapping
	curl -X POST 127.0.0.1:9200/index_test -d '
	{
			"mappings" : {
				"document": {
						"properties" : {
								"translations" : {
										"type" : "nested",
										"properties": {
												"locale_name" : {"type": "string" },
												"content"  : {"type": "string" }
										}
								}
						}
					}
			}
	}

## Inner Hits

	{
		"query" : {
			"nested": {
			 "path": "translations",
			 "query": {
				"query_string": {
				 "query": "translations.locale_name:'en' AND translations.verified:true"
				}
								},
			 "inner_hits": {
				"fields": "locale_name"
			 }
			}
		}
	}
