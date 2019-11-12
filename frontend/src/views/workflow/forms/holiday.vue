<template>
  <el-form
    :model="ruleForm"
    :rules="rules"
    ref="ruleForm"
    label-width="100px"
    class="demo-ruleForm"
  >
    <el-form-item label="请假开始时间" prop="datetime">
      <el-date-picker
        v-model="datetime"
        type="datetimerange"
        range-separator="至"
        start-placeholder="开始日期"
        end-placeholder="结束日期"
        @change="countHours"
      ></el-date-picker>
      <p>总时长：{{hours}} 小时</p>
    </el-form-item>
    <el-form-item label="请假类型" prop="type">
      <el-radio-group v-model="ruleForm.type">
        <el-radio label="1">年假</el-radio>
        <el-radio label="2">病假</el-radio>
        <el-radio label="3">产假</el-radio>
        <el-radio label="4">婚假</el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
      <el-button @click="resetForm('ruleForm')">重置</el-button>
    </el-form-item>
  </el-form>
</template>
<script>
export default {
  data() {
    return {
      datetime: [],
      hours: 0,
      ruleForm: {
        starttime: "",
        endtime: "",
        type: "3"
      },
      rules: {
        starttime: [{ required: true, message: "请填写内容", trigger: "blur" }],
        endtime: [{ required: true, message: "请填写内容", trigger: "blur" }],
        type: [{ required: true, message: "请填写内容", trigger: "change" }]
      }
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          alert("submit!");
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    countHours(val){
      console.log(val)
    }
  }
};
</script>