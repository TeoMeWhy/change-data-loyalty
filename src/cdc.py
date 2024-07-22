import datetime
import os

import pandas as pd
import sqlalchemy


class CDC:

    def __init__(self, path_new, path_old, tablename, id_field) -> None:

        engine_new = sqlalchemy.create_engine(f"sqlite:///{path_new}")
        engine_old = sqlalchemy.create_engine(f"sqlite:///{path_old}")
        
        self.df_new = pd.read_sql_table(tablename, engine_new)
        self.df_old = pd.read_sql_table(tablename, engine_old)

        self.tablename = tablename
        self.id_field = id_field
        
    def get_news(self):
        df_merge = self.df_new.merge(self.df_old,
                                on=[self.id_field],
                                how="left",
                                suffixes=["", "_old"],
                                indicator=True)

        news = df_merge[df_merge["_merge"] == 'left_only']
        news = news[self.df_new.columns]
        news["Op"] = "I"
        news['modified_date'] = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        return news


    def get_deleted(self):
        df_merge = self.df_old.merge(self.df_new,
                                on=[self.id_field],
                                how="left",
                                suffixes=["", "_new"],
                                indicator=True)

        deleted = df_merge[df_merge["_merge"] == 'left_only']
        deleted = deleted[self.df_new.columns]
        deleted["Op"] = "D"
        deleted['modified_date'] = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        return deleted


    def get_updates(self, news):
        df_merge = self.df_new.merge(self.df_old,
                            on=self.df_new.columns.tolist(),
                            how="left",
                            suffixes=["", "_old"],
                            indicator=True)
        
        update = df_merge[df_merge["_merge"] == 'left_only']
        update = update[~update[self.id_field].isin(news[self.id_field])]
        update = update[self.df_new.columns]
        update["Op"] = "U"
        update['modified_date'] = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        return update 

    def make_cdc(self):
        news = self.get_news()
        deleted = self.get_deleted()        
        update = self.get_updates(news)
        df_cdc = pd.concat([news,deleted,update], axis=0, ignore_index=True).reset_index(drop=True)
        return df_cdc

    def save(self, df):
        now = datetime.datetime.now().strftime("%Y%m%d_%H%M%S%f")
        path = f"../data/cdc/{self.tablename}/"
        filename = f"{path}/{now}.parquet"
        os.makedirs(path, exist_ok=True)
        df.to_parquet(filename)


    def execute(self):
        df = self.make_cdc()
        self.save(df)
