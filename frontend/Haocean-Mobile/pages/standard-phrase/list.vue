<!-- 规范用语列表页 - 左侧分类 + 右侧列表 -->
<template>
	<view class="t-standard-phrase-list">
		<!-- 左侧分类导航 -->
		<scroll-view scroll-y class="category-scroll">
			<view class="category-item"
				v-for="(item, index) in categoryList"
				:key="index"
				:class="{ active: currentCategoryIndex === index }"
				@click="handleCategoryClick(index, item)">
				<text class="category-name">{{ item.name }}</text>
				<view class="indicator" v-if="currentCategoryIndex === index"></view>
			</view>
		</scroll-view>

		<!-- 右侧条目列表 -->
		<scroll-view scroll-y class="item-scroll">
			<view class="current-category-name" v-if="currentCategory">{{ currentCategory.name }}</view>

			<view class="item-list">
				<view class="item-card" v-for="(item, index) in currentItemList" :key="index" @click="handleItemClick(item)">
					<view class="item-header">
						<text class="item-title">{{ item.title }}</text>
						<text class="item-scene" v-if="item.scene">{{ item.scene }}</text>
					</view>
					<view class="item-contents" v-if="item.contents && item.contents.length > 0">
						<view class="content-preview" v-for="(content, cIndex) in item.contents" :key="cIndex">
							<text class="content-text">{{ content.content }}</text>
						</view>
					</view>
					<view class="item-footer">
						<text class="content-count">共 {{ item.contents ? item.contents.length : 0 }} 条内容</text>
						<u-icon name="arrow-right" size="24" color="#999"></u-icon>
					</view>
				</view>

				<u-empty v-if="currentItemList.length === 0" text="暂无规范条目" mode="list"></u-empty>
			</view>
		</scroll-view>
	</view>
</template>

<script>
	import { listCategory, listItem } from '@/api/standardPhrase.js'

	export default {
		data() {
			return {
				supervisionTypeId: 0,
				supervisionTypeName: '',
				categoryList: [],
				currentCategory: null,
				currentCategoryIndex: 0,
				currentItemList: []
			}
		},
		onLoad(options) {
			if (options.supervisionTypeId) {
				this.supervisionTypeId = parseInt(options.supervisionTypeId)
				this.supervisionTypeName = options.supervisionTypeName || ''
				this.getCategoryList()
			}
		},
		methods: {
			// 获取类别列表
			getCategoryList() {
				listCategory({ supervisionTypeId: this.supervisionTypeId, isEnabled: 1 }).then(res => {
					this.categoryList = res.rows || res || []
					if (this.categoryList.length > 0) {
						this.currentCategory = this.categoryList[0]
						this.getItemList(this.currentCategory.id)
					}
				})
			},

			// 获取条目列表
			getItemList(categoryId) {
				listItem({ categoryId: categoryId, isEnabled: 1 }).then(res => {
					this.currentItemList = res.rows || res || []
				})
			},

			// 点击分类
			handleCategoryClick(index, item) {
				this.currentCategoryIndex = index
				this.currentCategory = item
				this.getItemList(item.id)
			},

			// 点击条目
			handleItemClick(item) {
				uni.navigateTo({
					url: '/pages/standard-phrase/detail?itemId=' + item.id + '&title=' + item.title
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-standard-phrase-list {
		display: flex;
		height: 100vh;
		background: #f5f5f5;

		.category-scroll {
			width: 200rpx;
			background: #f8f8f8;
			border-right: 1rpx solid #e8e8e8;

			.category-item {
				display: flex;
				align-items: center;
				justify-content: center;
				padding: 30rpx 20rpx;
				position: relative;
				background: #f8f8f8;
				border-bottom: 1rpx solid #e8e8e8;

				&.active {
					background: #fff;

					.category-name {
						color: #1890ff;
						font-weight: bold;
					}
				}

				.indicator {
					position: absolute;
					left: 0;
					top: 50%;
					transform: translateY(-50%);
					width: 6rpx;
					height: 40rpx;
					background: #1890ff;
					border-radius: 0 4rpx 4rpx 0;
				}

				.category-name {
					font-size: 28rpx;
					color: #333;
				}
			}
		}

		.item-scroll {
			flex: 1;
			overflow-y: auto;

			.current-category-name {
				padding: 20rpx 30rpx;
				background: #fff;
				font-size: 30rpx;
				font-weight: bold;
				color: #333;
				border-bottom: 1rpx solid #e8e8e8;
			}

			.item-list {
				padding: 20rpx;

				.item-card {
					background: #fff;
					border-radius: 16rpx;
					padding: 30rpx;
					margin-bottom: 20rpx;

					.item-header {
						display: flex;
						justify-content: space-between;
						align-items: center;
						margin-bottom: 16rpx;

						.item-title {
							font-size: 30rpx;
							font-weight: bold;
							color: #333;
							flex: 1;
						}

						.item-scene {
							font-size: 24rpx;
							color: #999;
							background: #f5f5f5;
							padding: 6rpx 12rpx;
							border-radius: 8rpx;
							margin-left: 16rpx;
						}
					}

					.item-contents {
						margin-bottom: 16rpx;

						.content-preview {
							background: #f9f9f9;
							padding: 16rpx;
							border-radius: 8rpx;
							margin-bottom: 12rpx;

							.content-text {
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

					.item-footer {
						display: flex;
						justify-content: space-between;
						align-items: center;
						padding-top: 16rpx;
						border-top: 1rpx solid #f0f0f0;

						.content-count {
							font-size: 24rpx;
							color: #999;
						}
					}
				}
			}
		}
	}
</style>
