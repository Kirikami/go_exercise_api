Using:

	1. Go to check out directory.
	2. Run `brew install glide`
	3. Run `glide install`
	4. Install mysql database
	5. Create database
	6. Change config in "config.json" and "dbconf.yml" for accessing to database
	7. Compile sources.
	8. Run server `go_exercise_api -config=<path_to_config_file>
	9. Open in browser `localhost:<port_from_config_file>/auth?provider=facebook`
	10. Copy JWT
	11. Send request for saving task, for example:
		
		`curl localhost:1223/v1/task -H "Authorization: Bearer <JWT>" -H "Content-Type: application/json" -X POST -d '{"title":"new task", "description":"new description", "priority": 6}'`
	
	12. Send request for updating task, for example:

		`curl 'localhost:1223/v1/task/1' -H "Authorization: Bearer <JWT>" -H "Content-Type: application/json" -X PUT -d '{"title":"updated task", "description":"updated description", "priority": 3}'`
	
	13. Send request for deleting task, for example: 

		`curl 'localhost:1223/v1/task/1' -H "Authorization: Bearer <JWT>" -H "Content-Type: application/json" -X DELETE`
	
	14. Send request for getting task, for example:

		`curl 'localhost:1223/v1/task/1 -H "Authorization: Bearer <JWT> -H "Content-Type: application/json" -X GET`

	15. Send request for getting all tasks, for example:

		`curl localhost:1223/v1/tasks -H "Authorization: Bearer <JWT> -H "Content-Type: application/json" -X GET`

Using with docker:
	1. Install docker
	2. Install docker-compose
	3. Go to checkout directory.
	5. Run `docker-compose build`
	6. Run `docker-compose up -d`
	7. Wait until server will go up.
	8. Steps 11-15 same as for default instalation address for accesing localhost:5000
