FROM python:3

WORKDIR /usr/src/app

COPY ./requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY ./get-models.py .
RUN ./get-models.py

COPY ./server.py .

CMD [ "./server.py" ]
