# 小程序商城后台

## 环境搭建

---

- 克隆代码

```bash
git clone https://github.com/hwdhy/shop_master_go.git
```

- 下载依赖包

```bash
go mod tidy
```

- 创建数据库

```bash
CREATE DATABASE shop;
```

- 导入数据

```bash
将pulic.sql文件导入数据库
```

- 配置小程序数据

```bash
修改config下的config.go,改为自己相应的配置
```

- 启动项目

```bash
go run main.go
```

