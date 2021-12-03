redis-server --appendonly yes --appendfsync always --daemonize yes && sleep 1

redis-cli < /entry_data/initialize_data.redis

redis-cli save

redis-cli shutdown

redis-server --appendonly yes --appendfsync always --daemonize no
