# Api

## User 

## Files

CREATE TABLE public.orders (
	order_uid varchar NOT NULL,
	track_number varchar,
	entry varchar,
	delivery json null,
	payment json NULL,
	items json NULL,
	locale varchar,
	internal_signature varchar,
	customer_id varchar,
	delivery_service varchar,
	shardkey varchar,
	sm_id int,
	date_created timestamptz,
	oof_shard varchar
);

select * from orders; 

insert into orders (order_uid, delivery) values ('00-1-1', '{"name": "название", "phone": "79999", "zip": "---"}'); 
