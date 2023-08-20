include services/admin/Makefile

.PHONY: env
env:
	cd deploy/docker && docker-compose up -d

