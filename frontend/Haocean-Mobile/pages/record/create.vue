<!-- 创建执法记录页面 -->
<template>
	<view class="t-record-create">
		<u-form ref="form" :model="form" :rules="rules" label-width="160">
			<u-form-item label="监管单位" prop="subjectId">
				<view class="selector" @click="showSubjectPicker = true">
					<text v-if="form.subjectId">{{ form.subjectName }}</text>
					<text v-else class="placeholder">请选择监管单位</text>
					<u-icon slot="right" name="arrow-right"></u-icon>
				</view>
			</u-form-item>

			<u-form-item label="检查日期" prop="checkDate">
				<u--input v-model="form.checkDate" type="date" placeholder="请选择检查日期" />
			</u-form-item>

			<u-form-item label="检查类型" prop="checkType">
				<u--input v-model="form.checkType" placeholder="请选择检查类型" disabled @click="showTypePicker = true" />
				<u-icon slot="right" name="arrow-right"></u-icon>
			</u-form-item>

			<u-form-item label="检查情况">
				<u--textarea v-model="form.checkSituation" placeholder="请输入检查情况" :maxlength="1000" count></u--textarea>
			</u-form-item>

			<u-form-item label="处理意见">
				<u--textarea v-model="form.processOpinion" placeholder="请输入处理意见" :maxlength="500" count></u--textarea>
			</u-form-item>

			<!-- 证据上传 -->
			<u-form-item label="证据材料">
				<u-upload :fileList="evidenceList" :maxCount="9" @chooseFile="handleChooseEvidence"></u-upload>
			</u-form-item>
		</u-form>

		<view class="action-btns">
			<u-button type="primary" text="保存" @click="handleSubmit" :loading="loading"></u-button>
			<u-button text="取消" @click="handleCancel"></u-button>
		</view>

		<!-- 单位选择器 -->
		<u-modal :show="showSubjectModal" :title="选择监管单位" @confirm="handleSubjectConfirm" @cancel="showSubjectModal = false">
			<view class="subject-selector">
				<u--input v-model="subjectKeyword" placeholder="搜索单位名称" @input="searchSubject"></u--input>
				<scroll-view scroll-y class="subject-list">
					<view class="subject-item" v-for="item in subjectOptions" :key="item.subjectId" @click="selectSubject(item)">
						{{ item.name }}
					</view>
				</scroll-view>
			</view>
		</u-modal>

		<!-- 检查类型选择器 -->
		<u-picker :show="showTypePicker" :list="checkTypeOptions" key-name="label" value-name="value" @confirm="handleTypeConfirm" @cancel="showTypePicker = false"></u-picker>
	</view>
</template>

<script>
	import { addRecord, updateRecord, getRecord, uploadEvidence } from '@/api/record.js'
	import { listSubject } from '@/api/subject.js'

	export default {
		data() {
			return {
				form: {
					recordId: undefined,
					subjectId: undefined,
					subjectName: '',
					industryId: undefined,
					industryName: '',
					checkDate: '',
					checkType: '',
					officialIds: [],
					checkSituation: '',
					processOpinion: ''
				},
				rules: {
					subjectId: [{ required: true, message: '请选择监管单位', trigger: ['blur', 'change'] }],
					checkDate: [{ required: true, message: '请选择检查日期', trigger: ['blur', 'change'] }],
					checkType: [{ required: true, message: '请选择检查类型', trigger: ['blur', 'change'] }]
				},
				subjectOptions: [],
				subjectKeyword: '',
				showSubjectModal: false,
				showTypePicker: false,
				checkTypeOptions: [
					{ label: '日常检查', value: '日常检查' },
					{ label: '专项检查', value: '专项检查' },
					{ label: '投诉举报', value: '投诉举报' },
					{ label: '双随机抽查', value: '双随机抽查' }
				],
				evidenceList: [],
				loading: false
			}
		},
		onLoad(options) {
			if (options.subjectId) {
				this.form.subjectId = parseInt(options.subjectId)
				this.loadSubjectInfo(options.subjectId)
			}
			if (options.recordId) {
				this.loadRecord(options.recordId)
			}
		},
		methods: {
			// 加载单位信息
			loadSubjectInfo(subjectId) {
				getSubject(subjectId).then(res => {
					const subject = res.data
					this.form.subjectName = subject.name
					this.form.industryId = subject.industryId
					this.form.industryName = subject.industryName
				})
			},

			// 加载记录
			loadRecord(recordId) {
				getRecord(recordId).then(res => {
					const record = res.data.record
					const evidences = res.data.evidences || []
					this.form = {
						...record,
						checkDate: record.checkDate || ''
					}
					this.evidenceList = evidences.map((e, index) => ({
						id: e.evidenceId,
						url: e.filePath,
						name: e.title
					}))
				})
			},

			// 选择单位
			selectSubject(item) {
				this.form.subjectId = item.subjectId
				this.form.subjectName = item.name
				this.form.industryId = item.industryId
				this.form.industryName = item.industryName
				this.showSubjectModal = false
			},

			// 搜索单位
			searchSubject() {
				if (this.subjectKeyword) {
					listSubject({ name: this.subjectKeyword, pageSize: 20 }).then(res => {
						this.subjectOptions = res.rows || []
					})
				} else {
					this.subjectOptions = []
				}
			},

			// 检查类型确认
			handleTypeConfirm(e) {
				this.form.checkType = e.value
			},

			// 选择证据
			handleChooseEvidence(files) {
				// 上传证据
				files.forEach(file => {
					const formData = new FormData()
					formData.append('file', file.url)
					formData.append('recordId', this.form.recordId || 0)
					formData.append('type', 'photo')
					uploadEvidence(formData).then(res => {
						uni.$u.toast('上传成功')
					})
				})
			},

			// 提交
			handleSubmit() {
				this.$refs.form.validate(valid => {
					if (valid) {
						this.loading = true
						const api = this.form.recordId ? updateRecord : addRecord
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
	.t-record-create {
		padding: 30rpx;

		.selector {
			display: flex;
			justify-content: space-between;
			align-items: center;
			width: 100%;

			.placeholder {
				color: #c0c4cc;
			}
		}

		.action-btns {
			margin-top: 60rpx;

			.u-button {
				margin-bottom: 20rpx;
			}
		}

		.subject-selector {
			.subject-list {
				max-height: 400rpx;

				.subject-item {
					padding: 20rpx;
					border-bottom: 1rpx solid #f0f0f0;

					&:active {
						background: #f5f5f5;
					}
				}
			}
		}
	}
</style>
