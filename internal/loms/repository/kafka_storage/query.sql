-- name: CreateMessage :exec
insert into kafka_outbox(event, entity_type, entity_id, data)
values
    ($1,$2, $3, $4);
