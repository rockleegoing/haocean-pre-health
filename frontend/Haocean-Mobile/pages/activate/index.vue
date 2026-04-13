<!-- 设备激活页面 -->
<template>
	<view class="t-activate">
		<view class="t-title">设备激活</view>
		<view class="t-subtitle">请输入激活码激活您的设备</view>

		<view class="activate-form">
			<!-- 激活码输入 -->
			<view class="form-item">
				<u--input label="激活码" v-model="activateCode" placeholder="请输入 8 位激活码" :maxlength="8" />
			</view>

			<!-- 设备信息展示 -->
			<view class="device-info" v-if="deviceInfo">
				<view class="info-item">
					<text class="label">设备型号：</text>
					<text class="value">{{ deviceInfo.deviceModel }}</text>
				</view>
				<view class="info-item">
					<text class="label">操作系统：</text>
					<text class="value">{{ deviceInfo.osType }} {{ deviceInfo.osVersion }}</text>
				</view>
			</view>

			<u-button type="primary" text="激活设备" @click="handleActivate" :loading="loading"></u-button>
		</view>
	</view>
</template>

<script>
	import { activateDevice } from '@/api/device.js'
	import { getSystemInfo } from '@/utils/system.js'

	export default {
		data() {
			return {
				activateCode: '',
				deviceInfo: null,
				loading: false
			}
		},
		onLoad() {
			this.getDeviceInfo()
		},
		methods: {
			// 获取设备信息
			getDeviceInfo() {
				const systemInfo = uni.getSystemInfoSync()
				this.deviceInfo = {
					deviceModel: systemInfo.model,
					osType: systemInfo.platform,
					osVersion: systemInfo.system,
					appVersion: '1.0.0'
				}
			},

			// 激活设备
			handleActivate() {
				if (!this.activateCode || this.activateCode.length !== 8) {
					uni.$u.toast('请输入 8 位激活码')
					return
				}

				this.loading = true
				const params = {
					activateCode: this.activateCode,
					...this.deviceInfo
				}

				activateDevice(params).then(res => {
					this.loading = false
					if (res.code === 200) {
						uni.$u.toast('激活成功')
						// 跳转到首页
						uni.switchTab({
							url: '/pages/index/index'
						})
					}
				}).catch(() => {
					this.loading = false
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-activate {
		padding: 40rpx;

		.t-title {
			font-size: 44rpx;
			font-weight: bold;
			text-align: center;
			margin-bottom: 20rpx;
		}

		.t-subtitle {
			font-size: 28rpx;
			color: #999;
			text-align: center;
			margin-bottom: 60rpx;
		}

		.activate-form {
			background: #fff;
			border-radius: 20rpx;
			padding: 40rpx;

			.form-item {
				margin-bottom: 30rpx;
			}

			.device-info {
				background: #f5f5f5;
				border-radius: 10rpx;
				padding: 20rpx;
				margin-bottom: 30rpx;

				.info-item {
					display: flex;
					justify-content: space-between;
					padding: 10rpx 0;

					.label {
						color: #666;
					}

					.value {
						color: #333;
						font-weight: 500;
					}
				}
			}
		}
	}
</style>
