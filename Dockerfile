FROM mongo

# Install  and Cron
RUN apt-get update && apt-get -y install python3

RUN pip3 install boto3

ADD run.py /run.py