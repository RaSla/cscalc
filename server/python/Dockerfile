FROM python:3.12-alpine
WORKDIR /app

COPY requirements.txt /app/
RUN set -ex \
  && pip install -r requirements.txt \
  && python --version \
  && pip --version

## Copy App in docker-image - for the Production
COPY server.py /app/
## or mount App-volume in docker-compose.yml - for the Development

CMD python server.py
