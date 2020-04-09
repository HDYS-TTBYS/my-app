FROM nginx:1.17.9-alpine
# nginxの設定を削除して、アプリで設定したファイルをコピー
COPY . /etc/nginx/conf.d/
#環境変数
ENV NGINX_LISTEN_PORT 80
ENV BACKEND_PORT http://go_app:8080
#実行権限
RUN chmod +x /etc/nginx/conf.d/setup.sh
# 80ポートオープンしてnginx
EXPOSE 80
CMD ["/etc/nginx/conf.d/setup.sh"]
