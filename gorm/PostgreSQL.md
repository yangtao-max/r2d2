## PostgreSQL

### install 安装
`brew install postgresql`

### init 初始化环境
`initdb /usr/local/var/postgres -E utf8`
> 指定 "/usr/local/var/postgres" 为 PostgreSQL 的配置数据存放目录，并且设置数据库数据编码是 utf8

### pg_ctl 启动PostgreSQL
`pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log start`

### pg_ctl 关闭PostgreSQL
`pg_ctl -D /usr/local/var/postgres stop -s -m fast`

### createuser 创建用户
`createuser pg_yt -P`
>Enter password for new role: 
>
>Enter it again:`

### createdb 创建数据库
`createdb ytdatebase -O pg_yt -E UTF8 -e`
>创建了一个名为 ytdatebase 的数据库，并指定 pg_yt 为改数据库的拥有者（owner）
>，数据库的编码（encoding）是 UTF8，参数 "-e" 是指把数据库执行操作的命令显示出来

### 连接数据库
`psql -U pg_yt -d ytdatebase -h 127.0.0.1`

### 数据库列表
`psql -l`

### 删除数据库
`psql -U pg_yt ytdatebase`

### SHOW databases,SHOW tables,Describe table
``` sql
//show databases
SELECT datname from pg_database;
//show tables
SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';
//desc table
SELECT column_name FROM information_schema.columns WHERE table_name ='table_name';
```