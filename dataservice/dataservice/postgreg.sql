CREATE TABLE T_DRG_Payer_DATA  (
  id   SERIAL PRIMARY KEY,
  payer_id   INT UNIQUE NOT NULL,
  payer_name TEXT NOT NULL,
  parent_id  INT
);

INSERT INTO T_DRG_Payer_DATA (payer_id, payer_name, parent_id)
VALUES (140,'Metropolitan Jewish Health System Elderplan',	320);

INSERT INTO T_DRG_Payer_DATA  (payer_id, payer_name, parent_id)
VALUES (842, 'Anthem Empire BlueCross BlueShield',606);

CREATE TABLE T_DRG_PLAN_DATA (
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

INSERT INTO T_DRG_PLAN_DATA ( payer_id,plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos)
VALUES (	140	,	17629	,'Elderplan Advantage For Nursing Home Residents'	,'Medicare'	,'Medicare SN'	,'Y'	,'N'	,'N'	);

INSERT INTO T_DRG_PLAN_DATA (payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos)
VALUES (140	,	17630	,'Elderplan Extra Help'	,'Medicare'	,'Medicare MA'	,'Y','N','N');

INSERT INTO T_DRG_PLAN_DATA (payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos)
VALUES	(141	,	14989	,'Anthem HealthSync'	,'Commercial'	,'Commercial'	,'Y'	,'N'	,'N'	);
INSERT INTO T_DRG_PLAN_DATA (payer_id, plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos)
VALUES (	141	,	4353	,'Blue Access'	,'Commercial'	,'Commercial'	,'Y'	,'N'	,'N'	);

INSERT INTO T_DRG_PLAN_DATA ( payer_id,plan_id, plan_name, plan_segment,plan_type,ppo,hmo,pos)
VALUES (141	,	15403	,'Empire BlueCross BlueShield Child Health Plus Medicaid'	,'Medicaid'	,'Managed Medicaid'	,'Y'	,'N'	,'N'	);

CREATE TABLE T_NCPDP_DATA (
  id       SERIAL PRIMARY KEY,
  pharmacy_id INT UNIQUE NOT NULL,
  pharmacy_name TEXT    NOT NULL,
  ncpdp       TEXT NOT NULL,
  npi_nbr         TEXT,
  phone_nbr       TEXT,
  fax_nbr          TEXT,
  created_Date     DATE not null,
  created_by       TEXT null,
  last_updated_date date,
  last_updated_by   TEXT
);

INSERT INTO T_NCPDP_DATA (
  pharmacy_id ,
  pharmacy_name ,
  ncpdp       ,
  npi_nbr         ,
  phone_nbr    ,
  fax_nbr       ,
  created_Date    ,
  created_by       
  ) VALUES(
    14262, --pharmacy_id ,
  'WALGREENS', --pharmacy_name ,
  '32079',--ncpdp       ,
  null, --npi_nbr         ,
  null, --phone_nbr    ,
  null, --fax_nbr       ,
  CURRENT_DATE, --created_Date    ,
  'DataService' --created_by   
  )

  INSERT INTO T_NCPDP_DATA (
  pharmacy_id ,
  pharmacy_name ,
  ncpdp       ,
  npi_nbr         ,
  phone_nbr    ,
  fax_nbr       ,
  created_Date    ,
  created_by       
  ) VALUES(
    14305, --pharmacy_id ,
  'MCKESSON SPECIALTY PHARMACY', --pharmacy_name ,
  '35271',--ncpdp       ,
  '1205035011', --npi_nbr         ,
  '4806634160', --phone_nbr    ,
  '8665583777', --fax_nbr       ,
  CURRENT_DATE, --created_Date    ,
  'DataService' --created_by   
  )

/*
14262	WALGREENS	32079			
14265	ACCREDO	9876123	9939911111	9919991112	8818828881
14305	MCKESSON SPECIALTY PHARMACY	35271	1205035011	4806634160	8665583777
*/