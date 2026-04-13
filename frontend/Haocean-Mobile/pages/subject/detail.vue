<!-- 监管单位详情页面 -->
<template>
	<view class="t-subject-detail">
		<!-- 单位基本信息 -->
		<view class="card">
			<view class="card-header">
				<text class="title">单位信息</text>
			</view>
			<view class="card-body">
				<view class="info-row">
					<text class="label">单位名称：</text>
					<text class="value">{{ subject.name }}</text>
				</view>
				<view class="info-row">
					<text class="label">行业分类：</text>
					<text class="value">{{ subject.industryName }}</text>
				</view>
				<view class="info-row">
					<text class="label">经营状态：</text>
					<u-tag :text="subject.status === 1 ? '正常' : '停业'" :type="subject.status === 1 ? 'success' : 'danger'" size="mini"></u-tag>
				</view>
				<view class="info-row">
					<text class="label">地址：</text>
					<text class="value">{{ subject.address || '暂无' }}</text>
				</view>
				<view class="info-row">
					<text class="label">联系人：</text>
					<text class="value">{{ subject.contactPerson || '暂无' }}</text>
				</view>
				<view class="info-row">
					<text class="label">联系电话：</text>
					<text class="value" @click="handleCall">{{ subject.contactPhone || '暂无' }}</text>
				</view>
				<view class="info-row">
					<text class="label">许可证号：</text>
					<text class="value">{{ subject.licenseNo || '暂无' }}</text>
				</view>
			</view>
		</view>

		<!-- 执法记录入口 -->
		<view class="card">
			<view class="card-header">
				<text class="title">执法记录</text>
			</view>
			<view class="card-body">
				<view class="record-entry" @click="handleViewRecords">
					<text>查看该单位的所有执法记录</text>
					<u-icon name="arrow-right"></u-icon>
				</view>
			</view>
		</view>

		<!-- 操作按钮 -->
		<view class="action-btns">
			<u-button text="编辑" @click="handleEdit"></u-button>
			<u-button type="primary" text="新建执法记录" @click="handleCreateRecord"></u-button>
		</view>
	</view>
</template>

<script>
	import { getSubject } from '@/api/subject.js'

	export default {
		data() {
			return {
				subjectId: 0,
				subject: {}
			}
		},
		onLoad(options) {
			if (options.id) {
				this.subjectId = parseInt(options.id)
				this.loadSubject()
			}
		},
		methods: {
			// 加载单位信息
			loadSubject() {
				getSubject(this.subjectId).then(res => {
					this.subject = res.data
				})
			},

			// 拨打电话
			handleCall() {
				if (this.subject.contactPhone) {
					uni.makePhoneCall({
						phoneNumber: this.subject.contactPhone
					})
				}
			},

			// 查看记录
			handleViewRecords() {
				uni.navigateTo({
					url: '/pages/record/list?subjectId=' + this.subjectId
				})
			},

			// 编辑
			handleEdit() {
				uni.navigateTo({
					url: '/pages/subject/add?id=' + this.subjectId
				})
			},

			// 新建执法记录
			handleCreateRecord() {
				uni.navigateTo({
					url: '/pages/record/create?subjectId=' + this.subjectId
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-subject-detail {
		padding: 30rpx;

		.card {
			background: #fff;
			border-radius: 16rpx;
			margin-bottom: 30rpx;

			.card-header {
				padding: 30rpx;
				border-bottom: 1rpx solid #f0f0f0;

				.title {
					font-size: 32rpx;
					font-weight: bold;
				}
			}

			.card-body {
				padding: 30rpx;

				.info-row {
					display: flex;
					align-items: center;
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

				.record-entry {
					display: flex;
					justify-content: space-between;
					align-items: center;
					padding: 20rpx 0;
					color: #1890ff;
				}
			}
		}

		.action-btns {
			margin-top: 40rpx;

			.u-button {
				margin-bottom: 20rpx;
			}
		}
	}
</style>
