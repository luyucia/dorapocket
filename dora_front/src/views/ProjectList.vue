<template>
    <div class="projects">
        <el-container>
            <el-header id="dora-header">
                <el-dropdown>
                    <i class="el-icon-setting" style="margin-right: 15px"></i>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item>查看</el-dropdown-item>
                        <el-dropdown-item>新增</el-dropdown-item>
                        <el-dropdown-item>删除</el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
                <span>luyulewin</span>
            </el-header>
            <el-main>
                <el-row :gutter="12">
                    <el-col v-for="(p,i) in projectList" :span="8" v-bind:key="p.id" class="dora-pj-card">
                        <el-card shadow="hover">
                            <div style="padding: 14px;">
                                <span>{{p.name}}</span>
                                <div class="bottom clearfix">
                                    <!--<time class="time">{{ p.id }}</time>-->

                                    <router-link :to="'/workspace/'+p.id"><el-button type="text" class="button">进入</el-button></router-link>
                                    <el-button type="text" class="button"
                                               @click="drawer=true;drawer_status='update';currentProject = p">编辑
                                    </el-button>
                                    <el-button type="text" class="button" @click="deleteItem(i)">删除</el-button>
                                </div>
                            </div>
                        </el-card>
                    </el-col>

                    <el-col :span="8" class="dora-pj-card">

                        <el-card shadow="hover">
              <span @click="drawer=true;drawer_status = 'add';currentProject={} ">
                <i class="el-icon-circle-plus"></i>
                添加
              </span>
                        </el-card>
                    </el-col>

                </el-row>
            </el-main>
        </el-container>

        <el-drawer title="我是标题" :visible.sync="drawer" :with-header="false">
            <span>编辑</span>
            <el-form ref="form" label-width="80px" class="dora-drawer-form">
                <el-form-item label="项目名称">
                    <el-input v-model="currentProject.name"></el-input>
                </el-form-item>
                <el-form-item label="项目描述">
                    <el-input v-model="currentProject.description"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button v-if="drawer_status == 'update'" type="primary" @click="saveEdit()">修改</el-button>
                    <el-button v-if="drawer_status == 'add'" type="primary" @click="createItem()">创建</el-button>
                    <el-button @click="drawer=false">取消</el-button>
                </el-form-item>
            </el-form>

        </el-drawer>
    </div>
</template>

<script>
    import project from '../request/project'

    export default {
        name: 'ProjectList',
        components: {},
        data() {
            return {
                projectList: [],
                drawer: false,
                drawer_status: 'add',
                currentProject: {},

            }
        },
        mounted() {
            let app = this

            project.List().then(response => {
                app.projectList = response.data.data
            })
        },
        methods: {
            saveEdit: function () {
                let app = this
                project.Update(this.currentProject).then(res => {
                    console.log(res)
                    if (res.data.code == 0) {
                        app.$message({
                            message: '修改成功',
                            type: 'success'
                        });
                    }
                })

                this.drawer = false
            },
            deleteItem: function (i) {
                let app = this

                this.$confirm('此操作将永久删除 是否继续?', '警告', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    project.Delete(app.projectList[i].id).then(res => {
                        if (res.data.code == 0) {
                            app.projectList.splice(i, 1)
                            app.$message({
                                type: 'success',
                                message: '删除成功!'
                            });
                        }
                    })

                }).catch(() => {
                    app.$message({
                        type: 'info',
                        message: '已取消删除'
                    });
                });
            },
            createItem: function () {
                let app = this
                project.Create(app.currentProject).then(res => {
                    if (res.data.code == 0) {
                        app.drawer = false
                        app.currentProject.id = res.data.data
                        app.projectList.push(app.currentProject)
                    }
                })


            }
        }
    }
</script>

<style>
    #dora-header {
        text-align: right;
        background-color: #409EFF;
        padding-top: 20px;
        font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;

    }

    .dora-pj-card {
        margin-top: 10px;
    }

    .dora-drawer-form {
        width: 85%;
        padding-left: 20px;
        padding-top: 20px;
    }
</style>