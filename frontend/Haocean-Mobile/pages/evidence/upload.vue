<!-- 证据上传页面 -->
<template>
	<view class="t-evidence-upload">
		<u-form ref="form" :model="form" label-width="160">
			<u-form-item label="证据类型">
				<u--input v-model="form.type" placeholder="请选择证据类型" disabled @click="showTypePicker = true" />
				<u-icon slot="right" name="arrow-right"></u-icon>
			</u-form-item>

			<u-form-item label="证据标题">
				<u--input v-model="form.title" placeholder="请输入证据标题" />
			</u-form-item>

			<u-form-item label="证据描述">
				<u--textarea v-model="form.description" placeholder="请输入证据描述" :maxlength="500" count></u--textarea>
			</u-form-item>

			<u-form-item label="选择文件">
				<u-upload :fileList="fileList" :maxCount="9" @chooseFile="handleChooseFile"></u-upload>
			</u-form-item>
		</u-form>

		<view class="action-btns">
			<u-button type="primary" text="上传" @click="handleSubmit" :loading="loading"></u-button>
		</view>

		<!-- 类型选择器 -->
		<u-picker :show="showTypePicker" :list="typeOptions" key-name="label" value-name="value" @confirm="handleTypeConfirm" @cancel="showTypePicker = false"></u-picker>
	</view>
</template>

<script>
	import { uploadEvidence } from '@/api/record.js'

	export default {
		data() {
			return {
				form: {
					recordId: 0,
					type: 'photo',
					title: '',
					description: ''
				},
				fileList: [],
				selectedFile: null,
				showTypePicker: false,
				typeOptions: [
					{ label: '照片', value: 'photo' },
					{ label: '录音', value: 'audio' },
					{ label: '录像', value: 'video' },
					{ label: '文档', value: 'document' }
				],
				loading: false
			}
		},
		onLoad(options) {
			if (options.recordId) {
				this.form.recordId = parseInt(options.recordId)
			}
		},
		methods: {
			// 选择类型
			handleTypeConfirm(e) {
				this.form.type = e.value
			},

			// 选择文件
			handleChooseFile(files) {
				this.selectedFile = files[0]
				this.fileList = files
				if (!this.form.title) {
					this.form.title = this.selectedFile.name || '证据文件'
				}
			},

			// 上传
			handleSubmit() {
				if (!this.selectedFile) {
					uni.$u.toast('请选择文件')
					return
				}

				this.loading = true
				const formData = new FormData()
				formData.append('file', this.selectedFile.url)
				formData.append('recordId', this.form.recordId)
				formData.append('type', this.form.type)
				formData.append('title', this.form.title)
				formData.append('description', this.form.description || '')

				uploadEvidence(formData).then(res => {
					this.loading = false
					uni.$u.toast('上传成功')
					setTimeout(() => {
						uni.navigateBack()
					}, 1500)
				}).catch(() => {
					this.loading = false
				})
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-evidence-upload {
		padding: 30rpx;

		.action-btns {
			margin-top: 60rpx;
		}
	}
</style>
