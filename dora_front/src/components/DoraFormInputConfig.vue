<template>
    <div>
        <el-form :inline="true">
            <el-form-item label="name">
                <el-input v-model="$store.state.layout[layout_index].component.name"></el-input>
            </el-form-item>
            <el-form-item label="label">
                <el-input v-model="$store.state.layout[layout_index].component.label"></el-input>
            </el-form-item>
            <el-form-item label="placeholder">
                <el-input v-model="$store.state.layout[layout_index].component.placeholder"></el-input>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    export default {
        name: 'DoraFormInputConfig',
        props: {
            layout_index: Number
        },
        methods: {
            // input:function (e) {
            //     let data = {
            //         label:e
            //     }
            //     let t = {}
            //     t.dora_id = this.dora_id
            //     t.data = data
            //     this.$store.commit('setComp',t)
            //     // $store.state.form_comp[dora_id].label
            //
            // }
        }

    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
