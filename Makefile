.PHONY: gomodgen deploy delete

gomodgen:
	GO111MODULE=on go mod init

deploy:
	gcloud functions deploy contactme --entry-point SendEmail --runtime go111 --trigger-http

delete:
	gcloud functions delete contactme --entry-point SendEmail --runtime go111 --trigger-http
