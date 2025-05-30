SELECT
    uuid AS idCliente,
    case when desc_email is not null then 1 else 0 end as flEmail,
    case when id_twitch is not null then 1 else 0 end as flTwitch,
    case when id_you_tube is not null then 1 else 0 end as flYouTube,
    case when id_blue_sky is not null then 1 else 0 end as flBlueSky,
    case when id_instagram is not null then 1 else 0 end as flInstagram,
    nr_points as qtdePontos,
    created_at AS dtCriacao,
    updated_at AS dtAtualizacao

FROM points.customers