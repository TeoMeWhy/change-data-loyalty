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
    exit(1)


DATA_FOLDER = os.getenv("DATA_FOLDER", "../../data/points")
KAGGLE_USERNAME = os.getenv("KAGGLE_USERNAME")
KAGGLE_KEY = os.getenv("KAGGLE_KEY")


def load_to_kaggle(kaggle_client):
    logging.info("Creating a new version of the dataset...")
    now = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    kaggle_client.dataset_create_version_cli(
        folder=DATA_FOLDER,
        version_notes=f"{now} - New version",
        convert_to_csv=False,
        quiet=False,
    )

def executer(kaggle_client):
    logging.info("Starting the dataset versioning process...")
    while True:
        load_to_kaggle(kaggle_client)
        logging.info("Version created successfully. Waiting for the next cycle...")
        time.sleep(60 * 60 * 6) 


def main():
    from kaggle.api.kaggle_api_extended import KaggleApi
    client = KaggleApi()
    client.authenticate()

    logging.basicConfig(level=logging.INFO)
    executer(client)

if __name__ == "__main__":
    main()