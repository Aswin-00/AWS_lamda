# Use the official AWS Lambda Python base image
FROM public.ecr.aws/lambda/python:3.9

# Copy function code
COPY app.py ${LAMBDA_TASK_ROOT}

# Set the CMD to your handler (filename.handler_function_name)
CMD ["app.lambda_handler"]
