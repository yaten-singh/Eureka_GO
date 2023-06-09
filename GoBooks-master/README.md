
# Golang Bookstore

![alt tag](https://upload.wikimedia.org/wikipedia/commons/thumb/2/23/Golang.png/300px-Golang.png)

##### Bookstore sample app showing classic CRUD data functions.

*Minimalist implementation, no web-framework, or ORM.*

Examples:
* Data layer separated from rendering layer to show example typical with the need for multiple different client types (web, mobile)
* JSON data serving -> marshal and unmarshal
* reusable model and db definitions
* HTML templates with inheritence

:+1: props to [practical-persistence]( http://www.alexedwards.net/blog/practical-persistence-sql) for Golang sql examples

Setup Postgres with...
---
```
CREATE TABLE books (
  isbn    char(14) NOT NULL,
  title   varchar(255) NOT NULL,
  author  varchar(255) NOT NULL,
  price   decimal(5,2) NOT NULL
);

INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);

ALTER TABLE books ADD PRIMARY KEY (isbn);
```
---

Clone this repo onto your system, then start dataserver with...

---
```
~/GoBooks$ go run bookstore_Dataserve.go
Bookstore: dataserver (port:4000)
```
---

now start webserver with...
---
```
~/GoBooks$ go run bookstore_Webserve.go
Bookstore: webserver (port:3000)
```
---

then browse http://127.0.0.1:3000/books

---
```
Bookstore index:

978-1503261969 	Del 	Emma 	Jayne Austen 	9.44
978-1505255607 	Del 	The Time Machine 	H. G. Wells 	5.99
978-1503379640 	Del 	The Prince 	Niccolò Machiavelli 	6.99
978-1470184841 	Del 	Metamorphosis 	Franz Kafka 	5.9
978-1470180925 	Del 	Siddhartha 	Hermann Hesse 	9.99

Add Book
```
---

#### endpoints:

 * LIST /books
 * SHOW /books/show?isbn=#
 * CREATE /books/create_book
 * DELETE /books/delete_book?isbn=#
 * UPDATE **TBD*

markdown created with http://jbt.github.io/markdown-editor





