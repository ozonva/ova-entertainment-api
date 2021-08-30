-- +goose Up
-- +goose StatementBegin
CREATE table entertainments (
       id bigserial primary key,
       user_id bigint not null,
       title varchar(50) not null,
       description text null,
       date timestamp default now()
);

INSERT INTO entertainments (user_id, title, description)
VALUES
(1, 'Экскурсии по пригородам Санкт-Петербурга', 'Петергоф, Пушкин, Кронштадт, Выборг'),
(1, 'Фабрика шедевров Арт-Тир', 'В этом интерактивном пространстве каждый может почувствовать себя художником и создать картину собственными руками'),
(2, 'VK Fest 2021', 'Один из крупнейших опэн-эйров страны снова в Петербурге!');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE entertainments;
-- +goose StatementEnd
