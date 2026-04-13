<!-- 规范用语详情页 - 规范条目列表 + 内容展示 -->
<template>
	<view class="t-standard-phrase-detail">
		<!-- 条目信息 -->
		<view class="item-header">
			<text class="item-title">{{ title }}</text>
			<text class="item-scene" v-if="itemInfo.scene">{{ itemInfo.scene }}</text>
		</view>

		<!-- 规范内容列表 -->
		<view class="content-section">
			<view class="section-title">规范内容</view>
			<view class="content-list">
				<view class="content-card" v-for="(content, index) in contentList" :key="index">
					<view class="content-number">
						<text>{{ index + 1 }}</text>
					</view>
					<view class="content-body">
						<text class="content-text">{{ content.content }}</text>
						<view class="content-footer" v-if="content.legalBasis || content.tips">
							<view class="legal-basis" v-if="content.legalBasis">
								<text class="label">法律依据：</text>
								<text class="text">{{ content.legalBasis }}</text>
							</view>
							<view class="tips" v-if="content.tips">
								<text class="label">提示要点：</text>
								<text class="text">{{ content.tips }}</text>
							</view>
						</view>
					</view>
				</view>
			</view>
		</view>

		<!-- 底部操作栏 -->
		<view class="bottom-bar">
			<view class="action-btn" @click="handleCopy">
				<u-icon name="copy" size="40" color="#1890ff"></u-icon>
				<text>复制</text>
			</view>
			<view class="action-btn" @click="handleShare">
				<u-icon name="share" size="40" color="#1890ff"></u-icon>
				<text>分享</text>
			</view>
			<view class="action-btn" @click="handlePrint">
				<u-icon name="printer" size="40" color="#1890ff"></u-icon>
				<text>打印</text>
			</view>
		</view>
	</view>
</template>

<script>
	import { getItem, listContent } from '@/api/standardPhrase.js'

	export default {
		data() {
			return {
				itemId: 0,
				title: '',
				itemInfo: {},
				contentList: []
			}
		},
		onLoad(options) {
			if (options.itemId) {
				this.itemId = parseInt(options.itemId)
				this.title = options.title || '规范详情'
				this.getItemDetail()
				this.getContentList()
			}
		},
		methods: {
			// 获取条目详情
			getItemDetail() {
				getItem(this.itemId).then(res => {
					this.itemInfo = res.data || res || {}
				})
			},

			// 获取内容列表
			getContentList() {
				listContent({ itemId: this.itemId }).then(res => {
					this.contentList = res.rows || res || []
				})
			},

			// 复制
			handleCopy() {
				if (this.contentList.length === 0) {
					uni.showToast({ title: '暂无内容可复制', icon: 'none' })
					return
				}
				let text = this.title + '\n\n'
				this.contentList.forEach((item, index) => {
					text += (index + 1) + '. ' + item.content + '\n'
					if (item.legalBasis) {
						text += '   法律依据：' + item.legalBasis + '\n'
					}
				})
				uni.setClipboardData({
					data: text,
					success: () => {
						uni.showToast({ title: '复制成功', icon: 'success' })
					}
				})
			},

			// 分享
			handleShare() {
				uni.showShareMenu()
			},

			// 打印
			handlePrint() {
				uni.showToast({ title: '打印功能待开发', icon: 'none' })
			}
		},
		onShareAppMessage() {
			return {
				title: this.title,
				path: '/pages/standard-phrase/detail?itemId=' + this.itemId
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-standard-phrase-detail {
		padding: 20rpx;
		background: #f5f5f5;
		min-height: 100vh;
		padding-bottom: 140rpx;

		.item-header {
			background: #fff;
			padding: 30rpx;
			border-radius: 16rpx;
			margin-bottom: 20rpx;

			.item-title {
				font-size: 36rpx;
				font-weight: bold;
				color: #333;
				display: block;
				margin-bottom: 12rpx;
			}

			.item-scene {
				font-size: 26rpx;
				color: #999;
				background: #f5f5f5;
				padding: 8rpx 16rpx;
				border-radius: 8rpx;
			}
		}

		.content-section {
			background: #fff;
			padding: 30rpx;
			border-radius: 16rpx;
			margin-bottom: 20rpx;

			.section-title {
				font-size: 32rpx;
				font-weight: bold;
				color: #333;
				margin-bottom: 20rpx;
				padding-left: 10rpx;
				border-left: 4rpx solid #1890ff;
			}

			.content-list {
				.content-card {
					display: flex;
					padding: 24rpx 0;
					border-bottom: 1rpx solid #f0f0f0;

					&:last-child {
						border-bottom: none;
					}

					.content-number {
						width: 60rpx;
						height: 60rpx;
						background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
						border-radius: 50%;
						display: flex;
						justify-content: center;
						align-items: center;
						margin-right: 20rpx;
						flex-shrink: 0;

						text {
							font-size: 28rpx;
							font-weight: bold;
							color: #fff;
						}
					}

					.content-body {
						flex: 1;

						.content-text {
							font-size: 30rpx;
							color: #333;
							line-height: 1.8;
							display: block;
							margin-bottom: 16rpx;
						}

						.content-footer {
							background: #f9f9f9;
							padding: 16rpx;
							border-radius: 8rpx;

							.legal-basis, .tips {
								margin-bottom: 12rpx;

								&:last-child {
									margin-bottom: 0;
								}

								.label {
									font-size: 24rpx;
									color: #999;
									font-weight: bold;
								}

								.text {
									font-size: 24rpx;
									color: #666;
									line-height: 1.6;
								}
							}
						}
					}
				}
			}
		}

		.bottom-bar {
			position: fixed;
			bottom: 0;
			left: 0;
			right: 0;
			height: 120rpx;
			background: #fff;
			display: flex;
			justify-content: space-around;
			align-items: center;
			box-shadow: 0 -2rpx 10rpx rgba(0, 0, 0, 0.05);
			padding-bottom: env(safe-area-inset-bottom);

			.action-btn {
				display: flex;
				flex-direction: column;
				align-items: center;
				color: #666;
				font-size: 24rpx;

				text {
					margin-top: 6rpx;
				}
			}
		}
	}
</style>
