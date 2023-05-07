start-db:
	pg_ctl -D $(PREFIX)/var/lib/postgresql start

stop-db:
	pg_ctl -D $(PREFIX)/var/lib/postgresql stop

inspect-db:
	psql -U erikrios -d blog -W

test:
	go test -v -cover -count=1 ./...
