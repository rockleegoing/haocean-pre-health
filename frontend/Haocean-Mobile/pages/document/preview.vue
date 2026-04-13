<!-- 文书预览页面 -->
<template>
	<view class="t-document-preview">
		<web-view :url="previewUrl"></web-view>
	</view>
</template>

<script>
	import { getTemplate } from '@/api/template.js'

	export default {
		data() {
			return {
				docId: 0,
				previewUrl: ''
			}
		},
		onLoad(options) {
			if (options.id) {
				this.docId = parseInt(options.id)
				this.loadDocument()
			}
		},
		methods: {
			// 加载文书
			loadDocument() {
				// 这里需要根据实际后端接口返回预览地址
				// 假设后端提供了预览接口
				const baseUrl = getApp().globalConfig.baseUrl
				this.previewUrl = `${baseUrl}/system/document/preview/${this.docId}`
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-document-preview {
		width: 100%;
		height: 100vh;
	}
</style>
