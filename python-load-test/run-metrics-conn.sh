#!/bin/bash

if [ -z $MAX_LOOP ]; then MAX_LOOP=5; fi
if [ -z $MAX_PARALEL ]; then MAX_PARALEL=100; fi
if [ -z $DELAY ]; then DELAY=0.5; MUTLIPLER=4; else MUTLIPLER=$(( 2 / $DELAY )); fi

FINAL_MAX_PARALEL=$(($MUTLIPLER * $MAX_PARALEL / $MAX_LOOP));

echo "[METRIC] START: update metrics conn, max: $FINAL_MAX_PARALEL, delay: $DELAY" 

psql --quiet "host=localhost port=5432 dbname=hasura user=postgres password=secret" -c "TRUNCATE metrics.connections"

for ((i=1; i<=$FINAL_MAX_PARALEL; i++))
do
	echo "[METRIC] updating $i"
	psql --quiet "host=localhost port=5432 dbname=hasura user=postgres password=secret" -f update-metrics-conn.sql
	sleep $DELAY
done

echo "[METRIC] END: update metrics conn" 
