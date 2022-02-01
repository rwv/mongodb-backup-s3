FROM mongo

# Install  and Cron
RUN apt-get update && apt-get -y install python3 python3-pip && rm -rf /var/lib/apt/lists/*

RUN pip3 install --no-cache-dir boto3

ADD run.py /run.py

CMD [ "python3", "/run.py" ]