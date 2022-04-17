#### WEB INIT

> ANT安装指南 ref https://pro.ant.design/docs/getting-started/
>
> 前置需求：nvm -- node版本管理工具 https://www.jianshu.com/p/6c32d2078a2d
>
> nvm v
> nvm list available
>
> 安装指定版本 [太高版本可能会出错]
> nvm install 14.18.3
>
> 查看当前node版本：
>
> nvm list
>
> 使用指定版本node:
>
> nvm use 14.18.3 -- 需要在管理员模式下
```shell
# 在管理员模式下
cd /root/
# 安装node，可选
nvm install 14.18.3
nvm use 14.18.3
nvm list
# 按照ant
yarn create umi web
cd web
npm install
# 运行
npm run start
# 编译
yarn run build

```



#### 错误处理

+ npm run start 时，报错

> opensslErrorStack: [ 'error:03000086:digital envelope routines::initialization error~

可能是高版本node的原因

降版本：

```shell
nvm install 14.18.3
nvm use 14.18.4
nvm list

# 重新启动
npm run start
http://localhost:8000

```



#### 其他命令

- npm run analyze

> The analyze script does the same thing as build, but he opens a page showing your dependency information. If you need to optimize performance and package size, you need it.

- npm run lint

> We offer a range of lint scripts, including TypeScript, less, css, md files. You can use this script to see what problems your code has. In commit we automatically run the related lint.

- npm run lint:fix 

> ame as lint, but the lint error is fixed automatically.

- npm run i18n-remove

> This script will attempt to remove all i18n code from the project, which is not good for complex run-time code and is used with caution.

- 更多命令

> https://umijs.org/docs/cli

