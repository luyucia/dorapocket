<template>
  <div  >
    <el-card class="box-card">
      <div  class="text item">
        {{$store.state.data}}
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'DoraPlainText',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .text {
    font-size: 14px;
  }

  /*.item {*/
    /*padding: 18px 0;*/
  /*}*/

  /*.box-card {*/
    /*width: 480px;*/
  /*}*/
</style>
