# igz_rates_user


Install aws cli
https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html

GOOS=linux go build -o my-lambda-binary main.go harvest_user_rates_API.go harvest_users_API.go

zip function.zip my-lambda-binary

aws lambda create-function --function-name harvestDefatulRatesPerUser --runtime go1.x --zip-file fileb://function.zip --handler my-lambda-binary --role arn:aws:iam::374208052150:role/service-role/lambdaTestRole

aws lambda update-function-code --function-name arn:aws:lambda:eu-west-1:374208052150:function:harvestDefatulRatesPerUser --zip-file fileb://function.zip 