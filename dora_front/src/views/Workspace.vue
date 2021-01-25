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
                        {{$store.state.layout}}
                    </el-col>
                </el-row>
            </el-aside>
            <el-main>
                <grid-layout
                        :layout.sync="$store.state.layout"
                        :col-num="12"

                        :is-draggable="true"
                        :is-resizable="true"
                        :is-mirrored="false"
                        :vertical-compact="true"
                        :margin="[10, 10]"
                        :autoSize = true
                        :use-css-transforms="true"
                >

                    <grid-item v-for="(item,i) in $store.state.layout"
                               :x="item.x"
                               :y="item.y"
                               :w="item.w"
                               :h="item.h"
                               :i="item.i"
                               :key="item.i" class="cell"  >
                        <!--<DoraFrom></DoraFrom>-->
                        {{i}}
                        <component class="comp" :is="item.component.type"  v-bind="item.component" :dora_id="item.component.type" ></component>
                        <!--<span @click="clickComponent(i)"><i class="el-icon-s-tools"></i></span>-->
                    </grid-item>
                </grid-layout>
            </el-main>
            <el-aside>
                <!--<el-input v-model="$store.state.form_comp[1].label"  ></el-input>-->
                <component :is="current.component_config" :layout_index="current.layout_index" ></component>
                <!--<DoraFormInputConfig></DoraFormInputConfig>-->
                {{$store.state.form_data}}
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
    import DoraFormInputConfig from "../components/DoraFormInputConfig";

    export default {
        name: 'Workspace',
        components: {
            DoraFormInputConfig,
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
                current:{},
                test:"xxx"

            }
        },
        methods: {
            addComponent: function (c) {
                let data_len = this.$store.state.layout.length
                let y = 0;
                let index = 0;
                if (data_len != 0) {
                    index = data_len;
                    y = this.$store.state.layout[data_len - 1].y + this.$store.state.layout[data_len - 1].h
                }

                let component = {
                    type:c,
                    id:'c'+index,
                }

                switch (c) {
                    case 'DoraFormInput':
                        component.label = 'label'
                        component.placeholder = 'label'
                        component.name = 'name'
                        break
                    case '':
                        break
                    default:


                }
                // this.$store.state.layout.push

                this.$store.state.layout.push({"x": 0, "y": y, "w": 12, "h": 2, "i": index, "component": component,"dora_id":'c'+index})


            },
            clickComponent:function (i) {
                // let cid = c.dora_id
                // let comp = this.$refs[cid]
                //组件名+Config 命名规则

                this.current.component_config = this.$store.state.layout[i].component.type+'Config'
                this.current.layout_index = i
                // this.current = c.component
                console.log(i)
                    // this.$refs[c.dora_id]


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
    .comp{
        /*margin: 10px;*/
    }
</style>