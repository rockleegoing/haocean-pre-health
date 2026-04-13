<!-- 添加监管单位页面 -->
<template>
	<view class="t-subject-add">
		<u-form ref="form" :model="form" :rules="rules" label-width="160">
			<u-form-item label="单位名称" prop="name">
				<u--input v-model="form.name" placeholder="请输入单位名称" />
			</u-form-item>

			<u-form-item label="行业分类" prop="industryId">
				<u--input v-model="form.industryName" placeholder="请选择行业分类" disabled @click="showIndustryPicker = true" />
				<u-icon slot="right" name="arrow-right"></u-icon>
			</u-form-item>

			<u-form-item label="地址" prop="address">
				<u--input v-model="form.address" placeholder="请输入地址" />
			</u-form-item>

			<u-form-item label="联系人" prop="contactPerson">
				<u--input v-model="form.contactPerson" placeholder="请输入联系人" />
			</u-form-item>

			<u-form-item label="联系电话" prop="contactPhone">
				<u--input v-model="form.contactPhone" placeholder="请输入联系电话" type="number" />
			</u-form-item>

			<u-form-item label="许可证号" prop="licenseNo">
				<u--input v-model="form.licenseNo" placeholder="请输入许可证号" />
			</u-form-item>

			<u-form-item label="经营状态">
				<u-switch v-model="form.status" :active-value="1" :inactive-value="0" slot="right"></u-switch>
			</u-form-item>
		</u-form>

		<view class="action-btns">
			<u-button type="primary" text="保存" @click="handleSubmit" :loading="loading"></u-button>
			<u-button text="取消" @click="handleCancel"></u-button>
		</view>

		<!-- 行业选择器 -->
		<u-picker :show="showIndustryPicker" :list="industryOptions" key-name="industryName" value-name="industryId" @confirm="handleIndustryConfirm" @cancel="showIndustryPicker = false"></u-picker>
	</view>
</template>

<script>
	import { addSubject, updateSubject, getSubject } from '@/api/subject.js'
	import { listIndustry } from '@/api/industry.js'

	export default {
		data() {
			return {
				form: {
					subjectId: undefined,
					name: '',
					industryId: 0,
					industryName: '',
					address: '',
					contactPerson: '',
					contactPhone: '',
					licenseNo: '',
					status: 1
				},
				rules: {
					name: [{ required: true, message: '单位名称不能为空', trigger: ['blur', 'change'] }],
					industryId: [{ required: true, message: '请选择行业分类', trigger: ['blur', 'change'] }]
				},
				industryOptions: [],
				showIndustryPicker: false,
				loading: false
			}
		},
		onLoad(options) {
			this.getIndustryList()
			if (options.id) {
				this.loadSubject(options.id)
			}
		},
		methods: {
			// 获取行业分类
			getIndustryList() {
				listIndustry({ isEnabled: 1 }).then(res => {
					this.industryOptions = res.rows || []
				})
			},

			// 加载单位信息
			loadSubject(id) {
				getSubject(id).then(res => {
					this.form = res.data
				})
			},

			// 选择行业
			handleIndustryConfirm(e) {
				const selected = e.value
				this.form.industryId = selected.industryId
				this.form.industryName = selected.industryName
				this.showIndustryPicker = false
			},

			// 提交
			handleSubmit() {
				this.$refs.form.validate(valid => {
					if (valid) {
						this.loading = true
						const api = this.form.subjectId ? updateSubject : addSubject
						api(this.form).then(res => {
							this.loading = false
							uni.$u.toast('保存成功')
							setTimeout(() => {
								uni.navigateBack()
							}, 1500)
						}).catch(() => {
							this.loading = false
						})
					}
				})
			},

			// 取消
			handleCancel() {
				uni.navigateBack()
			}
		}
	}
</script>

<style lang="scss" scoped>
	.t-subject-add {
		padding: 30rpx;

		.action-btns {
			margin-top: 60rpx;

			.u-button {
				margin-bottom: 20rpx;
			}
		}
	}
</style>
