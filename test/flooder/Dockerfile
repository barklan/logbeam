ARG DOCKER_IMAGE_PREFIX=
FROM ${DOCKER_IMAGE_PREFIX}tiangolo/uvicorn-gunicorn-fastapi:python3.8

ARG BUILDKIT_INLINE_CACHE=1

WORKDIR /app/

ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

RUN pip install wheel

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY ./ /app
ENV PYTHONPATH=/app

COPY ./entrypoint.local.sh /
RUN chmod +x /entrypoint.local.sh
