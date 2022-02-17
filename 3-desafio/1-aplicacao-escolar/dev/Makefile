STACK_NAME=school-dev

start:
	docker-compose -p ${STACK_NAME} up -d

stop:
	docker-compose -p ${STACK_NAME} stop

clean:
	docker-compose -p ${STACK_NAME} down -v

logs:
	docker-compose -p ${STACK_NAME} logs -f
