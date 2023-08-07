docker run -d --name yugabyte -p7000:7000 -p9000:9000 -p5433:5433 -p9042:9042 yugabytedb/yugabyte:2.19.0.0-b190 bin/yugabyted start \
 --base_dir=/tmp/ybd \
 --master_flags "ysql_num_shards_per_tserver=4" \
 --tserver_flags "ysql_num_shards_per_tserver=4,follower_unavailable_considered_failed_sec=30" \
 --daemon=false

docker run -d --name yugabyte -p7000:7000 -p9000:9000 -p5433:5433 -p9042:9042 yugabytedb/yugabyte:2.14.11.0-b35 bin/yugabyted start \
 --base_dir=/tmp/ybd \
 --master_flags "ysql_num_shards_per_tserver=4" \
 --tserver_flags "ysql_num_shards_per_tserver=4,follower_unavailable_considered_failed_sec=30" \
 --daemon=false


