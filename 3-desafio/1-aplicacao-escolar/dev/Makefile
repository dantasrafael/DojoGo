STACK_NAME=school-dev

start: build
	docker-compose -p ${STACK_NAME} up -d

stop:
	docker-compose -p ${STACK_NAME} stop

clean:
	docker-compose -p ${STACK_NAME} down -v
	docker rmi modulo-escolar financial

logs:
	docker-compose -p ${STACK_NAME} logs -f

build:
	docker buildx build -t modulo-escolar ../../ --build-arg SRC_PATH=/1-aplicacao-escolar/modulo-escolar
	docker buildx build -t financial ../../ --build-arg SRC_PATH=/1-aplicacao-escolar/financial
