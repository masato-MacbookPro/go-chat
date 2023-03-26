MYSQL_INFO:=-h 127.0.0.1 -P 3306 -u root
DB_NAME:=chat
DML_DIR:=./migrations/dummy

# ローカルデータ挿入のコマンド
.PHONY: seed
seed: ## seed
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_users.up.sql

# ローカルデータDLETEのコマンド
.PHONY: delete
delete: ## delete
	mysql $(MYSQL_INFO) $(DB_NAME) < $(DML_DIR)/dummy_users.down.sql