FROM python:3

WORKDIR /usr/src/app

COPY ./requirements.txt ./
RUN pip install torch --index-url https://download.pytorch.org/whl/cpu && \
    pip install --no-cache-dir -r requirements.txt

COPY ./get-models.py .
RUN ./get-models.py

COPY ./server.py /usr/src/app/server.py

EXPOSE 8080

CMD [ "/usr/src/app/server.py" ]
