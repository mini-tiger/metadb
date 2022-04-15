<template>
  <!-- <router-link to="/app-center/app-menu">应用中心</router-link> -->
  <div class="application">
    <!-- 头部信息 -->
    <div class="application-header">
      <!-- 头部文本框 -->
      <div class="header-left">
        <el-input
          v-model="state.search"
          size="small"
          class="search-input"
          placeholder="请输入"
          :input-style="{ borderRadius: '15px' }"
        />
        <svg-icon
          icon-class="search"
          class-name="search-icon"
          @click="searchApp"
        />
      </div>
      <div class="header-right">
        <!-- 头部右侧切换 -->
        <div class="header-tabs">
          <div
            v-for="(item, index) in state.tabChange.tab"
            :key="index"
            class="tabChangeSty"
            v-bind:class="{ tabChangeStActive: index === state.current }"
            @click="changeTab(item, index)"
          >
            {{ item.name }}
          </div>
        </div>
        <!-- 头部右侧排序 -->
        <div class="store-app">
          <div
            v-for="(item, index) in state.tabChange.store"
            :key="index"
            class="change-store"
          >
            {{ item.name }}
            <svg-icon
              icon-class="descendingOrder"
              class-name="descending-order order"
            />
            <svg-icon
              icon-class="ascendingOrder"
              class-name="ascending-order order"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- 应用展示 -->
    <div class="application-core">
      <div
        class="application-cation"
        v-for="(item, index) in state.applicationShow"
        :key="index"
      >
        <div class="application-title">{{ item.title }}</div>
        <div class="application-second">
          <div
            class="every-last"
            v-for="(itm, ind) in item.data"
            :key="ind"
            @click="appStoreCenter"
          >
            <img :src="itm.img" alt="" />
            <p>{{ itm.name }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from 'vue'
import { useRouter } from 'vue-router'
import * as types from '@/store/action-types'
import { useStore } from 'vuex'
export default defineComponent({
  name: 'AppHome',
  setup() {
    const store = useStore()
    const route = useRouter()
    store.dispatch(`app/${types.APP_CENTER_MENU}`)
    const state = reactive({
      search: '',
      tabChange: {
        tab: [
          { name: '全部', id: 1 },
          { name: '监控中心', id: 2 },
          { name: '运维中心', id: 2 },
          { name: '运营中心', id: 4 },
        ],
        store: [
          { name: '更新时间', id: 'updateTime' },
          { name: '发布时间', id: 'releaseTime' },
          { name: '使用次数', id: 'usageTime' },
        ],
      },
      applicationShow: [
        {
          title: '常用应用',
          data: [
            {
              name: 'CMDB',
              route: 'app-center',
              id: 1,
              img: 'https://img0.baidu.com/it/u=1721391133,702358773&fm=253&fmt=auto&app=120&f=JPEG?w=500&h=625',
            },
            {
              name: 'CMDB',
              route: 'app-center',
              id: 2,
              img: 'https://img0.baidu.com/it/u=1721391133,702358773&fm=253&fmt=auto&app=120&f=JPEG?w=500&h=625',
            },
            {
              name: 'CMDB',
              route: 'app-center',
              id: 3,
              img: 'https://img0.baidu.com/it/u=1721391133,702358773&fm=253&fmt=auto&app=120&f=JPEG?w=500&h=625',
            },
          ],
        },
        {
          title: '监控中心',
          data: [
            {
              name: '动环监控',
              route: 'app-center',
              id: 1,
              img: 'https://img0.baidu.com/it/u=1721391133,702358773&fm=253&fmt=auto&app=120&f=JPEG?w=500&h=625',
            },
            {
              name: 'CMDB',
              route: 'app-center',
              id: 2,
              img: 'https://img0.baidu.com/it/u=1721391133,702358773&fm=253&fmt=auto&app=120&f=JPEG?w=500&h=625',
            },
            {
              name: 'CMDB',
              route: 'app-center',
              id: 3,
              img: 'https://img0.baidu.com/it/u=1721391133,702358773&fm=253&fmt=auto&app=120&f=JPEG?w=500&h=625',
            },
          ],
        },
        {
          title: '运维中心',
          data: [
            {
              name: '事件管理',
              route: 'app-center',
              id: 1,
              img: 'https://img0.baidu.com/it/u=1721391133,702358773&fm=253&fmt=auto&app=120&f=JPEG?w=500&h=625',
            },
            {
              name: 'CMDB',
              route: 'app-center',
              id: 2,
              img: 'https://img0.baidu.com/it/u=1721391133,702358773&fm=253&fmt=auto&app=120&f=JPEG?w=500&h=625',
            },
            {
              name: 'CMDB',
              route: 'app-center',
              id: 3,
              img: 'https://img0.baidu.com/it/u=1721391133,702358773&fm=253&fmt=auto&app=120&f=JPEG?w=500&h=625',
            },
          ],
        },
      ],
      current: 0,
    })
    const changeTab = (data, index) => {
      console.log(data, index)
      state.current = index
    }
    // 文本框搜索
    const searchApp = () => {
      console.log(state.search, '------>')
    }
    const appStoreCenter = () => {
      route.push('/app-center/app-menu')
    }
    return {
      state,
      changeTab,
      searchApp,
      appStoreCenter,
    }
  },
})
</script>
<style lang="scss" scoped>
.application {
  width: 100%;
  // height: 100%;
  padding: 0 10px;
  background: #f5f6fa;
}
.application-header {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}
.header-left {
  position: relative;
}
.search-icon {
  position: absolute;
  right: 10px;
  top: 5px;
}
.search-input {
  width: 220px;
  height: 30px;
}
.header-tabs {
  display: flex;
  box-shadow: 0px 1px 4px 0px rgba(92, 107, 192, 0.2);
  border-radius: 4px;
  background: #f5f6fa;
  margin-right: 20px;

  .tabChangeSty {
    // border: 1px solid;
    width: 65px;
    height: 26px;
    font-size: 12px;
    color: #87879d;
    text-align: center;
    line-height: 26px;
    border-right: 1px solid #ebebf1;
    cursor: pointer;
  }
}
.header-right {
  display: flex;
}

.store-app {
  display: flex;
  .change-store {
    position: relative;
    width: 90px;
    height: 26px;
    font-size: 12px;
    color: #87879d;
    text-align: center;
    line-height: 26px;
    box-shadow: 0px 1px 4px 0px;
    margin-right: 20px;
    background: #f5f6fa;
    box-shadow: 0px 1px 4px 0px rgba(92, 107, 192, 0.2);
    border-radius: 4px;
    cursor: pointer;
  }
  .change-store:last-child {
    margin-right: 10px !important;
  }
  .order {
    position: absolute;
    width: 11px;
    height: 13px;
  }

  .descending-order {
    right: 6px;
    top: 9px;
  }
  .ascending-order {
    right: 6px;
    top: 3px;
  }
}

.application-core {
  height: 100%;
}
.application-cation {
  background: #fff;
  margin-bottom: 20px;
  border-radius: 8px;
  min-height: 180px;
}
.application-title {
  // width: 66px;
  // height: 20px;
  font-size: 16px;
  font-family: PingFangSC-Medium, PingFang SC;
  font-weight: 500;
  color: #302e6d;
  line-height: 16px;
  padding: 20px 0 20px 20px;
}
.application-second {
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-start;
  border-radius: 8px;
  margin-left: 40px;
  img {
    display: block;
    width: 60px;
    height: 60px;
    box-shadow: 0px 2px 4px 0px rgba(214, 214, 239, 0.52);
    border-radius: 8px;
    // border: 1px solid #d6d6ef;
  }
  .every-last {
    margin: 0 30px 20px 0;
    cursor: pointer;
  }
  p {
    width: 60px;
    height: 20px;
    margin-top: 10px;
    font-size: 12px;
    font-family: PingFangSC-Medium, PingFang SC;
    font-weight: 500;
    color: #302e6d;
    line-height: 17px;
    text-align: center;
  }
}
.tabChangeStActive {
  // width: 65px;
  // height: 26px;
  background: #ededfd;
  border-radius: 4px 0px 0px 4;
}
</style>
