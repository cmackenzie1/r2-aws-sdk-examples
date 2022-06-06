# r2-aws-sdk-examples

A collection of examples using the AWS SDK for S3 + Cloudflare R2.

In general, the following configuration settings must be specified when using any AWS S3 SDK.

- `region` set to `auto`
- `endpoint` set to `https://<account_id>.r2.cloudflarestorage.com`
- `sigining algorithm` set to `v4`

## Running an example

1. Create a file `.env` with the following properties specified

```bash
R2_ACCESS_KEY_ID=
R2_SECRET_ACCESS_KEY=
R2_ACCOUNT_ID=
```

2. Navigate to the target directory
3. Run `env $(cat ./.env | xargs) go run main.go` or `env $(cat ./.env | xargs) python main.py` (your command may
   differ)

## Languages

### Go

- ✅ `aws-sdk-go-v1`
- ✅ `aws-sdk-go-v2`

### Python

- ✅ `boto3`
    - `env $(cat ./.env | xargs) python main.py`

