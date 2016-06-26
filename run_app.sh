#!/usr/bin/env bash

echo 'Running server, waiting for database'

count=0

while [[ true ]] 
do
   if [[ $count == 60 ]]; then
	   exit 1
      echo "End of loop, error accessing database"
    break
    fi

   goose -env=docker up

   if [ $? == 0 ]
     then
     break
   fi
   echo "waiting for $count seconds..."
   ((count++))

   sleep 1
done

go_exercise_api -config=/go/src/github.com/kirikami/go_exercise_api/images/go/config.json
