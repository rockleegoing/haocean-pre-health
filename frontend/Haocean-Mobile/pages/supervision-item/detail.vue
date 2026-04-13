<template>
  <view class="container">
    <Navbar :title="detail.itemName || '事项详情'" :hideBtn="false" bgColor="#fff" :h5Show="false" @back="handleBack"></Navbar>

    <view v-if="detail.itemId" class="content">
      <!-- 基本信息 -->
      <view class="section">
        <view class="section-title">基本信息</view>
        <view class="info-list">
          <view class="info-item">
            <text class="info-label">事项编码：</text>
            <text class="info-value">{{ detail.itemNo || '-' }}</text>
          </view>
          <view class="info-item">
            <text class="info-label">事项名称：</text>
            <text class="info-value">{{ detail.itemName }}</text>
          </view>
          <view class="info-item">
            <text class="info-label">监管类型：</text>
            <text class="info-value">{{ detail.supervisionTypeName || '-' }}</text>
          </view>
          <view class="info-item" v-if="detail.regulatedObjects">
            <text class="info-label">监管对象：</text>
            <text class="info-value">{{ detail.regulatedObjects }}</text>
          </view>
        </view>
      </view>

      <!-- 监管要求 -->
      <view class="section" v-if="detail.checkPoints">
        <view class="section-title">监管要求</view>
        <view class="check-points">
          <view class="point-item" v-for="(point, index) in checkPointsList" :key="index">
            <text class="point-index">{{ index + 1 }}.</text>
            <text class="point-content">{{ point }}</text>
          </view>
        </view>
      </view>

      <!-- 关联规范用语 -->
      <view class="section" v-if="standardLanguageList.length > 0">
        <view class="section-title">对应规范用语</view>
        <view class="language-list">
          <view class="language-item" v-for="(item, index) in standardLanguageList" :key="index" @click="handleLanguageClick(item)">
            <view class="language-icon">
              <u-icon name="file-text" color="#2979ff" size="36rpx"></u-icon>
            </view>
            <view class="language-info">
              <view class="language-name">{{ item.title }}</view>
              <view class="language-type">{{ item.category || '' }}</view>
            </view>
            <u-icon name="arrow-right" color="#999" size="28rpx"></u-icon>
          </view>
        </view>
      </view>

      <!-- 关联法律法规 -->
      <view class="section" v-if="regulationList.length > 0">
        <view class="section-title">对应法律法规</view>
        <view class="regulation-list">
          <view class="regulation-item" v-for="(item, index) in regulationList" :key="index" @click="handleRegulationClick(item)">
            <view class="regulation-icon">
              <u-icon name="star" color="#ff9800" size="36rpx"></u-icon>
            </view>
            <view class="regulation-info">
              <view class="regulation-name">{{ item.title }}</view>
              <view class="regulation-type">{{ item.legalType || '' }}</view>
            </view>
            <u-icon name="arrow-right" color="#999" size="28rpx"></u-icon>
          </view>
        </view>
      </view>

      <!-- 法律依据 -->
      <view class="section" v-if="detail.legalBasis">
        <view class="section-title">法律依据</view>
        <view class="legal-basis">{{ detail.legalBasis }}</view>
      </view>
    </view>

    <u-loading-page v-show="loading"></u-loading-page>
  </view>
</template>

<script>
import Navbar from '@/components/navbar/Navbar'
import { getSupervisionItem } from '@/api/supervisionItem'

export default {
  components: {
    Navbar
  },
  data() {
    return {
      loading: false,
      detail: {},
      standardLanguageList: [],
      regulationList: [],
      checkPointsList: []
    }
  },
  onLoad(options) {
    if (options.id) {
      this.loadDetail(options.id)
    }
  },
  methods: {
    // 加载详情
    async loadDetail(id) {
      this.loading = true
      try {
        const res = await getSupervisionItem(id)
        this.detail = res.data

        // 解析检查要点
        if (this.detail.checkPoints) {
          try {
            const points = JSON.parse(this.detail.checkPoints)
            this.checkPointsList = Array.isArray(points) ? points : [this.detail.checkPoints]
          } catch (e) {
            this.checkPointsList = [this.detail.checkPoints]
          }
        }

        // 解析关联规范用语
        if (this.detail.standardLanguageIds) {
          try {
            const ids = JSON.parse(this.detail.standardLanguageIds)
            // 这里需要根据 ID 列表获取规范用语详情
            // 暂时显示占位数据
            this.standardLanguageList = ids.map((id, index) => ({
              id,
              title: `规范用语${index + 1}`,
              category: '检查中'
            }))
          } catch (e) {
            this.standardLanguageList = []
          }
        }

        // 解析关联法律法规
        if (this.detail.industryIds) {
          try {
            const ids = JSON.parse(this.detail.industryIds)
            // 这里需要根据 ID 列表获取法律法规详情
            // 暂时显示占位数据
            this.regulationList = ids.map((id, index) => ({
              id,
              title: `法律法规${index + 1}`,
              legalType: '行政法规'
            }))
          } catch (e) {
            this.regulationList = []
          }
        }
      } catch (error) {
        console.error('加载详情失败', error)
        uni.showToast({ title: '加载失败', icon: 'none' })
      } finally {
        this.loading = false
      }
    },

    // 返回
    handleBack() {
      uni.navigateBack()
    },

    // 点击查看规范用语
    handleLanguageClick(item) {
      uni.showToast({ title: '规范用语详情待开发', icon: 'none' })
    },

    // 点击查看法律法规
    handleRegulationClick(item) {
      uni.showToast({ title: '法律法规详情待开发', icon: 'none' })
    }
  }
}
</script>

<style lang="scss" scoped>
.container {
  background-color: #f5f5f5;
  min-height: 100vh;
}

.content {
  padding: 20rpx;
}

.section {
  margin-bottom: 20rpx;
  padding: 30rpx;
  background-color: #fff;
  border-radius: 16rpx;

  .section-title {
    font-size: 32rpx;
    font-weight: bold;
    margin-bottom: 24rpx;
    padding-bottom: 16rpx;
    border-bottom: 1px solid #f0f0f0;
  }
}

.info-list {
  .info-item {
    display: flex;
    padding: 12rpx 0;

    .info-label {
      font-size: 28rpx;
      color: #666;
      width: 160rpx;
      flex-shrink: 0;
    }

    .info-value {
      font-size: 28rpx;
      color: #333;
      flex: 1;
    }
  }
}

.check-points {
  .point-item {
    display: flex;
    padding: 12rpx 0;

    .point-index {
      font-size: 28rpx;
      color: #2979ff;
      font-weight: bold;
      margin-right: 12rpx;
      min-width: 40rpx;
    }

    .point-content {
      font-size: 28rpx;
      color: #333;
      flex: 1;
      line-height: 1.6;
    }
  }
}

.language-list, .regulation-list {
  .language-item, .regulation-item {
    display: flex;
    align-items: center;
    padding: 20rpx 0;
    border-bottom: 1px solid #f0f0f0;

    &:last-child {
      border-bottom: none;
    }

    .language-icon, .regulation-icon {
      margin-right: 20rpx;
    }

    .language-info, .regulation-info {
      flex: 1;

      .language-name, .regulation-name {
        font-size: 28rpx;
        color: #333;
        margin-bottom: 6rpx;
      }

      .language-type, .regulation-type {
        font-size: 24rpx;
        color: #999;
      }
    }
  }
}

.legal-basis {
  font-size: 28rpx;
  color: #333;
  line-height: 1.8;
}
</style>
