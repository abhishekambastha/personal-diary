#!/bin/bash

curl --request POST -H "Content-Type: application/json" --data @entry-post.json  http://localhost:8080/v1/entry/
