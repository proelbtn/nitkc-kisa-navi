TAG := proelbtn/school-eve-navi-front
BARGS := --tag ${TAG}
RARGS := --interactive --tty --publish 3000:80

all:

initialize:
	yarn

build:
	NODE_ENV=dev API_URL=/ yarn run build
	DOCKER_BUILDKIT=1 docker build ${BARGS} .

run:
	docker run ${RARGS} ${TAG}
