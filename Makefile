# Build an executable for running go server after database migrations pass
# 
SHELL := /bin/bash

run-server:
	count=1 ; while [[ $$count -le 60  ]] ; do \
		if [[ "$$count" == "60" ]]; then \
			echo "End of loop, error accessing database"; \
			exit 1; \
		fi;\
		goose -env=docker up; \
		if [[ $$? == "0" ]]; then \
			break; \
		fi; \
		echo "waiting for $$count seconds..."; \
		((count = count + 1)); \
		sleep 1; \
	done; \
	go_exercise_api -config=/go/src/github.com/kirikami/go_exercise_api/images/go/config.json
