<!-- 法律法规库搜索页 - 关键字搜索 + 类型筛选 -->
<template>
	<view class="t-law-search">
		<!-- 搜索栏 -->
		<view class="search-section">
			<u-search v-model="keyword" placeholder="搜索法律法规名称或内容" shape="round" @search="handleSearch" />
		</view>

		<!-- 筛选条件 -->
		<view class="filter-section">
			<view class="filter-item">
				<text class="filter-label">法律类型：</text>
				<u-picker-popup :show="showLegalTypePicker" :list="legalTypeList" @confirm="handleLegalTypeConfirm" @close="showLegalTypePicker = false">
					<view class="filter-value" @click="showLegalTypePicker = true">
						<text>{{ legalTypeName || '全部' }}</text>
						<u-icon name="arrow-down" size="24" color="#999"></u-icon>
					</view>
				</u-picker-popup>
			</view>
			<view class="filter-item">
				<text class="filter-label">监管类型：</text>
				<u-picker-popup :show="showSupervisionTypePicker" :list="supervisionTypeList" @confirm="handleSupervisionTypeConfirm" @close="showSupervisionTypePicker = false">
					<view class="filter-value" @click="showSupervisionTypePicker = true">
						<text>{{ supervisionTypeName || '全部' }}</text>
						<u-icon name="arrow-down" size="24" color="#999"></u-icon>
					</view>
				</u-picker-popup>
			</view>
		</view>

		<!-- 搜索历史 -->
		<view class="history-section" v-if="!hasSearched">
			<view class="history-header">
				<text class="history-title">搜索历史</text>
				<text class="clear-btn" @click="handleClearHistory">清除</text>
			</view>
			<view class="history-list">
				<view class="history-item" v-for="(item, index) in searchHistory" :key="index" @click="handleHistoryClick(item)">
					<u-icon name="clock" size="28" color="#999"></u-icon>
					<text class="history-text">{{ item }}</text>
				</view>
			</view>
		</view>

		<!-- 搜索结果 -->
		<view class="result-section" v-if="hasSearched">
			<view class="result-header">
				<text class="result-count">找到 {{ resultList.length }} 条结果</text>
			</view>

			<view class="result-list">
				<view class="result-item" v-for="(item, index) in resultList" :key="index" @click="handleItemClick(item)">
					<view class="result-icon">
						<u-icon name="book" size="40" color="#1890ff"></u-icon>
					</view>
					<view class="result-info">
						<text class="result-title">{{ highlightText(item.title) }}</text>
						<view class="result-tags" v-if="item.legalType || item.supervisionType">
							<u-tag v-if="item.legalType" :text="legalTypeMap[item.legalType]" type="primary" size="mini"></u-tag>
							<u-tag v-if="item.supervisionType" :text="supervisionTypeMap[item.supervisionType]" type="success" size="mini"></u-tag>
						</view>
						<text class="result-content">{{ item.content }}</text>
					</view>
					<u-icon name="arrow-right" size="24" color="#ccc"></u-icon>
				</view>
			</view>

			<u-empty v-if="resultList.length === 0" text="未找到相关结果" mode="search"></u-empty>
		</view>
	</view>
</template>

