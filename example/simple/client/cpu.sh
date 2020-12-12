#!/bin/bash

rm client
rm cpu_profile

go build 

./client 

go tool pprof client cpu_profile
