#You make my day
It's program - simplified UNIX cron for easy run in docker 


## How work with it?
All jobs configuring in crontab.conf.yml file.   
Run over docker compose:
``` 
git clone https://github.com/meklis/all-ok-cron.git
cd all-ok-cron
# Make changes in crontab.conf.yml and config you docker-compose file if necessary.
docker-compose up -d 
```

### Configuration file 
* cron.print_format - allowed parameters - log or stdout. Log print stdout and stderr from jobs in log format; stdout print all over fmt.Printf() with job name.    
* cron.jobs - crontab jobs
* logger.* - logger parameters 

### Example of  config file 
``` 
cron:
#Supported print format:  log | stdout
  print_format: stdout
  jobs:
  - name: Bar
    crontab: "* * * * *"
    job: /usr/bin/php /test/script.php sh 1
    print_output: true
  - name: Startup
    crontab: "@reboot"
    job: /usr/bin/php /test/script.php sh 1
    print_output: true
  - name: Foo
    crontab: "* * * * *"
    job: /usr/bin/php /test/script.php sh 1
    print_output: true


logger:
  console:
    enabled: true
    enable_color: false
    log_level: 6
    print_file: false

```

#Crontab
## Crontab syntax <a id="syntax"></a>

If you are not faimiliar with crontab syntax you might be better off with other packages for scheduling tasks. But if you are familiar with Linux and crontab, this package might be right for you.

Here are the few quick references about crontab simple but powerful syntax.

```
*     *     *     *     *        

^     ^     ^     ^     ^
|     |     |     |     |
|     |     |     |     +----- day of week (0-6) (Sunday=0)
|     |     |     +------- month (1-12)
|     |     +--------- day of month (1-31)
|     +----------- hour (0-23)
+------------- min (0-59)
```

### Examples

+ `* * * * *` run on every minute
+ `10 * * * *` run at 0:10, 1:10 etc
+ `10 15 * * *` run at 15:10 every day
+ `* * 1 * *` run on every minute on 1st day of month
+ `0 0 1 1 *` Happy new year schedule
+ `0 0 * * 1` Run at midnight on every Monday

### Lists

+ `* 10,15,19 * * *` run at 10:00, 15:00 and 19:00
+ `1-15 * * * *` run at 1, 2, 3...15 minute of each hour
+ `0 0-5,10 * * *` run on every hour from 0-5 and in 10 oclock

### Steps
+ `*/2 * * * *` run every two minutes
+ `10 */3 * * *` run every 3 hours on 10th min
+ `0 12 */2 * *` run at noon on every two days
+ `1-59/2 * * * *` run every two minutes, but on odd minutes



P.S. Thanks to Miloš Mileusnić (https://github.com/mileusna) for your work.