FROM mongo

# Install  and Cron
RUN apt-get update && apt-get -y install python3 python3-pip

RUN pip3 install boto3

ADD run.py /run.py

CMD [ "python3", "/run.py" ]