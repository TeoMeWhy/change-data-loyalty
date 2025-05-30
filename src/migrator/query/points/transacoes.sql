SELECT
	UUID AS idTransacao,
	id_customer AS idCliente,
	created_at AS dtCriacao,
	vl_points AS qtdePontos,
	desc_sys_origin AS descSistemaOrigem

FROM points.transactions