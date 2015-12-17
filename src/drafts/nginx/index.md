# Nginx Cheats

## Redirect http => https

This can be e.g. used in the listener config of an AWS ElasticLoadBalancer. Just forward all trafic on port 80 to instance port 8080 and all traffic on 443 to port 80.

  server {
    listen 8080;
      server_name  _;
      rewrite ^/(.*)$ https://example.com/$1 permanent;
  }

## Log to syslog

	access_log syslog:server=unix:/dev/log main;
