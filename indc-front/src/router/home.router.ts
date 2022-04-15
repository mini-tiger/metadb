//静态的路由
import * as types from '@/store/action-types'
import store from '@/store'
export default [
  {
    path: '/demo',
    redirect: '/demo/search',
    component: () =>
      import(/* webpackChunkName:'home' */ '@/views/demo/test.vue'),
    children: [
      {
        name: 'testSearch',
        path: '/demo/search',
        component: () => import('@/views/demo/testSearch.vue'),
      },
      {
        name: 'testButton',
        path: '/demo/button',
        component: () => import('@/views/demo/testButton.vue'),
      },
      {
        name: 'testCanvas',
        path: '/demo/canvas',
        component: () => import('@/views/demo/testCanvas.vue'),
      },
      {
        name: 'testTable',
        path: '/demo/table',
        component: () => import('@/views/demo/testTable.vue'),
      },
    ],
  },

  {
    path: '/',
    meta: {
      needLogin: true,
    },
    name: 'home',
    redirect: '/workspace',
    component: () =>
      import(/* webpackChunkName:'home' */ '@/views/home/Home.vue'),
    children: [
      {
        path: '/realTimeAlarm',
        meta: {
          needLogin: true,
          title: '监控列表',
        },
        component: () =>
          import(
            /* webpackChunkName:'cl' */ '@/views/app/monitoringCenter/AlarmList/index.vue'
          ),
      },
      {
        path: '/AlarmRuleConfiguration',
        name: 'AlarmRuleConfiguration',
        meta: {
          needLogin: true,
          title: '告警规则配置',
        },
        component: () =>
          import(
            /* webpackChunkName:'workspace' */ '@/views/app/monitoringCenter/AlarmRuleConfiguration/index.vue'
          ),
      },
      {
        path: '/NotificationRuleConfiguration',
        name: 'NotificationRuleConfiguration',
        meta: {
          needLogin: true,
          title: '告警规则配置',
        },
        component: () =>
          import(
            /* webpackChunkName:'workspace' */ '@/views/app/monitoringCenter/NotificationRuleConfiguration/index.vue'
          ),
      },
      {
        path: '/ShieldingAlarmConfiguration',
        name: 'ShieldingAlarmConfiguration',
        meta: {
          needLogin: true,
          title: '告警屏蔽规则配置',
        },
        component: () =>
          import(
            /* webpackChunkName:'workspace' */ '@/views/app/monitoringCenter/ShieldingAlarmConfiguration/index.vue'
          ),
      },
      {
        path: '/VoiceConfiguration',
        name: 'VoiceConfiguration',
        meta: {
          needLogin: true,
          title: '语音配置',
        },
        component: () =>
          import(
            /* webpackChunkName:'workspace' */ '@/views/app/monitoringCenter/VoiceConfiguration/index.vue'
          ),
      },
      {
        path: '/workspace',
        meta: {
          needLogin: true,
        },
        component: () =>
          import(
            /* webpackChunkName:'workspace' */ '@/views/home/Workspace.vue'
          ),
      },
      {
        path: '/report-center',
        meta: {
          needLogin: true,
        },
        component: () =>
          import(
            /* webpackChunkName:'report' */ '@/views/report/ReportCenter.vue'
          ),
      },
      {
        path: '/search-center',
        meta: {
          needLogin: true,
        },
        component: () =>
          import(
            /* webpackChunkName:'search' */ '@/views/search/SearchCenter.vue'
          ),
      },
      {
        path: '/app-center',
        name: 'appCenter',
        meta: {
          needLogin: true,
          title: '应用中心',
        },
        component: () =>
          import(
            /* webpackChunkName:'AppRender' */ '@/views/app/AppRender.vue'
          ),
        redirect: '/app-center/app-home',
        children: [
          {
            path: 'app-home',
            meta: {
              needLogin: true,
            },
            component: () =>
              import(
                /* webpackChunkName:'AppHome' */ '@/views/app/AppHome.vue'
              ),
          },
          {
            path: 'app-menu',
            name: 'appMenu',
            meta: {
              needLogin: true,
              title: '监控中心',
            },
            beforeEnter: async (to, from, next) => {
              // if ((store.state as any).user.hasPermission) {
              //   if (!(store.state as any).user.menuPermission) {
              await store.dispatch(`app/${types.APP_CENTER_MENU}`)
              await store.dispatch(`user/${types.ADD_ROUTES}`)
              //     next({ ...to, replace: true })
              //   } else {
              //     next()
              //   }
              // } else {
              //   next()
              // }
              next()
            },
            component: () =>
              import(
                /* webpackChunkName:'search' */ '@/views/app/AppContent.vue'
              ),
            children: [
              {
                path: 'cl-example',
                meta: {
                  needLogin: true,
                  title: 'cl实例',
                },
                component: () =>
                  import(
                    /* webpackChunkName:'realcolumn' */ '@/views/app/dataConfiguration/cl/realcolumn/index.vue'
                  ),
              },
              // {
              //   path: 'clExampleLook',
              //   meta: {
              //     needLogin: true,
              //     title: 'cl实例查看',
              //   },
              //   component: () =>
              //     import(
              //       /* webpackChunkName:'clExampleLook' */ '@/views/app/dataConfiguration/cl/realcolumn/cpns/RealcolumnLook.vue'
              //     ),
              // },
              {
                path: 'operationAudit',
                meta: {
                  needLogin: true,
                  title: '操作审计',
                },
                component: () =>
                  import(
                    /* webpackChunkName:'statement' */ '@/views/app/dataConfiguration/statement/OperationAudit.vue'
                  ),
              },
              {
                path: 'modelClassify',
                meta: {
                  needLogin: true,
                  title: 'CI模型分类管理',
                },
                component: () =>
                  import(
                    /* webpackChunkName:'cl' */ '@/views/app/dataConfiguration/model/modelClassify/index.vue'
                  ),
              },
              {
                path: 'relationConfig',
                meta: {
                  needLogin: true,
                  title: '关系类型配置',
                },
                component: () =>
                  import(
                    /* webpackChunkName:'relation' */ '@/views/app/dataConfiguration/configuration/RelationConfig.vue'
                  ),
              },
              {
                path: 'modelDetail',
                name: 'modelDetail',
                meta: {
                  needLogin: true,
                  title: 'CI模型详情',
                },
                component: () =>
                  import(
                    /* webpackChunkName:'realcolumn' */ '@/views/app/dataConfiguration/model/clModel/ModelDetail.vue'
                  ),
              },
            ],
          },
        ],
      },
      {
        path: 'cl-exampleLook',
        name: 'cl-exampleLook',
        meta: {
          needLogin: true,
          title: 'cl实例查看',
        },
        component: () =>
          import(
            /* webpackChunkName:'clExampleLook' */ '@/views/app/dataConfiguration/cl/realcolumn/cpns/RealcolumnLook.vue'
          ),
      },
      {
        path: 'cl-example',
        meta: {
          needLogin: true,
          title: 'cl实例',
        },
        component: () =>
          import(
            /* webpackChunkName:'realcolumn' */ '@/views/app/dataConfiguration/cl/realcolumn/index.vue'
          ),
      },
      {
        path: 'cl-relation',
        meta: {
          needLogin: true,
          title: 'cl关系',
        },
        component: () =>
          import(
            /* webpackChunkName:'realcolumn' */ '@/views/app/dataConfiguration/cl/relation/index.vue'
          ),
      },
      {
        path: 'model-relation',
        meta: {
          needLogin: true,
          title: 'cl模型关系',
        },
        component: () =>
          import(
            /* webpackChunkName:'realcolumn' */ '@/views/app/dataConfiguration/model/modelRelation/index.vue'
          ),
      },
      {
        path: 'model',
        meta: {
          needLogin: true,
          title: 'cl模型',
        },
        component: () =>
          import(
            /* webpackChunkName:'realcolumn' */ '@/views/app/dataConfiguration/model/clModel/index.vue'
          ),
      },
      {
        path: 'modelDetail',
        name: 'modelDetail',
        meta: {
          needLogin: true,
          title: 'CI模型详情',
        },
        component: () =>
          import(
            /* webpackChunkName:'realcolumn' */ '@/views/app/dataConfiguration/model/clModel/ModelDetail.vue'
          ),
      },
      {
        path: 'relationConfig',
        meta: {
          needLogin: true,
          title: '关系类型配置',
        },
        component: () =>
          import(
            /* webpackChunkName:'relation' */ '@/views/app/dataConfiguration/configuration/RelationConfig.vue'
          ),
      },
      {
        path: 'modelClassify',
        meta: {
          needLogin: true,
          title: 'CI模型分类管理',
        },
        component: () =>
          import(
            /* webpackChunkName:'cl' */ '@/views/app/dataConfiguration/model/modelClassify/index.vue'
          ),
      },
      {
        path: 'operationAudit',
        meta: {
          needLogin: true,
          title: '操作审计',
        },
        component: () =>
          import(
            /* webpackChunkName:'statement' */ '@/views/app/dataConfiguration/statement/OperationAudit.vue'
          ),
      },
      {
        path: 'emitTopology',
        meta: {
          needLogin: true,
          title: 'cl关系编辑',
        },
        component: () =>
          import(
            /* webpackChunkName:'statement' */ '@/views/app/dataConfiguration/cl/relation/emitTopology.vue'
          ),
      },
    ],
  },
  {
    path: '/login',
    component: () => import(/* webpackChunkName:'login' */ '@/views/Login.vue'),
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () =>
      import(/* webpackChunkName:'NotFound' */ '@/components/NotFound.vue'),
  },
]
