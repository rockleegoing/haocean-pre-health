<!-- 法律法规库首页 - 综合法律条例展示 + 监管类型网格入口 -->
<template>
	<view class="t-law-home">
		<!-- 搜索栏 -->
		<view class="search-section">
			<u-search v-model="keyword" placeholder="搜索法律法规" shape="round" @search="handleSearch" @change="handleSearchChange" />
		</view>

		<!-- 监管类型网格 -->
		<view class="grid-section">
			<view class="section-title">监管类型</view>
			<u-grid :col="3" :border="false">
				<u-grid-item v-for="(item, index) in supervisionTypeList" :key="index" :index="item.id" @click="handleSupervisionTypeClick">
					<view class="grid-item-content">
						<view class="icon-wrapper" :style="{ background: getIconColor(index) }">
							<u-icon :name="item.icon || 'folder'" size="40" color="#fff"></u-icon>
						</view>
						<text class="item-name">{{ item.name }}</text>
					</view>
				</u-grid-item>
			</u-grid>
		</view>

		<!-- 法律类型列表 -->
		<view class="legal-section">
			<view class="section-title">法律类型</view>
			<view class="legal-list">
				<view class="legal-item" v-for="(item, index) in legalTypeList" :key="index" @click="handleLegalTypeClick">
					<view class="legal-icon" :style="{ background: getLegalIconColor(index) }">
						<u-icon name="file-text" size="32" color="#fff"></u-icon>
					</view>
					<view class="legal-info">
						<text class="legal-name">{{ item.name }}</text>
						<text class="legal-count">共 {{ item.count || 0 }} 部法规</text>
					</view>
					<u-icon name="arrow-right" size="24" color="#ccc"></u-icon>
				</view>
			</view>
		</view>

		<!-- 最近浏览 -->
		<view class="recent-section" v-if="recentList.length > 0">
			<view class="section-title">最近浏览</view>
			<view class="recent-list">
				<view class="recent-item" v-for="(item, index) in recentList" :key="index" @click="handleRecentClick">
					<view class="recent-icon">
						<u-icon name="clock" size="30" color="#1890ff"></u-icon>
					</view>
					<view class="recent-info">
						<text class="recent-title">{{ item.title }}</text>
						<text class="recent-time">{{ item.time }}</text>
					</view>
					<u-icon name="arrow-right" size="24" color="#ccc"></u-icon>
				</view>
			</view>
			<view class="clear-btn" @click="handleClearRecent">
				<text>清除浏览记录</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { getHomeData, listLegalType, listSupervisionType } from '@/api/law.js'

	export default {
		data() {
			return {
				keyword: '',
				supervisionTypeList: [],
				legalTypeList: [],
				recentList: [],
				iconColors: [
					'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
					'linear-gradient(135deg, #f093fb 0%, #f5576c 100%)',
					'linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)',
					'linear-gradient(135deg, #43e97b 0%, #38f9d7 100%)',
					'linear-gradient(135deg, #fa709a 0%, #fee140 100%)',
					'linear-gradient(135deg, #a8edea 0%, #fed6e3 100%)',
					'linear-gradient(135deg, #ff9a9e 0%, #fecfef 100%)',
					'linear-gradient(135deg, #fbc2eb 0%, #a6c1ee 100%)',
					'linear-gradient(135deg, #fdcbf1 0%, #e6dee9 100%)',
					'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
					'linear-gradient(135deg, #f093fb 0%, #f5576c 100%)',
					'linear-gradient(135deg, #89f7fe 0%, #66a6ff 100%)'
				]
			}
		},
		onLoad() {
			this.getSupervisionTypeList()
			this.getLegalTypeList()
			this.loadRecentList()
		},
		methods: {
			// 获取监管类型列表
			getSupervisionTypeList() {
				listSupervisionType({ isEnabled: 1 }).then(res => {
					this.supervisionTypeList = res.rows || res || []
				})
			},

			// 获取法律类型列表
			getLegalTypeList() {
				getHomeData().then(res => {
					const data = res.data || res || {}
					this.legalTypeList = data.legalTypeCounts || []
					// 如果没有计数，设置默认值
					this.legalTypeList.forEach(item => {
						if (!item.count) item.count = 0
					})
				})
			},

			// 获取图标颜色
			getIconColor(index) {
				return this.iconColors[index % this.iconColors.length]
			},

			// 获取法律类型图标颜色
			getLegalIconColor(index) {
				return this.iconColors[index % this.iconColors.length]
			},

			// 点击监管类型
			handleSupervisionTypeClick(item) {
				uni.navigateTo({
					url: '/pages/law/list?supervisionType=' + item.code + '&supervisionTypeName=' + item.name
				})
			},

			// 点击法律类型
			handleLegalTypeClick(item) {
				uni.navigateTo({
					url: '/pages/law/list?legalType=' + item.type + '&legalTypeName=' + item.name
				})
			},

			// 搜索
			handleSearch(value) {
				if (!value) {
					uni.showToast({ title: '请输入搜索关键词', icon: 'none' })
					return
				}
				// 保存到搜索历史
				this.saveSearchHistory(value)
				uni.navigateTo({
					url: '/pages/law/search?keyword=' + value
				})
			},

			handleSearchChange(value) {
				this.keyword = value
			},

			// 加载最近浏览记录
			loadRecentList() {
				const recent = uni.getStorageSync('lawRecentList') || '[]'
				this.recentList = JSON.parse(recent)
			},

			// 点击最近浏览
			handleRecentClick(item) {
				uni.navigateTo({
					url: '/pages/law/detail?id=' + item.id
				})
			},

			// 清除浏览记录
			handleClearRecent() {
				uni.showModal({
					title: '提示',
					content: '确定清除浏览记录吗？',
					success: (res) => {
						if (res.confirm) {
							uni.removeStorageSync('lawRecentList')
							this.recentList = []
							uni.showToast({ title: '已清除', icon: 'success' })
						}
					}
				})
			},

			// 保存搜索历史
			saveSearchHistory(keyword) {
				let history = uni.getStorageSync('lawSearchHistory') || '[]'
				history = JSON.parse(history)
				// 移除已存在的
				const index = history.indexOf(keyword)
				if (index > -1) {
					history.splice(index, 1)
				}
				// 添加到开头
				history.unshift(keyword)
				// 只保留最近的 10 条
				if (history.length > 10) {
					history = history.slice(0, 10)
				}
				uni.setStorageSync('lawSearchHistory', JSON.stringify(history))
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-law-home {
		padding: 20rpx;
		background: #f5f5f5;
		min-height: 100vh;

		.search-section {
			background: #fff;
			padding: 20rpx;
			border-radius: 16rpx;
			margin-bottom: 20rpx;
		}

		.grid-section {
			background: #fff;
			padding: 30rpx 20rpx;
			border-radius: 16rpx;
			margin-bottom: 20rpx;

			.section-title {
				font-size: 32rpx;
				font-weight: bold;
				color: #333;
				margin-bottom: 20rpx;
				padding-left: 10rpx;
			}

			.grid-item-content {
				display: flex;
				flex-direction: column;
				align-items: center;
				padding: 20rpx;

				.icon-wrapper {
					width: 100rpx;
					height: 100rpx;
					border-radius: 20rpx;
					display: flex;
					justify-content: center;
					align-items: center;
					margin-bottom: 12rpx;
					box-shadow: 0 4rpx 12rpx rgba(0, 0, 0, 0.1);
				}

				.item-name {
					font-size: 26rpx;
					color: #666;
					text-align: center;
					line-height: 1.4;
				}
			}
		}

		.legal-section {
			background: #fff;
			padding: 30rpx 20rpx;
			border-radius: 16rpx;
			margin-bottom: 20rpx;

			.section-title {
				font-size: 32rpx;
				font-weight: bold;
				color: #333;
				margin-bottom: 20rpx;
				padding-left: 10rpx;
			}

			.legal-list {
				.legal-item {
					display: flex;
					align-items: center;
					padding: 24rpx 0;
					border-bottom: 1rpx solid #f0f0f0;

					&:last-child {
						border-bottom: none;
					}

					.legal-icon {
						width: 80rpx;
						height: 80rpx;
						border-radius: 16rpx;
						display: flex;
						justify-content: center;
						align-items: center;
						margin-right: 20rpx;
						flex-shrink: 0;
					}

					.legal-info {
						flex: 1;

						.legal-name {
							font-size: 30rpx;
							color: #333;
							display: block;
							margin-bottom: 8rpx;
							font-weight: bold;
						}

						.legal-count {
							font-size: 24rpx;
							color: #999;
						}
					}
				}
			}
		}

		.recent-section {
			background: #fff;
			padding: 30rpx 20rpx;
			border-radius: 16rpx;

			.section-title {
				font-size: 32rpx;
				font-weight: bold;
				color: #333;
				margin-bottom: 20rpx;
				padding-left: 10rpx;
			}

			.recent-list {
				.recent-item {
					display: flex;
					align-items: center;
					padding: 24rpx 0;
					border-bottom: 1rpx solid #f0f0f0;

					&:last-child {
						border-bottom: none;
					}

					.recent-icon {
						margin-right: 20rpx;
					}

					.recent-info {
						flex: 1;

						.recent-title {
							font-size: 30rpx;
							color: #333;
							display: block;
							margin-bottom: 8rpx;
						}

						.recent-time {
							font-size: 24rpx;
							color: #999;
						}
					}
				}
			}

			.clear-btn {
				margin-top: 20rpx;
				text-align: center;
				padding: 20rpx;
				border-top: 1rpx solid #f0f0f0;

				text {
					font-size: 28rpx;
					color: #999;
				}
			}
		}
	}
</style>
