<template>
  <view class="container">
    <Navbar title="监管事项" :hideBtn="true" bgColor="#fff" :h5Show="false"></Navbar>

    <!-- 搜索栏 -->
    <view class="search-bar">
      <u-search v-model="keyword" placeholder="搜索监管事项" @search="handleSearch" @clear="handleClear" />
    </view>

    <!-- 监管类型分类 -->
    <view class="category-section">
      <view class="section-title">监管事项分类</view>
      <u-grid :col="4">
        <u-grid-item v-for="(item, index) in categories" :key="index" @click="handleCategoryClick(item)">
          <u-icon :name="item.icon" :color="item.color" size="48rpx"></u-icon>
          <text class="category-name">{{ item.categoryName }}</text>
        </u-grid-item>
      </u-grid>
    </view>

    <!-- 最近浏览 -->
    <view class="recent-section">
      <view class="section-title">最近浏览</view>
      <view class="recent-list">
        <view class="recent-item" v-for="(item, index) in recentItems" :key="index" @click="handleItemClick(item)">
          <view class="item-icon">
            <u-icon name="file-text" color="#2979ff" size="40rpx"></u-icon>
          </view>
          <view class="item-info">
            <view class="item-name">{{ item.itemName }}</view>
            <view class="item-type">{{ item.supervisionTypeName }}</view>
          </view>
          <u-icon name="arrow-right" color="#999" size="28rpx"></u-icon>
        </view>
      </view>
    </view>
  </view>
</template>

<script>
import Navbar from '@/components/navbar/Navbar'
import { listSupervisionCategory, getSupervisionTree } from '@/api/supervisionItem'

export default {
  components: {
    Navbar
  },
  data() {
    return {
      keyword: '',
      categories: [],
      recentItems: []
    }
  },
  created() {
    this.loadCategories()
    this.loadRecentItems()
  },
  methods: {
    // 加载监管分类
    async loadCategories() {
      try {
        const res = await listSupervisionCategory()
        this.categories = res.data.map((item, index) => ({
          ...item,
          color: ['#2979ff', '#00c853', '#ff9100', '#ff5252', '#9c27b0', '#00bcd4', '#ff9800', '#4caf50', '#e91e63', '#3f51b5'][index % 10]
        }))
      } catch (error) {
        console.error('加载分类失败', error)
      }
    },

    // 加载最近浏览
    loadRecentItems() {
      const recent = uni.getStorageSync('recent_supervision_items') || []
      this.recentItems = recent.slice(0, 5)
    },

    // 处理搜索
    handleSearch() {
      if (!this.keyword.trim()) return
      uni.navigateTo({
        url: `/pages/supervision-item/list?keyword=${encodeURIComponent(this.keyword)}`
      })
    },

    // 清空搜索
    handleClear() {
      this.keyword = ''
    },

    // 点击分类
    handleCategoryClick(item) {
      uni.navigateTo({
        url: `/pages/supervision-item/list?categoryId=${item.categoryCode}&categoryName=${encodeURIComponent(item.categoryName)}`
      })
    },

    // 点击事项
    handleItemClick(item) {
      this.saveToRecent(item)
      uni.navigateTo({
        url: `/pages/supervision-item/detail?id=${item.itemId}`
      })
    },

    // 保存到最近浏览
    saveToRecent(item) {
      let recent = uni.getStorageSync('recent_supervision_items') || []
      // 移除已存在的
      recent = recent.filter(i => i.itemId !== item.itemId)
      // 添加到开头
      recent.unshift(item)
      // 限制最多保存 10 条
      if (recent.length > 10) recent = recent.slice(0, 10)
      uni.setStorageSync('recent_supervision_items', recent)
    }
  }
}
</script>

<style lang="scss" scoped>
.container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.search-bar {
  padding: 20rpx;
  background-color: #fff;
}

.category-section {
  margin: 20rpx;
  padding: 30rpx;
  background-color: #fff;
  border-radius: 16rpx;

  .section-title {
    font-size: 32rpx;
    font-weight: bold;
    margin-bottom: 24rpx;
  }

  .category-name {
    font-size: 24rpx;
    color: #666;
    margin-top: 12rpx;
    display: block;
    text-align: center;
  }
}

.recent-section {
  margin: 20rpx;
  padding: 30rpx;
  background-color: #fff;
  border-radius: 16rpx;

  .section-title {
    font-size: 32rpx;
    font-weight: bold;
    margin-bottom: 24rpx;
  }

  .recent-list {
    .recent-item {
      display: flex;
      align-items: center;
      padding: 24rpx 0;
      border-bottom: 1px solid #f0f0f0;

      &:last-child {
        border-bottom: none;
      }

      .item-icon {
        margin-right: 20rpx;
      }

      .item-info {
        flex: 1;

        .item-name {
          font-size: 28rpx;
          color: #333;
          margin-bottom: 8rpx;
        }

        .item-type {
          font-size: 24rpx;
          color: #999;
        }
      }
    }
  }
}
</style>
