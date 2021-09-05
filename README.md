# Dockerized starter template for Go + My SQL project.

- /docker/mysql/my.cnfの追加
例
参考にしたサイト：https://qiita.com/ucan-lab/items/b094dbfc12ac1cbee8cb#mycnf
```
# MySQLサーバーへの設定
[mysqld]
# 文字コード/照合順序の設定
character-set-server = utf8mb4
collation-server = utf8mb4_bin

# タイムゾーンの設定
default-time-zone = SYSTEM
log_timestamps = SYSTEM

# デフォルト認証プラグインの設定
default-authentication-plugin = mysql_native_password

# エラーログの設定
log-error = /var/log/mysql/mysql-error.log

# スロークエリログの設定
slow_query_log = 1
slow_query_log_file = /var/log/mysql/mysql-slow.log
long_query_time = 5.0
log_queries_not_using_indexes = 0

# 実行ログの設定
general_log = 1
general_log_file = /var/log/mysql/mysql-query.log

# mysqlオプションの設定
[mysql]
# 文字コードの設定
default-character-set = utf8mb4

# mysqlクライアントツールの設定
[client]
# 文字コードの設定
default-character-set = utf8mb4
```
- /config.iniの追加
例
```
[log]
log_file = godev.log

[db]
db_protocol = tcp
db_address = db
sql_driver = mysql
db_port = 3306
db_name = godev
db_user_name = TaroYamada
db_password = hogehoge

[app]
app_port = 8080
```
- /.envの追加
例
```
API_PORT=8080
DB_NAME=godev
DB_USER=TaroYamada
DB_PASS=hogehoge
TZ=Asia/Tokyo
DB_PORT=3306
```