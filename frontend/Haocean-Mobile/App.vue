<script>
import { ACCESS_TOKEN } from '@/store/mutation-types'
import storage from '@/utils/storage'

export default {
  onLaunch () {
    console.log('App Launch')
    // 先检查登录状态
    this.checkLoginStatus()
  },
  onShow () {
    // 每次显示时也检查登录状态
    this.checkLoginStatus()
  },
  onHide () {
  },
  methods: {
    // 检查登录状态
    checkLoginStatus() {
      const token = storage.get(ACCESS_TOKEN)

      // 获取当前页面路径
      const pages = getCurrentPages()
      const currentPage = pages.length > 0 ? pages[pages.length - 1].route : ''
      const isLoginPage = currentPage === 'pages/login/index'

      // 如果没有 token 且不在登录页，跳转到登录页
      if (!token && !isLoginPage) {
        console.log('未登录，跳转到登录页')
        uni.reLaunch({
          url: '/pages/login/index'
        })
        return
      }

      // 如果有 token 且在登录页，跳转到首页
      if (token && isLoginPage) {
        console.log('已登录，跳转到首页')
        uni.reLaunch({
          url: '/pages/index/index'
        })
        return
      }

      // 已登录且不在登录页，加载系统信息
      if (token && !isLoginPage) {
        this.$store.dispatch('SystemInfo')
      }
    }
  }
}
</script>

<style lang="scss">
@import "@/uni_modules/uview-ui/index.scss";
@import "@/static/style.scss";
</style>
