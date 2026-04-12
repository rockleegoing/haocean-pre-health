<template>
  <view class="mobile-item-container">
    <Navbar :title="userId ? '修改用户' : '新增用户'" bgColor="#fff" :h5Show="false"></Navbar>
    <u--form ref="form" :model="userInfo" :rules="rules" labelPosition="left" labelWidth="80">
		<u-form-item label="用户名字" prop="userName" borderBottom>
		  <u-input v-model="userInfo.userName" placeholder="请输入用户名字" :maxlength="30" border="none"></u-input>
		</u-form-item>
      <u-form-item label="用户昵称" prop="nickName" borderBottom>
        <u-input v-model="userInfo.nickName" placeholder="请输入用户昵称" :maxlength="30" border="none"></u-input>
      </u-form-item>
      <u-form-item label="归属部门" prop="deptId" borderBottom @click="actionShow = true;">
        <u--input v-model="noticeTypeName" disabled disabledColor="#ffffff" placeholder="请选择类型" border="none"></u--input>
				<u-icon slot="right" name="arrow-right"></u-icon>
      </u-form-item>
      <u-form-item label="状态" prop="status" borderBottom>
        <u-radio-group v-model="userInfo.status">
          <u-radio shape="circle" label="正常" name="0" checked></u-radio>
          <u-radio shape="circle" label="关闭" name="1"></u-radio>
        </u-radio-group>
      </u-form-item>
      <u-form-item label="备注" prop="remark">
        <u--textarea v-model="userInfo.remark" placeholder="请输入备注" :count="true" :maxlength="600" confirmType="done"></u--textarea>
      </u-form-item>
    </u--form>
    <u-action-sheet title="部门选择" :show="actionShow" @close="actionShow = false">
      <qian-tree ref="tkitree" confirmColor="#4e8af7" />
    </u-action-sheet>
    <u-row :gutter="16" style="margin-top: 36px;">
      <u-col :span="6">
        <u-button v-if="userId" type="error" text="删除" @click="del"></u-button>
        <u-button v-else icon="arrow-left" text="返回" plain @click="goBack()"></u-button>
      </u-col>
      <u-col :span="6">
        <u-button type="primary" text="提交" @click="submit"></u-button>
      </u-col>
    </u-row>
  </view>
</template>

<script setup>
import * as UserManageApi from '@/api/work/userManage'
import Navbar from '@/components/navbar/Navbar'
import qianTree from "@/components/qian-tree/qian-tree.vue"

export default {
  components: {
    Navbar,
    qianTree
  },
  data () {
    return {
      userId: undefined,
      userInfo: {
        nickName: '',
        status: '0',
        remark: ''
      },
      actionShow: false,
      actionTitle: '',
      noticeTypeName: null,
      rules: {
        nickName: [ { required: true, message: '请输入昵称', trigger: ['blur', 'change'] } ],
      }
    }
  },
  onLoad (params) {
    this.userId = params.id
    this.loadData()
  },
  methods: {
    loadData () {
      if (this.userId) {
        const app = this
        UserManageApi.userById(this.userId).then(res => {
          app.userInfo = res.data
		  this.noticeTypeName =  res.data.dept.deptName;
        })
      }
    },
    del () {
      NoticeApi.noticeDelete(this.userId).then(res => {
        uni.showToast({ title: '保存成功！' })
      })
    },
    submit () {
      this.$refs.noticeForm.validate().then(res => {
        if (this.userId) {
          NoticeApi.noticeModify(this.notice).then(res => {
            uni.showToast({ title: '提交成功！' })
          })
        } else {
          NoticeApi.noticeAdd(this.notice).then(res => {
            uni.showToast({ title: '提交成功！' })
          })
        }
      });
    },
    actionSelect (item) {
      this.noticeTypeName = item.name;
      this.userInfo.noticeType = item.value;
      this.$refs.noticeForm.validateField('noticeType');
    },
    goBack () {
      uni.navigateBack({ delta: 1});
    }
  }
}
</script>