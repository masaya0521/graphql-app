create table todo (
    id text not null,
    t_text text ,
    done boolean,
    user_id text,
    primary key(id),
    foreign key (user_id) references users(id)
);