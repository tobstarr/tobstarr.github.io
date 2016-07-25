# openssl

## encode

	openssl aes-256-cfb -pass file:pwfile

## decode

	openssl aes-256-cfb -d -pass file:pwfile

## check website certificate

	echo | openssl s_client -connect translations.phraseapp.io:443 2>/dev/null | openssl x509 -noout -dates
