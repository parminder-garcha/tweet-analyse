# tweet-analyse

## This is a small hobby project showing how we can link 3 components using go code, these being :- 

1) Inserting tweets into the postgres db. 
2) Analysing the tweets using openAI.
3) Outputting the tweet analyses to standart output and updating db for future data analysis.

The aim is to demonstrate rather than analyse hundereds/thousands of tweets, we can use the power of openAI to do the grunt work for us ! 

## For the code to run, if not already installed then download  :- 

* go @ https://go.dev/doc/install 
* docker @ https://docs.docker.com/get-docker/

## Getting started 

Clone the repo :- 
   ```sh
   git clone https://github.com/parminder-garcha/tweet-analyse.git
   ```
 Issue the following from the cli :-   
``` sh
docker-compose up
``` 

Docker command above will setup the Postgres db with initialisation data, after which you can see the response from openAI by running the executable :- 
``` sh
./tweets
```

## Roadmap

Rather than the hard coded SQL Inserts, the eventual plan will be, download public twitter tweets, on any subject (politics, sport etc) - and have these inserted 
into the database.  
