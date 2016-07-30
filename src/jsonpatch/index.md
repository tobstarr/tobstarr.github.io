# jsonpatch

# add field to array
	[{"path": "/env/1", "value": {"name": "b", "value": 2}, "op": "add"}]

# replace field
	[{"path": "/env/1/value", "value": 2, "op": "replace"}]

# reorder
	[{"path": "/env/1", "op": "remove"}, {"path": "/env/2", "value": {"name": "c", "value": 3}, "op": "add"}]
