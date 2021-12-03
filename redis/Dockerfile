FROM redis:latest
COPY initialize_data.redis /entry_data/
COPY entrypoint.sh /entry_data/
CMD ["sh", "/entry_data/entrypoint.sh"]