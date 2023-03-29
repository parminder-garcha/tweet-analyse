CREATE schema my_schema; 

CREATE TABLE IF NOT EXISTS my_schema.analyser (
  id       serial,
  prompt   text,
  "text"   text,
  response text,
  created  date default current_timestamp,
  PRIMARY KEY (id)
);

insert into my_schema.analyser(prompt, text) values('Classify the sentiment in these tweets:', 'I hate doing homework');
insert into my_schema.analyser(prompt, text) values('Classify the sentiment in these tweets:', 'I love going to school');
insert into my_schema.analyser(prompt, text) values('Classify the sentiment in these tweets:', 'Why are politicians incompetent');
insert into my_schema.analyser(prompt, text) values('Classify the sentiment in these tweets:', 'War and peace');
insert into my_schema.analyser(prompt, text) values('Classify the sentiment in these tweets:', 'What is 2 x 2');
insert into my_schema.analyser(prompt, text) values('Classify the sentiment in these tweets:', 'Trump was an interesting president');
