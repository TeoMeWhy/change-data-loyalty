# %%
import os
import dotenv
import time
import logging


import datetime

if dotenv.load_dotenv("../../.env"):
    print("Environment variables loaded successfully.")
else:
    print("Failed to load environment variables.")

KAGGLE_USERNAME = os.getenv("KAGGLE_USERNAME")
print(KAGGLE_USERNAME)

KAGGLE_KEY = os.getenv("KAGGLE_KEY")

from kaggle.api.kaggle_api_extended import KaggleApi
client = KaggleApi()
client.authenticate()

logging.basicConfig(level=logging.INFO)

logging.info("Starting versioning process...")
while True:

    logging.info("Creating a new version of the dataset...")
    now = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    client.dataset_create_version_cli(
        folder="../../data/points",
        version_notes=f"{now} - New version",
        convert_to_csv=False,
        quiet=False,
    )

    logging.info("Version created successfully. Waiting for the next cycle...")
    time.sleep(60 * 60 * 6)
