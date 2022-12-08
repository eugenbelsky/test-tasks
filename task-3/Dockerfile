FROM nginx:alpine

COPY ./default.conf /etc/nginx/conf.d/default.conf
COPY ./html /opt/html

RUN chmod -R 755 /opt/html && \
        chown -R nginx:nginx /var/cache/nginx && \
        chown -R nginx:nginx /var/log/nginx && \
        chown -R nginx:nginx /etc/nginx/conf.d &&\
        chown -R nginx:nginx /opt/html
RUN touch /var/run/nginx.pid && \
        chown -R nginx:nginx /var/run/nginx.pid
