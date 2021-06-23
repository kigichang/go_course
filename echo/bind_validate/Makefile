.PHONY: clean test

HOST=127.0.0.1:8080
API_VER=api/v1

users:
	curl --include http://${HOST}/${API_VER}/users
	curl --include http://${HOST}/${API_VER}/users?id=1

user:
	curl --include http://${HOST}/${API_VER}/user
	curl --include http://${HOST}/${API_VER}/user/0
	curl --include http://${HOST}/${API_VER}/user/1
	curl --include http://${HOST}/${API_VER}/user/1000

adduser1:
	curl --include --request POST \
		--header "Content-type: application/json" \
		--data '{"email":"joe@labstack.com"}' \
		http://${HOST}/${API_VER}/users

adduser2:
	curl --include --request POST \
		-d "name=Joe" \
		-d "email=joe@labstack.com" \
		http://${HOST}/${API_VER}/users