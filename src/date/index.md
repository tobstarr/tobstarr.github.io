# Date

## Use UTC as timezone

	TZ=UTC date

## Print ISO8601 nanoseconds (works only on linux)

	TZ=UTC date --iso-8601=ns

## Print date relative to now

	TZ=UTC date +"%Y-%m-%dT%H:%M:%S" -d "1 day ago"
