server {
    listen       {{.Port}};
    server_name   {{.Server_name}};

    charset utf-8;
    access_log  {{.Access_log}};

    location /(css|js|fonts|img)/ {
        access_log off;
        expires {{.Expires}};

        root "{{.Root}}";
        try_files $uri @backend;
    }

    location / {
        try_files /_not_exists_ @backend;
    }

    location @backend {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;

        proxy_pass {{.Proxy_pass}};
    }
}