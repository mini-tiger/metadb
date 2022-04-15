<template>
  <div>
    <el-button @click="handleLogin">登录</el-button>
    <el-button @click="handleTest">发送验证码</el-button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router'
import * as types from '../store/action-types'
import { useStore } from 'vuex'
import { ElMessage } from 'element-plus'
import axios from '@/utils/httpRequest'

const useLogin = (store, router) => {
  const handleLogin = async () => {
    try {
      await store.dispatch(`user/${types.USER_LOGIN}`)
      // let res = await store.dispatch(`user/${types.USER_LOGIN}`, {
      //   username: 'admin',
      //   password:
      //     'ND4iXOdcP1FF+LAcC8UygSCg04VJkQLVpeC3F9iVbDF2eXUG+/79oAh6E3u45jkapcxUv1h3dJlmCNj+n8IHJg==',
      //   captcha: '',
      // })
      // res && router.push('/')
      router.push('/')
    } catch (error) {
      return ElMessage('登录失败')
    }
  }
  return { handleLogin }
}
const useTest = () => {
  const handleTest = async () => {
    await axios.get(`/api/hzero/v1/users/email/send-captcha`, {
      email: '491028246@qq.com',
      userType: 'user',
      businessScope: 'i',
    })
  }
  return { handleTest }
}
export default defineComponent({
  name: 'login',
  setup() {
    const store = useStore()
    const router = useRouter()
    const { handleLogin } = useLogin(store, router)
    const { handleTest } = useTest()

    return {
      handleLogin,
      handleTest,
    }
  },
})
</script>
<style scoped lang="scss"></style>
