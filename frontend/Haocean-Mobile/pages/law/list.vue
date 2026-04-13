<!-- 法律法规库列表页 - 左侧分类树 + 右侧书本列表 -->
<template>
	<view class="t-law-list">
		<!-- 左侧分类导航 -->
		<scroll-view scroll-y class="category-scroll">
			<view class="category-header">
				<text>分类</text>
			</view>
			<view class="category-item"
				v-for="(item, index) in categoryList"
				:key="index"
				:class="{ active: currentCategoryIndex === index }"
				@click="handleCategoryClick(index, item)">
				<text class="category-name">{{ item.name }}</text>
				<view class="indicator" v-if="currentCategoryIndex === index"></view>
			</view>
		</scroll-view>

		<!-- 右侧书本列表 -->
		<scroll-view scroll-y class="book-scroll" refresher-enabled @refresherrefresh="handleRefresh">
			<view class="current-category-name" v-if="currentCategory">{{ currentCategory.name }}</view>

			<view class="book-list">
				<view class="book-card" v-for="(item, index) in bookList" :key="index" @click="handleBookClick(item)">
					<view class="book-header">
						<view class="book-icon">
							<u-icon name="book" size="40" color="#1890ff"></u-icon>
						</view>
						<view class="book-info">
							<text class="book-title">{{ item.title }}</text>
							<view class="book-tags" v-if="item.legalType || item.supervisionType">
								<u-tag v-if="item.legalType" :text="getLegalTypeName(item.legalType)" type="primary" size="mini"></u-tag>
								<u-tag v-if="item.supervisionType" :text="getSupervisionTypeName(item.supervisionType)" type="success" size="mini"></u-tag>
							</view>
						</view>
					</view>
					<view class="book-meta">
						<text class="meta-item">发布部门：{{ item.department || '未知' }}</text>
						<text class="meta-item">施行日期：{{ item.publishDate || '未知' }}</text>
					</view>
					<view class="book-footer">
						<text class="chapter-count">共 {{ item.chapterCount || 0 }} 章</text>
						<u-icon name="arrow-right" size="24" color="#999"></u-icon>
					</view>
				</view>

				<u-empty v-if="bookList.length === 0" text="暂无法律法规" mode="list"></u-empty>
			</view>
		</scroll-view>
	</view>
</template>

<script>
	import { listBook, listLegalType, listSupervisionType } from '@/api/law.js'

	export default {
		data() {
			return {
				legalType: '',
				legalTypeName: '',
				supervisionType: '',
				supervisionTypeName: '',
				categoryList: [],
				currentCategory: null,
				currentCategoryIndex: 0,
				bookList: [],
				legalTypeMap: {},
				supervisionTypeMap: {}
			}
		},
		onLoad(options) {
			if (options.legalType) {
				this.legalType = options.legalType
				this.legalTypeName = options.legalTypeName || ''
			}
			if (options.supervisionType) {
				this.supervisionType = options.supervisionType
				this.supervisionTypeName = options.supervisionTypeName || ''
			}
			this.getCategoryList()
			this.getTypeMapping()
		},
		methods: {
			// 获取分类列表
			getCategoryList() {
				// 如果有法律类型筛选，只显示该类型
				if (this.legalType) {
					this.categoryList = [{
						name: this.legalTypeName,
						type: this.legalType
					}]
					this.currentCategory = this.categoryList[0]
					this.getBookList(this.legalType, '')
				} else if (this.supervisionType) {
					// 如果有监管类型筛选，获取法律类型列表
					listLegalType({ isEnabled: 1 }).then(res => {
						const types = res.rows || res || []
						this.categoryList = types.map(item => ({
							name: item.name,
							type: item.code
						}))
						if (this.categoryList.length > 0) {
							this.currentCategory = this.categoryList[0]
							this.getBookList(this.categoryList[0].type, this.supervisionType)
						}
					})
				} else {
					// 没有筛选，获取所有法律类型
					listLegalType({ isEnabled: 1 }).then(res => {
						const types = res.rows || res || []
						this.categoryList = types.map(item => ({
							name: item.name,
							type: item.code
						}))
						if (this.categoryList.length > 0) {
							this.currentCategory = this.categoryList[0]
							this.getBookList(this.categoryList[0].type, '')
						}
					})
				}
			},

			// 获取类型映射
			getTypeMapping() {
				listLegalType({}).then(res => {
					const types = res.rows || res || []
					types.forEach(item => {
						this.legalTypeMap[item.code] = item.name
					})
				})
				listSupervisionType({}).then(res => {
					const types = res.rows || res || []
					types.forEach(item => {
						this.supervisionTypeMap[item.code] = item.name
					})
				})
			},

			// 获取书本列表
			getBookList(legalType, supervisionType) {
				const query = {}
				if (legalType) query.legalType = legalType
				if (supervisionType) query.supervisionType = supervisionType
				listBook(query).then(res => {
					this.bookList = res.data || res || []
				})
			},

			// 点击分类
			handleCategoryClick(index, item) {
				this.currentCategoryIndex = index
				this.currentCategory = item
				this.getBookList(item.type, this.supervisionType)
			},

			// 点击书本
			handleBookClick(item) {
				// 保存到最近浏览
				this.saveRecent(item)
				uni.navigateTo({
					url: '/pages/law/detail?id=' + item.regulationId
				})
			},

			// 获取法律类型名称
			getLegalTypeName(code) {
				return this.legalTypeMap[code] || code
			},

			// 获取监管类型名称
			getSupervisionTypeName(code) {
				return this.supervisionTypeMap[code] || code
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
			},

			// 下拉刷新
			handleRefresh() {
				setTimeout(() => {
					this.getCategoryList()
					uni.stopPullDownRefresh()
				}, 1000)
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-law-list {
		display: flex;
		height: 100vh;
		background: #f5f5f5;

		.category-scroll {
			width: 200rpx;
			background: #f8f8f8;
			border-right: 1rpx solid #e8e8e8;

			.category-header {
				padding: 30rpx 20rpx;
				background: #f0f0f0;
				text-align: center;

				text {
					font-size: 28rpx;
					font-weight: bold;
					color: #333;
				}
			}

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

		.book-scroll {
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

			.book-list {
				padding: 20rpx;

				.book-card {
					background: #fff;
					border-radius: 16rpx;
					padding: 30rpx;
					margin-bottom: 20rpx;

					.book-header {
						display: flex;
						align-items: flex-start;
						margin-bottom: 16rpx;

						.book-icon {
							width: 80rpx;
							height: 80rpx;
							background: #f5f5f5;
							border-radius: 16rpx;
							display: flex;
							justify-content: center;
							align-items: center;
							margin-right: 20rpx;
							flex-shrink: 0;
						}

						.book-info {
							flex: 1;

							.book-title {
								font-size: 30rpx;
								font-weight: bold;
								color: #333;
								display: block;
								margin-bottom: 12rpx;
							}

							.book-tags {
								display: flex;
								gap: 10rpx;
							}
						}
					}

					.book-meta {
						display: flex;
						justify-content: space-between;
						padding: 16rpx 0;
						border-top: 1rpx solid #f0f0f0;
						border-bottom: 1rpx solid #f0f0f0;

						.meta-item {
							font-size: 24rpx;
							color: #999;
						}
					}

					.book-footer {
						display: flex;
						justify-content: space-between;
						align-items: center;
						padding-top: 16rpx;

						.chapter-count {
							font-size: 24rpx;
							color: #999;
						}
					}
				}
			}
		}
	}
</style>
