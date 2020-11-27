
> fork from https://github.com/huacnlee/gobackup  && and cron job && docker image
 

 ## 参考文档

详细中文介绍： https://ruby-china.org/topics/34094

https://gobackup.github.io/

### Databases

- MySQL
- PostgreSQL
- Redis - `mode: sync/copy`
- MongoDB


### Cron Task Running

current cron with global config
### Archive

Use `tar` command to archive many file or path into a `.tar` file.

### Compressor

- Tgz - `.tar.gz`
- Uncompressed - `.tar`

### Encryptor

- OpenSSL - `aes-256-cbc` encrypt

### Storages

- Local
- FTP
- SCP - Upload via SSH copy
- [Amazon S3](https://aws.amazon.com/s3)
- [Alibaba Cloud Object Storage Service (OSS)](https://www.alibabacloud.com/product/oss)

```bash
$ gobackup -h
NAME:
   gobackup - Easy full stack backup operations on UNIX-like systems

USAGE:
   gobackup [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     perform
     start
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Configuration

GoBackup will seek config files in:

- ./gobackup.yml
- ~/.gobackup/gobackup.yml
- /etc/gobackup/gobackup.yml

Example config: [gobackup_test.yml](https://github.com/huacnlee/gobackup/blob/master/gobackup_test.yml)

```yml
models:
  gitlab:
    compress_with:
      type: tgz
    store_with:
      type: scp
      path: ~/backup
      host: your-host.com
      private_key: ~/.ssh/id_rsa
      username: ubuntu
      password: password
      timeout: 300
    databases:
      gitlab:
        type: mysql
        host: localhost
        port: 3306
        database: gitlab_production
        username: root
        password:
        additional_options: --single-transaction --quick
      gitlab_redis:
        type: redis
        mode: sync
        rdb_path: /var/db/redis/dump.rdb
        invoke_save: true
        password:
    archive:
      includes:
        - /home/git/.ssh/
        - /etc/mysql/my.conf
        - /etc/nginx/nginx.conf
        - /etc/nginx/conf.d
        - /etc/redis/redis.conf
        - /etc/logrotate.d/
      excludes:
        - /home/ubuntu/.ssh/known_hosts
        - /etc/logrotate.d/syslog
  gitlab_repos:
    store_with:
      type: local
      path: /data/backups/gitlab-repos/
    archive:
      includes:
        - /home/git/repositories
```

## Usage

```bash
$ gobackup perform
2017/09/08 06:47:36 ======== ruby_china ========
2017/09/08 06:47:36 WorkDir: /tmp/gobackup/1504853256396379166
2017/09/08 06:47:36 ------------- Databases --------------
2017/09/08 06:47:36 => database | Redis: mysql
2017/09/08 06:47:36 Dump mysql dump to /tmp/gobackup/1504853256396379166/mysql/ruby-china.sql
2017/09/08 06:47:36

2017/09/08 06:47:36 => database | Redis: redis
2017/09/08 06:47:36 Copying redis dump to /tmp/gobackup/1504853256396379166/redis
2017/09/08 06:47:36
2017/09/08 06:47:36 ----------- End databases ------------

2017/09/08 06:47:36 ------------- Compressor --------------
2017/09/08 06:47:36 => Compress with Tgz...
2017/09/08 06:47:39 -> /tmp/gobackup/2017-09-08T14:47:36+08:00.tar.gz
2017/09/08 06:47:39 ----------- End Compressor ------------

2017/09/08 06:47:39 => storage | FTP
2017/09/08 06:47:39 -> Uploading...
2017/09/08 06:47:39 -> upload /ruby_china/2017-09-08T14:47:36+08:00.tar.gz
2017/09/08 06:48:04 Cleanup temp dir...
2017/09/08 06:48:04 ======= End ruby_china =======
```

## Backup schedule

You may want run backup in scheduly, you need config scheduler:

example 

```code
scheduler:
  cron: "0 0/2 * * *"
models:
  first:
    compress_with:
      type: tgz
    store_with:
      type: local
      keep: 10
      path: /Users/dalong/mylearning/db-back/gobackup/mydb
    databases:
      demo-db:
        type: mysql
        host: 127.0.0.1
        port: 3306
        database: demo
        dumpPath: 
        username: root
        password: dalong
  second:
    compress_with:
      type: tgz
    store_with:
      type: s3
      keep: 20
      bucket: demo
      endpoint: http://localhost:9000
      path: backups
      access_key_id: minio
      secret_access_key: minio123
    databases:
      mydemodb:
        type: mysql
        host: 127.0.0.1
        port: 3306
        database: demo
        username: root
        password: dalong
```
## License

MIT
