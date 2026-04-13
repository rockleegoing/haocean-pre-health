<!-- 执法记录详情页面 -->
<template>
	<view class="t-record-detail">
		<!-- 基本信息 -->
		<view class="card">
			<view class="card-header">
				<text class="title">基本信息</text>
				<u-tag :text="getStatusText(record.status)" :type="getStatusType(record.status)" size="mini"></u-tag>
			</view>
			<view class="card-body">
				<view class="info-row">
					<text class="label">记录编号：</text>
					<text class="value">{{ record.recordNo }}</text>
				</view>
				<view class="info-row">
					<text class="label">监管单位：</text>
					<text class="value">{{ record.subjectName }}</text>
				</view>
				<view class="info-row">
					<text class="label">行业分类：</text>
					<text class="value">{{ record.industryName }}</text>
				</view>
				<view class="info-row">
					<text class="label">检查日期：</text>
					<text class="value">{{ record.checkDate }}</text>
				</view>
				<view class="info-row">
					<text class="label">检查类型：</text>
					<text class="value">{{ record.checkType }}</text>
				</view>
			</view>
		</view>

		<!-- 检查情况 -->
		<view class="card">
			<view class="card-header">
				<text class="title">检查情况</text>
			</view>
			<view class="card-body">
				<text class="content">{{ record.checkSituation || '暂无' }}</text>
			</view>
		</view>

		<!-- 处理意见 -->
		<view class="card">
			<view class="card-header">
				<text class="title">处理意见</text>
			</view>
			<view class="card-body">
				<text class="content">{{ record.processOpinion || '暂无' }}</text>
			</view>
		</view>

		<!-- 证据材料 -->
		<view class="card">
			<view class="card-header">
				<text class="title">证据材料 ({{ evidenceList.length }})</text>
			</view>
			<view class="card-body">
				<view class="evidence-list">
					<view class="evidence-item" v-for="(item, index) in evidenceList" :key="index" @click="handleViewEvidence(item)">
						<image v-if="item.type === 'photo'" :src="item.filePath" mode="aspectFill" class="evidence-img"></image>
						<view v-else class="evidence-file">
							<u-icon name="file" size="60"></u-icon>
							<text class="file-name">{{ item.title }}</text>
						</view>
						<text class="evidence-title">{{ item.title }}</text>
					</view>
				</view>
			</view>
		</view>

		<!-- 操作按钮 -->
		<view class="action-btns" v-if="record.status === 1">
			<u-button type="primary" text="上报" @click="handleSubmit"></u-button>
		</view>
	</view>
</template>

<script>
	import { getRecord, submitRecord } from '@/api/record.js'

	export default {
		data() {
			return {
				recordId: 0,
				record: {},
				evidenceList: []
			}
		},
		onLoad(options) {
			if (options.id) {
				this.recordId = parseInt(options.id)
				this.loadRecord()
			}
		},
		methods: {
			// 加载记录
			loadRecord() {
				getRecord(this.recordId).then(res => {
					this.record = res.data.record
					this.evidenceList = res.data.evidences || []
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

			// 查看证据
			handleViewEvidence(item) {
				if (item.type === 'photo') {
					uni.previewImage({
						urls: [item.filePath],
						current: 0
					})
				} else {
					uni.navigateTo({
						url: item.filePath
					})
				}
			},

			// 上报
			handleSubmit() {
				uni.$u.modal('确认上报', '确定要上报该执法记录吗？').then(() => {
					return submitRecord(this.recordId)
				}).then(() => {
					uni.$u.toast('上报成功')
					this.loadRecord()
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-record-detail {
		padding: 30rpx;

		.card {
			background: #fff;
			border-radius: 16rpx;
			margin-bottom: 30rpx;

			.card-header {
				padding: 30rpx;
				border-bottom: 1rpx solid #f0f0f0;
				display: flex;
				justify-content: space-between;
				align-items: center;

				.title {
					font-size: 32rpx;
					font-weight: bold;
				}
			}

			.card-body {
				padding: 30rpx;

				.info-row {
					display: flex;
					margin-bottom: 20rpx;

					.label {
						color: #666;
						width: 160rpx;
						flex-shrink: 0;
					}

					.value {
						color: #333;
						flex: 1;
					}
				}

				.content {
					color: #333;
					line-height: 1.6;
				}

				.evidence-list {
					display: flex;
					flex-wrap: wrap;
					gap: 20rpx;

					.evidence-item {
						width: 200rpx;
						text-align: center;

						.evidence-img {
							width: 200rpx;
							height: 200rpx;
							border-radius: 10rpx;
							margin-bottom: 10rpx;
						}

						.evidence-file {
							width: 200rpx;
							height: 200rpx;
							background: #f5f5f5;
							border-radius: 10rpx;
							display: flex;
							flex-direction: column;
							justify-content: center;
							align-items: center;
							margin-bottom: 10rpx;

							.file-name {
								font-size: 24rpx;
								color: #666;
								margin-top: 10rpx;
							}
						}

						.evidence-title {
							font-size: 24rpx;
							color: #666;
							overflow: hidden;
							text-overflow: ellipsis;
							white-space: nowrap;
						}
					}
				}
			}
		}

		.action-btns {
			margin-top: 40rpx;
		}
	}
</style>
