aws lambda create-function \
  --region "$REGION" \
  --function-name "$FUNCTION_NAME" \
  --memory 128 \
  --role arn:aws:iam::"$ACCOUNT_ID":role/"$EXECUTION_ROLE" \
  --runtime go1.x \
  --zip-file fileb://./handler.zip \
  --handler "$LAMBDA_HANDLER"
