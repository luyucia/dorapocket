<template>
    <div>
        <el-form :inline="true">
            <el-form-item :label="label">
                <el-input :placeholder="placeholder" v-model="$store.state.form_data[name]"></el-input>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    export default {
        name: 'DoraFormInput',
        props: {
            dora_id: String,
            label:String,
            name:String,
            placeholder:String
        },
        data() {
            return {
            }
        },
        beforeMount() {
            console.log(this.name)
            // let data = {}
            // data[this.dora_id] = {
            //     label:'文本框名称'
            // }
            // this.$store.commit('addComp',data)
        }

    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
