CREATE TABLE payer (
  id   SERIAL PRIMARY KEY,
  payer_id   INT UNIQUE NOT NULL,
  payer_name TEXT NOT NULL,
  parent_id  INT
);

INSERT INTO payer (payer_id, payer_name, parent_id)
VALUES (140,'Metropolitan Jewish Health System Elderplan',	320);

INSERT INTO payer (payer_id, payer_name, parent_id)
VALUES (842, 'Anthem Empire BlueCross BlueShield',606);

CREATE TABLE plan (
  id   SERIAL PRIMARY KEY,
  plan_id INT UNIQUE NOT NULL,
  payer_id INT  NOT NULL,
  plan_name TEXT NOT NULL,
  plan_segment TEXT,
  plan_type    TEXT,
  ppo          VARCHAR(1),
  hmo          VARCHAR(1),
  pos          VARCHAR(1)
);

INSERT INTO plan ( payer_id,plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos)
VALUES (	140	,	17629	,'Elderplan Advantage For Nursing Home Residents'	,'Medicare'	,'Medicare SN'	,'Y'	,'N'	,'N'	);

INSERT INTO plan (payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos)
VALUES (140	,	17630	,'Elderplan Extra Help'	,'Medicare'	,'Medicare MA'	,'Y','N','N');

INSERT INTO plan (payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos)
VALUES	(141	,	14989	,'Anthem HealthSync'	,'Commercial'	,'Commercial'	,'Y'	,'N'	,'N'	);
INSERT INTO plan (payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos)
VALUES (	141	,	4353	,'Blue Access'	,'Commercial'	,'Commercial'	,'Y'	,'N'	,'N'	);

INSERT INTO plan ( payer_id,plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos)
VALUES (141	,	15403	,'Empire BlueCross BlueShield Child Health Plus Medicaid'	,'Medicaid'	,'Managed Medicaid'	,'Y'	,'N'	,'N'	);