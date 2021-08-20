.PHONY: deploy preview
all: app

app:
	GOOS=linux GOARCH=amd64 go build

deploy: app
	sam deploy --region ap-northeast-1 --capabilities CAPABILITY_IAM --s3-bucket nana-lambda --stack-name freee-redirector

preview: app
	sam local start-api
