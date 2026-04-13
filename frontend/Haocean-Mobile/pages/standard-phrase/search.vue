<!-- 规范用语搜索页 -->
<template>
	<view class="t-standard-phrase-search">
		<!-- 搜索栏 -->
		<view class="search-section">
			<u-search v-model="keyword" placeholder="搜索规范用语内容或法律依据" shape="round" @search="handleSearch" />
		</view>

		<!-- 搜索结果 -->
		<view class="result-section" v-if="hasSearched">
			<view class="result-header">
				<text class="result-count">找到 {{ resultList.length }} 条结果</text>
			</view>

			<view class="result-list">
				<view class="result-item" v-for="(item, index) in resultList" :key="index" @click="handleItemClick(item)">
					<view class="result-path">
						<u-tag :text="getSupervisionTypeName(item)" type="primary" size="mini"></u-tag>
						<text class="path-category" v-if="item.item && item.item.category">{{ item.item.category.name }}</text>
						<text class="path-item" v-if="item.item">{{ item.item.title }}</text>
					</view>
					<view class="result-content">
						<text class="content-text" v-if="item.content">{{ highlightText(item.content) }}</text>
					</view>
					<view class="result-legal" v-if="item.legalBasis">
						<text class="legal-label">法律依据：</text>
						<text class="legal-text">{{ item.legalBasis }}</text>
					</view>
				</view>
			</view>

			<u-empty v-if="resultList.length === 0" text="未找到相关结果" mode="search"></u-empty>
		</view>

		<!-- 热门搜索 -->
		<view class="hot-section" v-else>
			<view class="section-title">热门搜索</view>
			<view class="hot-tags">
				<u-tag v-for="(tag, index) in hotKeywords" :key="index" :text="tag" type="info" plain size="small" @click="handleHotClick(tag)"></u-tag>
			</view>
		</view>
	</view>
</template>

<script>
	import { searchStandardPhrase, listSupervisionType } from '@/api/standardPhrase.js'

	export default {
		data() {
			return {
				keyword: '',
				hasSearched: false,
				resultList: [],
				supervisionTypeList: [],
				hotKeywords: ['执法证件', '检查目的', '权利义务', '整改意见', '法律依据', '行政处罚']
			}
		},
		onLoad() {
			this.getSupervisionTypeList()
		},
		methods: {
			// 获取监管类型列表（用于显示）
			getSupervisionTypeList() {
				listSupervisionType({ isEnabled: 1 }).then(res => {
					this.supervisionTypeList = res.rows || res || []
				})
			},

			// 搜索
			handleSearch(value) {
				if (!value) {
					uni.showToast({ title: '请输入搜索关键词', icon: 'none' })
					return
				}
				this.keyword = value
				this.hasSearched = true
				this.searchResult(value)
			},

			// 执行搜索
			searchResult(keyword) {
				searchStandardPhrase({ keyword: keyword }).then(res => {
					this.resultList = res.data || res || []
				})
			},

			// 获取监管类型名称
			getSupervisionTypeName(item) {
				if (item.item && item.item.category && item.item.category.supervisionType) {
					return item.item.category.supervisionType.name
				}
				return '规范用语'
			},

			// 高亮文本
			highlightText(text) {
				if (!this.keyword) return text
				return text
			},

			// 点击搜索结果
			handleItemClick(item) {
				if (item.itemId) {
					uni.navigateTo({
						url: '/pages/standard-phrase/detail?itemId=' + item.itemId
					})
				}
			},

			// 点击热门搜索
			handleHotClick(keyword) {
				this.keyword = keyword
				this.handleSearch(keyword)
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-standard-phrase-search {
		padding: 20rpx;
		background: #f5f5f5;
		min-height: 100vh;

		.search-section {
			background: #fff;
			padding: 20rpx;
			border-radius: 16rpx;
			margin-bottom: 20rpx;
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

					.result-path {
						display: flex;
						align-items: center;
						flex-wrap: wrap;
						margin-bottom: 16rpx;
						gap: 10rpx;

						.path-category {
							font-size: 26rpx;
							color: #666;
						}

						.path-item {
							font-size: 26rpx;
							color: #999;
						}
					}

					.result-content {
						margin-bottom: 16rpx;

						.content-text {
							font-size: 30rpx;
							color: #333;
							line-height: 1.6;
							display: block;
						}
					}

					.result-legal {
						background: #f9f9f9;
						padding: 12rpx;
						border-radius: 8rpx;

						.legal-label {
							font-size: 24rpx;
							color: #999;
							font-weight: bold;
						}

						.legal-text {
							font-size: 24rpx;
							color: #666;
						}
					}
				}
			}
		}

		.hot-section {
			background: #fff;
			padding: 30rpx;
			border-radius: 16rpx;

			.section-title {
				font-size: 32rpx;
				font-weight: bold;
				color: #333;
				margin-bottom: 20rpx;
			}

			.hot-tags {
				display: flex;
				flex-wrap: wrap;
				gap: 16rpx;
			}
		}
	}
</style>
