#!/bin/bash

clear

docker run --rm -i grafana/k6 run - <test.js