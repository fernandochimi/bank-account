DCMP = docker-compose

clean:
		find . -name "*.pyc" -exec rm -rf {} \;
		find . -type d -empty -delete;
		rm -rf htmlcov/
		rm -rf .coverage
		rm -rf *.log

build:
		${DCMP} build

start:
		${DCMP} up

startd:
		${DCMP} up -d

prune:
		docker container prune -f

stop:
		${DCMP} stop

execute:
		${MAKE} clean
		${MAKE} build
		${MAKE} start
		${MAKE} prune