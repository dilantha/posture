build:
	env GOOS=linux GOARCH=arm GOARM=7 go build

clean:
	rm posture

deploy: build
	ansible-playbook -i hosts deploy.yaml -vv
