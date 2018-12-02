FROM python

RUN pip install grpcio-tools

ADD . /data

WORKDIR /data

CMD [ "python", "main.py" ]