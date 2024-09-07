

rahulganjave@Rahuls-MBP GoRMSValidationAPI % mkdir -p db/migration
rahulganjave@Rahuls-MBP GoRMSValidationAPI % migrate create -ext sql -dir db/migration -seq init_schema

rahulganjave@Rahuls-MBP GoRMSValidationAPI % docker exec -it postgres14.5 /bin/sh

># createdb --username=root --owner=root MiddleOffice
># psql MiddleOffice
># dropdb MiddleOffice
># exit

rahulganjave@Rahuls-MBP GoRMSValidationAPI % docker exec -it postgres14.5 createdb --username=root --owner=root MiddleOffice

rahulganjave@Rahuls-MBP GoRMSValidationAPI % docker exec -it postgres14.5 psql -U root MiddleOffice

rahulganjave@Rahuls-MBP GoRMSValidationAPI % make createdb
rahulganjave@Rahuls-MBP GoRMSValidationAPI % make migrateup

rahulganjave@Rahuls-MBP GoRMSValidationAPI % go mod init github.com/pc/GoRMSValidationAPI
rahulganjave@Rahuls-MBP GoRMSValidationAPI % go mode tidy

rahulganjave@Rahuls-MBP GoRMSValidationAPI % go get github.com/stretchr/testify

https://github.com/kyleconroy/sqlc


Install gin framework
https://github.com/gin-gonic/gin

rahulganjave@Rahuls-MBP GoRMSValidationAPI % go get -u github.com/gin-gonic/gin

----

Viper: https://github.com/spf13/viper

rahulganjave@Rahuls-MBP GoRMSValidationAPI % go get github.com/spf13/viper

----
