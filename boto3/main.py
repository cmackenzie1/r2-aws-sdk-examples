import os

import boto3

access_key_id = os.getenv("R2_ACCESS_KEY_ID")
secret_access_key = os.getenv("R2_SECRET_ACCESS_KEY")
account_id = os.getenv("R2_ACCOUNT_ID")

r2 = boto3.resource(
    "s3",
    endpoint_url=f"https://{account_id}.r2.cloudflarestorage.com",
    aws_access_key_id=access_key_id,
    aws_secret_access_key=secret_access_key,
)

for bucket in r2.buckets.all():
    print(bucket.name)
