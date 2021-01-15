<template>
    <div class="home">
        <h1>{{ $route.params.projectId }}</h1>


        <el-container>
            <el-aside>
                <el-row class="tac">
                    <el-col :span="24">
                        <h5>组件</h5>
                        <el-menu
                                default-active="2"
                                class="el-menu-vertical-demo"
                        >
                            <!--<el-submenu index="1">-->
                            <!--<template slot="title">-->
                            <!--<i class="el-icon-location"></i>-->
                            <!--<span>表单</span>-->
                            <!--</template>-->
                            <!--<el-menu-item index="1-1" @click="addComponent('DoraTable')">表格</el-menu-item>-->
                            <!--<el-menu-item index="1-2" @click="addComponent('DoraMarkdown')">Markdown</el-menu-item>-->
                            <!--<el-menu-item index="1-3" @click="addComponent('DoraForm')">表单</el-menu-item>-->
                            <!--<el-menu-item index="1-4" @click="addComponent('DoraPlainText')">文本</el-menu-item>-->
                            <!--<el-menu-item index="1-4" @click="addComponent('DoraInput')">输入框</el-menu-item>-->
                            <!--</el-submenu>-->
                            <el-menu-item index="1-1" @click="addComponent('DoraTable')">表格</el-menu-item>
                            <el-menu-item index="1-2" @click="addComponent('DoraMarkdown')">Markdown</el-menu-item>
                            <el-menu-item index="1-3" @click="addComponent('DoraForm')">表单</el-menu-item>
                            <el-menu-item index="1-4" @click="addComponent('DoraPlainText')">文本</el-menu-item>
                            <el-menu-item index="1-4" @click="addComponent('DoraFormInput')">输入框</el-menu-item>

                        </el-menu>
                        {{layout}}
                    </el-col>
                </el-row>
            </el-aside>
            <el-main>
                <grid-layout
                        :layout.sync="layout"
                        :col-num="12"
                        :row-height="30"
                        :is-draggable="true"
                        :is-resizable="true"
                        :is-mirrored="false"
                        :vertical-compact="true"
                        :margin="[10, 10]"
                        :autoSize = "true"
                        :use-css-transforms="true"
                >

                    <grid-item v-for="item in layout"
                               :x="item.x"
                               :y="item.y"
                               :w="item.w"
                               :h="item.h"
                               :i="item.i"
                               :key="item.i" class="cell"  >
                        <!--<DoraFrom></DoraFrom>-->
                        <component :is="item.component" :ref="'c_'+item.i" ></component>
                        <span @click="clickComponent(item)">修改{{'c_'+item.i}}</span>
                    </grid-item>
                </grid-layout>
            </el-main>
            <el-aside>
                <el-input v-model="current.lable_name"  ></el-input>
            </el-aside>
        </el-container>

    </div>
</template>

<script>
    // @ is an alias to /src
    import VueGridLayout from 'vue-grid-layout';
    import DoraForm from "../components/DoraForm";
    import HelloWorld from "../components/HelloWorld";
    import DoraTable from "../components/DoraTable";
    import DoraMarkdown from "../components/DoraMarkdown";
    import DoraPlainText from "../components/DoraPlainText";
    import DoraFormInput from "../components/DoraFormInput";

    export default {
        name: 'Workspace',
        components: {
            DoraFormInput,
            DoraPlainText,
            DoraMarkdown,
            DoraTable,
            HelloWorld,
            DoraForm,
            GridLayout: VueGridLayout.GridLayout,
            GridItem: VueGridLayout.GridItem
        },
        data() {
            return {
                layout: [
                    // {"x": 0, "y": 0, "w": 12, "h": 4, "i": "0","component":"DoraForm"},
                    // {"x": 0, "y": 4, "w": 12, "h": 12, "i": "1","component":"DoraTable"},
                    // {"x": 0, "y": 10, "w": 12, "h": 4, "i": "2","component":"DoraPlainText"},

                ],
                current:{}

            }
        },
        methods: {
            addComponent: function (c) {
                let data_len = this.layout.length
                let y = 0;
                let index = 0;
                if (data_len != 0) {
                    index = data_len;
                    y = this.layout[data_len - 1].y + this.layout[data_len - 1].h
                }

                this.layout.push({"x": 0, "y": y, "w": 12, "h": 2, "i": index, "component": c})

            },
            clickComponent:function (c) {
                let cid = "c_"+c.i
                // let comp = this.$refs[cid]
                this.current = this.$refs[cid]
                // console.log(this.$refs[cid])
                // console.log(comp.lable_name)
                // this.form = comp.data
            }
        }
    }
</script>

<style>
    .cell {
        /*background-color: aliceblue;*/
        margin: 10px;
    }
</style>