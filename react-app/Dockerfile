FROM node:12.14.1 as builder
# 作業フォルダを作成しnpm設定
RUN mkdir /usr/src/app
WORKDIR /usr/src/app
ENV PATH /usr/src/app/node_modules/.bin:$PATH
# ソースを作業フォルダにコピーして、ビルド
COPY . /usr/src/app/
RUN npm install --silent
RUN npm run build

####################
FROM nginx:1.17.9-alpine
# nginxの設定を削除して、アプリで設定したファイルをコピー
COPY ./default.template.ini /etc/nginx/conf.d/
COPY ./setup.sh /etc/nginx/conf.d/
# 上記の生成されたアプリのビルド成果物をnginxのサンプルアプリが使用していたフォルダにコピー
COPY --from=builder /usr/src/app/build /usr/share/nginx/html
#環境変数
ENV NGINX_LISTEN_PORT 80
ENV BACKEND_PORT http://localhost:8081
#実行権限
RUN chmod +x /etc/nginx/conf.d/setup.sh
# 80ポートオープンしてnginx
EXPOSE 80
CMD ["/etc/nginx/conf.d/setup.sh"]
