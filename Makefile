deploy:
	GOOS=linux GOARCH=386 go build -o readit-api
	scp readit-api rd:/tmp/readit-api
	ssh rd mv /tmp/readit-api /var/www/readit-api
	ssh rd sudo restart readit
