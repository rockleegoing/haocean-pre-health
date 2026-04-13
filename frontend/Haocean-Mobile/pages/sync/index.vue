<!-- 数据同步页面 -->
<template>
	<view class="t-sync">
		<!-- 同步状态概览 -->
		<view class="sync-status">
			<view class="status-card">
				<view class="status-value" :style="{ color: pendingCount > 0 ? '#faad14' : '#52c41a' }">{{ pendingCount }}</view>
				<view class="status-label">待同步</view>
			</view>
			<view class="status-card">
				<view class="status-value" style="color: #1890ff">{{ syncingCount }}</view>
				<view class="status-label">同步中</view>
			</view>
			<view class="status-card">
				<view class="status-value" style="color: #999">{{ syncedCount }}</view>
				<view class="status-label">已同步</view>
			</view>
		</view>

		<!-- 同步按钮 -->
		<view class="sync-actions">
			<u-button type="primary" text="立即同步" @click="handleSync" :loading="syncing"></u-button>
		</view>

		<!-- 同步记录列表 -->
		<view class="sync-list">
			<view class="list-header">
				<text class="title">同步记录</text>
			</view>

			<view class="sync-item" v-for="(item, index) in syncList" :key="index">
				<view class="item-header">
					<text class="item-table">{{ item.tableName }}</text>
					<u-tag :text="getActionText(item.action)" :type="getActionType(item.action)" size="mini"></u-tag>
				</view>
				<view class="item-info">
					<text class="info-text">记录 ID: {{ item.recordId }}</text>
					<text class="info-text">同步类型：{{ getSyncTypeText(item.syncType) }}</text>
				</view>
				<view class="item-footer">
					<text class="time">{{ item.createTime }}</text>
					<u-tag v-if="item.status === 'pending'" text="待同步" type="warning" size="mini"></u-tag>
					<u-tag v-else-if="item.status === 'syncing'" text="同步中" type="primary" size="mini"></u-tag>
					<u-tag v-else-if="item.status === 'success'" text="成功" type="success" size="mini"></u-tag>
					<u-tag v-else text="失败" type="danger" size="mini"></u-tag>
				</view>
			</view>

			<u-empty v-if="syncList.length === 0" text="暂无同步记录" mode="list"></u-empty>
		</view>
	</view>
</template>

<script>
	import { checkSync, syncRecords, getSyncStatus } from '@/api/sync.js'
	import { listRecord } from '@/api/record.js'

	export default {
		data() {
			return {
				pendingCount: 0,
				syncingCount: 0,
				syncedCount: 0,
				syncing: false,
				syncList: []
			}
		},
		onLoad() {
			this.getSyncStatus()
		},
		onShow() {
			this.getSyncStatus()
		},
		methods: {
			// 获取同步状态
			getSyncStatus() {
				// 从本地缓存读取待同步数据
				const pendingData = uni.getStorageSync('pending_sync_data') || []
				this.pendingCount = pendingData.length
				this.syncingCount = 0
				this.syncedCount = 0

				// 加载同步记录列表
				this.loadSyncList()
			},

			// 加载同步列表
			loadSyncList() {
				// 从本地缓存读取同步记录
				const syncLogs = uni.getStorageSync('sync_logs') || []
				this.syncList = syncLogs.slice(0, 20) // 只显示最近 20 条
			},

			// 立即同步
			handleSync() {
				if (this.pendingCount === 0) {
					uni.$u.toast('暂无待同步数据')
					return
				}

				this.syncing = true

				// 获取待同步数据
				const pendingData = uni.getStorageSync('pending_sync_data') || []

				// 分类处理
				const records = pendingData.filter(item => item.type === 'record')
				const subjects = pendingData.filter(item => item.type === 'subject')

				let successCount = 0

				// 同步执法记录
				if (records.length > 0) {
					syncRecords(records.map(item => item.data)).then(res => {
						successCount += res.data?.successCount || 0
						this.handleSyncComplete()
					}).catch(() => {
						this.handleSyncComplete()
					})
				}

				// 同步监管单位
				if (subjects.length > 0) {
					syncSubjects(subjects.map(item => item.data)).then(res => {
						successCount += res.data?.successCount || 0
					})
				}
			},

			// 同步完成
			handleSyncComplete() {
				this.syncing = false
				uni.$u.toast('同步完成')
				this.getSyncStatus()
			},

			// 获取操作类型文本
			getActionText(action) {
				const actionMap = { create: '新增', update: '更新', delete: '删除' }
				return actionMap[action] || action
			},

			// 获取操作类型样式
			getActionType(action) {
				const typeMap = { create: 'primary', update: 'warning', delete: 'danger' }
				return typeMap[action] || 'info'
			},

			// 获取同步类型文本
			getSyncTypeText(syncType) {
				const typeMap = {
					'app_to_server': 'App→服务器',
					'server_to_app': '服务器→App'
				}
				return typeMap[syncType] || syncType
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-sync {
		padding: 30rpx;

		.sync-status {
			display: flex;
			justify-content: space-between;
			margin-bottom: 30rpx;

			.status-card {
				background: #fff;
				border-radius: 16rpx;
				padding: 30rpx;
				flex: 1;
				margin: 0 10rpx;
				text-align: center;

				&:first-child {
					margin-left: 0;
				}

				&:last-child {
					margin-right: 0;
				}

				.status-value {
					font-size: 48rpx;
					font-weight: bold;
					margin-bottom: 10rpx;
				}

				.status-label {
					font-size: 24rpx;
					color: #999;
				}
			}
		}

		.sync-actions {
			margin-bottom: 30rpx;
		}

		.sync-list {
			background: #fff;
			border-radius: 16rpx;
			padding: 30rpx;

			.list-header {
				margin-bottom: 20rpx;

				.title {
					font-size: 32rpx;
					font-weight: bold;
				}
			}

			.sync-item {
				padding: 20rpx 0;
				border-bottom: 1rpx solid #f0f0f0;

				&:last-child {
					border-bottom: none;
				}

				.item-header {
					display: flex;
					justify-content: space-between;
					align-items: center;
					margin-bottom: 10rpx;

					.item-table {
						font-size: 28rpx;
						font-weight: bold;
					}
				}

				.item-info {
					margin-bottom: 10rpx;

					.info-text {
						font-size: 24rpx;
						color: #666;
						margin-right: 20rpx;
					}
				}

				.item-footer {
					display: flex;
					justify-content: space-between;
					align-items: center;

					.time {
						font-size: 24rpx;
						color: #999;
					}
				}
			}
		}
	}
</style>
