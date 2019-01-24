# simple run
docker run --name some-nginx -v /home/trak/Documents/science/nginx/share/content:/usr/share/nginx/html:ro -d -p 80:80 nginx


# 
docker run --name some-nginx -v /home/trak/Documents/science/nginx/share/content:/usr/share/nginx/html:ro -v /home/trak/Documents/science/nginx/share/etc/nginx.conf:/etc/nginx/nginx.conf:ro -v /tmp/nginx_docker:/var/log/nginx:rw -d -p 80:80 nginx




### konfiguracja 
na localhost regulka jest w default-owej konfiguracji
na www.aaa.wp.pl jest regulka w nginx.conf
