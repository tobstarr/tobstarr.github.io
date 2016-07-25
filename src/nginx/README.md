# Nginx

## Disable daemon mode

	nginx -g 'daemon off;'

## Rewrite to https

	rewrite ^/(.*)$ https://translations.phraseapp.io/$1 permanent;
