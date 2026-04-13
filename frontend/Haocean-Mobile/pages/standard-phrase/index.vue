<!-- 规范用语首页 - 12 类监管类型网格展示 -->
<template>
	<view class="t-standard-phrase-home">
		<!-- 搜索栏 -->
		<view class="search-section">
			<u-search v-model="keyword" placeholder="搜索规范用语" @search="handleSearch" @change="handleSearchChange" />
		</view>

		<!-- 12 类监管类型网格 -->
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

		<!-- 常用规范用语 -->
		<view class="common-section">
			<view class="section-title">常用规范用语</view>
			<view class="common-list">
				<view class="common-item" v-for="(item, index) in commonPhrases" :key="index" @click="handleCommonPhraseClick">
					<view class="common-icon">
						<u-icon name="file-text" size="30" color="#1890ff"></u-icon>
					</view>
					<view class="common-info">
						<text class="common-title">{{ item.title }}</text>
						<text class="common-scene">{{ item.scene }}</text>
					</view>
					<u-icon name="arrow-right" size="24" color="#ccc"></u-icon>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import { listSupervisionType, listItem, searchStandardPhrase } from '@/api/standardPhrase.js'

	export default {
		data() {
			return {
				supervisionTypeList: [],
				keyword: '',
				commonPhrases: [],
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
			this.getCommonPhrases()
		},
		methods: {
			// 获取监管类型列表
			getSupervisionTypeList() {
				listSupervisionType({ isEnabled: 1 }).then(res => {
					this.supervisionTypeList = res.rows || res || []
				})
			},

			// 获取常用规范用语
			getCommonPhrases() {
				listItem({ categoryId: 1 }).then(res => {
					const items = res.rows || res || []
					this.commonPhrases = items.slice(0, 5)
				})
			},

			// 获取图标颜色
			getIconColor(index) {
				return this.iconColors[index % this.iconColors.length]
			},

			// 点击监管类型
			handleSupervisionTypeClick(item) {
				uni.navigateTo({
					url: '/pages/standard-phrase/list?supervisionTypeId=' + item.id + '&supervisionTypeName=' + item.name
				})
			},

			// 搜索
			handleSearch(value) {
				if (!value) {
					uni.showToast({ title: '请输入搜索关键词', icon: 'none' })
					return
				}
				uni.navigateTo({
					url: '/pages/standard-phrase/search?keyword=' + value
				})
			},

			handleSearchChange(value) {
				this.keyword = value
			},

			// 点击常用规范用语
			handleCommonPhraseClick(item) {
				uni.navigateTo({
					url: '/pages/standard-phrase/detail?itemId=' + item.id
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-standard-phrase-home {
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

		.common-section {
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

			.common-list {
				.common-item {
					display: flex;
					align-items: center;
					padding: 24rpx 0;
					border-bottom: 1rpx solid #f0f0f0;

					&:last-child {
						border-bottom: none;
					}

					.common-icon {
						margin-right: 20rpx;
					}

					.common-info {
						flex: 1;

						.common-title {
							font-size: 30rpx;
							color: #333;
							display: block;
							margin-bottom: 6rpx;
						}

						.common-scene {
							font-size: 24rpx;
							color: #999;
						}
					}
				}
			}
		}
	}
</style>
