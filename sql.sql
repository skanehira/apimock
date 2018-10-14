select * from endpoints;
select * from endpoints where id = '1';

insert into endpoints (id, url, method, response_status, response_headers, response_body, created_at, description) values ("3","https://localhost/users/","POST",200,"test: header","test body", date("2018-10-10"),"test descritpion");

