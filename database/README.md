# 初回だけ

docker volume create pq_todo

# 起動（コンテナイメージ作成処理で init フォルダ以下が/docker-entrypoint-initdb.d/にコピーされ、create と copy が実行される）

docker-compose up --build

# 停止

#docker-compose down
