## Configuration Step's :

```
go run main.go
curl http://localhost:9000
ngrok http 9000
```

## Postgres

```
\l \d \du \c \q

sudo -u postgres psql

sudo -u postgres createdb instagram

\c instagram

create table status(
    ID serial primary key,
    NAME text not null,
    URL text not null
);

\d status

grant all privileges on database instagram to saiashish;
```

## Connect To Postgres

```
host : localhost
port : 5432
username : saiashish
password: saiashish
database: instagram
```

## Insert first row via terminal or add corresponding rows via vstudio

```
insert into status(name,url) values('Sarthak','https://cdn.pixabay.com/photo/2016/08/24/23/08/cristiano-ronaldo-1618341__340.jpg');
```

## Select * command shortcut

```
TABLE status;
```

```
SELECT current_database();
\d
\d status
TABLE status;
```

```
select * from status order by id desc limit 3;
\q
```

## Posts table

```
create table posts(
    ID serial primary key,
    NAME text not null,
    URL text not null,
    PROFILE_URL text not null,
    TITLE text not null,
    DESCRIPTION text not null,
    COMMENTS_COUNT text not null,
    TIME text not null
);

```

## ALTER TABLE COMMAND

```
sudo -u postgres psql
\c instagram
\d status
alter table status add msg text ;
```

## MEDIA

```
instagram=# create table media(
ID serial primary key,
url text,
is_video int,
is_gallery int DEFAULT 0);

alter table media alter column is_video set  default 0 ;

```