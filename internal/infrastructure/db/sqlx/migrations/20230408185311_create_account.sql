-- +goose Up
-- +goose StatementBegin
create table account(
	id  varchar(100) not null primary key,
    description varchar(250) not null,
    due_date timestamp not null,
    payment_date timestamp not null,
    value numeric,
    type varchar(25) not null,
    status varchar(25) not null,
    owner_id varchar(250),
    account_group_id varchar(250),
    installments int not null,
    created_at timestamp,
    updated_at timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table account;
-- +goose StatementEnd
