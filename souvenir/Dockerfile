FROM python

RUN pip install grpcio-tools grpcio-reflection pymysql

ADD . /data

WORKDIR /data

CMD [ "python", "-u", "main.py" ]
