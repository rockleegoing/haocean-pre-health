<!-- 法律法规库详情页 - 书籍封面和基本信息 + 章节目录 + 条款内容 -->
<template>
	<view class="t-law-detail">
		<!-- 书籍信息 -->
		<view class="book-info-section">
			<view class="book-icon">
				<u-icon name="book" size="80" color="#1890ff"></u-icon>
			</view>
			<view class="book-meta">
				<text class="book-title">{{ bookInfo.title }}</text>
				<view class="book-tags">
					<u-tag v-if="bookInfo.legalType" :text="legalTypeName" type="primary" size="mini"></u-tag>
					<u-tag v-if="bookInfo.supervisionType" :text="supervisionTypeName" type="success" size="mini"></u-tag>
				</view>
				<view class="book-extra">
					<text class="extra-item">发布部门：{{ bookInfo.department || '未知' }}</text>
					<text class="extra-item">施行日期：{{ bookInfo.publishDate || '未知' }}</text>
				</view>
			</view>
		</view>

		<!-- 收藏按钮 -->
		<view class="favorite-btn" @click="handleFavorite">
			<u-icon :name="isFavorite ? 'star-fill' : 'star'" size="40" :color="isFavorite ? '#ffac2d' : '#999'"></u-icon>
			<text>{{ isFavorite ? '已收藏' : '收藏' }}</text>
		</view>

		<!-- 章节目录 -->
		<view class="chapter-section">
			<view class="section-title">章节目录</view>
			<view class="chapter-list">
				<view class="chapter-item" v-for="(chapter, index) in chapterList" :key="index" @click="handleChapterClick(chapter)">
					<view class="chapter-header">
						<view class="chapter-index">{{ index + 1 }}</view>
						<text class="chapter-title">{{ chapter.title }}</text>
						<u-icon :name="chapter.expanded ? 'arrow-up' : 'arrow-down'" size="24" color="#999"></u-icon>
					</view>
					<view class="chapter-content" v-if="chapter.expanded">
						<view class="article-list">
							<view class="article-item" v-for="(article, aIndex) in chapter.articles" :key="aIndex">
								<text class="article-title">{{ article.title }}</text>
								<text class="article-content">{{ article.content }}</text>
							</view>
						</view>
					</view>
				</view>
			</view>
		</view>

		<!-- 定性依据 -->
		<view class="basis-section" v-if="basisList.length > 0">
			<view class="section-title">定性依据</view>
			<view class="basis-list">
				<view class="basis-item" v-for="(item, index) in basisList" :key="index" @click="handleBasisClick(item)">
					<view class="basis-header">
						<text class="basis-title">{{ item.title }}</text>
						<u-icon name="arrow-right" size="24" color="#999"></u-icon>
					</view>
					<text class="basis-content">{{ item.content }}</text>
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
	import { getBookDetail, listLegalType, listSupervisionType } from '@/api/law.js'

	export default {
		data() {
			return {
				regulationId: 0,
				bookInfo: {},
				chapterList: [],
				basisList: [],
				legalTypeName: '',
				supervisionTypeName: '',
				isFavorite: false,
				legalTypeMap: {},
				supervisionTypeMap: {}
			}
		},
		onLoad(options) {
			if (options.id) {
				this.regulationId = parseInt(options.id)
				this.getBookDetail()
				this.getTypeMapping()
				this.checkFavorite()
			}
		},
		methods: {
			// 获取书本详情
			getBookDetail() {
				getBookDetail(this.regulationId).then(res => {
					const data = res.data || res || {}
					this.bookInfo = data.regulation || {}
					this.chapterList = (data.chapters || []).map(chapter => ({
						...chapter,
						expanded: false,
						articles: []
					}))
					this.basisList = data.basisList || []

					// 获取法律类型名称
					if (this.bookInfo.legalType) {
						this.legalTypeName = this.legalTypeMap[this.bookInfo.legalType] || this.bookInfo.legalType
					}
					// 获取监管类型名称
					if (this.bookInfo.supervisionType) {
						this.supervisionTypeName = this.supervisionTypeMap[this.bookInfo.supervisionType] || this.bookInfo.supervisionType
					}
				})
			},

			// 获取类型映射
			getTypeMapping() {
				listLegalType({}).then(res => {
					const types = res.rows || res || []
					types.forEach(item => {
						this.legalTypeMap[item.code] = item.name
					})
					if (this.bookInfo.legalType) {
						this.legalTypeName = this.legalTypeMap[this.bookInfo.legalType] || this.bookInfo.legalType
					}
				})
				listSupervisionType({}).then(res => {
					const types = res.rows || res || []
					types.forEach(item => {
						this.supervisionTypeMap[item.code] = item.name
					})
					if (this.bookInfo.supervisionType) {
						this.supervisionTypeName = this.supervisionTypeMap[this.bookInfo.supervisionType] || this.bookInfo.supervisionType
					}
				})
			},

			// 检查是否已收藏
			checkFavorite() {
				const favorites = uni.getStorageSync('lawFavorites') || '[]'
				const favoriteList = JSON.parse(favorites)
				this.isFavorite = favoriteList.some(item => item.id === this.regulationId)
			},

			// 收藏/取消收藏
			handleFavorite() {
				let favorites = uni.getStorageSync('lawFavorites') || '[]'
				let favoriteList = JSON.parse(favorites)

				if (this.isFavorite) {
					// 取消收藏
					const index = favoriteList.findIndex(item => item.id === this.regulationId)
					if (index > -1) {
						favoriteList.splice(index, 1)
						this.isFavorite = false
						uni.showToast({ title: '已取消收藏', icon: 'success' })
					}
				} else {
					// 添加收藏
					favoriteList.push({
						id: this.regulationId,
						title: this.bookInfo.title,
						time: this.formatTime(new Date())
					})
					this.isFavorite = true
					uni.showToast({ title: '收藏成功', icon: 'success' })
				}

				uni.setStorageSync('lawFavorites', JSON.stringify(favoriteList))
			},

			// 点击章节
			handleChapterClick(chapter) {
				// 切换展开状态
				chapter.expanded = !chapter.expanded

				// 如果展开且没有加载过条款，加载条款
				if (chapter.expanded && chapter.articles.length === 0) {
					// 这里可以调用 API 加载章节详细内容
					// getChapterContent(chapter.chapterId).then(res => { ... })
					// 暂时使用章节数据中的条款
					chapter.articles = chapter.articles || []
				}
			},

			// 点击定性依据
			handleBasisClick(item) {
				uni.showToast({ title: '查看定性依据详情', icon: 'none' })
			},

			// 复制
			handleCopy() {
				let text = this.bookInfo.title + '\n\n'
				this.chapterList.forEach((chapter, index) => {
					text += chapter.title + '\n'
					if (chapter.articles && chapter.articles.length > 0) {
						chapter.articles.forEach(article => {
							text += article.title + ': ' + article.content + '\n'
						})
					}
					text += '\n'
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
				uni.showShareMenu({
					withShareTicket: true
				})
			},

			// 打印
			handlePrint() {
				uni.showToast({ title: '打印功能待开发', icon: 'none' })
			},

			// 格式化时间
			formatTime(date) {
				const year = date.getFullYear()
				const month = (date.getMonth() + 1).toString().padStart(2, '0')
				const day = date.getDate().toString().padStart(2, '0')
				return `${year}-${month}-${day}`
			}
		},
		onShareAppMessage() {
			return {
				title: this.bookInfo.title,
				path: '/pages/law/detail?id=' + this.regulationId
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-law-detail {
		padding: 20rpx;
		background: #f5f5f5;
		min-height: 100vh;
		padding-bottom: 140rpx;

		.book-info-section {
			background: #fff;
			padding: 30rpx;
			border-radius: 16rpx;
			margin-bottom: 20rpx;
			display: flex;
			align-items: flex-start;

			.book-icon {
				width: 120rpx;
				height: 120rpx;
				background: #f5f5f5;
				border-radius: 24rpx;
				display: flex;
				justify-content: center;
				align-items: center;
				margin-right: 20rpx;
				flex-shrink: 0;
			}

			.book-meta {
				flex: 1;

				.book-title {
					font-size: 36rpx;
					font-weight: bold;
					color: #333;
					display: block;
					margin-bottom: 16rpx;
				}

				.book-tags {
					display: flex;
					gap: 10rpx;
					margin-bottom: 16rpx;
				}

				.book-extra {
					display: flex;
					flex-direction: column;
					gap: 8rpx;

					.extra-item {
						font-size: 24rpx;
						color: #999;
					}
				}
			}
		}

		.favorite-btn {
			background: #fff;
			padding: 20rpx 30rpx;
			border-radius: 16rpx;
			margin-bottom: 20rpx;
			display: flex;
			align-items: center;
			justify-content: center;
			gap: 10rpx;

			text {
				font-size: 28rpx;
				color: #666;
			}
		}

		.chapter-section {
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
				border-left: 4rpx solid #1890ff;
			}

			.chapter-list {
				.chapter-item {
					border-bottom: 1rpx solid #f0f0f0;

					&:last-child {
						border-bottom: none;
					}

					.chapter-header {
						display: flex;
						align-items: center;
						padding: 24rpx 0;

						.chapter-index {
							width: 50rpx;
							height: 50rpx;
							background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
							border-radius: 50%;
							display: flex;
							justify-content: center;
							align-items: center;
							margin-right: 20rpx;
							flex-shrink: 0;

							text {
								font-size: 24rpx;
								font-weight: bold;
								color: #fff;
							}
						}

						.chapter-title {
							flex: 1;
							font-size: 30rpx;
							color: #333;
							font-weight: bold;
						}
					}

					.chapter-content {
						background: #f9f9f9;
						border-radius: 8rpx;
						padding: 20rpx;
						margin-bottom: 20rpx;

						.article-list {
							.article-item {
								padding: 20rpx 0;
								border-bottom: 1rpx solid #e8e8e8;

								&:last-child {
									border-bottom: none;
								}

								.article-title {
									font-size: 28rpx;
									font-weight: bold;
									color: #333;
									display: block;
									margin-bottom: 12rpx;
								}

								.article-content {
									font-size: 26rpx;
									color: #666;
									line-height: 1.8;
									display: block;
								}
							}
						}
					}
				}
			}
		}

		.basis-section {
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
				border-left: 4rpx solid #1890ff;
			}

			.basis-list {
				.basis-item {
					padding: 24rpx 0;
					border-bottom: 1rpx solid #f0f0f0;

					&:last-child {
						border-bottom: none;
					}

					.basis-header {
						display: flex;
						justify-content: space-between;
						align-items: center;
						margin-bottom: 12rpx;

						.basis-title {
							font-size: 30rpx;
							font-weight: bold;
							color: #333;
						}
					}

					.basis-content {
						font-size: 26rpx;
						color: #666;
						line-height: 1.8;
						display: block;
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
