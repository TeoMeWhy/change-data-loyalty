import argparse
import subprocess

import pyarrow
import fastparquet
from tqdm import tqdm

from cdc import CDC

def download():
    cmd = [
        "kaggle",
        "datasets",
        "download",
        "-d",
        "teocalvo/teomewhy-loyalty-system",
        "--unzip",
        "--force",
        "--path",
        "../data/"
        ]
    
    subprocess.run(cmd)

# %%
def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--tables", "-t", nargs="+", default=["customers", "transactions", "transactions_product"])
    parser.add_argument("--id_fields", "-i", nargs="+", default=["idCustomer", "idTransaction", "idTransactionCart"])
    parser.add_argument("--new", "-n", default="../data/database.db")
    parser.add_argument("--old", "-o", default="../data/database_old.db")
    args = parser.parse_args()

    download()

    for i in range(len(args.tables)):
        cdc = CDC(path_new=args.new,
                  path_old=args.old,
                  tablename=args.tables[i],
                  id_field=args.id_fields[i])
        cdc.execute()

if __name__ == "__main__":
    main()