<script>
	import { searchRegulation, listLegalType, listSupervisionType } from '@/api/law.js'

	export default {
		data() {
			return {
				keyword: '',
				hasSearched: false,
				resultList: [],
				searchHistory: [],
				legalTypeList: [],
				supervisionTypeList: [],
				legalTypeCode: '',
				legalTypeName: '',
				supervisionTypeCode: '',
				supervisionTypeName: '',
				showLegalTypePicker: false,
				showSupervisionTypePicker: false,
				legalTypeMap: {},
				supervisionTypeMap: {}
			}
		},
		onLoad(options) {
			if (options.keyword) {
				this.keyword = options.keyword
				this.handleSearch(options.keyword)
			}
			this.loadSearchHistory()
			this.getTypeList()
		},
		methods: {
			// 获取类型列表
			getTypeList() {
				listLegalType({ isEnabled: 1 }).then(res => {
					const types = res.rows || res || []
					this.legalTypeList = types.map(item => ({
						label: item.name,
						value: item.code
					}))
					types.forEach(item => {
						this.legalTypeMap[item.code] = item.name
					})
				})
				listSupervisionType({ isEnabled: 1 }).then(res => {
					const types = res.rows || res || []
					this.supervisionTypeList = types.map(item => ({
						label: item.name,
						value: item.code
					}))
					types.forEach(item => {
						this.supervisionTypeMap[item.code] = item.name
					})
				})
			},

			// 加载搜索历史
			loadSearchHistory() {
				const history = uni.getStorageSync('lawSearchHistory') || '[]'
				this.searchHistory = JSON.parse(history)
			},

			// 搜索
			handleSearch(value) {
				if (!value) {
					uni.showToast({ title: '请输入搜索关键词', icon: 'none' })
					return
				}
				this.keyword = value
				this.hasSearched = true
				this.saveSearchHistory(value)
				this.searchResult(value)
			},

			// 执行搜索
			searchResult(keyword) {
				const query = { keyword: keyword }
				if (this.legalTypeCode) query.legalType = this.legalTypeCode
				if (this.supervisionTypeCode) query.supervisionType = this.supervisionTypeCode

				searchRegulation(query).then(res => {
					this.resultList = res.data || res || []
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
				this.searchHistory = history
			},

			// 清除搜索历史
			handleClearHistory() {
				uni.showModal({
					title: '提示',
					content: '确定清除搜索历史吗？',
					success: (res) => {
						if (res.confirm) {
							uni.removeStorageSync('lawSearchHistory')
							this.searchHistory = []
							uni.showToast({ title: '已清除', icon: 'success' })
						}
					}
				})
			},

			// 点击搜索历史
			handleHistoryClick(keyword) {
				this.keyword = keyword
				this.handleSearch(keyword)
			},

			// 选择法律类型
			handleLegalTypeConfirm(item) {
				this.legalTypeCode = item.value
				this.legalTypeName = item.label
				if (this.keyword) {
					this.searchResult(this.keyword)
				}
			},

			// 选择监管类型
			handleSupervisionTypeConfirm(item) {
				this.supervisionTypeCode = item.value
				this.supervisionTypeName = item.label
				if (this.keyword) {
					this.searchResult(this.keyword)
				}
			},

			// 高亮文本
			highlightText(text) {
				if (!this.keyword) return text
				return text
			},

			// 点击搜索结果
			handleItemClick(item) {
				// 保存到最近浏览
				this.saveRecent(item)
				uni.navigateTo({
					url: '/pages/law/detail?id=' + item.regulationId
				})
			},

			// 保存最近浏览
			saveRecent(item) {
				let recent = uni.getStorageSync('lawRecentList') || '[]'
				recent = JSON.parse(recent)
				// 移除已存在的
				const index = recent.findIndex(r => r.id === item.regulationId)
				if (index > -1) {
					recent.splice(index, 1)
				}
				// 添加到开头
				recent.unshift({
					id: item.regulationId,
					title: item.title,
					time: this.formatTime(new Date())
				})
				// 只保留最近的 10 条
				if (recent.length > 10) {
					recent = recent.slice(0, 10)
				}
				uni.setStorageSync('lawRecentList', JSON.stringify(recent))
			},

			// 格式化时间
			formatTime(date) {
				const year = date.getFullYear()
				const month = (date.getMonth() + 1).toString().padStart(2, '0')
				const day = date.getDate().toString().padStart(2, '0')
				const hour = date.getHours().toString().padStart(2, '0')
				const minute = date.getMinutes().toString().padStart(2, '0')
				return `${year}-${month}-${day} ${hour}:${minute}`
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-law-search {
		padding: 20rpx;
		background: #f5f5f5;
		min-height: 100vh;

		.search-section {
			background: #fff;
			padding: 20rpx;
			border-radius: 16rpx;
			margin-bottom: 20rpx;
		}

		.filter-section {
			background: #fff;
			padding: 20rpx 30rpx;
			border-radius: 16rpx;
			margin-bottom: 20rpx;

			.filter-item {
				display: flex;
				align-items: center;
				padding: 16rpx 0;
				border-bottom: 1rpx solid #f0f0f0;

				&:last-child {
					border-bottom: none;
				}

				.filter-label {
					font-size: 28rpx;
					color: #666;
					margin-right: 20rpx;
					width: 160rpx;
				}

				.filter-value {
					flex: 1;
					display: flex;
					justify-content: space-between;
					align-items: center;
					padding: 12rpx 0;

					text {
						font-size: 28rpx;
						color: #333;
					}
				}
			}
		}

		.history-section {
			background: #fff;
			padding: 30rpx 20rpx;
			border-radius: 16rpx;

			.history-header {
				display: flex;
				justify-content: space-between;
				align-items: center;
				margin-bottom: 20rpx;

				.history-title {
					font-size: 32rpx;
					font-weight: bold;
					color: #333;
				}

				.clear-btn {
					font-size: 26rpx;
					color: #999;
				}
			}

			.history-list {
				.history-item {
					display: flex;
					align-items: center;
					padding: 20rpx 0;
					border-bottom: 1rpx solid #f0f0f0;

					&:last-child {
						border-bottom: none;
					}

					u-icon {
						margin-right: 16rpx;
					}

					.history-text {
						font-size: 30rpx;
						color: #333;
					}
				}
			}
		}

		.result-section {
			.result-header {
				display: flex;
				justify-content: space-between;
				align-items: center;
				padding: 20rpx;
				background: #fff;
				border-radius: 16rpx;
				margin-bottom: 20rpx;

				.result-count {
					font-size: 28rpx;
					color: #999;
				}
			}

			.result-list {
				.result-item {
					background: #fff;
					border-radius: 16rpx;
					padding: 30rpx;
					margin-bottom: 20rpx;
					display: flex;
					align-items: flex-start;

					.result-icon {
						margin-right: 20rpx;
						flex-shrink: 0;
					}

					.result-info {
						flex: 1;

						.result-title {
							font-size: 30rpx;
							font-weight: bold;
							color: #333;
							display: block;
							margin-bottom: 12rpx;
						}

						.result-tags {
							display: flex;
							gap: 10rpx;
							margin-bottom: 12rpx;
						}

						.result-content {
							font-size: 26rpx;
							color: #666;
							line-height: 1.6;
							display: -webkit-box;
							-webkit-box-orient: vertical;
							-webkit-line-clamp: 2;
							overflow: hidden;
						}
					}
				}
			}
		}
	}
</style>
