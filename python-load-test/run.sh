#!/bin/bash

export MAX_LOOP=10
export MAX_PARALEL=100
bash run-metrics-conn.sh &
sleep 1
python wms-product-client.py
