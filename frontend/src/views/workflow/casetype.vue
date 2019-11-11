<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input
        v-model="listQuery.key"
        placeholder="请输入内容"
        clearable
        prefix-icon="el-icon-search"
        style="width: 200px;"
        class="filter-item"
        @keyup.enter.native="handleFilter"
        @clear="handleFilter"
      />
      <el-button-group>
        <el-button
          class="filter-item"
          type="primary"
          icon="el-icon-search"
          @click="handleFilter"
        >{{ "搜索" }}</el-button>
        <el-button
          v-if="permissionList.add"
          class="filter-item"
          type="success"
          icon="el-icon-edit"
          @click="handleCreate"
        >{{ "添加" }}</el-button>
        <el-button
          v-if="permissionList.del"
          class="filter-item"
          type="danger"
          icon="el-icon-delete"
          @click="handleBatchDel"
        >{{ "删除" }}</el-button>
      </el-button-group>
    </div>

    <el-table
      :key="tableKey"
      v-loading="listLoading"
      stripe
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
      @sort-change="sortChange"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" width="55" />
      <el-table-column label="名称" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.name }}</span>
        </template>
      </el-table-column>
      <el-table-column label="表单" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.form }}</span>
        </template>
      </el-table-column>
      <el-table-column label="状态" prop="status" sortable="custom" align="center">
        <template slot-scope="scope">
          <span>{{ scope.row.status | statusFilter }}</span>
        </template>
      </el-table-column>
      <el-table-column label="操作" align="center" width="330" class-name="small-padding fixed-width">
        <template slot-scope="{ row }">
          <el-button-group>
            <el-button
              v-if="permissionList.update"
              size="small"
              type="primary"
              @click="handleUpdate(row.id)"
            >{{ "编辑" }}</el-button>
            <el-button
              v-if="permissionList.del"
              size="small"
              type="danger"
              @click="handleDelete(row)"
            >{{ "删除" }}</el-button>
            <el-button
              v-if="permissionList.add"
              size="small"
              type="warning"
              @click="handleCreateFlow(row)"
            >{{ "编排步骤" }}</el-button>
          </el-button-group>
        </template>
      </el-table-column>
    </el-table>

    <div class="table-pagination">
      <pagination
        v-show="total > 0"
        :total="total"
        :page.sync="listQuery.page"
        :limit.sync="listQuery.limit"
        @pagination="getList"
      />
    </div>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogCaseVisible">
      <el-form
        ref="dataForm"
        v-loading="loading"
        element-loading-text="正在执行"
        element-loading-background="rgba(255,255,255,0.7)"
        :rules="rules"
        :model="temp"
        label-position="left"
        label-width="80px"
        style="width: 400px; margin-left:50px;"
      >
        <el-form-item label="名称" prop="name">
          <el-input v-model="temp.name" />
        </el-form-item>
        <el-form-item label="表单" prop="form">
          <el-select v-model="temp.form" placeholder="操作类型">
            <el-option
              v-for="item in caseforms"
              :key="item.code"
              :label="item.name"
              :value="item.code"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model.number="temp.status" type="number">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="2">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="备注" prop="memo">
          <el-input v-model="temp.memo" />
        </el-form-item>
      </el-form>
      <div
        v-if="
          dialogStatus !== 'detail' ? (loading === true ? false : true) : false
        "
        slot="footer"
        class="dialog-footer"
      >
        <el-button @click="dialogCaseVisible = false">{{ "取消" }}</el-button>
        <el-button
          type="primary"
          @click="dialogStatus === 'create' ? createData() : updateData()"
        >{{ "确定" }}</el-button>
      </div>
    </el-dialog>

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFlowVisible">
      <el-form
        ref="dynamicflowForm"
        element-loading-background="rgba(255,255,255,0.7)"
        :model="dynamicflowForm"
        label-position="left"
        label-width="100px"
      >
        <el-steps :active="active">
          <el-step title="提交工作流"></el-step>
          <el-step
            v-for="node in dynamicflowForm.nodes"
            :key="node.key"
            :title="node.name"
            :description="filterUsers[node.user_id]"
          ></el-step>
        </el-steps>
        <el-divider>
          <i class="el-icon-s-opportunity"></i>
        </el-divider>

        <el-form-item
          v-for="(node,index) in dynamicflowForm.nodes"
          :label="node.name + '步骤'"
          :value="node.user_id"
          :key="node.key"
          :prop="'nodes.' + index + '.value'"
          :rules="{required: true, message: node.name + '不能为空', trigger: 'blur'}"
        >
          <el-select v-model="node.user_id" placeholder="请选择用户">
            <el-option v-for="item in users" :key="item.id" :label="item.username" :value="item.id"></el-option>
          </el-select>
          <el-button plain type="danger" icon="el-icon-delete" @click.prevent="removeDomain(node)"></el-button>
        </el-form-item>
        <el-form-item>
          <el-button-group>
            <el-button type="success" @click="addDomain(1)">新增审核人</el-button>
            <el-button type="warning" @click="addDomain(2)">新增执行人</el-button>
            <el-button type="primary" @click="submitflowForm('dynamicflowForm')">提交</el-button>
          </el-button-group>
        </el-form-item>
      </el-form>
      <div
        v-if="
          dialogStatus !== 'detail' ? (loading === true ? false : true) : false
        "
        slot="footer"
        class="dialog-footer"
      >
        <el-button @click="dialogCaseVisible = false">{{ "取消" }}</el-button>
        <el-button
          type="primary"
          @click="dialogStatus === 'create' ? createData() : updateData()"
        >{{ "确定" }}</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { requestMenuButton } from "@/api/sys/menu";
