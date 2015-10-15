# AWK

## Filter all lines where the second word is some.host

	awk '{ if ($2 == "some.host") print $0 }'`
