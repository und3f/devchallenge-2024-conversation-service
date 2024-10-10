FROM python:3.10

WORKDIR /usr/src/app

COPY ./requirements.txt ./
RUN pip install torch --extra-index-url https://download.pytorch.org/whl/cpu && \
    pip install --no-cache-dir -r requirements.txt

COPY get-models.py models.py .
RUN ./get-models.py

COPY *.py /usr/src/app/

EXPOSE 8082

ENV HF_HUB_OFFLINE=1
CMD [ "/usr/src/app/server.py" ]
