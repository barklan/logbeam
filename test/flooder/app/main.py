import asyncio
import pathlib
import random
import sys
import time
from typing import Any, Dict

import faker
import fastapi as fa
from app.core import custom_logging

fake = faker.Faker()

config_path = pathlib.Path(__file__).with_name("logging_config.json")
log = custom_logging.CustomizeLogger.make_logger(config_path)

app = fa.FastAPI(
    title="Test app",
    debug=True,
)


@app.on_event("startup")
def repeat_random_log():
    while True:
        # time.sleep(random.randint(1, 2))
        time.sleep(6)
        for i in range(5):
            log.info(fake.text())
            log.error("This is the ERROR!")


@app.get("/info/{id}")
async def infolog(id: str):
    log.info(f"Hey! this is a test log! {id}")
    time.sleep(1)
    return "Hey!"


@app.get("/print/{id}")
async def justprint(id: str):
    print(f"This is print! {id}")
    return "Hey"


@app.get("/error/{id}")
async def errorlog(id: str):
    log.error(f"Hey! this is an er ror log! {id}")
    time.sleep(1)
    return "Hey!"
