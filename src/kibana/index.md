# Kibana

	sudo docker run --net=host -t -i -e ES_JAVA_OPTS="-Xms1g -Xmx1g" --rm elasticsearch:5.0.0
	sudo docker run -t -i --rm elasticsearch:2.3.3
	ip=172.17.0.2
	sudo docker run -t -i --rm --name some-kibana -e ELASTICSEARCH_URL=http://${ip}:9200 --net=host kibana:4.5.1
	sudo docker run -t -i --rm --name some-kibana -e ELASTICSEARCH_URL=http://127.0.0.1:9200 --net=host kibana:5.0.0
