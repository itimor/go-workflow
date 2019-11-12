<template>
  <el-form
    :model="ruleForm"
    :rules="rules"
    ref="ruleForm"
    label-width="100px"
    class="demo-ruleForm"
  >
    <el-form-item label="发布项目" prop="name">
      <el-input v-model="ruleForm.name"></el-input>
    </el-form-item>
    <el-form-item label="发布环境" prop="env">
      <el-select v-model="ruleForm.env" placeholder="请选择发布版本">
        <el-option label="正式" value="prod"></el-option>
        <el-option label="测试" value="test"></el-option>
      </el-select>
    </el-form-item>
    <el-form-item label="发布版本" prop="version">
      <el-input v-model="ruleForm.version"></el-input>
    </el-form-item>
    <el-form-item label="发布内容" prop="content">
      <el-input type="textarea" v-model="ruleForm.content"></el-input>
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
      ruleForm: {
        name: "",
        env: "test",
        version: "",
        content: ""
      },
      rules: {
        name: [{ required: true, message: "请填写内容", trigger: "blur" }],
        env: [{ required: true, message: "请填写内容", trigger: "change" }],
        version: [{ required: true, message: "请填写内容", trigger: "change" }],
        content: [{ required: true, message: "请填写内容", trigger: "blur" }]
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
    }
  }
};
</script>