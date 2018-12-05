FROM python

RUN pip install grpcio-tools grpcio-reflection

ADD . /data

WORKDIR /data

CMD [ "python", "main.py" ]
