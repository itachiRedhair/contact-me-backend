.PHONY: gomodgen deploy delete

gomodgen:
	GO111MODULE=on go mod init

deploy:
	gcloud functions deploy SendEmail --entry-point SendEmail --runtime go111 --trigger-http

delete:
	gcloud functions delete SendEmail --entry-point SendEmail --runtime go111 --trigger-http
