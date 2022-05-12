#! /usr/bin/bash
# set -e

bash ./prestart.sh

uvicorn app.main:app --host 0.0.0.0 --port 8000 --log-level info
