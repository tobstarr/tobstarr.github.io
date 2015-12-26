#!/bin/bash
set -xe

curl -d @template_not_analyzed.json -X PUT http://127.0.0.1:9200/_template/template_not_analyzed

