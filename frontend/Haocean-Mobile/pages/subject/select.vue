<!-- 监管单位选择页面 -->
<template>
	<view class="t-subject-select">
		<!-- 搜索栏 -->
		<view class="search-bar">
			<u-search v-model="keyword" placeholder="搜索单位名称" shape="round" @search="handleSearch" @custom="handleSearch"></u-search>
		</view>

		<!-- 行业分类筛选 -->
		<view class="filter-bar">
			<u-dropdown-menu>
				<u-dropdown-item v-model="industryId" :options="industryOptions" @change="handleIndustryChange"></u-dropdown-item>
			</u-dropdown-menu>
		</view>

		<!-- 单位列表 -->
		<view class="subject-list">
			<u-loadmore :status="loadStatus" />

			<view class="subject-item" v-for="(item, index) in subjectList" :key="index" @click="handleSelect(item)">
				<view class="subject-header">
					<text class="subject-name">{{ item.name }}</text>
					<u-tag :text="item.status === 1 ? '正常' : '停业'" :type="item.status === 1 ? 'success' : 'danger'" size="mini"></u-tag>
				</view>
				<view class="subject-info">
					<text class="info-text">行业：{{ item.industryName }}</text>
				</view>
				<view class="subject-info">
					<text class="info-text">地址：{{ item.address || '暂无' }}</text>
				</view>
				<view class="subject-info">
					<text class="info-text">联系人：{{ item.contactPerson || '暂无' }}</text>
					<text class="info-text">电话：{{ item.contactPhone || '暂无' }}</text>
				</view>
				<view class="subject-footer">
					<u-button text="查看详情" size="small" @click.stop="handleViewDetail(item)"></u-button>
				</view>
			</view>

			<u-loadmore :status="loadStatus" @loadmore="loadMore" />
		</view>

		<!-- 添加单位按钮 -->
		<view class="add-btn" @click="handleAdd">
			<u-icon name="plus" size="40" color="#fff"></u-icon>
		</view>
	</view>
</template>

<script>
	import { listSubject } from '@/api/subject.js'
	import { listIndustry } from '@/api/industry.js'

	export default {
		data() {
			return {
				keyword: '',
				industryId: 0,
				industryOptions: [],
				subjectList: [],
				queryParams: {
					pageNum: 1,
					pageSize: 10,
					industryId: 0
				},
				total: 0,
				loadStatus: 'loadmore'
			}
		},
		onLoad() {
			this.getIndustryList()
			this.getSubjectList()
		},
		methods: {
			// 获取行业分类
			getIndustryList() {
				listIndustry({ isEnabled: 1 }).then(res => {
					const industries = res.rows || []
					this.industryOptions = [
						{ label: '全部行业', value: 0 },
						...industries.map(item => ({ label: item.industryName, value: item.industryId }))
					]
				})
			},

			// 获取单位列表
			getSubjectList() {
				this.loadStatus = 'loading'
				listSubject(this.queryParams).then(res => {
					this.subjectList = res.rows || []
					this.total = res.total || 0
					this.loadStatus = this.subjectList.length < this.total ? 'loadmore' : 'nomore'
				}).catch(() => {
					this.loadStatus = 'loadmore'
				})
			},

			// 搜索
			handleSearch() {
				this.queryParams.pageNum = 1
				this.queryParams.name = this.keyword
				this.getSubjectList()
			},

			// 行业筛选
			handleIndustryChange(value) {
				this.queryParams.pageNum = 1
				this.queryParams.industryId = value
				this.getSubjectList()
			},

			// 加载更多
			loadMore() {
				if (this.subjectList.length >= this.total) {
					this.loadStatus = 'nomore'
					return
				}
				this.queryParams.pageNum++
				listSubject(this.queryParams).then(res => {
					this.subjectList = [...this.subjectList, ...(res.rows || [])]
					this.loadStatus = this.subjectList.length < this.total ? 'loadmore' : 'nomore'
				})
			},

			// 选择单位
			handleSelect(item) {
				// 回调给上级页面
				uni.$u.toast('已选择：' + item.name)
			},

			// 查看详情
			handleViewDetail(item) {
				uni.navigateTo({
					url: '/pages/subject/detail?id=' + item.subjectId
				})
			},

			// 添加单位
			handleAdd() {
				uni.navigateTo({
					url: '/pages/subject/add'
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-subject-select {
		padding: 20rpx;

		.search-bar {
			margin-bottom: 20rpx;
		}

		.filter-bar {
			margin-bottom: 20rpx;
		}

		.subject-list {
			.subject-item {
				background: #fff;
				border-radius: 16rpx;
				padding: 30rpx;
				margin-bottom: 20rpx;

				.subject-header {
					display: flex;
					justify-content: space-between;
					align-items: center;
					margin-bottom: 16rpx;

					.subject-name {
						font-size: 32rpx;
						font-weight: bold;
					}
				}

				.subject-info {
					display: flex;
					justify-content: space-between;
					margin-bottom: 10rpx;

					.info-text {
						font-size: 26rpx;
						color: #666;
					}
				}

				.subject-footer {
					margin-top: 20rpx;
					text-align: right;
				}
			}
		}

		.add-btn {
			position: fixed;
			right: 40rpx;
			bottom: 100rpx;
			width: 100rpx;
			height: 100rpx;
			background: linear-gradient(135deg, #1890ff, #096dd9);
			border-radius: 50%;
			display: flex;
			justify-content: center;
			align-items: center;
			box-shadow: 0 4rpx 20rpx rgba(24, 144, 255, 0.4);
			z-index: 99;
		}
	}
</style>
