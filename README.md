# Change Data loyalty

Uma lib simples para você poder criar arquivos de CDC a partir de um dataset no Kaggle.

## Sobre

Com base nos dados do kaggle (que você encontra [aqui](https://www.kaggle.com/datasets/teocalvo/teomewhy-loyalty-system)), criamos os arquivos de Change Data Capture (CDC).

A ideia é comprar os dois databases:
- _database.db_: Dados da versão atual do banco de dados do nosso sistema de pontos
- _database_old.db_: Dados da versão anterior do nosso sistema de pontos

Com bases nos arquivos criados, voc6e pode realizar ingestão de dados utilizando esses arquivos _.parquet_s.

- Playlist sobre consumo do CDC no datalake com Spark Streaming + Upsert: [Lago do Mago](https://www.youtube.com/playlist?list=PLvlkVRRKOYFTcLehYZ2Bd5hGIcLH0dJHE)

## Setup

1. Crie um ambiente novo com ajuda do anaconda:

```bash
conda create --name cdc-loyalty python=3.12
```

2. Instale as bibliotecas necessárias:

```bash
pip install -r requirements.txt
```

3. Configure a autenticação da api do Kaggle:

Confira esse artigo disponível no Kaggle para você criar sua API key [www.kaggle.com/docs/api](https://www.kaggle.com/docs/api)

## Execução

Para executar CDC, basta executar o comando abaixo:

```bash
cd src; python main.py
```

Este comando executará os seguintes passos:

1. Download dos dados do Kaggle
2. Identificação das mudanças existentes em cada tabela, gerando os arquivos de CDC no seguinte caminho: `data/cdc/nome_tabela.arquivo_com_data.parquet`
