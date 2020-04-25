#!/bin/sh
    envsubst '$$NGINX_LISTEN_PORT$$BACKEND_PORT' < /etc/nginx/conf.d/default.template.ini > /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'
