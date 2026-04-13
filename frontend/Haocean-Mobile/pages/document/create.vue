<!-- 文书制作页面 -->
<template>
	<view class="t-document-create">
		<!-- 选择模板 -->
		<view class="card">
			<view class="card-header">
				<text class="title">选择文书模板</text>
			</view>
			<view class="card-body">
				<view class="template-list">
					<view class="template-item" v-for="(item, index) in templateList" :key="index" @click="selectTemplate(item)">
						<view class="template-header">
							<text class="template-name">{{ item.templateName }}</text>
							<u-tag :text="item.version" type="primary" size="mini"></u-tag>
						</view>
						<view class="template-info">
							<text>分类：{{ item.categoryName }}</text>
						</view>
					</view>
				</view>
			</view>
		</view>

		<!-- 填写表单 -->
		<view class="card" v-if="selectedTemplate">
			<view class="card-header">
				<text class="title">填写信息</text>
			</view>
			<view class="card-body">
				<u-form ref="form" :model="formData" label-width="160">
					<u-form-item v-for="(field, index) in selectedTemplate.fields" :key="index" :label="field.name">
						<u--input v-model="formData[field.key]" :placeholder="'请输入' + field.name" />
					</u-form-item>
				</u-form>
			</view>
		</view>

		<!-- 操作按钮 -->
		<view class="action-btns">
			<u-button type="primary" text="生成文书" @click="handleGenerate" :loading="loading" :disabled="!selectedTemplate"></u-button>
		</view>
	</view>
</template>

<script>
	import { listTemplate, generateDocument } from '@/api/template.js'

	export default {
		data() {
			return {
				templateList: [],
				selectedTemplate: null,
				formData: {},
				loading: false
			}
		},
		onLoad(options) {
			if (options.recordId) {
				this.recordId = parseInt(options.recordId)
			}
			if (options.industryId) {
				this.getTemplateList(options.industryId)
			}
		},
		methods: {
			// 获取模板列表
			getTemplateList(industryId) {
				listTemplate({ industryId, isEnabled: 1 }).then(res => {
					this.templateList = res.rows || []
				})
			},

			// 选择模板
			selectTemplate(template) {
				this.selectedTemplate = template
				// 初始化表单数据
				if (template.fields && template.fields.length > 0) {
					template.fields.forEach(field => {
						this.formData[field.key] = ''
					})
				}
			},

			// 生成文书
			handleGenerate() {
				if (!this.selectedTemplate) {
					uni.$u.toast('请选择文书模板')
					return
				}

				this.loading = true
				generateDocument({
					recordId: this.recordId,
					templateId: this.selectedTemplate.templateId,
					fillData: this.formData
				}).then(res => {
					this.loading = false
					uni.$u.toast('生成成功')
					// 跳转到预览页面
					uni.navigateTo({
						url: '/pages/document/preview?id=' + res.data.docId
					})
				}).catch(() => {
					this.loading = false
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-document-create {
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

				.template-list {
					.template-item {
						padding: 20rpx;
						background: #f5f5f5;
						border-radius: 10rpx;
						margin-bottom: 15rpx;

						&:active {
							background: #e8e8e8;
						}

						.template-header {
							display: flex;
							justify-content: space-between;
							align-items: center;
							margin-bottom: 10rpx;

							.template-name {
								font-size: 30rpx;
								font-weight: bold;
							}
						}

						.template-info {
							font-size: 26rpx;
							color: #666;
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
