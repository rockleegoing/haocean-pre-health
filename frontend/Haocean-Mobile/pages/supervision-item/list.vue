<template>
  <view class="container">
    <Navbar :title="categoryName" :hideBtn="false" bgColor="#fff" :h5Show="false" @back="handleBack"></Navbar>

    <!-- 搜索栏 -->
    <view class="search-bar">
      <u-search v-model="keyword" placeholder="搜索事项" @search="handleSearch" @clear="handleClear" />
    </view>

    <!-- 事项列表 -->
    <view class="item-list">
      <view class="item-list-title">监管子项</view>
      <view class="list">
        <view class="list-item" v-for="(item, index) in itemList" :key="index" @click="handleItemClick(item)">
          <view class="item-left">
            <view class="item-index">{{ index + 1 }}.</view>
            <view class="item-content">
              <view class="item-name">{{ item.itemName }}</view>
              <view class="item-desc">{{ item.checkPoints || '暂无描述' }}</view>
            </view>
          </view>
          <u-icon name="arrow-right" color="#999" size="28rpx"></u-icon>
        </view>
      </view>

      <!-- 空状态 -->
      <u-empty v-if="itemList.length === 0" mode="list" text="暂无监管事项"></u-empty>
    </view>
  </view>
</template>

<script>
import Navbar from '@/components/navbar/Navbar'
import { getSupervisionChildren, listSupervisionItem } from '@/api/supervisionItem'

export default {
  components: {
    Navbar
  },
  data() {
    return {
      keyword: '',
      categoryId: '',
      categoryName: '监管事项',
      itemList: []
    }
  },
  onLoad(options) {
    if (options.categoryId) {
      this.categoryId = options.categoryId
      this.categoryName = decodeURIComponent(options.categoryName || '监管事项')
    }
    if (options.keyword) {
      this.keyword = decodeURIComponent(options.keyword)
      this.searchItems()
    } else {
      this.loadItems()
    }
  },
  methods: {
    // 加载事项列表
    async loadItems() {
      try {
        // 如果是分类，先获取分类下的事项
        if (this.categoryId) {
          const res = await listSupervisionItem({
            supervisionType: this.categoryId,
            pageNum: 1,
            pageSize: 100
          })
          this.itemList = res.rows || []
        }
      } catch (error) {
        console.error('加载事项失败', error)
        uni.showToast({ title: '加载失败', icon: 'none' })
      }
    },

    // 搜索事项
    async searchItems() {
      try {
        const res = await listSupervisionItem({
          itemName: this.keyword,
          pageNum: 1,
          pageSize: 100
        })
        this.itemList = res.rows || []
      } catch (error) {
        console.error('搜索失败', error)
      }
    },

    // 处理搜索
    handleSearch() {
      if (!this.keyword.trim()) return
      this.searchItems()
    },

    // 清空搜索
    handleClear() {
      this.keyword = ''
      this.loadItems()
    },

    // 返回
    handleBack() {
      uni.navigateBack()
    },

    // 点击事项
    handleItemClick(item) {
      uni.navigateTo({
        url: `/pages/supervision-item/detail?id=${item.itemId}`
      })
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

.item-list {
  margin: 20rpx;
  padding: 30rpx;
  background-color: #fff;
  border-radius: 16rpx;

  .item-list-title {
    font-size: 32rpx;
    font-weight: bold;
    margin-bottom: 24rpx;
  }

  .list {
    .list-item {
      display: flex;
      align-items: center;
      justify-content: space-between;
      padding: 24rpx 0;
      border-bottom: 1px solid #f0f0f0;

      &:last-child {
        border-bottom: none;
      }

      .item-left {
        display: flex;
        align-items: flex-start;
        flex: 1;
      }

      .item-index {
        font-size: 28rpx;
        color: #2979ff;
        font-weight: bold;
        margin-right: 16rpx;
        min-width: 40rpx;
      }

      .item-content {
        flex: 1;

        .item-name {
          font-size: 28rpx;
          color: #333;
          margin-bottom: 8rpx;
        }

        .item-desc {
          font-size: 24rpx;
          color: #999;
          overflow: hidden;
          text-overflow: ellipsis;
          display: -webkit-box;
          -webkit-line-clamp: 2;
          -webkit-box-orient: vertical;
        }
      }
    }
  }
}
</style>
