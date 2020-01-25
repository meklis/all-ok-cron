FROM meklis/all-ok-billing
LABEL maintainer="Max Boyar <max.boyar.a@gmail.com>"
RUN apt update && apt -y upgrade
ADD https://github.com/meklis/all-ok-cron/releases/download/1.0.0/all-ok-cron-amd64 /opt/all-ok-cron
COPY crontab.conf.yml /opt/crontab.conf.yml
RUN chmod +x /opt/all-ok-cron
ENTRYPOINT ["/opt/all-ok-cron", "-c", "/opt/crontab.conf.yml"]

