import Mock from 'mockjs'

export const userLogin = Mock.mock({
  code: '200',
  message: '成功',
  data: {
       token:'123456',
       authList:[{
        name: '一级菜单',
        id: 1,
        pid: -1,
        auth: 'app-center',
        path: 'appMenu',
      },
      {
        name: '二级菜单',
        id: 2,
        pid: 1,
        path: '/2',
      },
      {
        name: '三级菜单',
        id: 3,
        pid: 2,
        path: '/3',
      }],
       role:'admin',
       username:'田永刚'
  },
})
export const userValidate = Mock.mock({
  code: '200',
  message: '成功',
  data: {
       token:'100000000000',
       authList:[{
        name:'一级菜单',
        id:1,
        pid:-1,
        auth:'appMenu'
      },
      {
       name:'二级菜单',
        id:2,
        pid:1,
        path:'/'
     }],
       role:'admin',
       username:'田永刚'
  },
})
export const userlogout = Mock.mock({
  code: '200',
  message: '退出成功',
})