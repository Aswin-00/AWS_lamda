

create a role 


aws iam create-role \
  --role-name LambdaExecutionRole \
  --assume-role-policy-document '{
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Principal": {
          "Service": "lambda.amazonaws.com"
        },
        "Action": "sts:AssumeRole"
      }
    ]
  }'


attach policy extra policy

aws iam attach-role-policy \
  --role-name LambdaExecutionRole \
  --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole




show the policy
aws iam get-role --role-name LambdaExecutionRole --query "Role.Arn" --output text



create the lambda 
aws cloudformation deploy --template-file lambda-deployment.yml --stack-name LambdaStack
