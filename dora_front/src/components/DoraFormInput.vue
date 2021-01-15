<template>
    <div>
        <el-form :inline="true">
            <el-form-item :label="label">
                <el-input :placeholder="label_name"></el-input>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    export default {
        name: 'DoraFormInput',
        props: {
            label: String
        },
        data() {
            return {
                label_name:'label'
            }
        }

    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
