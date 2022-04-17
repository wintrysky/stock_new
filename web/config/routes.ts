export default [
  {
    path: '/user',
    layout: false,
    routes: [
      { path: '/user', routes: [{ name: '登录', path: '/user/login', component: './user/Login' }] },
      { component: './404' },
    ],
  },
  {
    path: '/',
    redirect: '/import',
  },
  {
    name: '上传基础数据',
    icon: 'table',
    path: '/import',
    component: './import',
  },
  { name: '查询', icon: 'table', path: '/stocks', component: './stocks' },
  { component: './404' },
];
