# 开发前端的规范

```
1、编写代码需要写注释，页面只是接口的请求和数据的展示（tsx，setup，以及setup用法糖格式不限制），页面逻辑过多，拆除方法到ts文件中
2、后端请求接口需要写到vuex中，相当于增加一个逻辑层，多个module之间可以互相引用属性和方法
3、css 颜色，尽量定义变量的形式，可以使用在css，js，模板中。
4、api文件夹里面是请求的方法和URL地址，命名尽量采取和modules文件夹下面统一
5、*.vue文件,components文件夹组件下面采用大驼峰命名
6、需要添加路由在router中，添加*.router.ts
7、书写格式eslint，prettier统一格式化
8、编写组件时，component写注释，譬如某个属性什么含义，复杂的可以注释参考xxx.vue
9、编写请求的需要定义一个常量去引用，参考 : action-types.ts
10、及时向开发要数据格式，采用mockjs进行开发，使用参考mock文件夹
```

# 自定义 ecahrts 使用方法

- 1. 首先引入 ecahrts import publicChart from '@/components/common/publicChart.vue'
     然后在 components 里面注册
- 2. 需要给自定义 echarts 传递一个对象 对象里面需要 echarts 的 option 配置(chartDate)和 自定义的标题(title) 和 tabs 切换(selectTime)
- 3. 子组件给父组件传递数据可以在 组件上面 @changeDate="getSelectDate" 然后在 setup 里面定义函数接受参数
     @changeDate="getSelectDate"
     const getSelectDate = (val) => {
     console.log(val, 'val')
     }
     ```javascript
     const stateOption = reactive({
       chartDate: {
         xAxis: {
           type: 'category',
           data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
         },
         yAxis: {
           type: 'value',
         },
         series: [
           {
             data: [120, 200, 150, 80, 70, 110, 130],
             type: 'bar',
             showBackground: true,
             backgroundStyle: {
               color: 'rgba(180, 180, 180, 0.2)',
             },
           },
         ],
       },
       title: '工单状态统计',
       selectTime: ['日', '周', '月'],
       isShowTabs: true, // 是否显示切换
     })
     ```
