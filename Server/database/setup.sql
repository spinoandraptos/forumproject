--this file can be loaded using psql -f to create the necessary tables within the forumdb database
--otherwise it will just hold info on the 4 main data structures stored in the database

--varchar(255) is used to restrict the username and password to 255 characters max as I do not wish for them to be overly long
--text is used for titles descriptions and contents as these can be very long depending on context and thus should not be restricted to a fixed length
--there is little difference between using varchar and text for long strings

--the ID column will be the primary key column of each table as ID is always unique whereas other fields can have potential overlaps
--the ID is also generated using the serial type which creates a sequence of integers that increments automatically
--columns like AuthorID and ThreadID are foreign keys which reference the primary key (ID) of table users and threads respectively
--threads is thus a child table of users and categories and comments a child table of users and threads table
--this referential integrity ensures that no one can insert rows in threads without a matching entry in users or categories
--finally on delete cascade is enabled so that when an user is deleted, all threads and comments that reference the user's id are also deleted and etc.

create table users (
  ID         serial primary key,
  Username   varchar(255) not null, 
  Password   varchar(255) not null,
  CreatedAt  timestamp not null,
  UpdatedAt  timestamp not null   
);

create table categories (
  ID           serial primary key,
  Title        text not null,
  Description  text not null
);

create table threads (
  ID          serial primary key,
  Title       text not null,
  Content     text not null,  
  AuthorID    integer references users(ID) on delete cascade,
  CategoryID  integer references categories(ID) on delete cascade,
  CreatedAt   timestamp not null,
  UpdatedAt   timestamp not null
);

create table comments (
  ID          serial primary key,
  Content     text not null,
  AuthorID    integer references users(ID) on delete cascade,
  ThreadID    integer references threads(ID) on delete cascade,
  CreatedAt   timestamp not null,
  UpdatedAt   timestamp not null
);