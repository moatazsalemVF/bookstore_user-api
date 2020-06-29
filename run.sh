export os_env_mysql_users_username=go
export os_env_mysql_users_password=go
export os_env_mysql_users_host=localhost:3306
export os_env_mysql_users_schema=users_db

go build
echo "Built"
./bookstore_user-api
