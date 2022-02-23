create table todo (
    id integer not null,
    t_text text ,
    done integer,
    user_id integer,
    primary key(id),
    foreign key (user_id) references users(id)
);