import os
import datetime
import time

import boto3

MONGODB_URI = os.getenv('MONGODB_URI')

S3_ACCESS_KEY_ID = os.getenv('S3_ACCESS_KEY_ID')
S3_SECRET_ACCESS_KEY = os.getenv('S3_SECRET_ACCESS_KEY')
S3_REGION = os.getenv('S3_REGION')
S3_ENDPOINT = os.getenv('S3_ENDPOINT')
S3_BUCKET = os.getenv('S3_BUCKET')


def login_s3():
    client = boto3.client('s3',
                          aws_access_key_id=S3_ACCESS_KEY_ID,
                          aws_secret_access_key=S3_SECRET_ACCESS_KEY,
                          endpoint_url=S3_ENDPOINT,
                          region_name=S3_REGION)
    return client


def dump_mongodb(target):
    # run mongodump
    os.system(f'mongodump --uri="{MONGODB_URI}" --gzip --archive={target}')


def upload_to_s3(source, target):
    client = login_s3()
    client.upload_file(source, S3_BUCKET, target)


def main():
    dump_mongodb('dump.gz')
    # upload to s3
    date_str = datetime.datetime.now().strftime('%Y%m%d%H%M%S')
    upload_to_s3('dump.gz', f'mongodb_dump_{date_str}.gz')


if __name__ == '__main__':
    while True:
        main()
        time.sleep(60 * 60 * 24)