import {
  requestList,
  requestDetail,
  requestUpdate,
  requestCreate,
  requestDelete,
  requestCreateSteps
} from "@/api/workflow/casetype";
import * as caseform from "@/api/workflow/caseform";
import * as user from "@/api/sys/user";
import Pagination from "@/components/Pagination"; // Secondary package based on el-pagination
import SelectTree from "@/components/TreeSelect";
import {
  checkAuthAdd,
  checkAuthDel,
  checkAuthView,
  checkAuthUpdate
} from "@/utils/permission";

export default {
  name: "CaseType",
  components: { Pagination, SelectTree },
  filters: {
    statusFilter(status) {
      const statusMap = {
        1: "启用",
        2: "禁用"
      };
      return statusMap[status];
    }
  },
  data() {
    return {
      caseforms: [],
      operationList: [],
      permissionList: {
        add: false,
        del: false,
        view: false,
        update: false
      },
      tableKey: 0,
      list: [],
      total: 0,
      listLoading: true,
      loading: true,
      listQuery: {
        page: 1,
        limit: 20,
        key: undefined,
        sort: "-id"
      },
      temp: {
        id: 0,
        name: "",
        form: "",
        status: 2,
        memo: ""
      },
      dialogCaseVisible: false,
      dialogFlowVisible: false,
      dialogStatus: "",
      textMap: {
        update: "编辑",
        create: "添加",
        detail: "详情"
      },
      rules: {
        name: [{ required: true, message: "请输入名称", trigger: "blur" }],
        form: [{ required: true, message: "请选择表单", trigger: "change" }]
      },
      multipleSelection: [],
      active: 1,
      users: [],
      dynamicflowForm: {},
      casetype_id: 0,
      filterUsers: {}
    };
  },
  computed: {},
  created() {
    this.getMenuButton();
    this.getList();
    this.getFormList();
    this.getUserList();
  },
  methods: {
    checkPermission() {
      this.permissionList.add = checkAuthAdd(this.operationList);
      this.permissionList.del = checkAuthDel(this.operationList);
      this.permissionList.view = checkAuthView(this.operationList);
      this.permissionList.update = checkAuthUpdate(this.operationList);
    },
    getMenuButton() {
      requestMenuButton("Menu")
        .then(response => {
          this.operationList = response.data;
        })
        .then(() => {
          this.checkPermission();
        });
    },
    getList() {
      this.listLoading = true;
      requestList(this.listQuery).then(response => {
        this.list = response.data.items;
        this.total = response.data.total;
        this.listLoading = false;
      });
    },
    getFormList() {
      caseform.requestList().then(response => {
        this.caseforms = response.data.items;
      });
    },
    getUserList() {
      user.requestList().then(response => {
        this.users = response.data.items;
        for (var user of this.users) {
          this.filterUsers[user.id] = user.username
        }
      });
    },
    handleFilter() {
      this.listQuery.page = 1;
      this.getList();
    },
    sortChange(data) {
      const { prop, order } = data;
      if (order === "ascending") {
        this.listQuery.sort = "+" + prop;
      } else if (order === "descending") {
        this.listQuery.sort = "-" + prop;
      } else {
        this.listQuery.sort = undefined;
      }
      this.handleFilter();
    },
    resetTemp() {
      this.temp = {
        id: 0,
        name: "",
        status: 1,
        memo: ""
      };
    },
    handleCreate() {
      this.resetTemp();
      this.dialogStatus = "create";
      this.dialogCaseVisible = true;
      this.loading = false;
      this.$nextTick(() => {
        this.$refs["dataForm"].clearValidate();
      });
    },
    createData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          this.loading = true;
          requestCreate(this.temp)
            .then(response => {
              this.temp.id = response.data.id;
              this.list.unshift(this.temp);
              this.dialogCaseVisible = false;
              this.$notify({
                title: "成功",
                message: "创建成功",
                type: "success",
                duration: 2000
              });
              this.total = this.total + 1;
              this.getList();
            })
            .catch(() => {
              this.loading = false;
            });
        }
      });
    },
    handleDetail(id) {
      this.loading = true;
      requestDetail(id).then(response => {
        this.loading = false;
        this.temp = response.data;
      });
      this.dialogStatus = "detail";
      this.dialogCaseVisible = true;
      this.$nextTick(() => {
        this.$refs["dataForm"].clearValidate();
      });
    },
    handleUpdate(id) {
      this.loading = true;
      requestDetail(id).then(response => {
        this.loading = false;
        this.temp = response.data;
      });
      this.dialogStatus = "update";
      this.dialogCaseVisible = true;
      this.$nextTick(() => {
        this.$refs["dataForm"].clearValidate();
      });
    },
    updateData() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          this.loading = true;
          const tempData = Object.assign({}, this.temp);
          requestUpdate(tempData)
            .then(() => {
              for (const v of this.list) {
                if (v.id === this.temp.id) {
                  const index = this.list.indexOf(v);
                  this.list.splice(index, 1, this.temp);
                  break;
                }
              }
              this.dialogCaseVisible = false;
              this.$notify({
                title: "成功",
                message: "更新成功",
                type: "success",
                duration: 2000
              });
              this.getList();
            })
            .catch(() => {
              this.loading = false;
            });
        }
      });
    },
    handleDelete(row) {
      var ids = [];
      ids.push(row.id);
      this.$confirm("是否确定删除?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          requestDelete(ids).then(() => {
            this.$message({
              message: "删除成功",
              type: "success"
            });
            this.total = this.total - 1;
            const index = this.list.indexOf(row);
            this.list.splice(index, 1);
            this.getList();
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    },
    handleCreateFlow(row) {
      this.dialogFlowVisible = true;

      this.casetype_id=row.id
      this.dynamicflowForm = {
        nodes: [
          {
            user_id: 0,
            name: "审核人",
            type: 1,
            casetype_id: this.casetype_id
          }
        ]
      }
    },
    removeDomain(item) {
      this.active -= 1;
      var index = this.dynamicflowForm.nodes.indexOf(item);
      if (index !== -1) {
        this.dynamicflowForm.nodes.splice(index, 1);
      }
    },
    addDomain(val) {
      this.active += 1;
      if (val === 1) {
        this.dynamicflowForm.nodes.push({
          user_id: 0,
          name: "审核人",
          type: val,
          casetype_id: this.casetype_id,
          key: Date.now()
        });
      } else {
        this.dynamicflowForm.nodes.push({
          user_id: 0,
          name: "执行人",
          type: val,
          casetype_id: this.casetype_id,
          key: Date.now()
        });
      }
    },
    submitflowForm(formName) {
      requestCreateSteps(this.dynamicflowForm.nodes);
    },
    handleSelectionChange(val) {
      this.multipleSelection = val;
    },
    handleBatchDel() {
      if (this.multipleSelection.length === 0) {
        this.$message({
          message: "未选中任何行",
          type: "warning",
          duration: 2000
        });
        return;
      }
      var ids = [];
      for (const v of this.multipleSelection) {
        ids.push(v.id);
      }
      this.$confirm("是否确定删除?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          requestDelete(ids).then(() => {
            this.$message({
              message: "删除成功",
              type: "success"
            });
            for (const row of this.multipleSelection) {
              this.total = this.total - 1;
              const index = this.list.indexOf(row);
              this.list.splice(index, 1);
            }
            this.getList();
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除"
          });
        });
    }
  }
};
</script>
