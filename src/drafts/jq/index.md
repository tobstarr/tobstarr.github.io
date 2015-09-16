# JQ

tutorial at: https://stedolan.github.io/jq/tutorial/

cat issues.json | jq '.[].title' -r

-r: raw output

# Iterate and return json

jq '.DBInstances[] | {id: .DBInstanceIdentifier, status: .DBInstanceStatus}'
