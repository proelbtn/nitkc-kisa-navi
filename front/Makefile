TAG := proelbtn/school-eve-navi-front
BARGS := --tag ${TAG}
RARGS := --interactive --tty --publish 8080:8080

all:

initialize:
	yarn

build:
	yarn run build
	DOCKER_BUILDKIT=1 docker build ${BARGS} .

run:
	docker run ${RARGS} ${TAG}
