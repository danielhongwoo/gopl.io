#!/bin/sh

head -n 4000 top-1m.csv | awk -F ',' '{print "http://"$2}' | xargs ./main echo -n
