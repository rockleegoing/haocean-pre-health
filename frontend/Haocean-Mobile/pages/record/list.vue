<!-- 执法记录列表页面 -->
<template>
	<view class="t-record-list">
		<!-- 搜索筛选 -->
		<view class="search-bar">
			<u-form :model="queryParams" inline>
				<u-form-item label="检查类型">
					<u-select v-model="queryParams.checkType" :list="checkTypeOptions" placeholder="全部" @change="handleQuery"></u-select>
				</u-form-item>
				<u-form-item label="状态">
					<u-select v-model="queryParams.status" :list="statusOptions" placeholder="全部" @change="handleQuery"></u-select>
				</u-form-item>
			</u-form>
		</view>

		<!-- 记录列表 -->
		<view class="list-container">
			<u-empty v-if="recordList.length === 0" text="暂无记录" mode="list"></u-empty>

			<view class="record-item" v-for="(item, index) in recordList" :key="index" @click="handleViewDetail(item)">
				<view class="record-header">
					<text class="record-no">{{ item.recordNo }}</text>
					<u-tag :text="getStatusText(item.status)" :type="getStatusType(item.status)" size="mini"></u-tag>
				</view>
				<view class="record-body">
					<view class="info-row">
						<text class="label">单位：</text>
						<text class="value">{{ item.subjectName }}</text>
					</view>
					<view class="info-row">
						<text class="label">行业：</text>
						<text class="value">{{ item.industryName }}</text>
					</view>
					<view class="info-row">
						<text class="label">类型：</text>
						<text class="value">{{ item.checkType }}</text>
					</view>
					<view class="info-row">
						<text class="label">日期：</text>
						<text class="value">{{ item.checkDate }}</text>
					</view>
					<view class="info-row">
						<text class="label">证据：</text>
						<text class="value">{{ item.evidenceCount }} 个</text>
						<text class="label" style="margin-left: 20rpx;">文书：</text>
						<text class="value">{{ item.documentCount }} 份</text>
					</view>
				</view>
				<view class="record-footer">
					<u-button v-if="item.status === 1" text="上报" size="small" @click.stop="handleSubmit(item)"></u-button>
					<u-button text="查看" size="small" type="primary" @click.stop="handleViewDetail(item)"></u-button>
				</view>
			</view>

			<u-loadmore :status="loadStatus" @loadmore="loadMore" />
		</view>

		<!-- 新建按钮 -->
		<view class="add-btn" @click="handleCreate">
			<u-icon name="plus" size="40" color="#fff"></u-icon>
		</view>
	</view>
</template>

<script>
	import { listRecord, submitRecord } from '@/api/record.js'

	export default {
		data() {
			return {
				recordList: [],
				queryParams: {
					pageNum: 1,
					pageSize: 10,
					subjectId: 0,
					checkType: '',
					status: -1
				},
				total: 0,
				loadStatus: 'loadmore',
				checkTypeOptions: [
					{ value: '', label: '全部' },
					{ value: '日常检查', label: '日常检查' },
					{ value: '专项检查', label: '专项检查' },
					{ value: '投诉举报', label: '投诉举报' },
					{ value: '双随机抽查', label: '双随机抽查' }
				],
				statusOptions: [
					{ value: -1, label: '全部' },
					{ value: 0, label: '草稿' },
					{ value: 1, label: '待上报' },
					{ value: 2, label: '已上报' },
					{ value: 3, label: '已审核' },
					{ value: 4, label: '已归档' }
				]
			}
		},
		onLoad(options) {
			if (options.subjectId) {
				this.queryParams.subjectId = parseInt(options.subjectId)
			}
			this.getRecordList()
		},
		onShow() {
			this.getRecordList()
		},
		methods: {
			// 获取记录列表
			getRecordList() {
				this.loadStatus = 'loading'
				listRecord(this.queryParams).then(res => {
					this.recordList = res.rows || []
					this.total = res.total || 0
					this.loadStatus = this.recordList.length < this.total ? 'loadmore' : 'nomore'
				}).catch(() => {
					this.loadStatus = 'loadmore'
				})
			},

			// 搜索
			handleQuery() {
				this.queryParams.pageNum = 1
				this.getRecordList()
			},

			// 加载更多
			loadMore() {
				if (this.recordList.length >= this.total) {
					this.loadStatus = 'nomore'
					return
				}
				this.queryParams.pageNum++
				listRecord(this.queryParams).then(res => {
					this.recordList = [...this.recordList, ...(res.rows || [])]
					this.loadStatus = this.recordList.length < this.total ? 'loadmore' : 'nomore'
				})
			},

			// 获取状态文本
			getStatusText(status) {
				const statusMap = { 0: '草稿', 1: '待上报', 2: '已上报', 3: '已审核', 4: '已归档' }
				return statusMap[status] || '未知'
			},

			// 获取状态类型
			getStatusType(status) {
				const typeMap = { 0: 'info', 1: 'warning', 2: 'success', 3: 'primary', 4: 'info' }
				return typeMap[status] || 'info'
			},

			// 查看详情
			handleViewDetail(item) {
				uni.navigateTo({
					url: '/pages/record/detail?id=' + item.recordId
				})
			},

			// 上报
			handleSubmit(item) {
				uni.$u.modal('确认上报', '确定要上报该执法记录吗？').then(() => {
					return submitRecord(item.recordId)
				}).then(() => {
					uni.$u.toast('上报成功')
					this.getRecordList()
				})
			},

			// 新建
			handleCreate() {
				uni.navigateTo({
					url: '/pages/record/create' + (this.queryParams.subjectId ? '?subjectId=' + this.queryParams.subjectId : '')
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-record-list {
		padding: 20rpx;

		.search-bar {
			background: #fff;
			border-radius: 16rpx;
			padding: 20rpx;
			margin-bottom: 20rpx;
		}

		.list-container {
			.record-item {
				background: #fff;
				border-radius: 16rpx;
				padding: 30rpx;
				margin-bottom: 20rpx;

				.record-header {
					display: flex;
					justify-content: space-between;
					align-items: center;
					margin-bottom: 20rpx;

					.record-no {
						font-size: 30rpx;
						font-weight: bold;
					}
				}

				.record-body {
					.info-row {
						display: flex;
						margin-bottom: 12rpx;

						.label {
							color: #666;
							width: 100rpx;
							flex-shrink: 0;
						}

						.value {
							color: #333;
							flex: 1;
						}
					}
				}

				.record-footer {
					display: flex;
					justify-content: flex-end;
					margin-top: 20rpx;
					gap: 20rpx;
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
