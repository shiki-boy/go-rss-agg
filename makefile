server: 
	@nodemon

kill-server:
	@sudo kill -9 `sudo lsof -t -i:8000`
